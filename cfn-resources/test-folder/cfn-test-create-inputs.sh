#!/usr/bin/env bash
# cfn-test-create-inputs.sh
# Tests Atlas access role creation and AWS IAM role authorization

set -euo pipefail

rm -rf inputs
mkdir inputs

# Set default project name if not provided
projectName="${1:-${PROJECT_NAME:-cfn-test-access-role-$(date +%s)-$RANDOM}}"
echo "Project name: $projectName"

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
echo -e "=====\nrun this command to clean up\n=====\natlas projects delete ${projectId} --force\n====="

# Get AWS region
keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
    keyRegion=$(aws configure get region)
fi
keyRegion="${keyRegion//-/_}"
keyRegion="${keyRegion^^}"

# Use dynamic role name to avoid conflicts
iamRoleName="mongodb-atlas-enc-role-${keyRegion}-$(date +%s)-${RANDOM}"

echo "IAM Role Name: ${iamRoleName}"

echo "==> Creating Atlas access role"
roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "Atlas roleId: ${roleID}"

echo "==> Getting Atlas account details"
atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAssumedRoleExternalId')

echo "Atlas AWS Account ARN: ${atlasAWSAccountArn}"
echo "Atlas External ID: ${atlasAssumedRoleExternalId}"

echo "==> Generating IAM role trust policy"
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
   '.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
   "$(dirname "$0")/role-policy-template.json" > "$(dirname "$0")/trust-policy.json"

echo "==> Creating AWS IAM role"
# Check if role exists, delete if found
awsRoleID=$(aws iam get-role --role-name "${iamRoleName}" 2>/dev/null | jq --arg roleName "${iamRoleName}" -r '.Role | select(.RoleName==$roleName) | .RoleId' || echo "")
if [ -n "$awsRoleID" ]; then
    echo "Deleting existing role: ${iamRoleName}"
    aws iam delete-role --role-name "${iamRoleName}"
fi

awsRoleID=$(aws iam create-role --role-name "${iamRoleName}" --assume-role-policy-document file://"$(dirname "$0")"/trust-policy.json | jq --arg roleName "${iamRoleName}" -r '.Role | select(.RoleName==$roleName) | .RoleId')
echo "Created AWS role: ${iamRoleName} (${awsRoleID})"

awsArn=$(aws iam get-role --role-name "${iamRoleName}" | jq --arg roleName "${iamRoleName}" -r '.Role | select(.RoleName==$roleName) | .Arn')
echo "AWS role ARN: ${awsArn}"

echo "==> Waiting for IAM role propagation (65 seconds)"
sleep 65

echo "==> Authorizing Atlas to assume AWS role"
atlas cloudProviders accessRoles aws authorize "${roleID}" --iamAssumedRoleArn "${awsArn}" --projectId "${projectId}"
echo "Successfully authorized Atlas roleId: ${roleID}"

# Save metadata to JSON file for cleanup (similar to stream-connection pattern)
jq -n \
  --arg projectId "$projectId" \
  --arg projectName "$projectName" \
  --arg roleId "$roleID" \
  --arg iamRoleName "$iamRoleName" \
  --arg awsArn "$awsArn" \
  '{
    ProjectId: $projectId,
    ProjectName: $projectName,
    AtlasRoleId: $roleId,
    IamRoleName: $iamRoleName,
    AwsRoleArn: $awsArn
  }' > "inputs/test-metadata.json"

echo ""
echo "==> Test completed successfully!"
echo "Project ID: ${projectId}"
echo "Atlas Role ID: ${roleID}"
echo "AWS Role Name: ${iamRoleName}"
echo "AWS Role ARN: ${awsArn}"