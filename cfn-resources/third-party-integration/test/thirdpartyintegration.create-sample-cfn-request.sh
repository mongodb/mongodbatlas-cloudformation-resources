#!/usr/bin/env bash
# thirdpartyintegration.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

jq --arg ProjectId "$MCLI_PROJECT_ID" \
	'.desiredResourceState.ProjectId?|=$ProjectId' \
	"$(dirname "$0")/thirdpartyintegration.sample-cfn-request.json"
