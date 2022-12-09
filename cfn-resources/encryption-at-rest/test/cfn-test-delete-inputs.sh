#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -x
echo "--------------------------------delete key and key policy document policy document starts ----------------------------"\n

policyContent=$(jq '.Statement[0].Resource[0]' "$(dirname "$0")/policy.json" )
echo "$policyContent"
keyID=$(echo ${policyContent##*/})
cleanedKeyID=$(echo "${keyID}" | sed 's/"//g')
echo $cleanedKeyID

aws kms schedule-key-deletion --key-id $cleanedKeyID --pending-window-in-days 7
echo "--------------------------------delete key and key policy document policy document ends ----------------------------"\n

trustPolicy=$(jq '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "add-policy.json" )
echo $trustPolicy
roleExternalID=$(echo ${trustPolicy##*/})
atlasAssumedRoleExternalID=$(echo "${roleExternalID}" | sed 's/"//g')
echo $atlasAssumedRoleExternalID

roleId=$(atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${atlasAssumedRoleExternalID}" -r '.awsIamRoles[] |select(.atlasAssumedRoleExternalId |test( $roleID)) |.roleId')
echo $roleId

atlas cloudProviders accessRoles aws deauthorize ${roleId} --force
echo "--------------------------------delete role starts ----------------------------"\n

aws iam delete-role-policy --role-name mongodb-test-enc-role --policy-name atlas-kms-role-policy
aws iam delete-role --role-name mongodb-test-enc-role
echo "--------------------------------delete role ends ----------------------------"\n

#mongocli iam projects delete "${projectId}" --force