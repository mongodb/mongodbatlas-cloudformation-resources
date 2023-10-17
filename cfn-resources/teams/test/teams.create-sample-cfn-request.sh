#!/usr/bin/env bash
# cloud-backup-restore-job.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

name="${1}"
userName="${2}"
jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg name "$name" \
	--arg userName "$userName" \
	'.desiredResourceState.Usernames?|=[$userName] |.desiredResourceState.OrgId?|=$org | .desiredResourceState.Name?|=$name' \
	"$(dirname "$0")/teams.sample-cfn-request.json"
