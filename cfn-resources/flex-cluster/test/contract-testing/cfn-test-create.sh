#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

projectName="cfn-test-bot-$((1 + RANDOM % 10000))"
clusterName="cfn-test-bot-$((1 + RANDOM % 10000))"

# create project
projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

echo "projectId: $projectId"
echo "projectName: $projectName"
echo "clusterName: $clusterName"

./test/cfn-test-create-inputs.sh $projectId $clusterName
