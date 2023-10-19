#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

if ! test -v AWS_DEFAULT_REGION; then
    region=$(aws configure get region)
else
  region=$AWS_DEFAULT_REGION
fi

jq --arg projId "$PROJECT_ID" \
	--arg region "$region" \
	'.desiredResourceState.ProjectId?|=$projId | .desiredResourceState.Region?|=$region' \
	"$(dirname "$0")/private-endpoint.sample-cfn-request.json"
