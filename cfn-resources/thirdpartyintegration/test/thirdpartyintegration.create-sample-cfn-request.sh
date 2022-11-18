#!/usr/bin/env bash
# thirdpartyintegration.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


  jq --arg pubkey "$MCLI_PUBLIC_API_KEY" \
     --arg pvtkey "$MCLI_PRIVATE_API_KEY" \
     --arg ProjectId "$MCLI_PROJECT_ID" \
     '.desiredResourceState.ProjectId?|=$ProjectId | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey' \
    "$(dirname "$0")/thirdpartyintegration.sample-cfn-request.json"
