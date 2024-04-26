#!/bin/bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
#      http://www.apache.org/licenses/LICENSE-2.0
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

set -euo pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#create apikey
org_id="$MONGODB_ATLAS_ORG_ID"

api_key_id=$(atlas organizations apikeys create --desc "cfn-boto-key-${CFN_TEST_TAG}" --role ORG_MEMBER --output json | jq -r '.id')

#create team
team_name="cfn-boto-team-${CFN_TEST_TAG}"
user_name=$(atlas organizations users list --output json | jq -r '.results[0].emailAddress')
echo "${user_name}"
team_id=$(atlas teams describe --name "$team_name" --output json | jq -r ".id")
if [ -z "$team_id" ]; then
	team_id=$(atlas team create "${team_name}" --username "${user_name}" --orgId "${org_id}" --output json | jq -r '.id')
fi

name="${1}"

jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg name "$name" \
	--arg key_id "$api_key_id" \
	--arg team_id "$team_id" \
	'.OrgId?|=$org |.Name?|=$name | .ProjectApiKeys[0].Key?|=$key_id | .ProjectTeams[0].TeamId?|=$team_id' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg name "${name}-update" \
	--arg key_id "$api_key_id" \
	--arg team_id "$team_id" \
	'.OrgId?|=$org | .Name?|=$name | .ProjectApiKeys[0].Key?|=$key_id | .ProjectTeams[0].TeamId?|=$team_id' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

jq --arg org "${org_id}" \
	--arg name "${name}-tags" \
	'.OrgId?|=$org |.Name?|=$name' \
	"test/inputs_2_create.template.json" >"inputs/inputs_2_create.json"

jq --arg org "${org_id}" \
	--arg name "${name}"-tags \
	'.OrgId?|=$org |.Name?|=$name' \
	"test/inputs_2_update.template.json" >"inputs/inputs_2_update.json"

ls -l inputs
echo "TODO: Delete the team and api_key created above"
