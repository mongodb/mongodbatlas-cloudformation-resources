#!/usr/bin/env bash
# cloud-backup-snapshot-export-job.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
	exit
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
projectId="${1}"
jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg projectId "$projectId" \
	'.desiredResourceState.GroupId?|=$projectId ' \
	"$(dirname "$0")/alert-configuration.sample-cfn-request.json"
