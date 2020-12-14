#!/usr/bin/env bash
# networkpeering.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <projectId> <cidr> <vpc>"
    exit 1
}

if [ "$#" -ne 3 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"
cidr="${2}"
vpc="${3}"
region=$(aws configure get region)
AWS_ACCOUNT_ID=$(aws sts get-caller-identity | jq -r '.Account')
#echo "AWS_ACCOUNT_ID=${AWS_ACCOUNT_ID}"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg region "$region" \
   --arg awsacctid "$AWS_ACCOUNT_ID" \
   --arg cidr "$cidr" \
   --arg vpc "$vpc" \
   '.desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.ProjectId?|=$projectId | .desiredResourceState.properties.AccepterRegionName?|=$region | .desiredResourceState.properties.AwsAccountId?|=$awsacctid| .desiredResourceState.properties.RouteTableCIDRBlock?|=$cidr| .desiredResourceState.properties.VpcId?|=$vpc' \
  "$(dirname "$0")/networkpeering.sample-cfn-request.json"
