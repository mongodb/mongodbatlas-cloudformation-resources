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

# export PROJECT_ID=5555555 APP_ID=44444444 DB_NAME=testDB COLLECTION_NAME=collection FUNC_NAME=funcn2 FUNC_ID=2222222 SERVICE_ID=111111

function usage {
	echo "usage: cfn-test-create-inputs.sh <projectId> <db_name> <coll_name> <trigger_name> <func_name> <func_id> <service_id> <app_id>"
	echo "or using environment variables to set params: 'export PROJECT_ID=5555555 APP_ID=44444444 DB_NAME=testDB COLLECTION_NAME=collection FUNC_NAME=funcn2 FUNC_ID=2222222 SERVICE_ID=111111'"
	echo "Creates a new DatabaseTrigger in given AppId within Project"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

# params start from #2 because cfn-testing-helper.sh calls these scripts with PROJECT_NAME as a param
project_Id="${2:-$PROJECT_ID}"
db_name="${3:-$DB_NAME}"
coll_name="${4:-$COLLECTION_NAME}"
trigger_name="cfn-test-trigger-$(date +%s)"
func_name="${5:-$FUNC_NAME}"
func_id="${6:-$FUNC_ID}"
service_id="${7:-$SERVICE_ID}"
app_id="${8:-$APP_ID}"

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
