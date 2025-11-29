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
clusterName=$projectName
echo "Creating required inputs"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Cant find project \"${projectName}\"\n"
fi
export MCLI_PROJECT_ID=$projectId

clusterId=$(atlas clusters list --projectId "${projectId}" --output json | jq --arg NAME "${clusterName}" -r '.results[]? | select(.name==$NAME) | .id')
if [ -z "$clusterId" ]; then
	echo "creating cluster.."
	atlas clusters create "${clusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
	atlas clusters watch "${clusterName}" --projectId "${projectId}"
	echo -e "Created Cluster \"${clusterName}\""
fi

policyId=$(atlas backups schedule describe "${clusterName}" --projectId "${projectId}" | jq -r '.policies[0].id')
echo "policyId: ${policyId}"

jq --arg group_id "$projectId" \
	--arg cluster_name "$clusterName" \
	--arg policy_id "$policyId" \
	'.ClusterName?|=$cluster_name |.ProjectId?|=$group_id| .Policies[0].ID?|=$policy_id' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg group_id "$projectId" \
	--arg cluster_name "$clusterName" \
	--arg policy_id "$policyId" \
	'.ClusterName?|=$cluster_name |.ProjectId?|=$group_id| .Policies[0].ID?|=$policy_id' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
echo "mongocli iam projects delete ${projectId} --force"
