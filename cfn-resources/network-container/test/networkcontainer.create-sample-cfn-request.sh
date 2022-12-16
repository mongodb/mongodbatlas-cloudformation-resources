#!/usr/bin/env bash
# networkcontainer.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi


  jq --arg pubkey "$MCLI_PUBLIC_API_KEY" \
     --arg pvtkey "$MCLI_PRIVATE_API_KEY" \
     --arg ProjectId "$MCLI_PROJECT_ID" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.ProjectId?|=$ProjectId' \
  "$(dirname "$0")/networkcontainer.sample-cfn-request.json"
