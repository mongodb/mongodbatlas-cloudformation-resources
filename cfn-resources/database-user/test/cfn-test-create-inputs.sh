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
# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

projectId=$(resolve_project_id "${projectName}")
if [ -z "$projectId" ]; then
	echo "Failed to create or find project ${projectName}" >&2
	exit 1
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
