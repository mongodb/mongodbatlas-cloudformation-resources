#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -euo pipefail

function usage {
	echo "usage:$0 "
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
workspaceName=$(jq -r '.WorkspaceName // .InstanceName' ./inputs/inputs_1_create.json)
processorName1=$(jq -r '.ProcessorName' ./inputs/inputs_1_create.json)
processorName2=$(jq -r '.ProcessorName' ./inputs/inputs_2_create.json)
processorName3=$(jq -r '.ProcessorName' ./inputs/inputs_3_create.json)

# Delete stream processors (if they exist)
for processorName in "$processorName1" "$processorName2" "$processorName3"; do
	if atlas streams processors delete "${processorName}" \
		--projectId "${projectId}" \
		--instance "${workspaceName}" \
		--force 2>/dev/null; then
		echo "deleted stream processor with name ${processorName}"
	else
		echo "failed to delete or stream processor '${processorName}' does not exist"
	fi
done

# Delete Sample connection (sample_stream_solar) if it exists
sampleConnectionName="sample_stream_solar"
if atlas streams connections delete "${sampleConnectionName}" \
	--projectId "${projectId}" \
	--instance "${workspaceName}" \
	--force 2>/dev/null; then
	echo "deleted sample stream connection with name ${sampleConnectionName}"
else
	echo "failed to delete or sample stream connection '${sampleConnectionName}' does not exist"
fi

# Get connection name from inputs_3 if it exists
if [ -f "./inputs/inputs_3_create.json" ]; then
	connectionName=$(jq -r '.Options.Dlq.ConnectionName // empty' ./inputs/inputs_3_create.json)
	if [ -n "$connectionName" ]; then
		if atlas streams connections delete "${connectionName}" \
			--projectId "${projectId}" \
			--instance "${workspaceName}" \
			--force 2>/dev/null; then
			echo "deleted stream connection with name ${connectionName}"
		else
			echo "failed to delete or stream connection '${connectionName}' does not exist"
		fi
	fi
fi

# Delete all clusters in the project (created for DLQ testing)
# The cluster name is not stored in input JSON, so we list and delete all clusters
# Clusters must be deleted before stream instance and project to avoid dependency conflicts
echo "Checking for clusters to delete in project ${projectId}..."
clusterList=$(atlas clusters list --projectId "${projectId}" --output json 2>/dev/null | jq -r '.results[]?.name // empty' 2>/dev/null || echo "")
if [ -n "$clusterList" ]; then
	while IFS= read -r clusterName; do
		if [ -n "$clusterName" ] && [ "$clusterName" != "null" ] && [ "$clusterName" != "" ]; then
			if atlas cluster delete "${clusterName}" --projectId "${projectId}" --force 2>/dev/null; then
				echo "deleting cluster with name ${clusterName}"
				# Wait for cluster deletion to complete
				atlas cluster watch "${clusterName}" --projectId "${projectId}" 2>/dev/null || true
			else
				echo "failed to delete or cluster '${clusterName}' does not exist"
			fi
		fi
	done <<< "$clusterList"
else
	echo "No clusters found in project"
fi

# Delete stream instance/workspace (after clusters are deleted)
if atlas streams instances delete "${workspaceName}" --projectId "${projectId}" --force 2>/dev/null; then
	echo "deleting stream instance/workspace with name ${workspaceName}"
	# Wait for deletion to complete
	atlas streams instances watch "${workspaceName}" --projectId "${projectId}" 2>/dev/null || true
else
	echo "failed to delete or stream instance/workspace '${workspaceName}' does not exist"
fi

#delete project
if atlas projects delete "$projectId" --force 2>/dev/null; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
