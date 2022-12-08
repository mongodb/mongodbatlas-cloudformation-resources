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

function usage {
    echo "Creates a new private endpoint role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
region="us-east-1"
projectName="${1}"
vpcId="${2}"
subnetId="${3}"

projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg groupId "$projectId" \
   --arg region "$region" \
   --arg vpcId "$vpcId" \
   --arg subnetId "$subnetId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$groupId | .Region?|=$region | .VpcId?|=$vpcId | .SubnetId?|=$subnetId' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg groupId "dsafasdgsdhsdfh" \
   --arg region "$region" \
   --arg vpcId "$vpcId" \
   --arg subnetId "$subnetId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$groupId | .Region?|=$region | .VpcId?|=$vpcId | .SubnetId?|=$subnetId' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"
