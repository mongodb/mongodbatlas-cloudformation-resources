#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

projectName="cfn-test-stream-pl-endpoint-$(date +%s)-$RANDOM"

echo "projectName: $projectName"

./test/cfn-test-create-inputs.sh "$projectName"
