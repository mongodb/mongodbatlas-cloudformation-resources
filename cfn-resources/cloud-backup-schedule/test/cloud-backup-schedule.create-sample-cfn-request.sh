#!/usr/bin/env bash
# cloud-backup-schedule.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#
set -o errexit
set -o nounset
set -o pipefail
  jq --arg pubkey "$MCLI_PUBLIC_API_KEY" \
     --arg pvtkey "$MCLI_PRIVATE_API_KEY" \
     --arg group_id "$PROJECT_ID" \
     --arg ClusterName "$CLUSTER_NAME" \
     '.ProjectId?|=$group_id | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ClusterName?|=$ClusterName' \
    "$(dirname "$0")/cloud-backup-schedule.sample-cfn-request.json"