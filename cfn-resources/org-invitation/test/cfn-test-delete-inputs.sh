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

teamId=$(jq -r '.TeamIds[0]' ./inputs/inputs_1_create.json)

#delete team
if mongocli iam team delete "$teamId" --force; then
	echo "$teamId team deletion OK"
else
	(echo "Failed cleaning team:$teamId" && exit 1)
fi
