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
clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)

# Get workspace name or instance name (workspace name takes precedence)
workspaceName=$(jq -r '.WorkspaceName // empty' ./inputs/inputs_1_create.json)
instanceName=$(jq -r '.InstanceName // empty' ./inputs/inputs_1_create.json)

# Use WorkspaceName if available, otherwise fall back to InstanceName
if [ -n "${workspaceName}" ] && [ "${workspaceName}" != "null" ] && [ "${workspaceName}" != "" ]; then
	workspaceOrInstanceName="${workspaceName}"
elif [ -n "${instanceName}" ] && [ "${instanceName}" != "null" ] && [ "${instanceName}" != "" ]; then
	workspaceOrInstanceName="${instanceName}"
else
	echo "Error: Neither WorkspaceName nor InstanceName found in inputs_1_create.json"
	exit 1
fi

if atlas cluster delete "${clusterName}" --projectId "${projectId}" --force; then
	echo "deleting cluster with name ${clusterName}"
else
	echo "failed to delete the cluster with name ${clusterName}"
fi

atlas cluster watch "${clusterName}" --projectId "${projectId}" && status=0 || status=$?
if [ "$status" -eq 0 ]; then
	echo "Cluster '${clusterName}' has been successfully watched until deletion."
fi

#delete stream workspace/instance (using instances delete for backward compatibility)
if atlas streams instances delete "${workspaceOrInstanceName}" --projectId "${projectId}" --force; then
	echo "deleting stream workspace/instance with name ${workspaceOrInstanceName}"
else
	echo "failed to delete the stream workspace/instance with name ${workspaceOrInstanceName}"
fi

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
