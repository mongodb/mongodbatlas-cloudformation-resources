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
jq --arg region "us-east-1" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.desiredResourceState.Name?|=$clusterName | .desiredResourceState.ProviderSettings.RegionName?|=$region | .desiredResourceState.ProjectId?|=$projectId ' \
   "$(dirname "$0")/cluster.sample-cfn-request.json"
