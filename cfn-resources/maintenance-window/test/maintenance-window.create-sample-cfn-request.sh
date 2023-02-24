#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

jq --arg projectID "$PROJECT_ID" \
	'.desiredResourceState.ProjectId?|=$projectID' \
	"$(dirname "$0")/maintenance-window.sample-cfn-request.json"
