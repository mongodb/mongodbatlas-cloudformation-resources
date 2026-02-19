#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
# It creates all required AWS resources (S3 bucket, IAM role, Cloud Provider Access role)
#

set -euo pipefail

rm -rf inputs
mkdir inputs

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
echo -e "=====\nrun this command to clean up\n=====\natlas projects delete ${projectId} --force\n====="

region=$AWS_DEFAULT_REGION
awsRegion=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
	region=$(aws configure get region)
	awsRegion=$region
fi
regionFormatted=$(echo "$region" | sed -e "s/-/_/g" | tr '[:lower:]' '[:upper:]')
echo "Using region: $region (formatted: $regionFormatted)"

# Use mongodb-atlas-cfn-test-* naming convention (allowed by AWS policy, same as log integration)
bucketTag="${CFN_TEST_TAG:-$(date +%Y%m%d%H%M%S)}"
bucketName="mongodb-atlas-cfn-test-export-bucket-${bucketTag}"
roleName="mongodb-atlas-cloud-backup-export-bucket-$(date +%s)-${RANDOM}"
policyName="atlas-cloud-backup-export-bucket-S3-role-policy-${regionFormatted}"

echo "Bucket name: ${bucketName}"
echo "Creating IAM role: ${roleName}"

# Create cloud provider access entry
roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "Created Atlas cloud provider access entry: ${roleID}"

# Get Atlas AWS Account ARN and External ID
atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId | test($roleID)) | .atlasAssumedRoleExternalId')

# Create trust policy
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
	"$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"

echo "--------------------------------AWS Role creation starts----------------------------"

# Check if role exists, delete if found (must remove inline policy first)
awsRoleId=$(aws iam get-role --role-name "${roleName}" 2>/dev/null | jq -r '.Role.RoleId' || echo "")
if [ -n "$awsRoleId" ]; then
	aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}" 2>/dev/null || true
	aws iam delete-role --role-name "${roleName}"
	echo "Deleted existing role"
fi

# Create IAM role
awsRoleId=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://"$(dirname "$0")/add-policy.json" | jq -r '.Role.RoleId')
echo "Created AWS IAM role: ${awsRoleId}"

# Get role ARN
awsRoleArn=$(aws iam get-role --role-name "${roleName}" | jq -r '.Role.Arn')

echo "--------------------------------AWS Role creation ends----------------------------"

# Wait for AWS IAM role to propagate (similar to encryption-at-rest / stream-connection pattern)
echo "Waiting for IAM role to propagate..."
sleep 65

# Authorize the role in Atlas
echo "--------------------------------Authorize MongoDB Atlas Role starts----------------------------"
atlas cloudProviders accessRoles aws authorize "${roleID}" --projectId "${projectId}" --iamAssumedRoleArn "${awsRoleArn}"
echo "Authorized role: ${roleName}"
echo "--------------------------------Authorize MongoDB Atlas Role ends----------------------------"

echo "--------------------------------Creating S3 Bucket----------------------------"
if aws s3 ls "s3://${bucketName}" 2>/dev/null; then
	aws s3 rb "s3://${bucketName}" --force
fi
aws s3 mb "s3://${bucketName}" --region "${awsRegion}"
echo "Created S3 bucket: ${bucketName}"

echo "--------------------------------Attaching S3 policy to IAM role----------------------------"
aws iam put-role-policy \
	--role-name "${roleName}" \
	--policy-name "${policyName}" \
	--policy-document "file://$(dirname "$0")/policy.json"
echo "--------------------------------attach mongodb Role to AWS Role ends----------------------------"

# Save role name for cleanup (delete script reads this when run from resource root)
echo "${roleName}" > "$(dirname "$0")/role-name.txt"

jq --arg projectId "$projectId" \
	--arg bucketName "$bucketName" \
	--arg iamRoleID "$roleID" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .ProjectId?|=$projectId | .BucketName?|=$bucketName | .IamRoleID?|=$iamRoleID' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

ls -l inputs
