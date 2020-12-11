#!/usr/bin/env bash
# cluster.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
    echo "usage:$0 <project_id> <cluster_name>"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"
clusterName="${2}"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg region "us-east-1" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey |  .desiredResourceState.properties.Name?|=$clusterName | .desiredResourceState.properties.ProviderSettings.RegionName?|=$region | .desiredResourceState.properties.ProjectId?|=$projectId ' \
   "$(dirname "$0")/cluster.sample-cfn-request.json"
