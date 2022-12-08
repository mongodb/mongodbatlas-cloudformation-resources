#!/usr/bin/env bash

echo "CustomDnsConfigurationClusterAws creation"
aws cloudformation deploy \
   --stack-name atlas-custom-dns-config-cluster-aws \
   --template-file customDnsConfigurationClusterAws.json \
    --no-fail-on-empty-changeset \
    --parameter-overrides PrivateKey="${ATLAS_PRIVATE_KEY}" PublicKey="${ATLAS_PUBLIC_KEY}" ProjectId="${ATLAS_GROUP_ID}" \
    "$@"