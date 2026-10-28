#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "Creates a new customdb role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

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

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_2_create.json" >"inputs/inputs_2_create.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_2_update.json" >"inputs/inputs_2_update.json"

ls -l inputs
echo "mongocli iam projects delete ${projectId} --force"
