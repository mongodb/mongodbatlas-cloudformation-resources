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

projectName="${1}"
MONGODB_ATLAS_PROFILE=${MONGODB_ATLAS_PROFILE:-"default"}
# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

projectId=$(resolve_project_id "${projectName}")
if [ -z "$projectId" ]; then
	echo "Failed to create or find project ${projectName}" >&2
	exit 1
fi

echo "Created project \"${projectName}\" with id: ${projectId}"

cd "$(dirname "$0")" || exit
WORDTOREMOVE="template."
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg ProjectId "$projectId" --arg Profile "${MONGODB_ATLAS_PROFILE}" \
		'.ProjectId?|=$ProjectId | .Profile?|=$Profile' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..