#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}

clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)
projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

#delete Cluster
if atlas clusters delete "$clusterName" --projectId "${projectId}" --force; then
	echo "$clusterName cluster deletion OK"
else
	(echo "Failed cleaning cluster:$clusterName" && exit 1)
fi

status="DELETING"
echo "Waiting for cluster to get deleted"
while [[ "${status}" == "DELETING" ]]; do
	sleep 30
	if atlas clusters describe "${clusterName}" --projectId "${projectId}"; then
		status=$(atlas clusters describe "${clusterName}" --projectId "${projectId}" --output=json | jq -r '.stateName')
	else
		status="DELETED"
	fi
	echo "status: ${status}"
done

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
