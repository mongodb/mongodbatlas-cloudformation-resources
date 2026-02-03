#!/usr/bin/env bash

# Run this script with the Makefile
# make create-test-resources
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

projectName="cfn-test-bot-$(date +%s)-$RANDOM"

# create project
projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

echo "projectId: $projectId"
echo "projectName: $projectName"

# Get the current user email for authorized email
# Use default test email for contract testing
authorizedEmail="test@example.com"

./test/cfn-test-create-inputs.sh "$projectName" "$authorizedEmail"
