#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
	echo "Creates a new customdb role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
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
	"$(dirname "$0")/inputs_1_invalid.json" >"inputs/inputs_1_invalid.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_2_create.json" >"inputs/inputs_2_create.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_2_update.json" >"inputs/inputs_2_update.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	'.ProjectId?|=$projectId |.Profile?|=$profile ' \
	"$(dirname "$0")/inputs_2_invalid.json" >"inputs/inputs_2_invalid.json"

ls -l inputs
