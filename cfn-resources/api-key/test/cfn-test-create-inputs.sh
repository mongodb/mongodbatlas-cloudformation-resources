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
	echo "Creates a template for org apikey creation"
}

if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi
# Check ATLAS_ORG_ID is set
if [ -z "${ATLAS_ORG_ID+x}" ];then
  echo "ATLAS_ORG_ID must be set"
  exit 1
fi

orgId="${ATLAS_ORG_ID}"


jq --arg orgId "$orgId" \
  --arg profile "$profile" \
	'.OrgId?|=$orgId | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

jq --arg orgId "$orgId" \
	--arg profile "$profile" \
  	'.OrgId?|=$orgId | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

jq --arg orgId " ( *&lkd" \
	--arg profile "$profile" \
  	'.OrgId?|=$orgId | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_1_invalid.json" >"inputs/inputs_1_invalid.json"


ls -l inputs
