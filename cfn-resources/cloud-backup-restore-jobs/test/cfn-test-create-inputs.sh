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

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi

region="us-east-1"

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "New project created \"${projectName}\"\n"
fi

ClusterName=$projectName
atlas clusters create "${ClusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json
atlas clusters watch "${ClusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${ClusterName}\""

SnapshotId=$(atlas backup snapshots create "${ClusterName}" --projectId "${projectId}" --desc "cfn unit test" --retention 3 --output=json | jq -r '.id')
sleep 300

jq --arg org "$ATLAS_ORG_ID" \
	--arg ClusterName "$ClusterName" \
	--arg group_id "$projectId" \
	--arg SnapshotId "$SnapshotId" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .SnapshotId?|=$SnapshotId | .ProjectId?|=$group_id | .ClusterName?|=$ClusterName' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg org "$ATLAS_ORG_ID" \
	--arg region "${region}- more B@d chars !@(!(@====*** ;;::" \
	--arg group_id "$projectId" \
	--arg ClusterName "$ClusterName" \
	--arg SnapshotId "$SnapshotId" \
  --arg profile "$profile" \
	'.Profile?|=$profile | .SnapshotId?|=$SnapshotId |.ProjectId?|=$group_id | .ClusterName?|=$ClusterName' \
	"$(dirname "$0")/inputs_1_invalid.template.json" >"inputs/inputs_1_invalid.json"

ls -l inputs
