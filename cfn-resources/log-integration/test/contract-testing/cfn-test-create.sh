#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

projectName="cfn-test-bot-$(date +%s)-$RANDOM"

# Set unique tag for S3 bucket to avoid conflicts in CI
export CFN_TEST_TAG="${projectName}"

echo "projectName: $projectName"

./test/cfn-test-create-inputs.sh "$projectName"
