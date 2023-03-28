#!/usr/bin/env bash
# private-endpoint-regional-mode.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#
set -o errexit
set -o nounset
set -o pipefail
jq --arg ProjectId "$ATLAS_PROJECT_ID" \
	'.desiredResourceState.ProjectId?|=$ProjectId ' \
	"$(dirname "$0")/private-endpoint-regional-mode.sample-cfn-request.json"
