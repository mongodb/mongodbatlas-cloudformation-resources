#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes test input files and cleans up AWS resources.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}

echo "--------------------------------delete S3 bucket and IAM role starts----------------------------"

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
echo "Check if a project is created $projectId"
export MCLI_PROJECT_ID=$projectId

region=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
	region=$(aws configure get region)
fi
# shellcheck disable=SC2001
region=$(echo "$region" | sed -e "s/-/_/g")
region=$(echo "$region" | tr '[:lower:]' '[:upper:]')

roleName="mongodb-atlas-logs-role-${region}"
policyName="atlas-logs-s3-policy-${region}"

trustPolicy=$(jq '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/trust-policy.json")
# shellcheck disable=SC2001
atlasAssumedRoleExternalID=$(echo "${trustPolicy}" | sed 's/"//g')

roleId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${atlasAssumedRoleExternalID}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($roleID)) | .roleId')

atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force
echo "--------------------------------deauthorize role ends----------------------------"
bucketName=$(jq -r '.BucketName' "./inputs/inputs_1_create.json")
aws s3 rb "s3://${bucketName}" --force

echo "--------------------------------delete IAM role starts----------------------------"
aws iam delete-role-policy --role-name "$roleName" --policy-name "$policyName"
aws iam delete-role --role-name "$roleName"
echo "--------------------------------delete IAM role ends----------------------------"

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
