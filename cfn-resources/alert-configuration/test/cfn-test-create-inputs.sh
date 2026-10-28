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
	echo "Creates a new project"
	exit 1
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

#set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

rm -rf inputs
mkdir inputs
projectName="${1}"
# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

projectId=$(resolve_project_id "${projectName}")
if [ -z "$projectId" ]; then
	echo "Failed to create or find project ${projectName}" >&2
	exit 1
fi
echo -e "Using project \"${projectName}\" with id: ${projectId}\n"
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

rm -rf inputs
mkdir inputs

jq --arg group_id "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$group_id | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg group_id "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$group_id | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

jq --arg group_id "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$group_id | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_2_create.template.json" >"inputs/inputs_2_create.json"

jq --arg group_id "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$group_id | .Profile?|=$profile' \
	"$(dirname "$0")/inputs_2_update.template.json" >"inputs/inputs_2_update.json"

ls -l inputs
