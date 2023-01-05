#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg groupID "$projectId" \
   '.desiredResourceState.GroupId?|=$groupID | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/auditing.sample-cfn-request.json"