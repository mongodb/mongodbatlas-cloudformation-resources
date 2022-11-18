#!/usr/bin/env bash
# cloud-backup-restore-job.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 3 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
projectId="${1}"
clusterName="${2}"
snapshotId="${3}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg projectId "$projectId" \
   --arg clusterName "$clusterName" \
   --arg snapshotId "$snapshotId" \
   ' .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey |   .desiredResourceState.SnapshotId?|=$snapshotId |   .desiredResourceState.ClusterName?|=$clusterName |  .desiredResourceState.ProjectId?|=$projectId ' \
   "$(dirname "$0")/cloud-backup-restore-job.sample-cfn-request.json"
