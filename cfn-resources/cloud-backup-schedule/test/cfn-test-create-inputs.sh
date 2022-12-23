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
    echo "Creates a new project and an Cluster for testing"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectName="${1}"
ClusterName=$projectName
echo "Creating required inputs"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Cant find project \"${projectName}\"\n"
fi
export MCLI_PROJECT_ID=$projectId


clusterId=$(atlas clusters create ${ClusterName} --projectId ${projectId} --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json | jq -r '.id')
sleep 900
echo -e "Created Cluster \"${ClusterName}\" with id: ${clusterId}\n"

name="${1}"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg group_id "$projectId" \
   --arg clusterName "$ClusterName" \
   '.ClusterName?|=$clusterName |.ProjectId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
  --arg pvtkey "$ATLAS_PRIVATE_KEY" \
  --arg group_id "$projectId" \
  --arg clusterName "$ClusterName" \
  '.ClusterName?|=$clusterName |.ProjectId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
  "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"

name="${name}- more B@d chars !@(!(@====*** ;;::"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg group_id "$projectId" \
   --arg clusterName "$ClusterName" \
     '.ClusterName?|=$clusterName |.ProjectId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"
ls -l inputs