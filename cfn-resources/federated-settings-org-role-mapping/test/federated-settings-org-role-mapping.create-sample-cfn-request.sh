#!/usr/bin/env bash


set -o errexit
set -o nounset
set -o pipefail

# federated-settings-org-role-mapping.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#
jq --arg pubkey "$MCLI_PUBLIC_API_KEY" \
   --arg pvtkey "$MCLI_PRIVATE_API_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
   '.FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey' \
  "$(dirname "$0")/federated-settings-org-role-mapping.sample-cfn-request.json"
