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

# Extract role info from metadata file (not from CFN input since AwsRoleArn is not part of the schema)
metadataFile="$(dirname "$0")/test-metadata.json"
if [ -f "$metadataFile" ]; then
	roleArn=$(jq -r '.awsRoleArn // empty' "$metadataFile")
	roleName=$(jq -r '.roleName // empty' "$metadataFile")
	policyName=$(jq -r '.policyName // empty' "$metadataFile")
	echo "Found test metadata file with role: ${roleName}"
else
	echo "Warning: test-metadata.json not found, skipping IAM role cleanup"
	roleArn=""
	roleName=""
	policyName=""
fi

# Deauthorize role from Atlas if trust policy exists
if [ -f "$(dirname "$0")/trust-policy.json" ] && [ -n "$roleArn" ]; then
	trustPolicy=$(jq '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/trust-policy.json")
	# shellcheck disable=SC2001
	atlasAssumedRoleExternalID=$(echo "${trustPolicy}" | sed 's/"//g')

	# Try to get roleId from Atlas (non-fatal if it fails)
	if atlas cloudProviders accessRoles list --projectId "${projectId}" --output json > /tmp/atlas_roles.json 2>&1; then
		roleId=$(jq --arg roleID "${atlasAssumedRoleExternalID}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($roleID)) | .roleId' /tmp/atlas_roles.json 2>/dev/null || echo "")
		rm -f /tmp/atlas_roles.json

		if [ -n "${roleId}" ] && [ "${roleId}" != "null" ] && [ "${roleId}" != "" ]; then
			echo "Deauthorizing role from Atlas: ${roleId}"
			if atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force; then
				echo "Successfully deauthorized role"
			else
				echo "Failed to deauthorize role (may already be deauthorized)"
			fi
			echo "--------------------------------deauthorize role ends----------------------------"
		else
			echo "Warning: Could not find Atlas role ID to deauthorize (may already be deauthorized)"
		fi
	else
		echo "Warning: Could not list Atlas roles (may be authentication issue or project already deleted)"
		rm -f /tmp/atlas_roles.json
	fi
else
	if [ -z "$roleArn" ]; then
		echo "Warning: No role ARN found in metadata, skipping Atlas role deauthorization"
	else
		echo "Warning: trust-policy.json not found, skipping Atlas role deauthorization"
	fi
fi
bucketName=$(jq -r '.BucketName' "./inputs/inputs_1_create.json")
if [ -n "$bucketName" ] && [ "$bucketName" != "null" ]; then
	echo "Deleting S3 bucket: ${bucketName}"
	aws s3 rb "s3://${bucketName}" --force || echo "Failed to delete S3 bucket (may already be deleted)"
else
	echo "Warning: No bucket name found in inputs"
fi

echo "--------------------------------delete IAM role starts----------------------------"
if [ -n "$roleName" ]; then
	echo "Deleting IAM role: ${roleName}"
	aws iam delete-role-policy --role-name "$roleName" --policy-name "$policyName" 2>/dev/null || echo "Role policy already deleted or doesn't exist"
	aws iam delete-role --role-name "$roleName" 2>/dev/null || echo "Role already deleted or doesn't exist"
	echo "Deleted IAM role: ${roleName}"
else
	echo "No IAM role to delete (not found in metadata)"
fi
echo "--------------------------------delete IAM role ends----------------------------"

# Clean up temporary test files
rm -f "$(dirname "$0")/trust-policy.json"
rm -f "$(dirname "$0")/s3-policy.json"
rm -f "$(dirname "$0")/test-metadata.json"
echo "Cleaned up temporary test files"

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	echo "Warning: Failed cleaning project:$projectId (may be authentication issue or already deleted)"
fi
