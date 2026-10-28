#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
echo "$projectName"
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

# shellcheck disable=SC2086
atlas clusters create "${ClusterName}" --projectId ${projectId} --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
atlas clusters watch "${ClusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${ClusterName}\""

atlas clusters loadSampleData "${ClusterName}" --projectId "${projectId}"

clusterName=${ClusterName}
collName="${2:-listingsAndReviews}"
dbName="${3:-sample_airbnb}"

WORDTOREMOVE="template."

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg cluster_name "$clusterName" \
		--arg coll_name "$collName" \
		--arg db_name "$dbName" \
		--arg project_id "$projectId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ClusterName?|=$cluster_name
		| .ProjectId?|=$project_id
		| .DbName?|=$db_name | .CollName?|=$coll_name' \
		"$inputFile" >"../inputs/$outputFile"
done

cd ..

ls -l inputs
