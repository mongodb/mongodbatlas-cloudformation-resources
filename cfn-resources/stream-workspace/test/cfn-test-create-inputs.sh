#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -euo pipefail

rm -rf inputs
mkdir inputs

projectName="${1:-${PROJECT_NAME:-test-stream-workspace}}"

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

# Try to find existing project first, or use provided project ID
if [ ${MONGODB_ATLAS_PROJECT_ID+x} ] && [ -n "${MONGODB_ATLAS_PROJECT_ID}" ]; then
	projectId="${MONGODB_ATLAS_PROJECT_ID}"
	echo -e "Using provided project ID: ${projectId}\n"
else
	projectId=$(resolve_project_id "${projectName}")
	if [ -z "$projectId" ]; then
		echo "Failed to create or find project ${projectName}" >&2
		exit 1
	fi
	echo -e "Using project \"${projectName}\" with id: ${projectId}\n"
fi

streamWorkspaceName="stream-workspace-$(date +%s)-$RANDOM"
cloudProvider="AWS"
region="VIRGINIA_USA"
tier="SP30"

WORDTOREMOVE="template."

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg project_id "$projectId" \
		--arg stream_workspace_name "$streamWorkspaceName" \
		--arg cloud_provider "$cloudProvider" \
		--arg region "$region" \
		--arg profile "$profile" \
		--arg tier "$tier" \
		'.Profile?|=$profile | .ProjectId?|=$project_id | .WorkspaceName?|=$stream_workspace_name | .DataProcessRegion.CloudProvider?|=$cloud_provider | .DataProcessRegion.Region?|=$region | .StreamConfig.Tier?|=$tier' \
		"$inputFile" >"../inputs/$outputFile"
done

cd ..

ls -l inputs
