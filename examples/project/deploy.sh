#!/bin/bash -e
name="$1"; shift
PvtKey="6d39c229-6a6f-423d-9b41-dda35547f57e"
PubKey="pzovkttc"
org_id="5fe4ea50d1a2b617175ee3d4"
region="us-east-1"

set -x
export AWS_DEFAULT_REGION="ap-northeast-2"
api_key_id=$(mongocli iam org apikey create --orgId "${org_id}" --desc "${name}" --role ORG_MEMBER --output json | jq -r '.id')
#
##create team
user_name=$(mongocli iam project users list --output json | jq -r '.[0].emailAddress')
team_id=$(mongocli iam team create "${name}" --username "${user_name}" --orgId "${org_id}" --output json | jq -r '.id')

aws cloudformation deploy \
    --stack-name stack-"$name" \
    --template-file project2.json \
    --no-fail-on-empty-changeset \
    --parameter-overrides Name="$name" PrivateKey="$PvtKey" PublicKey="$PubKey" OrgId="${org_id}" TeamId="${team_id}" KeyId="${api_key_id}" \
    "$@"

#Delete the keys
mongocli iam teams delete "$team_id" --force
mongocli iam organizations apiKeys accessLists delete "$api_key_id" --force