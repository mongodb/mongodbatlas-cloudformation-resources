#!/usr/bin/env bash
# datalakes.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail


function usage {
    echo "usage:$0 <project_name>"
    echo "usage:$0 <ExternalId>"
    echo "usage:$0 <IamAssumedRoleARN>"
    echo "usage:$0 <IamUserARN>"
    echo "usage:$0 <RoleId>"
    echo "usage:$0 <TestS3Bucket>"
    exit;
}

if [ "$#" -ne 6 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"
atlasAssumedRoleExternalId="${2}"
atlasAWSAccountArn="${3}"
awsArne="${4}"
roleID="${5}"
bucketName="${6}"

jq --arg org "$ATLAS_ORG_ID" \
   --arg projectId "$projectId" \
   --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
   --arg AWSAssumedArn "$awsArne" \
   --arg role "$roleID" \
   --arg bucketName "$bucketName" \
   '.desiredResourceState.TenantName?|=$bucketName |.desiredResourceState.CloudProviderConfig.Aws.TestS3Bucket?|=$bucketName |.desiredResourceState.CloudProviderConfig.Aws.RoleId?|=$role |.desiredResourceState.CloudProviderConfig.Aws.IamUserARN?|=$atlasAWSAccountArn |.desiredResourceState.CloudProviderConfig.Aws.ExternalId?|=$atlasAssumedRoleExternalId | .desiredResourceState.CloudProviderConfig.Aws.IamAssumedRoleARN?|=$AWSAssumedArn | .desiredResourceState.ProjectId?|=$projectId' \
   "$(dirname "$0")/datalakes.sample-cfn-request.json"

