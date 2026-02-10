#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
# Cleanup test resources

set -euo pipefail

function usage {
    echo "usage:$0 "
}

# Read metadata from JSON file (similar to stream-connection pattern)
if [ ! -f "./inputs/test-metadata.json" ]; then
    echo "Error: test-metadata.json not found. Run create-test-resources first."
    exit 1
fi

projectId=$(jq -r '.ProjectId' ./inputs/test-metadata.json)
projectName=$(jq -r '.ProjectName' ./inputs/test-metadata.json)
iamRoleName=$(jq -r '.IamRoleName' ./inputs/test-metadata.json)

echo "==> Reading test metadata"
echo "Project: ${projectName} (${projectId})"
echo "IAM Role: ${iamRoleName}"

# Deauthorize Atlas access role using the roleId from metadata
echo "==> Deauthorizing Atlas access role"
roleId=$(jq -r '.AtlasRoleId' ./inputs/test-metadata.json)

if [ -n "$roleId" ] && [ "$roleId" != "null" ]; then
    echo "Deauthorizing roleId: ${roleId}"
    atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force || echo "Failed to deauthorize role"
else
    echo "No Atlas role found in metadata"
fi

echo "==> Deleting AWS IAM role"
if [ -n "${iamRoleName:-}" ] && [ "${iamRoleName}" != "null" ] && [ "${iamRoleName}" != "" ]; then
    echo "Deleting IAM role: ${iamRoleName}"
    aws iam delete-role --role-name "${iamRoleName}" 2>/dev/null && echo "Deleted AWS role: ${iamRoleName}" || echo "AWS role not found or already deleted"
else
    echo "No IAM role name found, skipping deletion"
fi

echo "==> Deleting Atlas project"
if atlas projects delete "${projectId}" --force; then
    echo "$projectId project deletion OK"
else
    echo "Failed cleaning project: $projectId"
    exit 1
fi

echo "==> Cleaning up generated files"
rm -f "$(dirname "$0")/trust-policy.json"
rm -rf inputs

echo ""
echo "==> Cleanup completed!"