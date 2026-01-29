#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

projectName="cfn-test-bot-$(date +%s)-$RANDOM"
bucketName="atlas-logs-cfn-test-$RANDOM"
iamRoleId="65a1b2c3d4e5f6a7b8c9d0e1"

# create project
projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

echo "projectId: $projectId"
echo "projectName: $projectName"
echo "bucketName: $bucketName"
echo "iamRoleId: $iamRoleId (dummy 24-char hex format)"

./test/cfn-test-create-inputs.sh "$projectName" "$bucketName" "$iamRoleId"
