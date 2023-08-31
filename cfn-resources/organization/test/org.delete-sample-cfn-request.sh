#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

awsSecretName=$(jq -r '.desiredResourceState.AwsSecretName' ./org.sample-cfn-request.json)
if aws secretsmanager delete-secret --secret-id "${awsSecretName}" --force-delete-without-recovery;then
  echo "aws secret deleted with name : ${awsSecretName}"
else
  echo "aws secret delete failed with name : ${awsSecretName}"
  exit 1
fi

jq --arg empty "" \
	'.desiredResourceState.OrgOwnerId?|=$empty |
	 .desiredResourceState.Profile?|=$empty |
	 .desiredResourceState.Name?|= $empty |
	 .desiredResourceState.AwsSecretName?|=$empty' \
	"$(dirname "$0")/org.sample-cfn-request.json" > sample.temp && mv sample.temp org.sample-cfn-request.json

