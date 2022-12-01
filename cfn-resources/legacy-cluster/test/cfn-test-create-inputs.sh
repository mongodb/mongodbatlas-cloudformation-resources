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
    echo "usage:$0 <project/cluster_name>"
    echo "Creates a new project and cluster by that name for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
projectName="${1}"
projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')
echo "Created project \"${projectName}\" with id: ${projectId}"

clusterName="${projectName}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg region "us-east-1" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |  .Name?|=$clusterName | .ProviderSettings.RegionName?|=$region | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_create.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg region "us-east-1" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |  .Name?|=$clusterName | .ProviderSettings.RegionName?|=$region | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_update.json" > "inputs/inputs_1_update.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg region "us-east-1" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |  .Name?|=$clusterName | .ProviderSettings.RegionName?|=$region | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_invalid.json" > "inputs/inputs_1_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"
