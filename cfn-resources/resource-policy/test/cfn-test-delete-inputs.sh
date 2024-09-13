#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}

orgId=$(jq -r '.OrgId' ./inputs/inputs_1_create.json)

#delete org
if atlas organization delete "$orgId" --force; then
	echo "$orgId organization deletion OK"
else
	(echo "Failed cleaning organization:$orgId" && exit 1)
fi
