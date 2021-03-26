#!/usr/bin/env bash
# networkcontainer.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <projectId>"
    exit 1
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.ProjectId?|=$projectId' \
  "$(dirname "$0")/networkcontainer.sample-cfn-request.json"
