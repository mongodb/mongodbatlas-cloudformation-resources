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

projectName="cfn-bot-cloud-user-test"
# create ProjectId
if [ ${#projectName} -gt 22 ];then
  projectName=${projectName:0:21}
fi

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

username=""
mobileNumber="1234567890"
password=""

jq --arg orgId "${orgId}" \
   --arg profile "${profile}" \
   --arg projectId "${projectId}" \
   --arg username "${username}" \
   --arg mobileNumber "${mobileNumber}" \
   --arg password  "${password}" \
	  '.desiredResourceState.EmailAddress?|=$username | .desiredResourceState.Username?|=$username |.desiredResourceState.MobileNumber?|=$mobileNumber | .desiredResourceState.Password?|=$password | .desiredResourceState.OrgId?|=$orgId | .desiredResourceState.Profile?|=$profile | .desiredResourceState.ProjectAssignments[0].ProjectId?|=$projectId' \
	  "$(dirname "$0")/cloud-user-cfn-invoke-request.json" > sample.temp && mv sample.temp cloud-user-cfn-invoke-request.json

echo " NOTE: Delete the projects once tested using apikey.delete-sample-cfn-request.sh."