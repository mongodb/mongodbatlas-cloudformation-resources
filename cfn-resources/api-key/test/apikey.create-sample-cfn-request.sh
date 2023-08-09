#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

profile="default"
orgId="${MONGODB_ATLAS_ORG_ID}"

jq --arg orgId "$orgId" \
    --arg profile "$profile" \
   '.desiredResourceState.OrgId?|=$orgId | .desiredResourceState.Profile?|=$profile' \
   "$(dirname "$0")/apikey.sample-cfn-request.json"