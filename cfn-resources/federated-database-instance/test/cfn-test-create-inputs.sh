#!/usr/bin/env bash
# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# cfn-test-create-inputs-with-clusters.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

#project_id
projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
	keyRegion=$(aws configure get region)
fi

# shellcheck disable=SC2001
keyRegionUnderScore=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegionUnderScore=$(echo "$keyRegionUnderScore" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"

echo -e "--------------------------------create aws bucket document starts ----------------------------\n"
bucketName="mongodb-atlas-cfn-test-df-${keyRegion}"
aws s3 rb "s3://${bucketName}" --force
aws s3 mb "s3://${bucketName}" --output json
echo -e "--------------------------------create aws bucket document  ends ----------------------------\n"

roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo -e "--------------------------------Mongo CLI Role creation ends ----------------------------\n"

echo -e "--------------------------------create key and key policy document starts ----------------------------\n"
roleName="mongodb-atlas-df-role-${keyRegionUnderScore}"
policyName="mongodb-atlas-df-bucket-role-policy-${keyRegionUnderScore}"
echo "roleName: ${roleName} , policyName: ${policyName}"

atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId == $roleID) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId == $roleID) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json

echo -e "--------------------------------AWS Role creation starts ----------------------------\n"

awsRoleID=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
if [ -z "$awsRoleID" ]; then
	awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document "file://$(dirname "$0")/add-policy.json" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
	echo -e "No role found, hence creating the role. Created id: ${awsRoleID}\n"
else
	aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}"
	aws iam delete-role --role-name "${roleName}"
	awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document "file://$(dirname "$0")/add-policy.json" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
	echo -e "FOUND id: ${awsRoleID}\n"
fi
echo -e "--------------------------------AWS Role creation ends ----------------------------\n"

awsArn=$(aws iam get-role --role-name "${roleName}" | jq -r '.Role | .Arn')

aws iam put-role-policy --role-name "${roleName}" --policy-name "${policyName}" --policy-document "file://$(dirname "$0")/policy.json"
echo -e "--------------------------------attach mongodb  Role to AWS Role ends ----------------------------\n"

echo -e "--------------------------------Role Id : ${roleID} ----------------------------\n"

# shellcheck disable=SC2001
awsArne=$(echo "${awsArn}" | sed 's/"//g')

# AWS IAM Takes time for IAM put role policy
sleep 65
atlas cloudProviders accessRoles aws authorize "${roleID}" --iamAssumedRoleArn "${awsArne}" --projectId "${projectId}"

echo -e "--------------------------------authorize mongodb  Role ends ----------------------------\n"

jq --arg projectId "$projectId" \
	--arg role "$roleID" \
	--arg name "${projectName}" \
	--arg bucketName "$bucketName" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .TenantName?|=$name | .CloudProviderConfig.TestS3Bucket?|=$bucketName |.CloudProviderConfig.RoleId?|=$role | .ProjectId?|=$projectId' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg projectId "$projectId" \
	--arg role "$roleID" \
	--arg name "${projectName}" \
	--arg bucketName "$bucketName" \
	--arg profile "$profile" \
	'.Profile?|=$profile | .TenantName?|=$name | .CloudProviderConfig.TestS3Bucket?|=$bucketName |.CloudProviderConfig.RoleId?|=$role | .ProjectId?|=$projectId' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
