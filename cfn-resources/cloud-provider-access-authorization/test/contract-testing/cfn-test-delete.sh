#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

projectId=""
if [ -f inputs/project-id.txt ]; then
	projectId=$(cat inputs/project-id.txt)
elif [ -f inputs/inputs_1_create.json ]; then
	projectId=$(jq -r '.ProjectId // empty' inputs/inputs_1_create.json)
fi

cleanupStatus=0
./test/cfn-test-delete-inputs.sh || cleanupStatus=$?

if [ -n "$projectId" ]; then
	if atlas projects delete "$projectId" --force; then
		echo "$projectId project deletion OK"
		# Prevent retries from using a stale project ID in the generated inputs.
		: >inputs/project-id.txt
	else
		echo "Failed cleaning project: $projectId" >&2
		cleanupStatus=1
	fi
fi
exit "$cleanupStatus"
