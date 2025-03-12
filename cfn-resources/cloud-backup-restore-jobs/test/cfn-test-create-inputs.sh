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
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "New project created \"${projectName}\"\n"
fi

ClusterName=$projectName

clusterId=$(atlas clusters list --projectId "${projectId}" --output json | jq --arg NAME "${ClusterName}" -r '.results[]? | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
	echo "creating cluster.."
	atlas clusters create "${ClusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 8.0 --diskSizeGB 10 --output=json
	atlas clusters watch "${ClusterName}" --projectId "${projectId}"
	echo -e "Created Cluster \"${ClusterName}\""
fi

echo -e "Created Cluster \"${ClusterName}\""

SnapshotId=$(atlas backup snapshots create "${ClusterName}" --projectId "${projectId}" --desc "cfn unit test" --retention 3 --output=json | jq -r '.id')
sleep 300

jq --arg cluster_name "$ClusterName" \
	--arg group_id "$projectId" \
	--arg SnapshotId "$SnapshotId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .SnapshotId?|=$SnapshotId | .ProjectId?|=$group_id | .InstanceName?|=$cluster_name' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

ls -l inputs
