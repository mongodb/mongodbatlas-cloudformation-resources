#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes AWS and Atlas resources created for `cfn test`.
#
echo "--------------------------------delete key and key policy document policy document starts ----------------------------"

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
echo "Check if a project is created $projectId"
export MCLI_PROJECT_ID=$projectId

keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
	keyRegion=$(aws configure get region)
fi
# shellcheck disable=SC2001
keyRegion=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"

roleName="mongodb-atlas-enc-role-${keyRegion}"
policyName="mongodb-atlas-kms-policy-${keyRegion}"

policyContent=$(jq '.Statement[0].Resource[0]' "$(dirname "$0")/policy.json")
echo "$policyContent"
keyID="${policyContent##*/}"
# shellcheck disable=SC2001
cleanedKeyID=$(echo "${keyID}" | sed 's/"//g')
echo "$cleanedKeyID"

aws kms schedule-key-deletion --key-id "$cleanedKeyID" --pending-window-in-days 7
echo "--------------------------------delete key and key policy document policy document ends ----------------------------"
pwd
trustPolicy=$(jq '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "$(dirname "$0")/add-policy.json")
echo "$trustPolicy"
roleExternalID="${trustPolicy##*/}"
# shellcheck disable=SC2001
atlasAssumedRoleExternalID=$(echo "${roleExternalID}" | sed 's/"//g')
echo "$atlasAssumedRoleExternalID"

roleId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${atlasAssumedRoleExternalID}" -r '.awsIamRoles[] |select(.atlasAssumedRoleExternalId |test( $roleID)) |.roleId')
echo "$roleId"

atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force
echo "--------------------------------delete role starts ----------------------------"

aws iam delete-role-policy --role-name "$roleName" --policy-name "$policyName"
aws iam delete-role --role-name "$roleName"
echo "--------------------------------delete role ends ----------------------------"
