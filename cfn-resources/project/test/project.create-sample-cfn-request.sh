#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

name="${1}"
jq --arg profile "$MONGODB_ATLAS_PROFILE" \
   --arg org "$MONGODB_ATLAS_ORG_ID" \
   --arg name "$name" \
   '.desiredResourceState.properties.OrgId?|=$org
   | .desiredResourceState.properties.Profile?|=$profile
   | .desiredResourceState.properties.Name?|=$name' \
   "$(dirname "$0")/project.sample-cfn-request.json"
