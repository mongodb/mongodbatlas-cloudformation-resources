#!/usr/bin/env bash
# cloud-backup-schedule.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#
set -o errexit
set -o nounset
set -o pipefail
  jq --arg group_id "$projectId" \
     --arg ClusterName "$ClusterName" \
     '.desiredResourceState.ProjectId?|=$group_id | .desiredResourceState.ClusterName?|=$ClusterName' \
    "$(dirname "$0")/cloud-backup-schedule.sample-cfn-request.json"