#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

projectName="cfn-test-bot-$(date +%s)-$RANDOM"

# create project
projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

echo "projectId: $projectId"
echo "projectName: $projectName"

./test/cfn-test-create-inputs.sh "$projectName"
