#!/usr/bin/env bash
# cloudprovideraccessauthorization.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project id> <role id> <iam assumed role arn>"
	echo "This test authorizes a Cloud Provider Access role, you can edit the test/ inputs if needed."
}

if [ "$#" -ne 3 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectID="${1}"
roleID="${2}"
iamAssumedRoleArn="${3}"

jq --arg projectID "$projectID" \
	--arg roleID "$roleID" \
	--arg arn "$iamAssumedRoleArn" \
	'.desiredResourceState.ProjectId?|=$projectID | .desiredResourceState.RoleId?|=$roleID | .desiredResourceState.IamAssumedRoleArn?|=$arn' \
	"$(dirname "$0")/cloudprovideraccessauthorization.sample-cfn-request.json"
