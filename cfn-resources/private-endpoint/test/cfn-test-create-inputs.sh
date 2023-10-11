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
	echo "Creates a new private endpoint role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectName="${1:-$PROJECT_NAME}"
vpcId="${2:-$AWS_VPC_ID}"
subnetId="${3:-$AWS_SUBNET_ID}"

if ! test -v AWS_DEFAULT_REGION; then
	region=$(aws configure get region)
else
	region=$AWS_DEFAULT_REGION
fi

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"

jq --arg groupId "$projectId" \
	--arg region "$region" \
	--arg vpcId "$vpcId" \
	--arg subnetId "$subnetId" \
	'.GroupId?|=$groupId | .Region?|=$region | .PrivateEndpoints[0].VpcId?|=$vpcId | .PrivateEndpoints[0].SubnetIds[0]?|=$subnetId' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

echo "mongocli iam projects delete ${projectId} --force"
