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

projectName="${1:-$PROJECT_NAME}"

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

streamInstanceName="stream-$RANDOM"
cloudProvider="AWS"
region="US_EAST_1"
tier="SP30"

jq --arg project_id "$projectId" \
	--arg stream_instance_name "$streamInstanceName" \
	--arg cloud_provider "$cloudProvider" \
	--arg region "$region" \
	--arg profile "$profile" \
	--arg tier "$tier" \
	'.Profile?|=$profile | .GroupId?|=$project_id | .Name?|=$stream_instance_name | .DataProcessRegion.CloudProvider?|=$cloud_provider | .DataProcessRegion.Region?|=$region | .StreamConfig.Tier?|=$tier' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg project_id "$projectId" \
	--arg stream_instance_name "$streamInstanceName" \
	--arg cloud_provider "$cloudProvider" \
	--arg region "$region" \
	--arg profile "$profile" \
	--arg tier "$tier" \
	'.Profile?|=$profile | .GroupId?|=$project_id | .Name?|=$stream_instance_name | .DataProcessRegion.CloudProvider?|=$cloud_provider | .DataProcessRegion.Region?|=$region | .StreamConfig.Tier?|=$tier' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

