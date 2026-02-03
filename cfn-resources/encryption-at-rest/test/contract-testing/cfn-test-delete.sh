#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
projectName=$(atlas projects list --output json | jq --arg projectId "${projectId}" -r '.results[] | select(.id==$projectId) | .name')

echo "Deleting resources for project: $projectName (ID: $projectId)"

./test/cfn-test-delete-inputs.sh "$projectName"

if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project: $projectId" && exit 1)
fi
