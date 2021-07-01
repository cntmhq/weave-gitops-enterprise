package capi

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	parsed := mustParseFile(t, "testdata/template3.yaml")

	b, err := Render(parsed.Spec, map[string]string{"CLUSTER_NAME": "testing"})
	if err != nil {
		t.Fatal(err)
	}

	want := `---
apiVersion: cluster.x-k8s.io/v1alpha3
kind: Cluster
metadata:
  name: testing
---
apiVersion: infrastructure.cluster.x-k8s.io/v1alpha3
kind: AWSMachineTemplate
metadata:
  name: testing-md-0
`
	if diff := cmp.Diff(want, writeMultiDoc(t, b)); diff != "" {
		t.Fatalf("rendering failure:\n%s", diff)
	}
}

func TestRenderWithCRD(t *testing.T) {
	parsed := mustParseFile(t, "testdata/template0.yaml")

	b, err := Render(parsed.Spec, map[string]string{"CLUSTER_NAME": "testing"})
	if err != nil {
		t.Fatal(err)
	}

	want := `---
apiVersion: cluster.x-k8s.io/v1alpha3
kind: Cluster
metadata:
  name: testing
spec:
  clusterNetwork:
    pods:
      cidrBlocks:
      - 192.168.0.0/16
  controlPlaneRef:
    apiVersion: controlplane.cluster.x-k8s.io/v1alpha3
    kind: KubeadmControlPlane
    name: testing-control-plane
  infrastructureRef:
    apiVersion: infrastructure.cluster.x-k8s.io/v1alpha3
    kind: AWSCluster
    name: testing
`
	if diff := cmp.Diff(want, writeMultiDoc(t, b)); diff != "" {
		t.Fatalf("rendering failure:\n%s", diff)
	}
}

func TestRender_unknown_parameter(t *testing.T) {
	parsed := mustParseFile(t, "testdata/template3.yaml")

	_, err := Render(parsed.Spec, map[string]string{})
	assert.EqualError(t, err, "processing template: value for variables [CLUSTER_NAME] is not set. Please set the value using os environment variables or the clusterctl config file")
}

func writeMultiDoc(t *testing.T, objs [][]byte) string {
	t.Helper()
	var out bytes.Buffer
	for _, v := range objs {
		if _, err := out.Write([]byte("---\n")); err != nil {
			t.Fatal(err)
		}
		if _, err := out.Write(v); err != nil {
			t.Fatal(err)
		}
	}
	return out.String()
}