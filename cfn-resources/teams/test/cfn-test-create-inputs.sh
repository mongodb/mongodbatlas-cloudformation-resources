#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail
set -x

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
name="${1}"

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

userName=$(atlas organizations users list --orgId "$ATLAS_ORG_ID" --output json | jq -r '.results' | jq -r '.[0].emailAddress')
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
	--arg pvtkey "$ATLAS_PRIVATE_KEY" \
	--arg org "$ATLAS_ORG_ID" \
	--arg userName "$userName" \
	'.Usernames?|=[$userName]|.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
	--arg pvtkey "$ATLAS_PRIVATE_KEY" \
	--arg org "$ATLAS_ORG_ID" \
	'.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

name="${name}- more B@d chars !@(!(@====*** ;;::"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
	--arg pvtkey "$ATLAS_PRIVATE_KEY" \
	--arg org "$ATLAS_ORG_ID" \
	--arg userName "$userName" \
	'.Usernames?|=[$userName]|.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
	"$(dirname "$0")/inputs_1_invalid.template.json" >"inputs/inputs_1_invalid.json"

ls -l inputs
