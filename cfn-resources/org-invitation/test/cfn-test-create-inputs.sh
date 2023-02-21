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
set -x

function usage {
	echo "usage:$0 <project_name>"
	echo "Creates a new project and an Cluster for testing"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

team_name="cfn-boto-team-${CFN_TEST_TAG}"
#user_name=$(atlas organizations users list --orgId "$ATLAS_ORG_ID" --output json | jq -r '.[0].emailAddress')
user_name=$(atlas organizations users list --orgId "$ATLAS_ORG_ID" --output json | jq -r '.results' | jq -r '.[0].emailAddress')
team_id=$(atlas teams create "${team_name}" --username "${user_name}" --orgId "$ATLAS_ORG_ID" --output json | jq -r '.id')

username="cfntest@mongodb.com"

jq --arg orgID "$ATLAS_ORG_ID" \
	--arg team_id "$team_id" \
	--arg username "$username" \
	'.OrgId?|=$orgID |.TeamIds[0]?|=$team_id |.Username?|=$username' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

#inputs_1_update.json
jq --arg orgID "$ATLAS_ORG_ID" \
	--arg team_id "$team_id" \
	--arg username "$username" \
	'.OrgId?|=$orgID |.TeamIds[0]?|=$team_id |.Username?|=$username' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

#inputs_1_invalid.json
username="(*&)(*&*&)(*&(*&"
jq --arg orgID "$ATLAS_ORG_ID" \
	--arg team_id "$team_id" \
	--arg username "$username" \
	'.OrgId?|=$orgID |.TeamIds[0]?|=$team_id |.Username?|=$username' \
	"$(dirname "$0")/inputs_1_invalid.template.json" >"inputs/inputs_1_invalid.json"
