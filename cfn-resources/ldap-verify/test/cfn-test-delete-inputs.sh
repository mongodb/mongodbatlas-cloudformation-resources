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
}


projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

cluster_names=$(atlas clusters list --projectId "$projectId" --output=json | jq -r '.results[].name')

# Get the first cluster name from the list
first_cluster_name=$(echo "$cluster_names" | head -n 1)

if [ -n "$first_cluster_name" ]; then
    echo "Deleting cluster: $first_cluster_name"

    # Delete Cluster
    if atlas clusters delete "$first_cluster_name" --projectId "$projectId" --force; then
        echo "$first_cluster_name cluster deletion OK"
    else
        (echo "Failed cleaning cluster:$first_cluster_name" && exit 1)
    fi

    status="DELETING"
    echo "Waiting for cluster to get deleted"
    while [[ "${status}" == "DELETING" ]]; do
        sleep 30
        if atlas clusters describe "$first_cluster_name" --projectId "$projectId"; then
            status=$(atlas clusters describe "$first_cluster_name" --projectId "$projectId" --output=json | jq -r '.stateName')
        else
            status="DELETED"
        fi
        echo "status: ${status}"
    done
else
    echo "No clusters found in the project."
fi

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
