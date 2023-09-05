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
orgName="cfn-bot-org-test-$((1 + RANDOM % 1000))"

if [ -z "${CLOUD_DEV_ORG_OWNER_ID+x}" ];then
  export MONGODB_ATLAS_ORG_OWNER_ID=${CLOUD_DEV_ORG_OWNER_ID}
  echo "MONGODB_ATLAS_ORG_OWNER_ID is not set, exiting..."
  exit 1
fi

if [ -z ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "MONGODB_ATLAS_PROFILE is not set, exiting..."
    exit 1
fi

./test/cfn-test-create-inputs.sh $orgName
