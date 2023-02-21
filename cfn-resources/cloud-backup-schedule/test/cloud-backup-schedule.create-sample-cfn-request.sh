#!/usr/bin/env bash
# cloud-backup-schedule.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#
set -o errexit
set -o nounset
set -o pipefail


projectId="${1}"
clusterName="${2}"

  jq --arg projectId "$projectId" \
     --arg clusterName "$clusterName" \
     '.desiredResourceState.ProjectId?|=$projectId | .desiredResourceState.ClusterName?|=$clusterName' \
    "$(dirname "$0")/cloud-backup-schedule.sample-cfn-request.json"