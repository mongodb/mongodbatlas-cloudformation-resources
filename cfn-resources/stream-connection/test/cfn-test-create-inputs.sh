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

# Single IAM role for both CREATE and UPDATE scenarios (following Terraform pattern)
iamRoleName="mongodb-atlas-streams-lambda-$(date +%s)-${RANDOM}"

echo "Creating IAM role: ${iamRoleName}"

# Create cloud provider access entry
roleId=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "Created Atlas cloud provider access entry: ${roleId}"

# Get Atlas AWS Account ARN and External ID
atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleId}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleId}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAssumedRoleExternalId')

# Create trust policy
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
	"$(dirname "$0")/lambda-role-policy-template.json" >"$(dirname "$0")/lambda-trust-policy.json"

echo "--------------------------------AWS IAM Role creation starts ----------------------------"

# Check if role exists, delete if found
awsRoleId=$(aws iam get-role --role-name "${iamRoleName}" 2>/dev/null | jq --arg roleName "${iamRoleName}" -r '.Role | select(.RoleName==$roleName) | .RoleId' || echo "")
if [ -n "$awsRoleId" ]; then
	aws iam delete-role --role-name "${iamRoleName}"
	echo "Deleted existing role"
fi

# Create IAM role
awsRoleId=$(aws iam create-role --role-name "${iamRoleName}" --assume-role-policy-document file://"$(dirname "$0")"/lambda-trust-policy.json | jq --arg roleName "${iamRoleName}" -r '.Role | select(.RoleName==$roleName) | .RoleId')
echo "Created AWS IAM role: ${awsRoleId}"

# Get role ARN
awsArn=$(aws iam get-role --role-name "${iamRoleName}" | jq --arg roleName "${iamRoleName}" -r '.Role | select(.RoleName==$roleName) | .Arn')

echo "--------------------------------AWS IAM Role creation ends ----------------------------"

# Wait for AWS IAM role to propagate (similar to encryption-at-rest pattern)
echo "Waiting for IAM role to propagate..."
sleep 65

# Authorize the role in Atlas
echo "--------------------------------Authorize MongoDB Atlas Role starts ----------------------------"
atlas cloudProviders accessRoles aws authorize "${roleId}" --iamAssumedRoleArn "${awsArn}" --projectId "${projectId}"
echo "Authorized role: ${iamRoleName}"
echo "--------------------------------Authorize MongoDB Atlas Role ends ----------------------------"

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
	--arg role_arn "$awsArn" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .Aws.RoleArn=$role_arn' \
	"$(dirname "$0")/inputs_4_create.json" >"inputs/inputs_4_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg role_arn "$awsArn" \
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

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_7_create.json" >"inputs/inputs_7_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_7_update.json" >"inputs/inputs_7_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_8_create.json" >"inputs/inputs_8_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name' \
	"$(dirname "$0")/inputs_8_update.json" >"inputs/inputs_8_update.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg role_arn "$awsArn" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .Aws.RoleArn=$role_arn' \
	"$(dirname "$0")/inputs_9_create.json" >"inputs/inputs_9_create.json"

jq --arg workspace_name "$workspaceName" \
	--arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg role_arn "$awsArn" \
	'.Profile?|=$profile
   | .ProjectId?|=$project_id
   | .WorkspaceName?|=$workspace_name
   | .Aws.RoleArn=$role_arn' \
	"$(dirname "$0")/inputs_9_update.json" >"inputs/inputs_9_update.json"
