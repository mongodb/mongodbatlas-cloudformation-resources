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


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$ATLAS_PROJECT_ID" \
   --arg regionName "$ATLAS_REGION_NAME" \
   '.ProjectId?|=$projectId | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .RegionName?|=$regionName' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$ATLAS_PROJECT_ID" \
   --arg regionName "$ATLAS_REGION_NAME" \
   '.ProjectId?|=$projectId | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .RegionName?|=$regionName' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$ATLAS_PROJECT_ID" \
   --arg regionName "$ATLAS_REGION_NAME" \
   '.ProjectId?|=$projectId | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .RegionName?|=$regionName' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"

ls -l inputs

echo "TODO: Delete the team and api_key created above"
