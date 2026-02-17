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
	echo "Generates test input files for mongodb employee access grant"
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
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

# Create a cluster for the employee access grant test
clusterName="cfn-test-cluster-$(date +%s)-$RANDOM"
echo "Creating cluster: $clusterName (this may take several minutes...)"

# Create a basic cluster (M10 tier) and wait for it to be ready
atlas clusters create "$clusterName" \
	--projectId "$projectId" \
	--provider AWS \
	--region US_EAST_1 \
	--tier M10 \
	--mdbVersion 7.0 \
	--diskSizeGB 10 \
	--output json

# Wait for cluster using atlas watch command
atlas clusters watch "$clusterName" --projectId "$projectId"

echo "Cluster is ready!"
echo "Using cluster: $clusterName in project: $projectId"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg clusterName "$clusterName" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ProjectId?|=$projectId | .ClusterName?|=$clusterName' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
