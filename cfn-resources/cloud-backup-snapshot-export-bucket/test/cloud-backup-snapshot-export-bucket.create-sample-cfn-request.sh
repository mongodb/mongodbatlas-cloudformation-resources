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

projectId="${1}"
iamRoleID="${2}"
bucketName="${3}"

jq --arg projectId "$projectId" \
   --arg iamRoleID "$iamRoleID" \
   --arg bucketName "$bucketName" \
   '.desiredResourceState.ProjectId?|=$projectId | .desiredResourceState.IamRoleID?|=$iamRoleID | .desiredResourceState.BucketName?|=$bucketName ' \
   "$(dirname "$0")/cloud-backup-snapshot-export-bucket.sample-cfn-request.json"
