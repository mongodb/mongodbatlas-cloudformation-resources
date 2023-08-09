#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail
set -x

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
region="us-east-1"

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Cant find project \"${projectName}\"\n"
fi
export MCLI_PROJECT_ID=$projectId
ClusterName=$projectName
atlas clusters create "${ClusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json
atlas clusters watch "${ClusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${ClusterName}\""

SnapshotId=$(atlas backup snapshots create "${ClusterName}" --desc "cfn unit test" --retention 3 --output=json | jq -r '.id')
sleep 300

jq --arg org "$ATLAS_ORG_ID" \
	--arg ClusterName "$ClusterName" \
	--arg group_id "$projectId" \
	--arg SnapshotId "$SnapshotId" \
	'.SnapshotId?|=$SnapshotId | .ProjectId?|=$group_id | .InstanceName?|=$ClusterName' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

ls -l inputs
