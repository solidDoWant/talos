#!/bin/bash

set -eou pipefail

source ./hack/test/e2e.sh

PROVISIONER=qemu
CLUSTER_NAME=e2e-${PROVISIONER}

case "${CI:-false}" in
  true)
    REGISTRY="127.0.0.1:5000"
    REGISTRY_ADDR=`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' registry`
    QEMU_FLAGS="--registry-mirror ${REGISTRY}=http://${REGISTRY_ADDR}:5000"
    INSTALLER_TAG="${TAG}"
    docker tag ${INSTALLER_IMAGE} 127.0.0.1:5000/autonomy/installer:"${TAG}"
    docker push 127.0.0.1:5000/autonomy/installer:"${TAG}"
    ;;
  *)
    QEMU_FLAGS=
    INSTALLER_TAG="latest"
    ;;
esac

case "${CUSTOM_CNI_URL:-false}" in
  false)
    CUSTOM_CNI_FLAG=
    ;;
  *)
    CUSTOM_CNI_FLAG="--custom-cni-url=${CUSTOM_CNI_URL}"
    ;;
esac

function create_cluster {
  "${TALOSCTL}" cluster create \
    --provisioner "${PROVISIONER}" \
    --name "${CLUSTER_NAME}" \
    --masters=3 \
    --mtu 1500 \
    --memory 2048 \
    --cpus 2.0 \
    --cidr 172.20.1.0/24 \
    --install-image ${REGISTRY:-docker.io}/autonomy/installer:${INSTALLER_TAG} \
    --with-init-node=false \
    --crashdump \
    ${QEMU_FLAGS} \
    ${CUSTOM_CNI_FLAG}

  "${TALOSCTL}" config node 172.20.1.2
}

function destroy_cluster() {
  "${TALOSCTL}" cluster destroy --name "${CLUSTER_NAME}"
}

create_cluster
get_kubeconfig
run_talos_integration_test
run_kubernetes_integration_test
