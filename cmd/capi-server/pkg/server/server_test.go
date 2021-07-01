package server

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	capiv1 "github.com/weaveworks/wks/cmd/capi-server/api/v1alpha1"
	"github.com/weaveworks/wks/cmd/capi-server/pkg/capi"
	"github.com/weaveworks/wks/cmd/capi-server/pkg/git"
	capiv1_protos "github.com/weaveworks/wks/cmd/capi-server/pkg/protos"
	"github.com/weaveworks/wks/cmd/capi-server/pkg/templates"
	"google.golang.org/protobuf/testing/protocmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/yaml"
)

func TestListTemplates(t *testing.T) {
	testCases := []struct {
		name             string
		clusterState     []runtime.Object
		expected         []*capiv1_protos.Template
		err              error
		expectedErrorStr string
	}{
		{
			name:     "no configmap",
			err:      errors.New("configmap capi-templates not found in default namespace"),
			expected: []*capiv1_protos.Template{},
		},
		{
			name: "no templates",
			clusterState: []runtime.Object{
				makeTemplateConfigMap(),
			},
			expected: []*capiv1_protos.Template{},
		},
		{
			name: "1 template",
			clusterState: []runtime.Object{
				makeTemplateConfigMap("template1", makeTemplate(t)),
			},
			expected: []*capiv1_protos.Template{
				{
					Name:        "cluster-template-1",
					Description: "this is test template 1",
					Body:        "eyJhcGlWZXJzaW9uIjoiZm9vdmVyc2lvbiIsImtpbmQiOiJmb29raW5kIiwibWV0YWRhdGEiOnsibGFiZWxzIjp7Im5hbWUiOiIke0NMVVNURVJfTkFNRX0ifX19",
					Objects: []*capiv1_protos.TemplateObject{
						{
							ApiVersion: "fooversion",
							Kind:       "fookind",
							Parameters: []string{"CLUSTER_NAME"},
						},
					},
					Parameters: []*capiv1_protos.Parameter{
						{
							Name:        "CLUSTER_NAME",
							Description: "This is used for the cluster naming.",
						},
					},
				},
			},
		},
		{
			name: "2 templates",
			clusterState: []runtime.Object{
				makeTemplateConfigMap("template2", makeTemplate(t, func(ct *capiv1.CAPITemplate) {
					ct.ObjectMeta.Name = "cluster-template-2"
					ct.Spec.Description = "this is test template 2"
				}), "template1", makeTemplate(t)),
			},
			expected: []*capiv1_protos.Template{
				{
					Name:        "cluster-template-1",
					Description: "this is test template 1",
					Body:        "eyJhcGlWZXJzaW9uIjoiZm9vdmVyc2lvbiIsImtpbmQiOiJmb29raW5kIiwibWV0YWRhdGEiOnsibGFiZWxzIjp7Im5hbWUiOiIke0NMVVNURVJfTkFNRX0ifX19",
					Objects: []*capiv1_protos.TemplateObject{
						{
							ApiVersion: "fooversion",
							Kind:       "fookind",
							Parameters: []string{"CLUSTER_NAME"},
						},
					},
					Parameters: []*capiv1_protos.Parameter{
						{
							Name:        "CLUSTER_NAME",
							Description: "This is used for the cluster naming.",
						},
					},
				},
				{
					Name:        "cluster-template-2",
					Description: "this is test template 2",
					Body:        "eyJhcGlWZXJzaW9uIjoiZm9vdmVyc2lvbiIsImtpbmQiOiJmb29raW5kIiwibWV0YWRhdGEiOnsibGFiZWxzIjp7Im5hbWUiOiIke0NMVVNURVJfTkFNRX0ifX19",
					Objects: []*capiv1_protos.TemplateObject{
						{
							ApiVersion: "fooversion",
							Kind:       "fookind",
							Parameters: []string{"CLUSTER_NAME"},
						},
					},
					Parameters: []*capiv1_protos.Parameter{
						{
							Name:        "CLUSTER_NAME",
							Description: "This is used for the cluster naming.",
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := createServer(tt.clusterState, "capi-templates", "default", nil)

			listTemplatesRequest := new(capiv1_protos.ListTemplatesRequest)

			listTemplatesResponse, err := s.ListTemplates(context.Background(), listTemplatesRequest)
			if err != nil {
				if tt.err == nil {
					t.Fatalf("failed to read the templates:\n%s", err)
				}
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Fatalf("failed to read the templates:\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, listTemplatesResponse.Templates, protocmp.Transform()); diff != "" {
					t.Fatalf("templates didn't match expected:\n%s", diff)
				}
			}
		})
	}
}

func TestListTemplateParams(t *testing.T) {
	testCases := []struct {
		name             string
		clusterState     []runtime.Object
		expected         []*capiv1_protos.Parameter
		err              error
		expectedErrorStr string
	}{
		{
			name: "1 parameter err",
			err:  errors.New("error looking up template cluster-template-1: configmap capi-templates not found in default namespace"),
		},
		{
			name: "1 parameter",
			clusterState: []runtime.Object{
				makeTemplateConfigMap("template1", makeTemplate(t)),
			},
			expected: []*capiv1_protos.Parameter{
				{
					Name:        "CLUSTER_NAME",
					Description: "This is used for the cluster naming.",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := createServer(tt.clusterState, "capi-templates", "default", nil)

			listTemplateParamsRequest := new(capiv1_protos.ListTemplateParamsRequest)
			listTemplateParamsRequest.TemplateName = "cluster-template-1"

			listTemplateParamsResponse, err := s.ListTemplateParams(context.Background(), listTemplateParamsRequest)
			if err != nil {
				if tt.err == nil {
					t.Fatalf("failed to read the templates:\n%s", err)
				}
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Fatalf("got the wrong error:\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, listTemplateParamsResponse.Parameters, protocmp.Transform()); diff != "" {
					t.Fatalf("templates didn't match expected:\n%s", diff)
				}
			}
		})
	}
}

func TestRenderTemplate(t *testing.T) {
	testCases := []struct {
		name             string
		clusterState     []runtime.Object
		expected         string
		err              error
		expectedErrorStr string
	}{
		{
			name: "render template",
			clusterState: []runtime.Object{
				makeTemplateConfigMap("template1", makeTemplate(t)),
			},
			expected: "apiVersion: fooversion\nkind: fookind\nmetadata:\n  labels:\n    name: test-cluster\n",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := createServer(tt.clusterState, "capi-templates", "default", nil)

			renderTemplateRequest := &capiv1_protos.RenderTemplateRequest{
				TemplateName: "cluster-template-1",
				Values: &capiv1_protos.ParameterValues{
					Values: map[string]string{
						"CLUSTER_NAME": "test-cluster",
					},
				},
			}

			renderTemplateResponse, err := s.RenderTemplate(context.Background(), renderTemplateRequest)
			if err != nil {
				if tt.err == nil {
					t.Fatalf("failed to read the templates:\n%s", err)
				}
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Fatalf("got the wrong error:\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, renderTemplateResponse.RenderedTemplate, protocmp.Transform()); diff != "" {
					t.Fatalf("templates didn't match expected:\n%s", diff)
				}
			}
		})
	}
}

func TestCreatePullRequest(t *testing.T) {
	testCases := []struct {
		name         string
		clusterState []runtime.Object
		provider     git.Provider
		req          *capiv1_protos.CreatePullRequestRequest
		expected     string
		err          error
	}{
		{
			name: "validation errors",
			req:  &capiv1_protos.CreatePullRequestRequest{},
			err:  errors.New("6 errors occurred:\ntemplate name must be specified\nparameter values must be specified\nhead branch must be specified\ntitle must be specified\ndescription must be specified\ncommit message must be specified"),
		},
		{
			name: "pull request failed",
			clusterState: []runtime.Object{
				makeTemplateConfigMap("template1", makeTemplate(t)),
			},
			provider: NewFakeGitProvider("", errors.New("oops")),
			req: &capiv1_protos.CreatePullRequestRequest{
				TemplateName: "cluster-template-1",
				ParameterValues: map[string]string{
					"CLUSTER_NAME": "foo",
				},
				RepositoryUrl: "https://github.com/org/repo.git",
				HeadBranch:    "feature-01",
				BaseBranch:    "main",
				Title:         "New Cluster",
				Description:   "Creates a cluster through a CAPI template",
				CommitMessage: "Add cluster manifest",
			},
			err: errors.New("oops"),
		},
		{
			name: "create pull request",
			clusterState: []runtime.Object{
				makeTemplateConfigMap("template1", makeTemplate(t)),
			},
			provider: NewFakeGitProvider("https://github.com/org/repo/pull/1", nil),
			req: &capiv1_protos.CreatePullRequestRequest{
				TemplateName: "cluster-template-1",
				ParameterValues: map[string]string{
					"CLUSTER_NAME": "foo",
				},
				RepositoryUrl: "https://github.com/org/repo.git",
				HeadBranch:    "feature-01",
				BaseBranch:    "main",
				Title:         "New Cluster",
				Description:   "Creates a cluster through a CAPI template",
				CommitMessage: "Add cluster manifest",
			},
			expected: "https://github.com/org/repo/pull/1",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := createServer(tt.clusterState, "capi-templates", "default", tt.provider)
			ctx := context.Background()

			createPullRequestResponse, err := s.CreatePullRequest(ctx, tt.req)

			if err != nil {
				if tt.err == nil {
					t.Fatalf("failed to create a pull request:\n%s", err)
				}
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Fatalf("got the wrong error:\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, createPullRequestResponse.WebUrl, protocmp.Transform()); diff != "" {
					t.Fatalf("pull request url didn't match expected:\n%s", diff)
				}
			}
		})
	}
}

func createServer(clusterState []runtime.Object, configMapName, namespace string, provider git.Provider) capiv1_protos.ClustersServiceServer {
	scheme := runtime.NewScheme()
	schemeBuilder := runtime.SchemeBuilder{
		corev1.AddToScheme,
		capiv1.AddToScheme,
	}
	schemeBuilder.AddToScheme(scheme)

	cl := fake.NewClientBuilder().
		WithScheme(scheme).
		WithRuntimeObjects(clusterState...).
		Build()

	s := NewClusterServer(&templates.ConfigMapLibrary{
		Client:        cl,
		ConfigMapName: configMapName,
		Namespace:     namespace,
	}, provider)

	return s
}

func TestToTemplate(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected *capiv1_protos.Template
		err      error
	}{
		{
			name:     "empty",
			value:    "",
			expected: &capiv1_protos.Template{},
		},
		{
			name: "Basics",
			value: `
apiVersion: capi.weave.works/v1alpha1
kind: CAPITemplate
metadata:
  name: foo
`,
			expected: &capiv1_protos.Template{
				Name: "foo",
			},
		},
		{
			name: "Params and Objects",
			value: makeTemplate(t, func(ct *capiv1.CAPITemplate) {
				ct.ObjectMeta.Name = "cluster-template-1"
				ct.Spec.Description = "this is test template 1"
				ct.Spec.ResourceTemplates = []capiv1.CAPIResourceTemplate{
					{
						RawExtension: rawExtension(`{
							"apiVersion": "fooversion",
							"kind": "fookind",
							"metadata": {
								"labels": {
								"name": "${CLUSTER_NAME}",
								"region": "${REGION}"
								}
							}
						}`),
					},
				}
			}),
			expected: &capiv1_protos.Template{
				Name:        "cluster-template-1",
				Description: "this is test template 1",
				Body:        "eyJhcGlWZXJzaW9uIjoiZm9vdmVyc2lvbiIsImtpbmQiOiJmb29raW5kIiwibWV0YWRhdGEiOnsibGFiZWxzIjp7Im5hbWUiOiIke0NMVVNURVJfTkFNRX0iLCJyZWdpb24iOiIke1JFR0lPTn0ifX19",
				Objects: []*capiv1_protos.TemplateObject{
					{
						ApiVersion: "fooversion",
						Kind:       "fookind",
						Parameters: []string{"CLUSTER_NAME", "REGION"},
					},
				},
				Parameters: []*capiv1_protos.Parameter{
					{
						Name:        "CLUSTER_NAME",
						Description: "This is used for the cluster naming.",
					},
					{
						Name: "REGION",
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToTemplateResponse(mustParseBytes(t, tt.value))
			if err != nil {
				if tt.err == nil {
					t.Fatalf("failed to read the templates:\n%s", err)
				}
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Fatalf("got the wrong error:\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, result, protocmp.Transform()); diff != "" {
					t.Fatalf("templates didn't match expected:\n%s", diff)
				}
			}
		})
	}
}

func mustParseBytes(t *testing.T, data string) *capiv1.CAPITemplate {
	t.Helper()
	parsed, err := capi.ParseBytes([]byte(data), "no-key-provided")
	if err != nil {
		t.Fatal(err)
	}
	return parsed
}

func makeTemplateConfigMap(s ...string) *corev1.ConfigMap {
	data := make(map[string]string)
	for i := 0; i < len(s); i += 2 {
		data[s[i]] = s[i+1]
	}
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "capi-templates",
			Namespace: "default",
		},
		Data: data,
	}
}

func makeTemplate(t *testing.T, opts ...func(*capiv1.CAPITemplate)) string {
	t.Helper()
	basicRaw := `
{
  "apiVersion": "fooversion",
  "kind": "fookind",
  "metadata": {
    "labels": {
      "name": "${CLUSTER_NAME}"
    }
  }
}`
	ct := &capiv1.CAPITemplate{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CAPITemplate",
			APIVersion: "capi.weave.works/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster-template-1",
		},
		Spec: capiv1.CAPITemplateSpec{
			Description: "this is test template 1",
			Params: []capiv1.TemplateParam{
				{
					Name:        "CLUSTER_NAME",
					Description: "This is used for the cluster naming.",
				},
			},
			ResourceTemplates: []capiv1.CAPIResourceTemplate{
				{
					RawExtension: rawExtension(basicRaw),
				},
			},
		},
	}
	for _, o := range opts {
		o(ct)
	}
	b, err := yaml.Marshal(ct)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func rawExtension(s string) runtime.RawExtension {
	return runtime.RawExtension{
		Raw: []byte(s),
	}
}

func NewFakeGitProvider(url string, err error) git.Provider {
	return &FakeGitProvider{
		url: url,
		err: err,
	}
}

type FakeGitProvider struct {
	url string
	err error
}

func (p *FakeGitProvider) WriteFilesToBranchAndCreatePullRequest(ctx context.Context, req git.WriteFilesToBranchAndCreatePullRequestRequest) (*git.WriteFilesToBranchAndCreatePullRequestResponse, error) {
	if p.err != nil {
		return nil, p.err
	}
	return &git.WriteFilesToBranchAndCreatePullRequestResponse{WebURL: p.url}, nil
}