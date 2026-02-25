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
	echo "Deletes the cluster and project created for testing"
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)

echo "Cleaning up contract test resources..."
echo "Project ID: $projectId"
echo "Cluster Name: $clusterName"

# Delete cluster
if atlas clusters delete "$clusterName" --projectId "$projectId" --force; then
	echo "$clusterName cluster deletion OK"
else
	(echo "Failed cleaning cluster: $clusterName" && exit 1)
fi

# Wait for cluster deletion with proper status checking
status="DELETING"
echo "Waiting for cluster to get deleted"
while [[ "${status}" == "DELETING" ]]; do
	sleep 30
	if atlas clusters describe "${clusterName}" --projectId "${projectId}" --output json 2>/dev/null; then
		status=$(atlas clusters describe "${clusterName}" --projectId "${projectId}" --output json | jq -r '.stateName')
	else
		status="DELETED"
	fi
	echo "status: ${status}"
done

# Delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project: $projectId" && exit 1)
fi
