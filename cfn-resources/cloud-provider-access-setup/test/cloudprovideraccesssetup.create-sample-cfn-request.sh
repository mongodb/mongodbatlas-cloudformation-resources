#!/usr/bin/env bash
# cloudprovideraccesssetup.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_id>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"

jq --arg projectId "$projectId" \
	'.desiredResourceState.ProjectId?|=$projectId' \
	"$(dirname "$0")/cloudprovideraccesssetup.sample-cfn-request.json"
