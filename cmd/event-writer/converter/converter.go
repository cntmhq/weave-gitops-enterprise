package converter

import (
	"bytes"
	"encoding/json"
	"fmt"

	// sourcev1 "github.com/fluxcd/source-controller/api/v1beta1"
	"github.com/weaveworks/wks/cmd/event-writer/database/models"
	"github.com/weaveworks/wks/cmd/event-writer/messages"

	// workspacesv1 "github.com/weaveworks/wks/cmd/workspaces-controller/api/v1beta1"
	"gorm.io/datatypes"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
)

// ConvertEvent returns a models.Event from a k8s core api v1.Event object
func ConvertEvent(event v1.Event) (models.Event, error) {
	eventJSONbytes, err := SerializeEventToJSON(&event)
	if err != nil {
		return models.Event{}, err
	}
	eventJSON, err := datatypes.JSON.MarshalJSON(eventJSONbytes)
	if err != nil {
		return models.Event{}, err
	}

	flattenedLabels := SerializeStringMap(event.ObjectMeta.Labels)
	flattenedAnnotations := SerializeStringMap(event.ObjectMeta.Annotations)

	result := models.Event{
		Name:        event.ObjectMeta.Name,
		Namespace:   event.ObjectMeta.Namespace,
		Labels:      flattenedLabels,
		Annotations: flattenedAnnotations,
		ClusterName: event.ClusterName,
		Reason:      event.Reason,
		Message:     event.Message,
		Type:        event.Type,
		RawEvent:    eventJSON,
	}
	return result, nil
}

// ConvertCluster returns a models.Cluster from a NATS message with cluster info
func ConvertCluster(cluster messages.Cluster) (models.Cluster, error) {
	result := models.Cluster{
		Name:              cluster.Name,
		Namespace:         cluster.Namespace,
		Labels:            cluster.Labels,
		Annotations:       cluster.Annotations,
		ControlPlaneNodes: cluster.ControlPlaneNodes,
		WorkerNodes:       cluster.WorkerNodes,
		CNIInfo:           cluster.CNIInfo,
		CSIInfo:           cluster.CSIInfo,
		CRIInfo:           cluster.CRIInfo,
		Version:           cluster.Version,
		IngressInfo:       cluster.IngressInfo,
	}
	return result, nil
}

// SerializeStringMap flattens a string to string map to a string given a printf format
func SerializeStringMap(m map[string]string) string {
	format := "%s:%s,"
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, format, key, value)
	}
	return b.String()
}

// SerializeEventToJSON serializes a v1.Event object to a byte array
func SerializeEventToJSON(e *v1.Event) ([]byte, error) {
	output := bytes.NewBufferString("")
	encoder := json.NewEncoder(output)
	encoder.Encode(e)
	return output.Bytes(), nil
}

// DeserializeJSONToEvent constructs a v1.Event from a byte array and returns a pointer to it
func DeserializeJSONToEvent(b []byte) (*v1.Event, error) {
	decoder := serializer.NewCodecFactory(scheme.Scheme).UniversalDecoder()
	e := &v1.Event{}
	err := runtime.DecodeInto(decoder, b, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
