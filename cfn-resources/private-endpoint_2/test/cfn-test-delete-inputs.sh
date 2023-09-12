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

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
interfaceEndpointId=$(jq -r '.InterfaceEndpointId' ./inputs/inputs_1_create.json)

# Delete the VPC endpoint
aws ec2 delete-vpc-endpoints --vpc-endpoint-ids "$interfaceEndpointId"

PRIVATE_ENDPOINTS=$(atlas privateEndpoints aws list --projectId $projectId --output json)

# Check if there are any private endpoints
if [[ "$PRIVATE_ENDPOINTS" == "[]" ]]; then
  echo "No private endpoints found."
else
  # Extract the IDs of the private endpoints and delete them one by one
  for ID in $(echo "$PRIVATE_ENDPOINTS" | jq -r '.[].id'); do
    echo "Deleting private endpoint with ID: $ID"
    atlas privateEndpoints aws delete $ID --projectId $projectId --force
  done
fi

while true; do
  # List private endpoints and store the result in a variable
  PRIVATE_ENDPOINTS=$(atlas privateEndpoints aws list --projectId $projectId --output json)
  if [[ "$PRIVATE_ENDPOINTS" == "[]" ]]; then
    echo "No private endpoints found."
    break  # Exit the loop when no private endpoints are found
  fi

  # Add a sleep to avoid continuous polling (adjust the duration as needed)
  sleep 10  # Sleep for 10 seconds before checking again (adjust as needed)
done


#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
