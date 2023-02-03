#!/usr/bin/env bash
set -x
function usage {
    echo "usage:$0 <project_name>"
    echo "Creates a new encryption key for the the project "
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
rm -rf inputs
mkdir inputs

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_CONNECTED_ORG_ID" \
   --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
   '.FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_CONNECTED_ORG_ID" \
   --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
   '.FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey ' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_CONNECTED_ORG_ID" \
   --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
   '.FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey ' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"

ls -l inputs
#mongocli iam projects delete "${projectId}" --force



#mongocli atlas cloudProviders accessRoles aws authorize 63721b924ad9a46eeef105ae --iamAssumedRoleArn "arn:aws:iam::816546967292:role/mongodb-test-role"