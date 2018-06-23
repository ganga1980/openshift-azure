#!/bin/bash -ex

if [[ $# -eq 0 ]]; then
    echo usage: $0 resourcegroup
    exit 1
fi

RESOURCEGROUP=$1

rm -rf _out
mkdir _out
go generate ./...

go run cmd/create/create.go <<EOF
TenantID: $AZURE_TENANT_ID
SubscriptionID: $AZURE_SUBSCRIPTION_ID
ClientID: $AZURE_CLIENT_ID
ClientSecret: $AZURE_CLIENT_SECRET
Location: eastus
ResourceGroup: $RESOURCEGROUP
VMSize: Standard_D4s_v3
ComputeCount: 1
InfraCount: 1
ImageResourceGroup: images
ImageResourceName: centos7-3.10-201806231427
EOF

helm template pkg/helm/chart -f _out/values.yaml --output-dir _out

# poor man's helm create (without tiller running)
oc delete -Rf _out/osa/templates || true
oc create -Rf _out/osa/templates

az group create -n $RESOURCEGROUP -l eastus
az group deployment create -g $RESOURCEGROUP --template-file _out/azuredeploy.json
