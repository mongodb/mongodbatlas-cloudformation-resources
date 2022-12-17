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

region="${1}"
vpcId="${2}"
subnetId="${3}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg region "$region" \
   --arg vpcId "$vpcId" \
   --arg subnetId "$subnetId" \
   '.desiredResourceState.GroupId?|=$org | .desiredResourceState.VpcId?|=$vpcId | .desiredResourceState.SubnetId?|=$subnetId | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.Region?|=$region' \
   "$(dirname "$0")/private-endpoint.sample-cfn-request.json"
