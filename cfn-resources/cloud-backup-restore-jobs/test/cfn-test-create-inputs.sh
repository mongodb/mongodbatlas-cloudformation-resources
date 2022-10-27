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
region="us-east-1"

#project_id
if [[   -v ATLAS_EXISTING_PROJECT_NAME ]]; then
      projectName=$ATLAS_EXISTING_PROJECT_NAME
      echo "Created project \"${projectName}\" with id:"
else
      projectName="${1}"
      echo "Createdddddd project \"${projectName}\" with id:"
fi

projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"
ClusterName="${projectName}"
clusterId=$(mongocli atlas clusters list --output json  | jq --arg NAME ${ClusterName} -r '.results[] | select(.name==$NAME) | .name')
if [ -z "$clusterId" ]; then
    clusterId=$(mongocli atlas cluster create ${ClusterName} --projectId ${projectId} --provider AWS --region US_EAST_1 --members 3 --backup --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json | jq -r '.name')
    sleep 20m
    echo -e "Created Cluster \"${ClusterName}\" with id: ${clusterId}\n"
else
    echo -e "FOUND Cluster \"${ClusterName}\" with id: ${clusterId}\n"
fi

SnapshotId=$(mongocli atlas backup snapshots list cfntest --output json  | jq --arg ID "6321433" -r '.results[] | select(.id==$ID) | .id')
if [ -z "$SnapshotId" ]; then
    SnapshotId=$(mongocli atlas backup snapshots create ${ClusterName} --desc "cfn unit test" --retention 3 --output=json | jq -r '.id')
    sleep 5m
    echo -e "Created Cluster \"${ClusterName}\" with id: ${

    }\n"
else
        echo -e "FOUND Cluster \"${ClusterName}\" with id: ${SnapshotId}\n"
    fi


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg group_id "$projectId" \
   --arg SnapshotId "$SnapshotId" \
   '.SnapshotId?|=$SnapshotId | .ProjectId?|=$group_id | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ClusterName?|=$ClusterName' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg region "${region}- more B@d chars !@(!(@====*** ;;::" \
   --arg group_id "$projectId" \
   --arg ClusterName "$ClusterName" \
   --arg SnapshotId "$SnapshotId" \
   '.SnapshotId?|=$SnapshotId |.ProjectId?|=$group_id | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ClusterName?|=$ClusterName' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

echo "mongocli iam projects delete ${projectId} --force"

ls -l inputs
