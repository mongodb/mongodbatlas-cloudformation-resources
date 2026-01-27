#!/usr/bin/env bash

# This tool generates json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

serviceAccountName="cfn-test-service-account-$(date +%s)-$RANDOM"

echo "serviceAccountName: $serviceAccountName"

./test/cfn-test-create-inputs.sh "$serviceAccountName"
