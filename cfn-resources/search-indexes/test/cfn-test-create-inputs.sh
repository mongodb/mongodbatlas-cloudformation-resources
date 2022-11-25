#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectId="$ATLAS_ORG_ID"


cluster_name="$CLUSTER_NAME"
index_name="$INDEX_NAME"

#create team
#team_name="cfn-boto-team-${CFN_TEST_TAG}"
#user_name=$(mongocli iam project users list --output json | jq -r '.[0].emailAddress')
#team_id=$(mongocli iam team create "${team_name}" --username "${user_name}" --orgId "${org_id}" --output json | jq -r '.id')
#mongocli atlas cluster create "${cluster_name}" --projectId "${project_id}" --provider AWS --region US_EAST_1 --tier M10


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$projectId" \
   --arg cluster  "$cluster_name" \
   --arg name "$index_name" \
   '.GroupId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |.ClusterName?|=$cluster |.Name?|=$name' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$projectId" \
   '.GroupId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$projectId" \
   --arg cluster  "$cluster_name" \
   --arg name "$index_name" \
   '.GroupId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |.ClusterName?|=$cluster |.Name?|=$name' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"


ls -l inputs