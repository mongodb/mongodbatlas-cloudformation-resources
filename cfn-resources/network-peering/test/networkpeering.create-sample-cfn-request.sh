#!/usr/bin/env bash
# networkpeering.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <projectId> <containerid> <cidr> <vpc>"
    exit 1
}

if [ "$#" -ne 4 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"
containerId="${2}"
cidr="${3}"
vpc="${4}"
region=$(aws configure get region)
AWS_ACCOUNT_ID=$(aws sts get-caller-identity | jq -r '.Account')
echo "AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg region "$region" \
   --arg awsacctid "${AWS_ACCOUNT_ID}" \
   --arg containerid "$containerId" \
   --arg cidr "$cidr" \
   --arg vpc "$vpc" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.ProjectId?|=$projectId | .desiredResourceState.properties.AccepterRegionName?|=$region | .desiredResourceState.properties.AWSAccountId?|=$awsacctid| .desiredResourceState.properties.ContainerId?|=$containerid | .desiredResourceState.properties.RouteTableCidrBlock?|=$cidr| .desiredResourceState.properties.VPCId?|=$vpc' \
  "$(dirname "$0")/networkpeering.sample-cfn-request.json"
