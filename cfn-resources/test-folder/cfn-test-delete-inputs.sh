#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
# Cleanup test resources

set -e
set -x

if [ "$#" -ne 1 ]; then
    echo "usage: $0 <project_name>"
    exit 1
fi

projectName="${1}"

echo "==> Getting project details"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    echo "Project not found: ${projectName}"
    exit 1
fi

export MCLI_PROJECT_ID=$projectId
echo "Found project: ${projectName} (${projectId})"

# Get AWS region
keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
    keyRegion=$(aws configure get region)
fi
keyRegion=${keyRegion//-/_}
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')

roleName="mongodb-atlas-enc-role-${keyRegion}"

# Get role details from trust policy file
if [ -f "$(dirname "$0")/trust-policy.json" ]; then
    echo "==> Reading trust policy"
    atlasAssumedRoleExternalId=$(jq -r '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/trust-policy.json")
    
    if [ -n "$atlasAssumedRoleExternalId" ] && [ "$atlasAssumedRoleExternalId" != "null" ]; then
        echo "External ID: ${atlasAssumedRoleExternalId}"
        
        echo "==> Deauthorizing Atlas access role"
        roleId=$(atlas cloudProviders accessRoles list --output json | jq --arg externalId "${atlasAssumedRoleExternalId}" -r '.awsIamRoles[] | select(.atlasAssumedRoleExternalId | test($externalId)) | .roleId' || echo "")
        
        if [ -n "$roleId" ]; then
            echo "Deauthorizing roleId: ${roleId}"
            atlas cloudProviders accessRoles aws deauthorize "${roleId}" --force || echo "Failed to deauthorize role"
        else
            echo "No Atlas role found to deauthorize"
        fi
    fi
fi

echo "==> Deleting AWS IAM role"
aws iam delete-role --role-name "$roleName" 2>/dev/null && echo "Deleted AWS role: ${roleName}" || echo "AWS role not found or already deleted"

echo "==> Deleting Atlas project"
atlas projects delete "${projectId}" --force && echo "Deleted project: ${projectName} (${projectId})" || echo "Failed to delete project"

echo "==> Cleaning up generated files"
rm -f "$(dirname "$0")/trust-policy.json"

echo ""
echo "==> Cleanup completed!"
