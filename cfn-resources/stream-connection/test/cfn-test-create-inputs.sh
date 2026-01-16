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

# AWS IAM role creation and authorization for Lambda connections
echo "--------------------------------AWS Lambda IAM Role creation starts ----------------------------"

# Role names for CREATE and UPDATE scenarios
iamRoleNameCreate="mongodb-atlas-streams-lambda-$(date +%s)-${RANDOM}"
iamRoleNameUpdate="mongodb-atlas-streams-lambda-$(date +%s)-${RANDOM}-updated"
policyName="atlas-lambda-invoke-policy"

echo "Creating IAM roles: ${iamRoleNameCreate} and ${iamRoleNameUpdate}"

# Create first cloud provider access entry (for CREATE role)
roleIdCreate=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "Created Atlas cloud provider access entry for CREATE role: ${roleIdCreate}"

# Create second cloud provider access entry (for UPDATE role)
roleIdUpdate=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "Created Atlas cloud provider access entry for UPDATE role: ${roleIdUpdate}"

# Get Atlas AWS Account ARN and External ID for CREATE role
atlasAWSAccountArnCreate=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleIdCreate}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAWSAccountArn')
atlasAssumedRoleExternalIdCreate=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleIdCreate}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAssumedRoleExternalId')

# Get Atlas AWS Account ARN and External ID for UPDATE role
atlasAWSAccountArnUpdate=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleIdUpdate}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAWSAccountArn')
atlasAssumedRoleExternalIdUpdate=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleIdUpdate}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAssumedRoleExternalId')

# Create trust policy for CREATE role
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalIdCreate" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArnCreate" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
	"$(dirname "$0")/lambda-role-policy-template.json" >"$(dirname "$0")/lambda-trust-policy-create.json"

# Create trust policy for UPDATE role
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalIdUpdate" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArnUpdate" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
	"$(dirname "$0")/lambda-role-policy-template.json" >"$(dirname "$0")/lambda-trust-policy-update.json"

echo "--------------------------------AWS IAM Role creation starts ----------------------------"

# Check if CREATE role exists, delete if found
awsRoleIdCreate=$(aws iam get-role --role-name "${iamRoleNameCreate}" 2>/dev/null | jq --arg roleName "${iamRoleNameCreate}" -r '.Role | select(.RoleName==$roleName) | .RoleId' || echo "")
if [ -n "$awsRoleIdCreate" ]; then
	aws iam delete-role-policy --role-name "${iamRoleNameCreate}" --policy-name "${policyName}" 2>/dev/null || true
	aws iam delete-role --role-name "${iamRoleNameCreate}"
	echo "Deleted existing CREATE role"
fi

# Create CREATE role
awsRoleIdCreate=$(aws iam create-role --role-name "${iamRoleNameCreate}" --assume-role-policy-document file://"$(dirname "$0")"/lambda-trust-policy-create.json | jq --arg roleName "${iamRoleNameCreate}" -r '.Role | select(.RoleName==$roleName) | .RoleId')
echo "Created AWS IAM role for CREATE: ${awsRoleIdCreate}"

# Check if UPDATE role exists, delete if found
awsRoleIdUpdate=$(aws iam get-role --role-name "${iamRoleNameUpdate}" 2>/dev/null | jq --arg roleName "${iamRoleNameUpdate}" -r '.Role | select(.RoleName==$roleName) | .RoleId' || echo "")
if [ -n "$awsRoleIdUpdate" ]; then
	aws iam delete-role-policy --role-name "${iamRoleNameUpdate}" --policy-name "${policyName}" 2>/dev/null || true
	aws iam delete-role --role-name "${iamRoleNameUpdate}"
	echo "Deleted existing UPDATE role"
fi

# Create UPDATE role
awsRoleIdUpdate=$(aws iam create-role --role-name "${iamRoleNameUpdate}" --assume-role-policy-document file://"$(dirname "$0")"/lambda-trust-policy-update.json | jq --arg roleName "${iamRoleNameUpdate}" -r '.Role | select(.RoleName==$roleName) | .RoleId')
echo "Created AWS IAM role for UPDATE: ${awsRoleIdUpdate}"

# Get role ARNs
awsArnCreate=$(aws iam get-role --role-name "${iamRoleNameCreate}" | jq --arg roleName "${iamRoleNameCreate}" -r '.Role | select(.RoleName==$roleName) | .Arn')
awsArnUpdate=$(aws iam get-role --role-name "${iamRoleNameUpdate}" | jq --arg roleName "${iamRoleNameUpdate}" -r '.Role | select(.RoleName==$roleName) | .Arn')

# Attach Lambda permissions to both roles
aws iam put-role-policy --role-name "${iamRoleNameCreate}" --policy-name "${policyName}" --policy-document file://"$(dirname "$0")"/lambda-permissions-template.json
aws iam put-role-policy --role-name "${iamRoleNameUpdate}" --policy-name "${policyName}" --policy-document file://"$(dirname "$0")"/lambda-permissions-template.json
echo "Attached Lambda invoke permissions to both roles"

echo "--------------------------------AWS IAM Role creation ends ----------------------------"

# Wait for AWS IAM role to propagate (similar to encryption-at-rest pattern)
echo "Waiting for IAM roles to propagate..."
sleep 65

# Authorize the roles in Atlas
echo "--------------------------------Authorize MongoDB Atlas Roles starts ----------------------------"
atlas cloudProviders accessRoles aws authorize "${roleIdCreate}" --iamAssumedRoleArn "${awsArnCreate}" --projectId "${projectId}"
echo "Authorized CREATE role: ${iamRoleNameCreate}"

atlas cloudProviders accessRoles aws authorize "${roleIdUpdate}" --iamAssumedRoleArn "${awsArnUpdate}" --projectId "${projectId}"
echo "Authorized UPDATE role: ${iamRoleNameUpdate}"
echo "--------------------------------Authorize MongoDB Atlas Roles ends ----------------------------"

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
	--arg role_arn "$awsArnCreate" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .Aws.RoleArn=$role_arn' \
	"$(dirname "$0")/inputs_4_create.json" >"inputs/inputs_4_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg role_arn "$awsArnUpdate" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .Aws.RoleArn=$role_arn' \
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
