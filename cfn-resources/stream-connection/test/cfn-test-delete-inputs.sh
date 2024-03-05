#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -euo pipefail

function usage {
	echo "usage:$0 "
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)

if atlas cluster delete "${clusterName}" --projectId "${projectId}" --force; then
	echo "deleted cluster with name ${clusterName}"
else
	echo "failed to delete the cluster with name ${clusterName}"
fi

atlas cluster watch "${clusterName}" --projectId "${projectId}"

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
