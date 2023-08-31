#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail
set -x


# setting projectName
orgName="cfn-bot-org-test-$((1 + RANDOM % 10000))"

./test/cfn-test-create-inputs.sh $orgName
