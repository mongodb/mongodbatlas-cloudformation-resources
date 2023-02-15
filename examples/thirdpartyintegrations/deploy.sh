#!/usr/bin/env bash

echo "Thirdparty Webhook creation"
aws cloudformation deploy \
	--stack-name atlas-thirdpartyintegration-webhook \
	--template-file webhook.json \
	--no-fail-on-empty-changeset \
	--parameter-overrides PrivateKey="${ATLAS_PRIVATE_KEY}" PublicKey="${ATLAS_PUBLIC_KEY}" ProjectId="${ATLAS_GROUP_ID}" \
	"$@"
