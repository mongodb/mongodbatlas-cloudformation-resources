#!/usr/bin/env bash

echo "PrivateEndPointRegionalMode creation"
aws cloudformation deploy \
	--stack-name atlas-private-endpoint-regional-mode \
	--template-file privateEndpointRegionalMode.json \
	--no-fail-on-empty-changeset \
	--parameter-overrides PrivateKey="${MONGODB_ATLAS_PRIVATE_KEY}" PublicKey="${MONGODB_ATLAS_PUBLIC_KEY}" ProjectId="${ATLAS_GROUP_ID}" \
	"$@"
