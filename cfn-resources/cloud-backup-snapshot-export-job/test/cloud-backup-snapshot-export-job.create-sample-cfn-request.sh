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

jq --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg project_Id "$projectId" \
   --arg SnapshotId "$snapshotId" \
    --arg ExportBucketId "$ExportBucketId" \
   '.desiredResourceState.ExportBucketId?|=$ExportBucketId |.desiredResourceState.SnapshotId?|=$SnapshotId | .desiredResourceState.ProjectId?|=$project_Id | .desiredResourceState.ClusterName?|=$ClusterName' \
  "$(dirname "$0")/cloud-backup-restore-job.sample-cfn-request.json"
