#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.

set -euo pipefail

function usage {
	echo "usage:$0 "
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
workspaceName=$(jq -r '.WorkspaceName' ./inputs/inputs_1_create.json)

# delete stream workspace (using instances delete for backward compatibility)
if atlas streams instances delete "${workspaceName}" --projectId "${projectId}" --force; then
	echo "deleting stream workspace with name ${workspaceName}"
else
	echo "failed to delete the stream workspace with name ${workspaceName}"
fi
