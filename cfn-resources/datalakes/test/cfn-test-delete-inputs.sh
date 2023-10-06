#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail
set -x

function usage {
	echo "usage:$0 "
}

project_id=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

# Get a list of clusters for the current project
clusters_response=$(atlas clusters list --projectId "$project_id" 2>&1)

# Check if the response contains the "GROUP_NOT_FOUND" error message
if [[ "$clusters_response" == *"GROUP_NOT_FOUND"* ]]; then
    echo "Project with ID $project_id not found. Continuing with the next project."
fi

total_count=$(echo "$clusters_response" | jq -r '.totalCount')

# Check if there are clusters for the project
if [ "$total_count" -eq 0 ]; then
    echo "No clusters found for project with ID: $project_id"
else
    # Extract cluster names from the response
    mapfile -t cluster_names < <(echo "$clusters_response" | jq -r '.results[].name')

    # Delete each cluster within the project
    for cluster_name in "${cluster_names[@]}"; do
        echo "Deleting cluster with name: $cluster_name"
        atlas clusters delete "$cluster_name" --projectId "$project_id" --force

        echo "Waiting for cluster to get deleted"
        status=$(atlas clusters describe "${cluster_name}" --projectId "${project_id}" --output=json | jq -r '.stateName')
        echo "status: ${status}"

        while [[ "${status}" == "DELETING" ]]; do
            sleep 30
            if atlas clusters describe "${cluster_name}" --projectId "${project_id}"; then
                status=$(atlas clusters describe "${cluster_name}" --projectId "${project_id}" --output=json | jq -r '.stateName')
            else
                status="DELETED"
            fi
            echo "status: ${status}"
        done
    done
fi

#delete project
if atlas projects delete "$project_id" --force; then
	echo "$project_id project deletion OK"
else
	(echo "Failed cleaning project:$project_id" && exit 1)
fi
