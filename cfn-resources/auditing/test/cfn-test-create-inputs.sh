#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

echo "$ATLAS_PUBLIC_KEY"
echo "$ATLAS_PRIVATE_KEY"
echo "$projectId"

function usage {
    echo "Creates a new customdb role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
projectName="${1}"

projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"

clusterName="${projectName}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg groupId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$groupId ' \
   "$(dirname "$0")/inputs_1_create.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_update.json" > "inputs/inputs_1_update.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_invalid.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$projectId ' \
   "$(dirname "$0")/inputs_2_create.json" > "inputs/inputs_2_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$projectId ' \
   "$(dirname "$0")/inputs_2_update.json" > "inputs/inputs_2_update.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$projectId ' \
   "$(dirname "$0")/inputs_2_invalid.json" > "inputs/inputs_2_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"
