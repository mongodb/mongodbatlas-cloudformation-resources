#!/usr/bin/env bash
# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
echo "Came inside create inputs to test"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Cant find project \"${projectName}\"\n"
fi

atlas clusters create "${clusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
atlas clusters watch "${clusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${clusterName}\""

rm -rf inputs
mkdir inputs
jq --arg group_id "$projectId" \
	--arg clusterName "$clusterName" \
	'.InstanceName?|=$clusterName |.ProjectId?|=$group_id' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

ls -l inputs
