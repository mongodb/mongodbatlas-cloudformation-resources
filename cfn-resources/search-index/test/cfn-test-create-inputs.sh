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

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

ClusterName="${projectName}"
clusterId=$(atlas clusters list --projectId "${projectId}" --output json | jq --arg NAME "${ClusterName}" -r '.results[]? | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
	atlas clusters create "${ClusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
	atlas clusters watch "${ClusterName}" --projectId "${projectId}"
	echo -e "Created Cluster \"${ClusterName}\""

	atlas clusters loadSampleData "${ClusterName}" --projectId "${projectId}"
fi

cluster_name=${ClusterName}
db_name="${4:-sample_airbnb}"
coll_name="${5:-listingsAndReviews}"
index_name="search-$(date +%s)-$RANDOM"
u_index_name="${index_name}"
WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	index_name="${u_index_name}"
	jq --arg org "$projectId" \
		--arg cluster "$cluster_name" \
		--arg name "$index_name" \
		--arg db "$db_name" \
		--arg coll "$coll_name" \
		'.CollectionName?|=$coll |.Database?|=$db |.ProjectId?|=$org |.ClusterName?|=$cluster |.Name?|=$name' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
