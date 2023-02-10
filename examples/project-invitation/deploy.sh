#!/usr/bin/env bash

username="testing@mongodb.com"
role="GROUP_CLUSTER_MANAGER"

# AWS CFN Deploy
aws cloudformation deploy \
	--stack-name project-invitation-test \
	--template-file project-invitation.json \
	--no-fail-on-empty-changeset \
	--parameter-overrides Role="$role" Username="$username" PrivateKey="${ATLAS_PRIVATE_KEY}" PublicKey="${ATLAS_PUBLIC_KEY}" ProjectId="${ATLAS_GROUP_ID}" \
	"$@"
