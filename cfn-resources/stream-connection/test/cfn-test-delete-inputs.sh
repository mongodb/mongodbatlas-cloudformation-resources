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

# Delete AWS Lambda IAM role if it exists
echo "--------------------------------delete AWS Lambda IAM role starts ----------------------------"

# Check if Lambda input files exist
if [ -f "./inputs/inputs_4_create.json" ]; then
	echo "Found Lambda connection inputs, cleaning up IAM role..."
	
	# Extract role ARN from CREATE input file (same role used for both CREATE and UPDATE)
	roleArn=$(jq -r '.Aws.RoleArn // empty' ./inputs/inputs_4_create.json)
	# Extract role name from ARN (everything after the last '/')
	iamRoleName=$(echo "${roleArn}" | awk -F'/' '{print $NF}')
	
	# Get external ID from trust policy file and find roleId in Atlas
	if [ -f "$(dirname "$0")/lambda-trust-policy.json" ]; then
		atlasAssumedRoleExternalId=$(jq -r '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/lambda-trust-policy.json")
		roleId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg extId "${atlasAssumedRoleExternalId}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($extId)) | .roleId')
		
		if [ -n "${roleId}" ] && [ "${roleId}" != "null" ]; then
			echo "Deauthorizing role from Atlas: ${roleId}"
			atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force || echo "Failed to deauthorize role"
		fi
	fi
	
	# Delete IAM role
	if [ -n "${iamRoleName}" ] && [ "${iamRoleName}" != "null" ] && [ "${iamRoleName}" != "" ]; then
		echo "Deleting IAM role: ${iamRoleName}"
		aws iam delete-role --role-name "${iamRoleName}" 2>/dev/null || echo "Role already deleted or doesn't exist"
	fi
	
	# Clean up temporary file
	rm -f "$(dirname "$0")/lambda-trust-policy.json"
	
	echo "Cleaned up Lambda IAM role and temporary files"
else
	echo "No Lambda connection inputs found, skipping IAM role cleanup"
fi
echo "--------------------------------delete AWS Lambda IAM role ends ----------------------------"

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
