#!/usr/bin/env bash
# custom-dns-config-cluster-aws.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#
set -o errexit
set -o nounset
set -o pipefail


projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

jq --arg ProjectId "$projectId" \
     '.desiredResourceState.ProjectId?|=$ProjectId' \
    "$(dirname "$0")/custom-dns-config-cluster-aws.sample-cfn-request.json"