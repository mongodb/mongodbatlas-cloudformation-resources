#!/usr/bin/env bash
# cfn-test-create-inputs.sh
# Tests Atlas access role creation and AWS IAM role authorization

set -e

if [ "$#" -ne 1 ]; then
    echo "usage: $0 <project_name>"
    exit 1
fi

projectName="${1}"

# Get or create project
echo "==> Getting/Creating Atlas project"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
    echo "Created project: ${projectName} (${projectId})"
else
    echo "Found project: ${projectName} (${projectId})"
fi

export MCLI_PROJECT_ID=$projectId

# Get AWS region
keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
    keyRegion=$(aws configure get region)
fi
keyRegion=${keyRegion//-/_}
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')

roleName="mongodb-atlas-enc-role-${keyRegion}"

echo "==> Creating Atlas access role"
roleID=$(atlas cloudProviders accessRoles aws create --output json | jq -r '.roleId')
echo "Atlas roleId: ${roleID}"

echo "==> Getting Atlas account details"
atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAssumedRoleExternalId')

echo "Atlas AWS Account ARN: ${atlasAWSAccountArn}"
echo "Atlas External ID: ${atlasAssumedRoleExternalId}"

echo "==> Generating IAM role trust policy"
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
   '.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
   "$(dirname "$0")/role-policy-template.json" > "$(dirname "$0")/trust-policy.json"

echo "==> Creating AWS IAM role"
# Check if role exists
awsRoleID=$(aws iam get-role --role-name "${roleName}" 2>/dev/null | jq -r '.Role.RoleId' || echo "")
if [ -n "$awsRoleID" ]; then
    echo "Deleting existing role: ${roleName}"
    aws iam delete-role --role-name "${roleName}"
fi

awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://"$(dirname "$0")"/trust-policy.json | jq -r '.Role.RoleId')
echo "Created AWS role: ${roleName} (${awsRoleID})"

awsArn=$(aws iam get-role --role-name "${roleName}" | jq -r '.Role.Arn')
echo "AWS role ARN: ${awsArn}"

echo "==> Waiting for IAM role propagation (65 seconds)"
sleep 65

echo "==> Authorizing Atlas to assume AWS role"
atlas cloudProviders accessRoles aws authorize "${roleID}" --iamAssumedRoleArn "${awsArn}"
echo "Successfully authorized Atlas roleId: ${roleID}"

echo ""
echo "==> Test completed successfully!"
echo "Project ID: ${projectId}"
echo "Atlas Role ID: ${roleID}"
echo "AWS Role Name: ${roleName}"
echo "AWS Role ARN: ${awsArn}"
