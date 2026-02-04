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
	echo "Creates a new encryption key for the the project "
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
rm -rf inputs
mkdir inputs

# set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

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

echo "roleName: ${roleName} , policyName: ${policyName}"

echo "--------------------------------create key and key policy document starts ----------------------------"

keyARN=$(aws kms create-key | jq '.KeyMetadata|.Arn')
# shellcheck disable=SC2089
prefix='{ "Version": "2012-10-17", "Statement": ['
echo "--------------------------------printing key  starts ----------------------------"
echo "$keyARN"
# shellcheck disable=SC2001
cleanedkeyARN=$(echo "${keyARN}" | sed 's/"//g')
echo "$cleanedkeyARN"
echo "--------------------------------printing key  ends ----------------------------"

policyContent=$(jq --arg cleanedkeyARN "$cleanedkeyARN" '.Statement[0]|.Resource[0]?|=$cleanedkeyARN' "$(dirname "$0")/key-policy-template.json")
suffix=']}'
policyDocument="${prefix} ${policyContent} ${suffix}"
echo "$policyDocument" >"$(dirname "$0")"/policy.json

policyContent=$(jq '.Statement[0].Resource[0]' "$(dirname "$0")/policy.json")
echo "$policyContent"
keyID="${policyContent##*/}"
# shellcheck disable=SC2001
cleanedKeyID=$(echo "${keyID}" | sed 's/"//g')
echo "$cleanedKeyID"

echo "--------------------------------create key and key policy document policy document ends ----------------------------"

echo "$policyDocument"
echo "--------------------------------policy document finished ----------------------------"

roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "--------------------------------Mongo CLI Role creation ends ----------------------------"

atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json
echo "--------------------------------AWS Role creation ends ----------------------------"

awsRoleID=$(aws iam get-role --role-name "${roleName}" 2>/dev/null | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId' || true)
if [ -z "$awsRoleID" ]; then
	awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://"$(dirname "$0")"/add-policy.json | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
	echo -e "No role found, hence creating the role. Created id: ${awsRoleID}\n"
else
	aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}" 2>/dev/null || true
	aws iam delete-role --role-name "${roleName}"
	awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://"$(dirname "$0")"/add-policy.json | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
	echo -e "FOUND id: ${awsRoleID}\n"
fi
echo "--------------------------------AWS Role creation ends ----------------------------"

awsArn=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.Arn')

aws iam put-role-policy --role-name "${roleName}" --policy-name "${policyName}" --policy-document file://"$(dirname "$0")"/policy.json
echo "--------------------------------attach mongodb  Role to AWS Role ends ----------------------------"

# shellcheck disable=SC2001
awsArne=$(echo "${awsArn}" | sed 's/"//g')
# shellcheck disable=SC2086
#TODO Needs change to while loop using get operation
sleep 65

atlas cloudProviders accessRoles aws authorize "${roleID}" --projectId "${projectId}" --iamAssumedRoleArn "${awsArne}"
echo "--------------------------------authorize mongodb  Role ends ----------------------------"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	--arg KMS_KEY "$cleanedKeyID" \
	--arg KMS_ROLE "${roleID}" \
	--arg region "$keyRegion" \
	'.Profile?|=$profile | .ProjectId?|=$projectId | .AwsKmsConfig.CustomerMasterKeyID?|=$KMS_KEY | .AwsKmsConfig.RoleID?|=$KMS_ROLE | .AwsKmsConfig.Region?|=$region ' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	--arg KMS_KEY "$cleanedKeyID" \
	--arg KMS_ROLE "${roleID}" \
	--arg region "$keyRegion" \
	'.Profile?|=$profile | .ProjectId?|=$projectId | .AwsKmsConfig.CustomerMasterKeyID?|=$KMS_KEY | .AwsKmsConfig.RoleID?|=$KMS_ROLE | .AwsKmsConfig.Region?|=$region ' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
