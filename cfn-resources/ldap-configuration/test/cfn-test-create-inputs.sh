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

jq --arg project_id "$projectId" \
	--arg bindPassword "$bindPassword" \
	--arg bindUsername "$bindUsername" \
	--arg hostname "$hostname" \
	'.ProjectId?|=$project_id | .BindPassword?|=$bindPassword | .BindUsername?|=$bindUsername | .Hostname?|=$hostname' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg project_id "$projectId" \
	--arg bindPassword "$bindPassword" \
	--arg bindUsername "$bindUsername" \
	--arg hostname "$hostname" \
	'.ProjectId?|=$project_id | .BindPassword?|=$bindPassword | .BindUsername?|=$bindUsername | .Hostname?|=$hostname' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
