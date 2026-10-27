#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o nounset
set -o pipefail
WORDTOREMOVE="template."
function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectName="${1}"
MONGODB_ATLAS_PROFILE=${MONGODB_ATLAS_PROFILE:-"default"}
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	# Atlas project lists are eventually consistent: when this script follows a project
	# creation (contract testing creates the project first), the lookup above can miss
	# the new project and the create below fails with HTTP 409. Re-resolve by name.
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id') ||
		projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id') ||
		projectId=""
	if [ -z "$projectId" ]; then
		echo "Failed to create or find project ${projectName}" >&2
		exit 1
	fi
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Check if a project is created $projectId"

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg ProjectId "$projectId" --arg Profile "${MONGODB_ATLAS_PROFILE}" \
		'.ProjectId?|=$ProjectId | .Profile?|=$Profile' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..

ls -l inputs
#mongocli iam projects delete "${projectId}" --force
