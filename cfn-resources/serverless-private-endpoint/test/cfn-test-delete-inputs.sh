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

instanceName=$(jq -r '.InstanceName' ./inputs/inputs_1_create.json)

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

# Delete the Atlas Serverless instance
delete_output=$(atlas serverless delete "$instanceName" --projectId "$projectId" --force 2>&1)

# Check if the instance deletion was initiated successfully
if [[ $delete_output =~ Serverless\ instance\ \'$instanceName\'\ deleted ]]; then
    echo "Deletion initiated for instance $instanceName."

    # Loop until the instance is deleted
    while true; do
      if atlas serverless describe "${instanceName}" --projectId "${projectId}"; then
        echo "Waiting for instance $instanceName to be deleted..."
        sleep 10  # Adjust the sleep interval as needed
      else
        echo "Instance $instanceName has been deleted."
        break
      fi
    done
else
    echo "Failed to initiate deletion for instance $instanceName. Error: $delete_output"
fi

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
