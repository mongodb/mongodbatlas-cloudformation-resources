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
	exit 1
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

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

ClusterName="${projectName}"
clusterId=$(atlas clusters list --projectId "${projectId}" --output json | jq --arg NAME "${ClusterName}" -r '.results[]? | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
	atlas clusters create "${ClusterName}" --projectId "${projectId}" --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
	atlas clusters watch "${ClusterName}" --projectId "${projectId}"
	echo -e "Created Cluster \"${ClusterName}\""

	sampleJobId=$(atlas clusters sampleData load "${ClusterName}" --projectId "${projectId}" --output=json | jq -r '._id')
	echo "Waiting for sample data load job ${sampleJobId} to complete..."
	atlas clusters sampleData watch "${sampleJobId}" --projectId "${projectId}"
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
		--arg profile "$profile" \
		'.Profile?|=$profile |.CollectionName?|=$coll |.Database?|=$db |.ProjectId?|=$org |.ClusterName?|=$cluster |.Name?|=$name' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
