#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


jq --arg profile "$ATLAS_PROFILE" \
   --arg groupID "$projectId" \
   '.desiredResourceState.ProjectId?|=$groupID | .desiredResourceState.Profile?|=$profile' \
   "$(dirname "$0")/custom-db-role.sample-cfn-request.json"
