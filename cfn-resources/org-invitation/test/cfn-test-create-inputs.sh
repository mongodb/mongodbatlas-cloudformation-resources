#!/usr/bin/env bash

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

team_name="cfn-boto-team-${CFN_TEST_TAG}"
user_name=$(atlas organizations users list --orgId "$MONGODB_ATLAS_ORG_ID" --output json | jq -r '.results' | jq -r '.[0].emailAddress')
team_id=$(atlas teams create "${team_name}" --username "${user_name}" --orgId "$MONGODB_ATLAS_ORG_ID" --output json | jq -r '.id')

username="cfntest@mongodb.com"

jq --arg orgID "$MONGODB_ATLAS_ORG_ID" \
	--arg team_id "$team_id" \
	--arg username "$username" \
	'.OrgId?|=$orgID |.TeamIds[0]?|=$team_id |.Username?|=$username' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg orgID "$MONGODB_ATLAS_ORG_ID" \
	--arg team_id "$team_id" \
	--arg username "$username" \
	'.OrgId?|=$orgID |.TeamIds[0]?|=$team_id |.Username?|=$username' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"
