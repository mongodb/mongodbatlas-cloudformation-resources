#!/usr/bin/env bash
org_id="$ATLAS_ORG_ID"
echo "$org_id"
mongocli iam org apikey create --orgId "${org_id}" --desc "cfn-test-bot3" --role ORG_MEMBER > orgid_key.json
cat orgid_key.json
api_key_id=$(cat orgid_key.json | jq -r '.id')
echo "$api_key_id"


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg name "$name" \
   --arg key_id $api_key_id \
   '.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .Name?|=$name | .ProjectApiKeys[0].Key?|=$key_id' \
   "project/test/inputs_1_update.template.json" > "project/inputs/inputs_1_update.json"
