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

projId="${1}"
instanceName="${2}"
vpcId="${3}"
subnetId="${4}"
jq --arg projId "$projId" \
   --arg instanceName "$instanceName" \
   --arg vpcId "$vpcId" \
   --arg subnetId "$subnetId" \
   '.desiredResourceState.ProjectId?|=$projId | .desiredResourceState.InstanceName?|=$instanceName | .desiredResourceState.AwsPrivateEndpointConfigurationProperties.VpcId?|=$vpcId | .desiredResourceState.AwsPrivateEndpointConfigurationProperties.SubnetIds[0]?|=$subnetId' \
   "$(dirname "$0")/serverless-private-endpooint.sample-cfn-request.json"
