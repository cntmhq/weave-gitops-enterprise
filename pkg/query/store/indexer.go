package store

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	bleve "github.com/blevesearch/bleve/v2"
	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Indexer
type Indexer interface {
	IndexReader
	IndexWriter
}

//counterfeiter:generate . IndexWriter
type IndexWriter interface {
	Add(ctx context.Context, objects []models.Object) error
	Remove(ctx context.Context, objects []models.Object) error
	RemoveByQuery(ctx context.Context, q string) error
}

type Facets map[string][]string

//counterfeiter:generate . IndexReader
type IndexReader interface {
	Search(ctx context.Context, query Query, opts QueryOption) (Iterator, error)
	ListFacets(ctx context.Context) (Facets, error)
}

var indexFile = "index.db"

func NewIndexer(s Store, path string) (Indexer, error) {
	idxFileLocation := filepath.Join(path, indexFile)
	mapping := bleve.NewIndexMapping()

	// We do this to get our facets to appear correctly.
	// A keyword analyzer gives us the exact value of the field,
	// rather than splitting up things like `flux-system` into `flux` and `system`.
	// https://github.com/blevesearch/bleve/issues/639
	// Following this issue, we can probably expand this to include other fields.
	objMapping := bleve.NewDocumentMapping()
	nsMapping := bleve.NewTextFieldMapping()
	nsMapping.Analyzer = "keyword"

	clusterNameMapping := bleve.NewTextFieldMapping()
	clusterNameMapping.Analyzer = "keyword"
	objMapping.AddFieldMappingsAt("cluster", clusterNameMapping)

	objMapping.AddFieldMappingsAt("namespace", nsMapping)

	mapping.AddDocumentMapping("object", objMapping)

	index, err := bleve.New(idxFileLocation, mapping)
	if err != nil {
		return nil, fmt.Errorf("failed to create indexer: %w", err)
	}

	return &bleveIndexer{
		idx:   index,
		store: s,
	}, nil
}

type bleveIndexer struct {
	idx   bleve.Index
	store Store
}

func (i *bleveIndexer) Add(ctx context.Context, objects []models.Object) error {
	batch := i.idx.NewBatch()

	for _, obj := range objects {
		if err := batch.Index(obj.GetID(), obj); err != nil {
			return fmt.Errorf("failed to index object: %w", err)
		}
	}

	return i.idx.Batch(batch)
}

func (i *bleveIndexer) Remove(ctx context.Context, objects []models.Object) error {
	for _, obj := range objects {
		if err := i.idx.Delete(obj.GetID()); err != nil {
			return fmt.Errorf("failed to delete object: %w", err)
		}
	}

	return nil
}

func (i *bleveIndexer) RemoveByQuery(ctx context.Context, q string) error {
	query := bleve.NewQueryStringQuery(q)
	req := bleve.NewSearchRequest(query)

	result, err := i.idx.Search(req)
	if err != nil {
		return fmt.Errorf("failed to search index: %w", err)
	}

	for _, hit := range result.Hits {
		if err := i.idx.Delete(hit.ID); err != nil {
			return fmt.Errorf("failed to delete object: %w", err)
		}
	}

	return nil
}

func (i *bleveIndexer) Search(ctx context.Context, q Query, opts QueryOption) (Iterator, error) {
	// Match all by default.
	// Conjunction queries will return results that match all of the subqueries.
	query := bleve.NewConjunctionQuery(bleve.NewMatchAllQuery())

	terms := q.GetTerms()

	if terms != "" {
		tq := bleve.NewMatchQuery(terms)
		query.AddQuery(tq)
	}

	filters := q.GetFilters()

	if len(filters) > 0 {
		filterString := strings.Join(q.GetFilters(), " ")

		qs := bleve.NewQueryStringQuery(strings.ToLower(filterString))
		query.AddQuery(qs)
	}

	req := bleve.NewSearchRequest(query)

	sortBy := "name"
	tmpl := "-%v"

	if opts != nil {
		// `-` reverses the order
		if !opts.GetAscending() {
			tmpl = "%v"
		}

		sort := opts.GetOrderBy()
		if sort != "" {
			sortBy = sort
		}

		if opts.GetOffset() > 0 {
			req.From = int(opts.GetOffset())
		}
	}

	req.SortBy([]string{fmt.Sprintf(tmpl, sortBy)})

	searchResults, err := i.idx.Search(req)
	if err != nil {
		return nil, fmt.Errorf("failed to search for objects: %w", err)
	}

	iter := &indexerIterator{
		result: searchResults,
		s:      i.store,
		mu:     sync.Mutex{},
		index:  -1,
		opts:   opts,
	}

	return iter, nil
}

func (i *bleveIndexer) ListFacets(ctx context.Context) (Facets, error) {
	query := bleve.NewMatchAllQuery()

	req := bleve.NewSearchRequest(query)

	req.AddFacet("Kind", bleve.NewFacetRequest("kind", 100))
	req.AddFacet("Namespace", bleve.NewFacetRequest("namespace", 100))
	req.AddFacet("Cluster", bleve.NewFacetRequest("cluster", 100))
	// req.AddFacet("Namespace", bleve.NewFacetRequest("namespace.facet", 100))

	searchResults, err := i.idx.Search(req)
	if err != nil {
		return nil, fmt.Errorf("failed to search for objects: %w", err)
	}

	facets := map[string][]string{}

	for k, v := range searchResults.Facets {
		facets[k] = []string{}

		for _, t := range v.Terms.Terms() {
			facets[k] = append(facets[k], t.Term)
		}
	}

	return facets, nil
}

type indexerIterator struct {
	result *bleve.SearchResult
	mu     sync.Mutex
	index  int
	s      Store
	opts   QueryOption
}

func (i *indexerIterator) Next() bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.index++
	return i.index < len(i.result.Hits)
}

func (i *indexerIterator) Row() (models.Object, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	result := i.result.Hits[i.index]

	id := result.ID

	return i.s.GetObjectByID(context.Background(), id)
}

func (i *indexerIterator) Close() error {
	return nil
}

func (i *indexerIterator) All() ([]models.Object, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	ids := []string{}

	for _, hit := range i.result.Hits {
		ids = append(ids, hit.ID)
	}

	iter, err := i.s.GetObjects(context.Background(), ids, i.opts)

	if err != nil {
		return nil, fmt.Errorf("failed to get objects: %w", err)
	}

	return iter.All()
}
