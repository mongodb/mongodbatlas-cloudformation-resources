#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail
projectIds=()
projectIds["${#projectIds[@]}"]=$(jq -r '.desiredResourceState.ProjectAssignments[0].ProjectId' ./cloud-user-cfn-invoke-request.json)

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
	'.desiredResourceState.OrgId?|=$orgId | .desiredResourceState.Profile?|=$profile | .desiredResourceState.ProjectAssignments[0].ProjectId?|=$projectId1' \
	"$(dirname "$0")/cloud-user-cfn-invoke-request.json" > sample.temp && mv sample.temp cloud-user-cfn-invoke-request.json

