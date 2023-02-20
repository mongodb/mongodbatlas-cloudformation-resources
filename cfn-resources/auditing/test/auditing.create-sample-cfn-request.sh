#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


jq --arg groupID "$projectId" \
   '.desiredResourceState.GroupId?|=$groupID' \
   "$(dirname "$0")/auditing.sample-cfn-request.json"