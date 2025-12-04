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

# Run this script with the Makefile
# make create-test-resources
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -eu

if [ -z "${MONGODB_ATLAS_PUBLIC_API_KEY+x}" ] || [ -z "${MONGODB_ATLAS_PRIVATE_API_KEY+x}" ]; then
	echo "Error: MONGODB_ATLAS_PUBLIC_API_KEY  and MONGODB_ATLAS_PRIVATE_API_KEY environment variables must be set"
	exit 1
fi

org_id="${MONGODB_ATLAS_ORG_ID}"
team_id="${MONGODB_ATLAS_TEAM_ID}"
profile="${MONGODB_ATLAS_PROFILE}"
project_name="Project-$(date +%s%3N)"

if [ -z "${MONGODB_ATLAS_ORG_API_KEY_ID+x}" ]; then
	api_key_id=$(atlas organizations apikeys create --desc "Created as part of the contract testing: ${project_name}" --role ORG_MEMBER --output json | jq -r '.id')
	echo "created api_key id $api_key_id"
else
	api_key_id="${MONGODB_ATLAS_ORG_API_KEY_ID}"
	echo "using existing api key $api_key_id"
fi

rm -rf "inputs" && mkdir "inputs"

jq --arg org "${org_id}" \
	--arg name "${project_name}" \
	--arg profile "${profile}" \
	--arg key_id "${api_key_id}" \
	--arg team_id "${team_id}" \
	'.OrgId?|=$org |.Name?|=$name |.Profile?|=$profile | .ProjectApiKeys[0].Key?|=$key_id | .ProjectTeams[0].TeamId?|=$team_id' \
	"test/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg org "${org_id}" \
	--arg name "${project_name}-update" \
	--arg profile "${profile}" \
	--arg key_id "$api_key_id" \
	--arg team_id "$team_id" \
	'.OrgId?|=$org | .Name?|=$name |.Profile?|=$profile | .ProjectApiKeys[0].Key?|=$key_id | .ProjectTeams[0].TeamId?|=$team_id' \
	"test/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

jq --arg org "${org_id}" \
	--arg name "${project_name}" \
	--arg profile "${profile}" \
	'.OrgId?|=$org |.Name?|=$name |.Profile?|=$profile' \
	"test/inputs_2_create.template.json" >"inputs/inputs_2_create.json"

jq --arg org "${org_id}" \
	--arg name "${project_name}" \
	--arg profile "${profile}" \
	'.OrgId?|=$org |.Name?|=$name |.Profile?|=$profile' \
	"test/inputs_2_update.template.json" >"inputs/inputs_2_update.json"
