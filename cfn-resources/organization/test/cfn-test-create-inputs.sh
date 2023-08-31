#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
	echo "Creates a template for org apikey creation"
}

if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="dev-cloud-profile"
#if [ ${MONGODB_ATLAS_PROFILE+x} ];then
#    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
#    profile=${MONGODB_ATLAS_PROFILE}
#fi

orgOwnerId="${MONGODB_ATLAS_ORG_OWNER_ID}"

# create ProjectId
orgName="${1}"

# create aws secret key
awsSecretName="cfn/atlas/profile/${orgName}"
if aws secretsmanager create-secret --name "${awsSecretName}" --secret-string "atlas api-keys goes here";then
  echo "aws secret created with name : ${awsSecretName}"
else
  echo "aws secret create failed with name : ${awsSecretName}"
  exit 1
fi

# Create
jq --arg orgOwnerId "$orgOwnerId" \
     --arg profile "$profile" \
      --arg orgName "$orgName" \
      --arg awsSecretName "$awsSecretName" \
	'.OrgOwnerId?|=$orgOwnerId | .Profile?|=$profile
	| .Name?|=$orgName | .AwsSecretName?|=$awsSecretName ' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

orgName="${orgName}-update"

# Update
jq --arg orgOwnerId "$orgOwnerId" \
     --arg profile "$profile" \
      --arg orgName "$orgName" \
      --arg awsSecretName "$awsSecretName" \
	'.OrgOwnerId?|=$orgOwnerId | .Profile?|=$profile
	 |.Name?|=$orgName | .AwsSecretName?|=$awsSecretName ' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"


ls -l inputs
