#!/usr/bin/env bash
# cloud-backup-snapshot-export-job.create-sample-cfn-request.sh
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
projectId="${1}"
ClusterName="${2}"
snapshotId="${3}"
ExportBucketId="${4}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg group_id "$projectId" \
   --arg SnapshotId "$SnapshotId" \
    --arg ExportBucketId "$ExportBucketId" \
   '.desiredResourceState.ExportBucketId?|=$ExportBucketId |.desiredResourceState.SnapshotId?|=$SnapshotId | .desiredResourceState.GroupId?|=$group_id | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.ClusterName?|=$ClusterName' \
  "$(dirname "$0")/cloud-backup-restore-job.sample-cfn-request.json"
