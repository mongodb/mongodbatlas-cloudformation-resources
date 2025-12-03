#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

# setting orgName
orgName="cfn-bot-org-test-$(date +%s)"

if [ -z ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "MONGODB_ATLAS_PROFILE is not set, exiting..."
	exit 1
fi

./test/cfn-test-create-inputs.sh "$orgName"
