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

#delete endpoint
endpointId=$(jq -r '.EndpointId' ./inputs/inputs_1_create.json)
if atlas privateEndpoints aws delete "${endpointId}" --projectId "$projectId"; then
  echo "endpoint $endpointId  deleted"
fi

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
