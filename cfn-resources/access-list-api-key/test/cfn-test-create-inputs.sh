#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail
set -x

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
Profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
echo "profile set to ${MONGODB_ATLAS_PROFILE}"
Profile=${MONGODB_ATLAS_PROFILE}
fi
# Check MONGODB_ATLAS_ORG_ID is set
if [ -z "${MONGODB_ATLAS_ORG_ID+x}" ];then
  echo "MONGODB_ATLAS_ORG_ID must be set"
  exit 1
fi

OrgId="${MONGODB_ATLAS_ORG_ID}"

IpAddress="203.0.113.11"
# Create an organization API key with organization owner access in the organization with the ID 5a1b39eec902201990f12345:
APIUserId=$(atlas organizations apiKeys create --role ORG_READ_ONLY --desc "cfn bot access-list testing" --orgId "${OrgId}" --output json | jq -r '.id')


jq  --arg OrgId "$OrgId" \
  	--arg IpAddress "$IpAddress" \
  	--arg APIUserId "$APIUserId" \
  	--arg Profile "$Profile" \
  	'.OrgId?|=$OrgId | .IpAddress?|=$IpAddress | .APIUserId?|=$APIUserId | .Profile?|=$Profile' \
	"$(dirname "$0")/input_1_create.json" >"inputs/inputs_1_create.json"

ls -l inputs
