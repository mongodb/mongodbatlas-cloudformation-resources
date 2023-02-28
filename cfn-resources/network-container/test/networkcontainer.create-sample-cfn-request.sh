#!/usr/bin/env bash
# networkcontainer.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

jq --arg ProjectId "$MCLI_PROJECT_ID" \
	'.desiredResourceState.ProjectId?|=$ProjectId' \
	"$(dirname "$0")/networkcontainer.sample-cfn-request.json"
