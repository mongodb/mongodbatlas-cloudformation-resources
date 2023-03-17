#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#
# Run this script with the Makefile
# make create-test-file
#
set -eu

api_key_id=$(jq -r '.ProjectApiKeys[0] | .Key' ./inputs/inputs_1_create.json)
project_name=$(jq -r '.Name' ./inputs/inputs_1_create.json)

if [ -z "${MONGODB_ATLAS_PUBLIC_API_KEY+x}" ] || [ -z "${MONGODB_ATLAS_PRIVATE_API_KEY+x}" || [ -z "${MONGODB_ATLAS_ORG_ID+x}" ]; then
	echo "Error: MONGODB_ATLAS_PUBLIC_API_KEY, MONGODB_ATLAS_PRIVATE_API_KEY and MONGODB_ATLAS_ORG_ID environment variables must be set"
	exit 1
fi

#delete apikey
atlas organizations apiKeys delete "${api_key_id}" --force

#delete project
projectId=$(atlas projects list --output json | jq --arg NAME "${project_name}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	echo -e "No project found with \"${project_name}"
else
	echo -e "project found with ${project_name} and id ${projectId}, deleting"
	atlas projects delete "${projectId}" --force
fi
