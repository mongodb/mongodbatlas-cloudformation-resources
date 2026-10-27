#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
# It creates an Atlas project, registers a Cloud Provider Access role, and
# creates an AWS IAM role whose trust policy pins the external id returned by
# the registration. The role is intentionally left unauthorized: the resource
# under test performs the authorization.

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
	echo "Creates a project, a Cloud Provider Access role registration and an assumable AWS IAM role"
}

# Cloud Tag Policy: the IAM role created below must carry the required tags.
# shellcheck source=/dev/null
source "$(dirname "$0")/../../tag-policy.sh"
require_tag_policy_vars

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

# Keep cleanup possible if provisioning fails before the input files are complete.
printf '%s\n' "$projectId" >inputs/project-id.txt

echo "Check if a project is created $projectId"

echo "--------------------------------register Cloud Provider Access role starts ----------------------------"
roleId=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo "Created Cloud Provider Access role: ${roleId}"

roles=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json)
atlasAWSAccountArn=$(echo "$roles" | jq --arg roleId "${roleId}" -r '.awsIamRoles[] | select(.roleId == $roleId) | .atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(echo "$roles" | jq --arg roleId "${roleId}" -r '.awsIamRoles[] | select(.roleId == $roleId) | .atlasAssumedRoleExternalId')
echo "atlasAWSAccountArn: ${atlasAWSAccountArn}"
echo "atlasAssumedRoleExternalId: ${atlasAssumedRoleExternalId}"
echo "--------------------------------register Cloud Provider Access role ends ----------------------------"

echo "--------------------------------create AWS IAM role starts ----------------------------"
jq --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	--arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' \
	"$(dirname "$0")/trust-policy-template.json" >"$(dirname "$0")/trust-policy.json"

roleName="mongodb-atlas-cpa-role-$(date +%s)-${RANDOM}"
echo "roleName: ${roleName}"

# Record the name before creation, including when the response or next lookup fails.
printf '%s\n' "$roleName" >"$(dirname "$0")/role-name.txt"
aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://"$(dirname "$0")"/trust-policy.json --tags Key=mongodb-owner,Value="${MONGODB_TAG_OWNER}" Key=mongodb-env,Value="${MONGODB_TAG_ENV}" >/dev/null
awsArn=$(aws iam get-role --role-name "${roleName}" | jq -r '.Role.Arn')
echo "Created AWS IAM role: ${awsArn}"

echo "--------------------------------create AWS IAM role ends ----------------------------"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	--arg roleId "$roleId" \
	--arg arn "$awsArn" \
	'.Profile?|=$profile | .ProjectId?|=$projectId | .RoleId?|=$roleId | .IamAssumedRoleArn?|=$arn' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg projectId "$projectId" \
	--arg profile "$profile" \
	--arg roleId "$roleId" \
	--arg arn "$awsArn" \
	'.Profile?|=$profile | .ProjectId?|=$projectId | .RoleId?|=$roleId | .IamAssumedRoleArn?|=$arn' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
