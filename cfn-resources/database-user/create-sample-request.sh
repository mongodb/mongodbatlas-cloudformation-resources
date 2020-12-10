#!/usr/bin/env bash
# databaseuser.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <projectId> <databaseuser.username>"
}

#if [ "$#" -ne 2 ]; then usage
#if [[ "$*" == help ]]; then usage
projectId="${1}"
username="${2}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg username "$username" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.Username?|=$username | .desiredResourceState.properties.ProjectId?|=$projectId' \
   "databaseuser.json"
