#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#project_id
projectName="${1}"
bindPassword="$LDAP_BIND_PASSWORD"
bindUsername="$LDAP_BIND_USER_NAME"
hostname="$LDAP_HOST_NAME"

# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

projectId=$(resolve_project_id "${projectName}")
if [ -z "$projectId" ]; then
	echo "Failed to create or find project ${projectName}" >&2
	exit 1
fi
echo -e "Using project \"${projectName}\" with id: ${projectId}\n"
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

#An active cluster is needed to test the LDAP verify
ClusterName="${projectName}"
clusterId=$(atlas clusters list --projectId "${projectId}" --output json | jq --arg NAME "${ClusterName}" -r '.results[]? | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
	echo "creating cluster.."
	clusterId=$(atlas clusters create "${ClusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json | jq -r '.id')
fi

status=$(atlas clusters describe "${ClusterName}" --projectId "${projectId}" --output=json | jq -r '.stateName')
echo "status: ${status}"

while [[ "${status}" != "IDLE" ]]; do
	sleep 30
	status=$(atlas clusters describe "${ClusterName}" --projectId "${projectId}" --output=json | jq -r '.stateName')
	if [ -z "$status" ]; then
		status="timeout"
	fi
	echo "status: ${status}"
done

jq --arg projectId "$projectId" \
	--arg bindPassword "$bindPassword" \
	--arg bindUsername "$bindUsername" \
	--arg hostname "$hostname" \
	'.ProjectId?|=$projectId | .BindPassword?|=$bindPassword | .BindUsername?|=$bindUsername | .HostName?|=$hostname' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

ls -l inputs
