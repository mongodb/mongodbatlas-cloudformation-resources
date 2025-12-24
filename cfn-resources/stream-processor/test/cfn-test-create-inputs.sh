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

# Create Stream Instance/Workspace (this is a LONG-RUNNING operation, can take 10-30+ minutes)
workspaceName="stream-workspace-$(date +%s)-$RANDOM"
cloudProvider="AWS"

echo -e "Creating Stream Instance/Workspace \"${workspaceName}\" (this may take 10-30+ minutes)...\n"
atlas streams instances create "${workspaceName}" --projectId "${projectId}" --region VIRGINIA_USA --provider ${cloudProvider}
echo -e "Waiting for Stream Instance/Workspace \"${workspaceName}\" to be ready...\n"
# Poll until the stream instance is ready (watch command doesn't exist for stream instances)
while true; do
	hostnames=$(atlas streams instances describe "${workspaceName}" --projectId "${projectId}" --output json 2>/dev/null | jq -r '.hostnames[]? // empty' 2>/dev/null | head -1)
	if [ -n "$hostnames" ]; then
		echo -e "Stream Instance/Workspace \"${workspaceName}\" is ready\n"
		break
	fi
	sleep 10
done

# For inputs_3 (DLQ testing), we need a cluster and stream connection
# Create cluster for DLQ connection (if needed)
clusterName="cluster-$(date +%s)-$RANDOM"
connectionName="stream-connection-$(date +%s)-$RANDOM"

echo -e "Creating Cluster \"${clusterName}\" for DLQ connection...\n"
atlas clusters create "${clusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
atlas clusters watch "${clusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${clusterName}\"\n"

echo -e "Creating Stream Connection \"${connectionName}\" for DLQ...\n"
# Create temporary JSON file for connection configuration using jq (consistent with rest of script)
connectionConfig=$(mktemp).json
jq -n \
	--arg type "Cluster" \
	--arg clusterName "${clusterName}" \
	'{
		"type": $type,
		"clusterName": $clusterName,
		"dbRoleToExecute": {
			"role": "atlasAdmin",
			"type": "BUILT_IN"
		}
	}' > "${connectionConfig}"
atlas streams connections create "${connectionName}" \
	--projectId "${projectId}" \
	--instance "${workspaceName}" \
	--file "${connectionConfig}" \
	--output=json
rm -f "${connectionConfig}"
echo -e "Created Stream Connection \"${connectionName}\"\n"

# Create Sample connection for inputs_1 and inputs_2 (sample_stream_solar)
sampleConnectionName="sample_stream_solar"
echo -e "Creating Sample Stream Connection \"${sampleConnectionName}\" for inputs_1 and inputs_2...\n"
sampleConnectionConfig=$(mktemp).json
jq -n \
	--arg type "Sample" \
	'{
		"type": $type
	}' > "${sampleConnectionConfig}"
# Check if connection already exists
if atlas streams connections describe "${sampleConnectionName}" --projectId "${projectId}" --instance "${workspaceName}" --output json >/dev/null 2>&1; then
	echo "Sample connection \"${sampleConnectionName}\" already exists, skipping creation"
else
	atlas streams connections create "${sampleConnectionName}" \
		--projectId "${projectId}" \
		--instance "${workspaceName}" \
		--file "${sampleConnectionConfig}" \
		--output=json
	echo -e "Created Sample Stream Connection \"${sampleConnectionName}\"\n"
fi
rm -f "${sampleConnectionConfig}"

# Reuse the Cluster connection from inputs_3 for inputs_1 and inputs_2 sink (saves time/resources)
# No need to create a separate cluster - we'll use the same connectionName

# Generate input files
# Reuse connectionName from inputs_3 for inputs_1 and inputs_2 sink (saves creating another cluster)
# Also set InstanceName from WorkspaceName for primary identifier (both fields required)
jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg sink_connection_name "$connectionName" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .InstanceName?|=$workspace_name
   | .Pipeline?|=gsub("SINK_CONNECTION_PLACEHOLDER"; $sink_connection_name)' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg sink_connection_name "$connectionName" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .InstanceName?|=$workspace_name
   | .Pipeline?|=gsub("SINK_CONNECTION_PLACEHOLDER"; $sink_connection_name)' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg sink_connection_name "$connectionName" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .InstanceName?|=$workspace_name
   | .Pipeline?|=gsub("SINK_CONNECTION_PLACEHOLDER"; $sink_connection_name)' \
	"$(dirname "$0")/inputs_2_create.json" >"inputs/inputs_2_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg sink_connection_name "$connectionName" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .InstanceName?|=$workspace_name
   | .Pipeline?|=gsub("SINK_CONNECTION_PLACEHOLDER"; $sink_connection_name)' \
	"$(dirname "$0")/inputs_2_update.json" >"inputs/inputs_2_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg connection_name "$connectionName" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .InstanceName?|=$workspace_name
   | .Options.Dlq.ConnectionName?|=$connection_name
   | .Pipeline?|=gsub("CONNECTION_NAME_PLACEHOLDER"; $connection_name)' \
	"$(dirname "$0")/inputs_3_create.json" >"inputs/inputs_3_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg connection_name "$connectionName" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .InstanceName?|=$workspace_name
   | .Options.Dlq.ConnectionName?|=$connection_name
   | .Pipeline?|=gsub("CONNECTION_NAME_PLACEHOLDER"; $connection_name)' \
	"$(dirname "$0")/inputs_3_update.json" >"inputs/inputs_3_update.json"

echo -e "Test input files generated successfully in inputs/ directory\n"
