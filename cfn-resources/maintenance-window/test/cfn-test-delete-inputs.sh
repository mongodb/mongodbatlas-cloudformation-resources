#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
	echo "Deletes the project created for testing"
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

if [ -z "$projectId" ]; then
	echo "Error: ProjectId not found in inputs/inputs_1_create.json"
	exit 1
fi

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
