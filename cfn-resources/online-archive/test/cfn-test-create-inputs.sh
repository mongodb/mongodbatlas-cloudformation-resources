#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

rm -rf inputs
mkdir inputs

projectName="${1}"
echo $projectName
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="
export MCLI_PROJECT_ID=$projectId

ClusterName="${projectName}"

clusterId=$(atlas clusters create ${ClusterName} --projectId ${projectId} --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json | jq -r '.id')
sleep 600
echo -e "Created Cluster \"${ClusterName}\" with id: ${clusterId}\n"

if [ -z "$clusterId" ]; then
    echo -e "Error Can't find Cluster \"${ClusterName}\""
    exit 1
fi

atlas clusters loadSampleData ${ClusterName} --projectId ${projectId}


clusterName=${ClusterName}
collName="${2:-listingsAndReviews}"
dbName="${3:-sample_airbnb}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg cluster_name "$clusterName" \
   --arg coll_name "$collName" \
   --arg db_name "$dbName" \
   --arg project_id "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey
   | .ClusterName?|=$cluster_name
   | .ProjectId?|=$project_id
   | .DbName?|=$db_name | .CollName?|=$coll_name' \
   "$(dirname "$0")/inputs_1_create.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
  --arg cluster_name "$clusterName" \
  --arg coll_name "$collName" \
  --arg db_name "$dbName" \
  --arg project_id "$projectId" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey
    | .ClusterName?|=$cluster_name
    | .ProjectId?|=$project_id
    | .DbName?|=$db_name | .CollName?|=$coll_name' \
   "$(dirname "$0")/inputs_1_update.json" > "inputs/inputs_1_update.json"

#SET INVALID NAME
clusterName="^%LKJ)(*J_ {+_+O_)"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg clusterName "$clusterName" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |  .ClusterName?|=$clusterName ' \
   "$(dirname "$0")/inputs_1_invalid.json" > "inputs/inputs_1_invalid.json"
