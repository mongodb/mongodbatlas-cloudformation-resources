#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

projectId=$(jq -r '.RoleAssignments[] | select(.ProjectId != null) | .ProjectId' ./inputs/inputs_1_create.json | head -n 1)

if [ -n "$projectId" ] && [ "$projectId" != "null" ]; then
	if atlas projects delete "$projectId" --force; then
		echo "$projectId project deletion OK"
	else
		(echo "Failed cleaning project: $projectId" && exit 1)
	fi
else
	echo "No project to delete (no ProjectId found in test inputs)"
fi
