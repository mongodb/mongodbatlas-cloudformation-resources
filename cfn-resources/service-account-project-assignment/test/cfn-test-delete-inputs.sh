#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
orgId=$(jq -r '.OrgId' ./inputs/inputs_1_create.json)
clientId=$(jq -r '.ClientId' ./inputs/inputs_1_create.json)

# Delete service account
if atlas api serviceAccounts deleteOrgServiceAccount --orgId "$orgId" --clientId "$clientId" 2>/dev/null; then
	echo "$clientId service account deletion OK"
else
	(echo "Failed cleaning service account: $clientId" && exit 1)
fi

# Delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project: $projectId" && exit 1)
fi
