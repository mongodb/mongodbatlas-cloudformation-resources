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
	echo "usage: cfn-test-create-inputs.sh <projectId> <endpoint>"
	echo "Creates a new Search Index"
}

if [ "$#" -ne 1 ]; then usage; fi
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

endpoint="${2:-$AWS_VPC_ENDPOINT}"
WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg proj "$projectId" \
		--arg endpoint_id "$endpoint" \
		'.ProjectId?|=$proj |.EndpointId?|=$endpoint_id' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
