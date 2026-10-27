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
	echo "Generates test input files for project service account"
	exit 0
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
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

serviceAccountName="cfn-test-sa-$(date +%s)-$RANDOM"
echo "serviceAccountName: $serviceAccountName"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg serviceAccountName "$serviceAccountName" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ProjectId?|=$projectId |.Name?|=$serviceAccountName' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
