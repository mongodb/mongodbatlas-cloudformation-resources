#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

function usage {
	echo "usage:$0 <project_name>"
	echo "Creates a new encryption key for the the project "
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
rm -rf inputs
mkdir inputs

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

roleName="mongodb-test-enc-role-${keyRegion}"
policyName="atlas-kms-role-policy-${keyRegion}"

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
# shellcheck disable=SC2116
keyID=$(echo "${policyContent##*/}")
# shellcheck disable=SC2001
cleanedKeyID=$(echo "${keyID}" | sed 's/"//g')
echo "$cleanedKeyID"

echo "--------------------------------create key and key policy document policy document ends ----------------------------"

echo "$policyDocument"
echo "--------------------------------policy document finished ----------------------------"

roleID=$(atlas cloudProviders accessRoles aws create --output json | jq -r '.roleId')
echo "--------------------------------Mongo CLI Role creation ends ----------------------------"

atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json
echo "--------------------------------AWS Role creation ends ----------------------------"

awsRoleID=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
if [ -z "$awsRoleID" ]; then
	awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://"$(dirname "$0")"/add-policy.json | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
	echo -e "No role found, hence creating the role. Created id: ${awsRoleID}\n"
else
	aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}"
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

atlas cloudProviders accessRoles aws authorize "${roleID}" --iamAssumedRoleArn "${awsArne}"
echo "--------------------------------authorize mongodb  Role ends ----------------------------"

jq --arg projectId "$projectId" \
	--arg KMS_KEY "$cleanedKeyID" \
	--arg KMS_ROLE "${roleID}" \
	--arg region "$keyRegion" \
	'.AwsKmsConfig.CustomerMasterKeyID?|=$KMS_KEY | .AwsKmsConfig.RoleID?|=$KMS_ROLE | .ProjectId?|=$projectId | .AwsKmsConfig.Region?|=$region ' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg projectId "$projectId" \
	--arg KMS_KEY "$cleanedKeyID" \
	--arg KMS_ROLE "${roleID}" \
	--arg region "$keyRegion" \
	'.AwsKmsConfig.CustomerMasterKeyID?|=$KMS_KEY | .AwsKmsConfig.RoleID?|=$KMS_ROLE | .ProjectId?|=$projectId | .AwsKmsConfig.Region?|=$region ' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
