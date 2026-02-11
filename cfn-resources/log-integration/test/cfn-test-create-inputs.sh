#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
# It creates all required AWS resources (S3 bucket, IAM role, Cloud Provider Access role)
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage: $0 <project_name>"
	echo "Creates S3 bucket, Cloud Provider Access role, IAM role, and generates test input files for log integration"
	exit 0
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

region=$AWS_DEFAULT_REGION
awsRegion=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
	region=$(aws configure get region)
	awsRegion=$region
fi

regionFormatted=$(echo "$region" | sed -e "s/-/_/g" | tr '[:lower:]' '[:upper:]')
echo "Using region: $region (formatted: $regionFormatted)"

# Use dynamic role name to avoid conflicts in CI (matches test-folder pattern)
roleName="mongodb-atlas-logs-role-${regionFormatted}-$(date +%s)-${RANDOM}"
policyName="atlas-logs-s3-policy-${regionFormatted}"
bucketTag="${CFN_TEST_TAG:-$(date +%Y%m%d%H%M%S)}"
bucketName="mongodb-atlas-cfn-test-logs-${bucketTag}"

echo "Bucket name: ${bucketName}"
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "--------------------------------Creating Cloud Provider Access Role----------------------------"
roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "--------------------------------Mongo CLI Role creation ends----------------------------"

atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId == $roleID) | .atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] | select(.roleId == $roleID) | .atlasAssumedRoleExternalId')

jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS = $atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"] = $atlasAssumedRoleExternalId' \
	"$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/trust-policy.json"
echo "--------------------------------AWS Role creation starts----------------------------"
awsRoleID=$(aws iam get-role --role-name "${roleName}" 2>/dev/null | jq -r '.Role.RoleId' || echo "")
if [ -z "$awsRoleID" ]; then
	awsRoleID=$(aws iam create-role \
		--role-name "${roleName}" \
		--assume-role-policy-document "file://$(dirname "$0")/trust-policy.json" | jq -r '.Role.RoleId')
	echo -e "No role found, hence creating the role. Created id: ${awsRoleID}\n"
else
	aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}" 2>/dev/null || true
	aws iam delete-role --role-name "${roleName}"
	awsRoleID=$(aws iam create-role \
		--role-name "${roleName}" \
		--assume-role-policy-document "file://$(dirname "$0")/trust-policy.json" | jq -r '.Role.RoleId')
	echo -e "FOUND role id, deleted and recreated with new trust policy. Created id: ${awsRoleID}\n"
fi
echo "--------------------------------AWS Role creation ends----------------------------"

awsRoleArn=$(aws iam get-role --role-name "${roleName}" | jq -r '.Role.Arn')

echo "--------------------------------Creating S3 Bucket----------------------------"
if aws s3 ls "s3://${bucketName}" 2>/dev/null; then
	aws s3 rb "s3://${bucketName}" --force
fi
aws s3 mb "s3://${bucketName}" --region "${awsRegion}"
echo "Created S3 bucket: ${bucketName}"
echo "--------------------------------Attaching S3 policy to IAM role----------------------------"
bucketArn="arn:aws:s3:::${bucketName}"
jq --arg bucketArn "$bucketArn" \
	--arg bucketArnWildcard "${bucketArn}/*" \
	'.Statement[0].Resource[0] = $bucketArn | .Statement[0].Resource[1] = $bucketArnWildcard' \
	"$(dirname "$0")/s3-policy-template.json" >"$(dirname "$0")/s3-policy.json"

aws iam put-role-policy \
	--role-name "${roleName}" \
	--policy-name "${policyName}" \
	--policy-document "file://$(dirname "$0")/s3-policy.json"
echo "--------------------------------attach mongodb Role to AWS Role ends----------------------------"

# shellcheck disable=SC2086
sleep 30

atlas cloudProviders accessRoles aws authorize "${roleID}" \
	--projectId "${projectId}" \
	--iamAssumedRoleArn "${awsRoleArn}"
echo "--------------------------------authorize mongodb Role ends----------------------------"
rm -rf inputs
mkdir inputs

# Store AWS role ARN in a separate metadata file for cleanup (not part of CFN schema)
cat > "$(dirname "$0")/test-metadata.json" <<EOF
{
  "awsRoleArn": "${awsRoleArn}",
  "roleName": "${roleName}",
  "policyName": "${policyName}"
}
EOF

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg bucketName "$bucketName" \
		--arg iamRoleId "$roleID" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ProjectId?|=$projectId | .BucketName?|=$bucketName | .IamRoleId?|=$iamRoleId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
