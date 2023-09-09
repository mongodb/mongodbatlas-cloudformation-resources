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

tenantName=$(jq -r '.TenantName' ./inputs/inputs_1_create.json)

if atlas dataFederation delete "${tenantName}" --projectId "${projectId}" --force ;then
  echo "deleted tenant with name ${tenantName}"
else
  echo "failed to delete the tenant with name ${tenantName}"
fi

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
