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

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

name="serverless-cfn-bot"

jq --arg org "MONGODB_ATLAS_ORG_ID" \
	--arg name "$name" \
	--arg project_id "$projectId" \
	'.ProjectID?|=$project_id | .Name?|=$name ' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg name "${name}- more B@d chars !@(!(@====*** ;;::" \
	--arg project_id "$projectId" \
	'.ProjectID?|=$project_id | .Name?|=$name' \
	"$(dirname "$0")/inputs_1_invalid.template.json" >"inputs/inputs_1_invalid.json"

jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg name "${name}" \
	--arg project_id "$projectId" \
	'.ProjectID?|=$project_id | .Name?|=$name ' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
