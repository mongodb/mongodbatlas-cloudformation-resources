#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -euo pipefail

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

streamInstanceName="stream-$(date +%s%3N)"
cloudProvider="AWS"
region="VIRGINIA_USA"
tier="SP30"

WORDTOREMOVE="template."

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg project_id "$projectId" \
		--arg stream_instance_name "$streamInstanceName" \
		--arg cloud_provider "$cloudProvider" \
		--arg region "$region" \
		--arg profile "$profile" \
		--arg tier "$tier" \
		'.Profile?|=$profile | .ProjectId?|=$project_id | .InstanceName?|=$stream_instance_name | .DataProcessRegion.CloudProvider?|=$cloud_provider | .DataProcessRegion.Region?|=$region | .StreamConfig.Tier?|=$tier' \
		"$inputFile" >"../inputs/$outputFile"
done

cd ..

ls -l inputs
