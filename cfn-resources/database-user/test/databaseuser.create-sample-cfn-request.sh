#!/usr/bin/env bash
# databaseuser.create-sample-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <projectId> <databaseuser.username> <databaseuser.password>"
	exit 1
}

if [ "$#" -ne 3 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId="${1}"
username="${2}"
password="${3}"

jq --arg ProjectId "$projectId" \
	--arg username "$username" \
	--arg password "$password" \
	'.desiredResourceState.properties.Username?|=$username | .desiredResourceState.properties.Password?|=$password | .desiredResourceState.properties.ProjectId?|=$projectId' \
	"$(dirname "$0")/databaseuser.sample-cfn-request.json"
