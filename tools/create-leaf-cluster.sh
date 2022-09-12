#!/bin/bash
set -e

if [ $# -eq 0 ]
then
    echo "Please, provide the cluster name (e.g. leaf-cluster-01)"
    exit
fi

if ! command -v vcluster &> /dev/null
then
    echo "vcluster could not be found"
    exit
fi

echo "Creating cluster..."
vcluster create --connect=false $1

echo "Waiting cluster config..."
until kubectl get secret -n vcluster-$1 vc-$1; do sleep 5; done

echo "Creating GitopsCluster secret..."
kubectl get secret -n vcluster-$1 vc-$1 --template={{.data.config}} \
 | base64 -D \
 | sed "s/localhost:8443/$1.vcluster-$1/g" \
 | kubectl create secret -n vcluster-$1 generic $1-config --from-file=value=/dev/stdin

echo "Creating GitopsCluster resource..."
cat <<EOF | kubectl apply -f -
apiVersion: gitops.weave.works/v1alpha1
kind: GitopsCluster
metadata:
  name: $1
  namespace: vcluster-$1
  # Signals that this cluster should be bootstrapped.
  labels:
    weave.works/capi: bootstrap
spec:
  secretRef:
    name: $1-config
EOF
