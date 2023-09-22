#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
  echo "usage:$0 "
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
interfaceEndpointId=$(jq -r '.EndpointId' ./inputs/inputs_1_create.json)

echo "STEP 1 DELETING UNUSED AWS PRIVATE ENDPOINT"

aws ec2 delete-vpc-endpoints --vpc-endpoint-ids "$interfaceEndpointId"

# Delete the VPC endpoint
echo "STEP 1 DELETING ATLAS PRIVATE ENDPOINTS"

PRIVATE_ENDPOINTS=$(atlas privateEndpoints aws list --projectId "$projectId" --output json)
# Check if there are any private endpoints
if [[ "$PRIVATE_ENDPOINTS" == "[]" ]]; then
  echo "No private endpoints found."
else
  for ATLAS_PRIVATE_ENDPOINT_SERVICE in $(echo "$PRIVATE_ENDPOINTS" | jq -r '.[].id'); do
    # Delete the VPC endpoint
    echo "STEP 1.a DELETING INTERFACES FOR SERVICE $ATLAS_PRIVATE_ENDPOINT_SERVICE"
    ENDPOINT_OUTPUT=$(atlas privateEndpoints aws describe "$ATLAS_PRIVATE_ENDPOINT_SERVICE" --projectId "$projectId" --output json)
    if ! echo "$ENDPOINT_OUTPUT" | jq -e '.interfaceEndpoints' > /dev/null; then
      echo "No interfaceEndpoints found for $ATLAS_PRIVATE_ENDPOINT_SERVICE"
      echo "STEP 1.d deleting privateEndpoint Service "
      atlas privateEndpoints aws delete "$ATLAS_PRIVATE_ENDPOINT_SERVICE" --projectId "$projectId" --force
    else
      interfaceEndpoints=$(echo "$ENDPOINT_OUTPUT" | jq -r '.interfaceEndpoints[]')
      for interfaceId in $(echo "$ENDPOINT_OUTPUT" | jq -r '.interfaceEndpoints[]'); do
        echo "STEP 1.b DELETING INTERFACE $interfaceId FOR SERVICE $ATLAS_PRIVATE_ENDPOINT_SERVICE"
        atlas privateEndpoints aws interface delete "$interfaceId" --endpointServiceId "$ATLAS_PRIVATE_ENDPOINT_SERVICE" --projectId "$projectId" --force

        sleep 20 #waiting until the connection gets rejected
        echo "STEP 1.c DELETING aws private endpoint $interfaceId FOR SERVICE $ATLAS_PRIVATE_ENDPOINT_SERVICE"
        aws ec2 delete-vpc-endpoints --vpc-endpoint-ids "$interfaceId"
      done
    fi
  done
fi

sleep 10

echo "STEP 2 waiting until all private endpoints get deleted"
while true; do
  PRIVATE_ENDPOINTS=$(atlas privateEndpoints aws list --projectId "$projectId" --output json)

  # Check if PRIVATE_ENDPOINTS is an empty array
  if [[ "$PRIVATE_ENDPOINTS" == "[]" ]]; then
    echo "ALL PRIVATE ENDPOINTS HAVE BEEN DELETED"
    break
  fi

  # Optional: Add a delay before checking again (e.g., sleep for 10 seconds)
  sleep 10
done

echo "STEP 3 deleting project"
# delete project
if atlas projects delete "$projectId" --force; then
  echo "$projectId project deletion OK"
else
  (echo "Failed cleaning project:$projectId" && exit 1)
fi