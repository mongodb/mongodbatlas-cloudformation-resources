#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#


function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#create apikey
org_id="$ATLAS_ORG_ID"

api_key_id=$(mongocli iam project apikey create --desc "cfn-boto-key-${CFN_TEST_TAG}" --role GROUP_OWNER --output json | jq -r '.id')

#create team
team_name="cfn-boto-team-${CFN_TEST_TAG}"
user_name=$(mongocli iam project users list --output json | jq -r '.[0].emailAddress')
team_id=$(mongocli iam teams describe --name $team_name --output json | jq -r ".id")
if [ -z "$team_id" ]; then
  team_id=$(mongocli iam team create "${team_name}" --username "${user_name}" --orgId "${org_id}" --output json | jq -r '.id')
fi


name="${1}-test"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg name "$name" \
   --arg key_id "$api_key_id" \
   --arg team_id "$team_id" \
   '.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .Name?|=$name | .ProjectApiKeys[0].Key?|=$key_id | .ProjectTeams[0].TeamId?|=$team_id' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg name "${name}- more B@d chars !@(!(@====*** ;;::" \
   '.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .Name?|=$name' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg name "${name}" \
   --arg key_id "$api_key_id" \
   --arg team_id "$team_id" \
   '.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .Name?|=$name | .ProjectApiKeys[0].Key?|=$key_id | .ProjectTeams[0].TeamId?|=$team_id'\
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"

ls -l inputs

echo "TODO: Delete the team and api_key created above"
