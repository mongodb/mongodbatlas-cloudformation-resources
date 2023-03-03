#!/usr/bin/env bash
# cloudprovideraccess.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
    echo "usage:$0 <project_id> <cluster_name>"
}

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

jq --arg projectId "$projectId" \
   '.desiredResourceState.ProjectId?|=$projectId ' \
   "$(dirname "$0")/cloudprovideraccess.sample-cfn-request.json" > "test.request.json"
