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

# Delete AWS Lambda IAM roles if they exist
echo "--------------------------------delete AWS Lambda IAM roles starts ----------------------------"

# Check if Lambda input files exist
if [ -f "./inputs/inputs_4_create.json" ]; then
	echo "Found Lambda connection inputs, cleaning up IAM roles..."
	
	policyName="atlas-lambda-invoke-policy"
	
	# Extract role ARN from CREATE input file
	roleArnCreate=$(jq -r '.Aws.RoleArn // empty' ./inputs/inputs_4_create.json)
	# Extract role name from ARN (everything after the last '/')
	iamRoleNameCreate=$(echo "${roleArnCreate}" | awk -F'/' '{print $NF}')
	
	# Extract role ARN from UPDATE input file
	roleArnUpdate=$(jq -r '.Aws.RoleArn // empty' ./inputs/inputs_4_update.json)
	# Extract role name from ARN (everything after the last '/')
	iamRoleNameUpdate=$(echo "${roleArnUpdate}" | awk -F'/' '{print $NF}')
	
	# Get external IDs from trust policy files and find roleIds in Atlas
	if [ -f "$(dirname "$0")/lambda-trust-policy-create.json" ]; then
		atlasAssumedRoleExternalIdCreate=$(jq -r '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/lambda-trust-policy-create.json")
		roleIdCreate=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg extId "${atlasAssumedRoleExternalIdCreate}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($extId)) | .roleId')
		
		if [ -n "${roleIdCreate}" ] && [ "${roleIdCreate}" != "null" ]; then
			echo "Deauthorizing CREATE role from Atlas: ${roleIdCreate}"
			atlas cloudProviders accessRoles aws deauthorize "${roleIdCreate}" --projectId "${projectId}" --force || echo "Failed to deauthorize CREATE role"
		fi
	fi
	
	if [ -f "$(dirname "$0")/lambda-trust-policy-update.json" ]; then
		atlasAssumedRoleExternalIdUpdate=$(jq -r '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/lambda-trust-policy-update.json")
		roleIdUpdate=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg extId "${atlasAssumedRoleExternalIdUpdate}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($extId)) | .roleId')
		
		if [ -n "${roleIdUpdate}" ] && [ "${roleIdUpdate}" != "null" ]; then
			echo "Deauthorizing UPDATE role from Atlas: ${roleIdUpdate}"
			atlas cloudProviders accessRoles aws deauthorize "${roleIdUpdate}" --projectId "${projectId}" --force || echo "Failed to deauthorize UPDATE role"
		fi
	fi
	
	# Delete CREATE IAM role
	if [ -n "${iamRoleNameCreate}" ] && [ "${iamRoleNameCreate}" != "null" ] && [ "${iamRoleNameCreate}" != "" ]; then
		echo "Deleting CREATE IAM role: ${iamRoleNameCreate}"
		aws iam delete-role-policy --role-name "${iamRoleNameCreate}" --policy-name "${policyName}" 2>/dev/null || echo "Policy already deleted or doesn't exist"
		aws iam delete-role --role-name "${iamRoleNameCreate}" 2>/dev/null || echo "Role already deleted or doesn't exist"
	fi
	
	# Delete UPDATE IAM role
	if [ -n "${iamRoleNameUpdate}" ] && [ "${iamRoleNameUpdate}" != "null" ] && [ "${iamRoleNameUpdate}" != "" ]; then
		echo "Deleting UPDATE IAM role: ${iamRoleNameUpdate}"
		aws iam delete-role-policy --role-name "${iamRoleNameUpdate}" --policy-name "${policyName}" 2>/dev/null || echo "Policy already deleted or doesn't exist"
		aws iam delete-role --role-name "${iamRoleNameUpdate}" 2>/dev/null || echo "Role already deleted or doesn't exist"
	fi
	
	# Clean up temporary files
	rm -f "$(dirname "$0")/lambda-trust-policy-create.json"
	rm -f "$(dirname "$0")/lambda-trust-policy-update.json"
	
	echo "Cleaned up Lambda IAM roles and temporary files"
else
	echo "No Lambda connection inputs found, skipping IAM role cleanup"
fi
echo "--------------------------------delete AWS Lambda IAM roles ends ----------------------------"

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
