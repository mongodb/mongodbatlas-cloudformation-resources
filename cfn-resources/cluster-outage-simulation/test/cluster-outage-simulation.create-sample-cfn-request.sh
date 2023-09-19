#!/usr/bin/env bash
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$1 <project_name>"
    echo "usage:$2 <cluster_name>"
}

projectId="${1}"
clusterName="${2}"
jq --arg profile "$ATLAS_PROFILE" \
   --arg projectId "$projectId" \
   --arg clusterName "$clusterName" \
   '.desiredResourceState.Profile?|=$profile |   .desiredResourceState.ClusterName?|=$clusterName  | .desiredResourceState.ProjectId?|=$projectId ' \
   "$(dirname "$0")/cluster-outage-simulation.sample-cfn-request.json"
