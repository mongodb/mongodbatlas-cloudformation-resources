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

# Use existing project ID if set, otherwise try to find or create project
if [ -n "${MONGODB_ATLAS_PROJECT_ID:-}" ]; then
	projectId="${MONGODB_ATLAS_PROJECT_ID}"
	echo -e "Using existing project ID from MONGODB_ATLAS_PROJECT_ID: ${projectId}\n"
else
	projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
	if [ -z "$projectId" ]; then
		projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
		echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
	else
		echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
	fi
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

workspaceName="stream-workspace-$(date +%s)-$RANDOM"
cloudProvider="AWS"
clusterName="cluster-$(date +%s)-$RANDOM"


atlas streams instances create "${workspaceName}" --projectId "${projectId}" --region VIRGINIA_USA --provider ${cloudProvider}
echo -e "Created StreamWorkspace \"${workspaceName}\""

atlas clusters create "${clusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
atlas clusters watch "${clusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${clusterName}\""

jq --arg cluster_name "$clusterName" \
	--arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .ClusterName?|=$cluster_name
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

jq --arg cluster_name "$clusterName" \
	--arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .ClusterName?|=$cluster_name
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_2_create.json" >"inputs/inputs_2_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_2_update.json" >"inputs/inputs_2_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_3_create.json" >"inputs/inputs_3_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_3_update.json" >"inputs/inputs_3_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_4_create.json" >"inputs/inputs_4_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_4_update.json" >"inputs/inputs_4_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_5_create.json" >"inputs/inputs_5_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_5_update.json" >"inputs/inputs_5_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_6_create.json" >"inputs/inputs_6_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_6_update.json" >"inputs/inputs_6_update.json"
	