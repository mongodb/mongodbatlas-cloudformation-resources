#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

# Delete S3 bucket first (using base script)
./test/cfn-test-delete-inputs.sh

# Delete project if inputs file exists
if [ -f "./inputs/inputs_1_create.json" ]; then
	projectId=$(jq -r '.ProjectId // empty' ./inputs/inputs_1_create.json)

	if [ -n "$projectId" ] && [ "$projectId" != "null" ]; then
		# delete project
		if atlas projects delete "$projectId" --force; then
			echo "$projectId project deletion OK"
		else
			echo "⚠️  Warning: Failed cleaning project: $projectId"
		fi
	fi
fi
