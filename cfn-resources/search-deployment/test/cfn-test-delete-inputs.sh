#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}

clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)
projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

# TEMPORARY: Skip deletion of test cluster/project (only for today's testing - 2025-12-29)
# TODO: Remove this after testing is complete
# Note: TEST_PROJECT and TEST_CLUSTER should be set via environment variables if needed
# This section can be removed after testing is complete
if [[ -n "${TEST_CLUSTER:-}" && -n "${TEST_PROJECT:-}" ]]; then
	if [[ "$clusterName" == "${TEST_CLUSTER}" && "$projectId" == "${TEST_PROJECT}" ]]; then
		echo "SKIPPING deletion of test cluster '$clusterName' and project '$projectId' (preserved for testing)"
		exit 0
	fi
fi

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

#delete project (skip if it's the test project)
if [[ -n "${TEST_PROJECT:-}" && "$projectId" == "$TEST_PROJECT" ]]; then
	echo "SKIPPING deletion of test project '$projectId' (preserved for testing)"
else
	if atlas projects delete "$projectId" --force; then
		echo "$projectId project deletion OK"
	else
		(echo "Failed cleaning project:$projectId" && exit 1)
	fi
fi
