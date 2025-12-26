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

# Try to find existing project first, or use provided project ID
if [ ${MONGODB_ATLAS_PROJECT_ID+x} ] && [ -n "${MONGODB_ATLAS_PROJECT_ID}" ]; then
	projectId="${MONGODB_ATLAS_PROJECT_ID}"
	echo -e "Using provided project ID: ${projectId}\n"
else
	projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
	if [ -z "$projectId" ]; then
		# Try to create project, but if IP is not on access list, use default project ID
		if projectId=$(atlas projects create "${projectName}" --output=json 2>/dev/null | jq -r '.id'); then
			echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
		else
			# Fallback to project ID from environment variable
			if [ -z "${MONGODB_ATLAS_PROJECT_ID:-}" ]; then
				echo -e "ERROR: Could not create project and MONGODB_ATLAS_PROJECT_ID is not set. Please set MONGODB_ATLAS_PROJECT_ID environment variable.\n"
				exit 1
			fi
			projectId="${MONGODB_ATLAS_PROJECT_ID}"
			echo -e "Could not create project (IP may not be on access list). Using project ID from environment: ${projectId}\n"
		fi
	else
		echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
	fi
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
