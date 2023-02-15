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
projectName="${1}"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"

rm -rf inputs
mkdir inputs
name="${1}"
jq --arg group_id "$projectId" \
	'.ProjectId?|=$group_id' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

echo "mongocli iam projects delete ${projectId} --force"
ls -l inputs
