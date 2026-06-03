#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

# setting projectName
projectName="ct-api-key-$(date +%s | tail -c 5)${RANDOM}"

./test/cfn-test-create-inputs.sh "$projectName"
