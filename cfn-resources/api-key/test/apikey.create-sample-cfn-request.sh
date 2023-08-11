#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

profile="default"
orgId="${MONGODB_ATLAS_ORG_ID}"

projectName="cfn-bot-apikey-test"
# create ProjectId
if [ ${#projectName} -gt 22 ];then
  projectName=${projectName:0:21}
fi

project1="${projectName}-1"
project2="${projectName}-2"

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


## TEST-1
# Create assigns 2 projects
jq --arg orgId "$orgId" \
  --arg profile "$profile" \
  --arg projectId1 "$projectId1" \
  --arg projectId2 "$projectId2" \
	'.desiredResourceState.OrgId?|=$orgId | .desiredResourceState.Profile?|=$profile | .desiredResourceState.ProjectAssignments[0].ProjectId?|=$projectId1 | .desiredResourceState.ProjectAssignments[1].ProjectId?|=$projectId2' \
	"$(dirname "$0")/apikey.sample-cfn-request.json" > sample.temp && mv sample.temp apikey.sample-cfn-request.json


echo " NOTE: Delete the projects once tested using apikey.delete-sample-cfn-request.sh."