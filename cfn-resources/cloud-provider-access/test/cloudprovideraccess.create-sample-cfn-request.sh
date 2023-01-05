#!/usr/bin/env bash
# cloudprovideraccess.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

export ATLAS_USER=htcxduya
export ATLAS_USER_KEY=e2fba4fd-aeaf-4cb6-9f25-f3a3d876f29f
export ATLAS_PUBLIC_KEY=htcxduya
export ATLAS_PRIVATE_KEY=e2fba4fd-aeaf-4cb6-9f25-f3a3d876f29f
export ATLAS_ORG_ID=5fe4ea50d1a2b617175ee3d4
export ATLAS_PROJECT_ID=625454459c4e6108393d650d
export ATLAS_REGION_NAME=us-east-1

function usage {
    echo "usage:$0 <project_id> <cluster_name>"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId=625454459c4e6108393d650d
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   '.desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.ProjectId?|=$projectId ' \
   "$(dirname "$0")/cloudprovideraccess.sample-cfn-request.json" > "test.request.json"
