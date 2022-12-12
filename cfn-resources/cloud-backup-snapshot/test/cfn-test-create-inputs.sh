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


echo "Came inside create inputs to test"
 if [ -z "${ClusterName}" ] || [ -z "${ProjectName}" ]; then
          echo "Test" "$ClusterName" "$ProjectName"
          echo "Error is testing cloud-provider-snapshots, we need ClusterName and SnapshotId provided in OtherParams during Automation. Kindly provide these values.
          Example: 'ClusterName'='cluster-123','ProjectName'='Project-123'"
          exit 1
 fi

#if [ -z "$ProjectName" ]; then
#   projectName="${1}"
#  else
#   projectName="$ProjectName"
#fi

echo "Project Name  $ProjectName"
projectId=$(atlas projects list --output json | jq --arg NAME "${ProjectName}" -r '.results[] | select(.name==$NAME) | .id')
#if [ -z "$projectId" ]; then
#    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
#    echo -e "Cant find project \"${projectName}\"\n"
#fi
#echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

export MCLI_PROJECT_ID=$projectId

clusterId=$(atlas clusters describe "${ClusterName}"  --output json | jq -r '.id')

if [ -z "$clusterId" ]; then
    echo -e "Error Can't find Cluster \"${ClusterName}\""
    exit 1
fi
#SnapshotId=$(atlas backup snapshots list "${ClusterName}" --output json  | jq --arg ID "$SnapshotId" -r '.results[] | select(.id==$ID) | .id')
#if [ -z "$SnapshotId" ]; then
#    echo -e "Error Can't find SnapshotId \"${SnapshotId}\""
#    exit 1
#fi
rm -rf inputs
mkdir inputs
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg group_id "$projectId" \
   --arg clusterName "$ClusterName" \
   '.ClusterName?|=$clusterName |.GroupId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

ClusterName="${ClusterName}- more B@d chars !@(!(@====*** ;;::"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg group_id "$projectId" \
   --arg clusterName "$ClusterName" \
   '.ClusterName?|=$clusterName |.GroupId?|=$group_id |.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"
ls -l inputs
