#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

awsSecretName=$(jq -r '.desiredResourceState.AwsSecretName' ./apikey.sample-cfn-request.json)
if aws secretsmanager delete-secret --secret-id "${awsSecretName}" --force-delete-without-recovery;then
  echo "aws secret deleted with name : ${awsSecretName}"
else
  echo "aws secret delete failed with name : ${awsSecretName}"
  exit 1
fi

projectIds=()
projectIds["${#projectIds[@]}"]=$(jq -r '.desiredResourceState.ProjectAssignments[0].ProjectId' ./apikey.sample-cfn-request.json)
projectIds["${#projectIds[@]}"]=$(jq -r '.desiredResourceState.ProjectAssignments[1].ProjectId' ./apikey.sample-cfn-request.json)

echo "${projectIds[@]}"
for projectId in "${projectIds[@]}"
do
  #delete project
  if atlas projects delete "$projectId" --force; then
    echo "$projectId project deletion OK"
  else
    (echo "Failed cleaning project:$projectId" && exit 1)
  fi
done



jq --arg orgId "" \
  --arg profile "" \
  --arg projectId1 "" \
  --arg projectId2 "" \
	'.desiredResourceState.OrgId?|=$orgId |
	 .desiredResourceState.Profile?|=$profile |
	  .desiredResourceState.ProjectAssignments[0].ProjectId?|=$projectId1 |
	   .desiredResourceState.ProjectAssignments[1].ProjectId?|=$projectId2' \
	"$(dirname "$0")/apikey.sample-cfn-request.json" > sample.temp && mv sample.temp apikey.sample-cfn-request.json

