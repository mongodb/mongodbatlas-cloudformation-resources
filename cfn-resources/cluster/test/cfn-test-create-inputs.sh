#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project/cluster_name>"
	echo "Creates a new project and cluster by that name for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

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

echo "Check if a project is created $projectId"

region="us-east-1"

clusterName="${projectName}"

jq --arg region "$region" \
	--arg clusterName "$clusterName" \
	--arg projectId "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .Name?|=$clusterName | .ProjectId?|=$projectId ' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg region "$region" \
	--arg clusterName "$clusterName" \
	--arg projectId "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .Name?|=$clusterName | .ProjectId?|=$projectId ' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

jq --arg region "$region" \
	--arg clusterName "$clusterName" \
	--arg projectId "$projectId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .Name?|=$clusterName | .ProjectId?|=$projectId ' \
	"$(dirname "$0")/inputs_1_invalid.template.json" >"inputs/inputs_1_invalid.json"
