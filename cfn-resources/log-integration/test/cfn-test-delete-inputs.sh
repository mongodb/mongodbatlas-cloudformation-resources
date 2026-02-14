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

region=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
	region=$(aws configure get region)
fi
# shellcheck disable=SC2001
region=$(echo "$region" | sed -e "s/-/_/g")
region=$(echo "$region" | tr '[:lower:]' '[:upper:]')

policyName="atlas-logs-s3-policy-${region}"

# Get role name from metadata file
if [ -f "$(dirname "$0")/role-name.txt" ]; then
    roleName=$(cat "$(dirname "$0")/role-name.txt")
    echo "Found role name from metadata: ${roleName}"
else
    # Fallback: list AWS roles and find the one matching our pattern
    roleName=$(aws iam list-roles --output json | jq -r ".Roles[] | select(.RoleName | startswith(\"mongodb-atlas-logs-role-${region}\")) | .RoleName" | head -1)
    if [ -z "$roleName" ]; then
        echo "Warning: Could not find AWS IAM role to delete"
        roleName="mongodb-atlas-logs-role-${region}"  # fallback to static
    fi
    echo "Discovered role name: ${roleName}"
fi

echo "Using role name: ${roleName}"

trustPolicy=$(jq '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/trust-policy.json")
# shellcheck disable=SC2001
atlasAssumedRoleExternalID=$(echo "${trustPolicy}" | sed 's/"//g')

roleId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${atlasAssumedRoleExternalID}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($roleID)) | .roleId')

atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force
echo "--------------------------------deauthorize role ends----------------------------"
bucketName=$(jq -r '.BucketName' "./inputs/inputs_1_create.json")
aws s3 rb "s3://${bucketName}" --force

echo "--------------------------------delete IAM role starts----------------------------"
aws iam delete-role-policy --role-name "$roleName" --policy-name "$policyName" 2>/dev/null || echo "Policy already deleted or doesn't exist"
aws iam delete-role --role-name "$roleName" 2>/dev/null || echo "Role already deleted or doesn't exist"

# Clean up metadata files
rm -f "$(dirname "$0")/role-name.txt"
rm -f "$(dirname "$0")/trust-policy.json"
rm -f "$(dirname "$0")/s3-policy.json"

echo "--------------------------------delete IAM role ends----------------------------"

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
