#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
echo "$projectName"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

ClusterName="${projectName}"

# shellcheck disable=SC2086
# atlas clusters create "${ClusterName}" --projectId ${projectId} --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json
# atlas clusters watch "${ClusterName}" --projectId "${projectId}"
# echo -e "Created Cluster \"${ClusterName}\""

# atlas clusters loadSampleData "${ClusterName}" --projectId "${projectId}"

clusterName=${ClusterName}
streamInstanceName="stream-$RANDOM"
cloudProvider="AWS"
region="US_EAST_1"
tier="SP30"

jq --arg cluster_name "$clusterName" \
	--arg project_id "$projectId" \
	--arg stream_instance_name "$streamInstanceName" \
	--arg cloud_provider "$cloudProvider" \
	--arg region "$region" \
	--arg profile "$profile" \
	--arg tier "$tier" \
	'.Profile?|=$profile | .ClusterName?|=$cluster_name | .GroupId?|=$project_id | .Name?|=$stream_instance_name | .DataProcessRegion.CloudProvider?|=$cloud_provider | .DataProcessRegion.Region?|=$region | .StreamConfig.Tier?|=$tier' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

jq --arg cluster_name "$clusterName" \
	--arg project_id "$projectId" \
	--arg stream_instance_name "$streamInstanceName" \
	--arg cloud_provider "$cloudProvider" \
	--arg region "$region" \
	--arg profile "$profile" \
	--arg tier "$tier" \
	'.Profile?|=$profile | .ClusterName?|=$cluster_name | .GroupId?|=$project_id | .Name?|=$stream_instance_name | .DataProcessRegion[0].CloudProvider?|=$cloud_provider | .DataProcessRegion[0].Region?|=$region | .StreamConfig[0].Tier?|=$tier' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

