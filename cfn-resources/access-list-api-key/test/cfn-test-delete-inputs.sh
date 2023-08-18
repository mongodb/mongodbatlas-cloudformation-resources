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

OrgId=$(jq -r '.OrgId' ./inputs/inputs_1_create.json)
ApiUserId=$(jq -r '.ApiUserId' ./inputs/inputs_1_create.json)



#delete access list entry
if atlas organizations apiKeys delete "$ApiUserId" --orgId "$OrgId" --force; then
	echo "$OrgId access list deletion OK"
else
	(echo "Failed cleaning organizations:$OrgId" && exit 1)
fi
