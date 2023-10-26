#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

userName=$(atlas organizations users list --orgId "$MONGODB_ATLAS_ORG_ID" --output json | jq -r '.results' | jq -r '.[0].emailAddress')
jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg userName "$userName" \
	'.Usernames?|=[$userName]|.OrgId?|=$org' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	'.OrgId?|=$org' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"
