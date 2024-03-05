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
	echo "Creates a new project and an Cluster for testing"
}

if [ "$#" -ne 2 ]; then usage; fi
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
clusterName="${projectName}"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "Found project \"${projectName}\" with id: ${projectId}\n"
fi

clusterId=$(atlas clusters list --projectId "${projectId}" --output json | jq --arg NAME "${clusterName}" -r '.results[]? | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
	atlas clusters create "${clusterName}" --projectId "${projectId}" --provider AWS --region US_EAST_1 --tier M10 --mdbVersion 7.0 --output=json
	atlas clusters watch "${clusterName}" --projectId "${projectId}"
	echo -e "Created Cluster \"${clusterName}\""
fi

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg clusterName "$clusterName" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ProjectId?|=$projectId |.ClusterName?|=$clusterName' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
