#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb and AWS resources used for `cfn test` as inputs.
# Run from resource root (e.g. make delete-test-resources).
#

set -euo pipefail

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
bucketName=$(jq -r '.BucketName' ./inputs/inputs_1_create.json)
roleId=$(jq -r '.IamRoleID' ./inputs/inputs_1_create.json)

echo "Deleting resources for projectId: ${projectId}"

# Get IAM role name from file saved by create script
scriptDir="$(dirname "$0")"
if [ -f "${scriptDir}/role-name.txt" ]; then
	roleName=$(cat "${scriptDir}/role-name.txt")
	echo "Using role name from role-name.txt: ${roleName}"
else
	echo "role-name.txt not found, skipping IAM role cleanup"
	roleName=""
fi

keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
	keyRegion=$(aws configure get region)
fi
# shellcheck disable=SC2001
keyRegion=$(echo "$keyRegion" | sed -e "s/-/_/g" | tr '[:lower:]' '[:upper:]')
policyName="atlas-cloud-backup-export-bucket-S3-role-policy-${keyRegion}"

# Deauthorize Atlas role
if [ -n "$roleId" ] && [ "$roleId" != "null" ]; then
	echo "Deauthorizing Atlas role: ${roleId}"
	atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force || echo "Failed to deauthorize role"
fi

# Delete AWS IAM role
if [ -n "$roleName" ] && [ "$roleName" != "" ]; then
	echo "--------------------------------delete AWS IAM role starts----------------------------"
	aws iam delete-role-policy --role-name "$roleName" --policy-name "$policyName" 2>/dev/null || true
	aws iam delete-role --role-name "$roleName" 2>/dev/null || true
	echo "--------------------------------delete AWS IAM role ends----------------------------"
fi

# Delete S3 bucket
if [ -n "$bucketName" ] && [ "$bucketName" != "null" ]; then
	echo "Deleting S3 bucket: ${bucketName}"
	aws s3 rb "s3://${bucketName}" --force 2>/dev/null || true
fi

# Delete project
if atlas projects delete "$projectId" --force; then
	echo "${projectId} project deletion OK"
else
	echo "Failed cleaning project: ${projectId}"
	exit 1
fi
