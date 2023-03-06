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

project_Id="${2:-63f4df9e1c744217893c19f7}"
db_name="${3:-sample_mflix}"
coll_name="${4:-comments}"
trigger_name="salestrigger-${RANDOM}"
func_name="${5:-myFuncInTrigger}"
func_id="${6:-64020cfc5ef00ef0be9584e4}"
service_id="${7:-64020111d61a66df7ea2c732}"
app_id="${7:-64020111d61a66df7ea2c72b}"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg db "$db_name" \
		--arg coll "$coll_name" \
		--arg funcName "$func_name" \
		--arg funcId "$func_id" \
		--arg appId "$app_id" \
		--arg serviceId "$service_id" \
		--arg projectId "$project_Id" \
		--arg triggerName "$trigger_name" \
		'.AppId?|=$appId |.DatabaseTrigger.ServiceId?|=$serviceId | .ProjectId?|=$projectId |
 .EventProcessors.FUNCTION.FuncConfig.FunctionName?|=$funcName |.EventProcessors.FUNCTION.FuncConfig.FunctionId?|=$funcId|
 .DatabaseTrigger.Collection?|=$coll |.DatabaseTrigger.Database?|=$db | .Name?|=$triggerName' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
