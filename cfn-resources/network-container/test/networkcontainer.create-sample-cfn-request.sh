#!/usr/bin/env bash
# networkcontainer.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <projectId> <Region-(Atlas format)> <AtlasCIDRBlock>"
    exit 1
}

if [ "$#" -ne 3 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"
region="${2}"
cidr="${3}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg region "$region" \
   --arg cidr "$cidr" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.RegionName?|=$region | .desiredResourceState.properties.AtlasCIDRBlock?|=$cidr | .desiredResourceState.properties.ProjectId?|=$projectId' \
  "$(dirname "$0")/networkcontainer.sample-cfn-request.json"
