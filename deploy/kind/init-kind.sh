#!/bin/bash

which kind
if [[ $? != 0 ]]; then
    echo 'kind not available in $PATH, installing latest kind'
    # Install latest kind
    curl -s https://api.github.com/repos/kubernetes-sigs/kind/releases/latest \
    | grep "browser_download_url.*kind-linux-amd64" \
    | cut -d : -f 2,3 \
    | tr -d \" \
    | wget -qi -
    mv kind-linux-amd64 deploy/kind/kind && chmod +x deploy/kind/kind
    export PATH=$PATH:$PWD/deploy/kind
fi

cluster_created=$($PWD/deploy/kind/kind get clusters 2>&1)
if [[ "$cluster_created" == "No kind clusters found." ]]; then
    $PWD/deploy/kind/kind create cluster --config $PWD/deploy/kind/config.yaml
else
    echo "Cluster '$cluster_created' already present"
fi
