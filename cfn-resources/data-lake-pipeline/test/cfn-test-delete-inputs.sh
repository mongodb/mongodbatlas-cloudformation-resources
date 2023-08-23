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
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}

projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

clusterName=$(jq -r '.Source.ClusterName' ./inputs/inputs_1_create.json)

#delete Cluster
if atlas clusters delete "$clusterName" --projectId "${projectId}" --force; then
	echo "$clusterName cluster deletion OK"
else
	(echo "Failed cleaning cluster:$clusterName" && exit 1)
fi

status="DELETING"
echo "Waiting for cluster to get deleted"
while [[ "${status}" == "DELETING" ]]; do
	sleep 60
	if atlas clusters describe "${clusterName}" --projectId "${projectId}"; then
		status=$(atlas clusters describe "${clusterName}" --projectId "${projectId}" --output=json | jq -r '.stateName')
	else
		status="DELETED"
	fi
	echo "status: ${status}"
done


#delete Project
if atlas projects delete "$projectId"  --force; then
  echo "$projectId is deleted"
else
  (echo "Failed cleaning project:$projectId" && exit 1)
fi





