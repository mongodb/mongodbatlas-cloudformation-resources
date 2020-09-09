#!/usr/bin/env bash

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
   "sample-request.databaseuser.json"
