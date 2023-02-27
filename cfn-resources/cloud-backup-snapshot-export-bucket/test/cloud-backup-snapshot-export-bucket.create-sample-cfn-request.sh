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

groupId="${1}"
iamRoleID="${2}"
bucketName="${3}"

jq --arg profile "$ATLAS_PROFILE" \
   --arg groupId "$groupId" \
   --arg iamRoleID "$iamRoleID" \
   --arg bucketName "$bucketName" \
   '.desiredResourceState.Profile?|=$profile | .desiredResourceState.GroupId?|=$groupId | .desiredResourceState.IamRoleID?|=$iamRoleID | .desiredResourceState.BucketName?|=$bucketName ' \
   "$(dirname "$0")/cloud-backup-snapshot-export-bucket.sample-cfn-request.json"
