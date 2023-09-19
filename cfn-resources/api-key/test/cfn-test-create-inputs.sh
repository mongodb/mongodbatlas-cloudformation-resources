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
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi
# Check MONGODB_ATLAS_ORG_ID is set
if [ -z "${MONGODB_ATLAS_ORG_ID+x}" ];then
  echo "MONGODB_ATLAS_ORG_ID must be set"
  exit 1
fi

orgId="${MONGODB_ATLAS_ORG_ID}"

# create ProjectId
projectName="${1}"
if [ ${#projectName} -gt 22 ];then
  projectName=${projectName:0:21}
fi

# create aws secret key
awsSecretName="mongodb/atlas/apikey/${projectName}"
if aws secretsmanager create-secret --name "${awsSecretName}" --secret-string "atlas api-keys goes here";then
  echo "aws secret created with name : ${awsSecretName}"
else
  echo "aws secret create failed with name : ${awsSecretName}"
  exit 1
fi



project1="${projectName}-1"
project2="${projectName}-2"
project3="${projectName}-3"

projectId1=$(atlas projects list --output json | jq --arg NAME "${project1}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId1" ]; then
	projectId1=$(atlas projects create "${project1}" --output=json | jq -r '.id')
	echo -e "Created project \"${project1}\" with id: ${projectId1}\n"
else
	echo -e "FOUND project \"${project1}\" with id: ${projectId1}\n"
fi

projectId2=$(atlas projects list --output json | jq --arg NAME "${project2}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId2" ]; then
	projectId2=$(atlas projects create "${project2}" --output=json | jq -r '.id')
	echo -e "Created project \"${project2}\" with id: ${projectId2}\n"
else
	echo -e "FOUND project \"${project2}\" with id: ${projectId2}\n"
fi

projectId3=$(atlas projects list --output json | jq --arg NAME "${project3}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId3" ]; then
	projectId3=$(atlas projects create "${project3}" --output=json | jq -r '.id')
	echo -e "Created project \"${project3}\" with id: ${projectId3}\n"
else
	echo -e "FOUND project \"${project3}\" with id: ${projectId3}\n"
fi

# Create assigns 2 projects
jq --arg orgId "$orgId" \
  --arg profile "$profile" \
  --arg awsSecretName "$awsSecretName" \
  --arg projectId1 "$projectId1" \
  --arg projectId2 "$projectId2" \
	'.OrgId?|=$orgId | .Profile?|=$profile | .AwsSecretName?|=$awsSecretName |
	 .ProjectAssignments[0].ProjectId?|=$projectId1 |
	 .ProjectAssignments[1].ProjectId?|=$projectId2' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"


# Update with un-assign 1, update 1 and assign 1 Project.
jq --arg orgId "$orgId" \
	--arg profile "$profile" \
	--arg projectId2 "$projectId2" \
	--arg projectId3 "$projectId3" \
  	'.OrgId?|=$orgId | .Profile?|=$profile |
  	 .ProjectAssignments[0].ProjectId?|=$projectId2 |
  	 .ProjectAssignments[1].ProjectId?|=$projectId3' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"


ls -l inputs
