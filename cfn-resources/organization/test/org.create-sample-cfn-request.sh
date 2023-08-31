#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

profile="dev-cloud-profile"
orgOwnerId="${MONGODB_ATLAS_ORG_OWNER_ID}"

orgName="cfn-bot-org-test"

# create aws secret key
awsSecretName="mongodb/atlas/apikey/${orgName}"
if aws secretsmanager describe-secret --secret-id "${awsSecretName}";then
  echo "aws secret already exists with name : ${awsSecretName}"
elif aws secretsmanager create-secret --name "${awsSecretName}" --secret-string "atlas org api-keys goes here";then
  echo "aws secret created with name : ${awsSecretName}"
else
  echo "aws secret create failed with name : ${awsSecretName}"
  exit 1
fi
## TEST-1
jq --arg orgOwnerId "$orgOwnerId" \
  --arg profile "$profile" \
   --arg awsSecretName "$awsSecretName" \
   --arg orgName "$orgName" \
	'.desiredResourceState.OrgOwnerId?|=$orgOwnerId |
	.desiredResourceState.Profile?|=$profile |
	.desiredResourceState.Name?|=$orgName |
	.desiredResourceState.AwsSecretName?|=$awsSecretName ' \
	"$(dirname "$0")/org.sample-cfn-request.json" > sample.temp && mv sample.temp org.sample-cfn-request.json


echo " NOTE: Delete the projects once tested using org.delete-sample-cfn-request.sh."