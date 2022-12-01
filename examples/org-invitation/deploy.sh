#!/bin/bash -e
name="$1"; shift

set -x
export AWS_DEFAULT_REGION="ap-northeast-2"
username="govardhan.pagidi@mongodb.com"
role="ORG_READ_ONLY"
team_id=$(mongocli iam team create "${name}" --username "${username}" --orgId "$ATLAS_ORG_ID" --output json | jq -r '.id')

# AWS CFN Deploy
aws cloudformation deploy \
    --stack-name stack-"$name" \
    --template-file org-invitation-sample.json \
    --no-fail-on-empty-changeset \
    --parameter-overrides Role="$role" TeamId="$team_id" Username="$username" PrivateKey="${ATLAS_PRIVATE_KEY}" PublicKey="${ATLAS_PUBLIC_KEY}" OrgId="${ATLAS_ORG_ID}" \
    "$@"


#if mongocli iam team delete "$team_id" --force
#then
#    echo "$team_id team deletion OK"
#else
#    (echo "Failed cleaning team:$team_id" && exit 1)
#fi