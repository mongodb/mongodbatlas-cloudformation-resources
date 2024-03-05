#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -euo pipefail

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1:-$PROJECT_NAME}"
echo "$projectName"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

instanceName="stream-instance-$RANDOM"
cloudProvider="AWS"
clusterName="cluster-$RANDOM"

atlas streams instances create "${instanceName}" --projectId "${projectId}" --region VIRGINIA_USA --provider ${cloudProvider}
echo -e "Created StreamInstance \"${instanceName}\""

atlas clusters create "${clusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 6.0 --diskSizeGB 10 --output=json
atlas clusters watch "${clusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${clusterName}\""

jq --arg cluster_name "$clusterName" \
	--arg instance_name "$instanceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .ClusterName?|=$cluster_name
   | .ProjectId?|=$project_id
   | .InstanceName?|=$instance_name' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

jq --arg cluster_name "$clusterName" \
	--arg instance_name "$instanceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .ClusterName?|=$cluster_name
   | .ProjectId?|=$project_id
   | .InstanceName?|=$instance_name' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

jq --arg instance_name "$instanceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .InstanceName?|=$instance_name' \
	"$(dirname "$0")/inputs_2_create.json" >"inputs/inputs_2_create.json"

jq --arg instance_name "$instanceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .InstanceName?|=$instance_name' \
	"$(dirname "$0")/inputs_2_update.json" >"inputs/inputs_2_update.json"
