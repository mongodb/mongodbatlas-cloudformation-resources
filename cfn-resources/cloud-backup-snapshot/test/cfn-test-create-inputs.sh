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
projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

ClusterName="${projectName}"
clusterId=$(mongocli atlas clusters list --output json  | jq --arg NAME ${ClusterName} -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
    clusterId=$(mongocli atlas cluster create ${ClusterName} --projectId ${projectId} --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json | jq -r '.id')
    sleep 20m
    echo -e "Created Cluster \"${ClusterName}\" with id: ${clusterId}\n"
else
    echo -e "FOUND Cluster \"${ClusterName}\" with id: ${clusterId}\n"
fi

rm -rf inputs
mkdir inputs
name="${1}"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg group_id "$projectId" \
   --arg clusterName "$ClusterName" \
   '.ClusterName?|=$clusterName |.ProjectId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

name="${name}- more B@d chars !@(!(@====*** ;;::"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg group_id "$projectId" \
   --arg clusterName "$ClusterName" \
     '.ClusterName?|=$clusterName |.ProjectId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"
ls -l inputs
