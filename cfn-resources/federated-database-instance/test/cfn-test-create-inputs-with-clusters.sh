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

# cfn-test-create-inputs.sh
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

#project_id

projectName="${1}"
cluster1="${1}-cluster1"
cluster2="${1}-cluster2"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "Found project \"${projectName}\" with id: ${projectId}\n"
fi

keyRegion="us-east-1"
if [ -z "$keyRegion" ]; then
	keyRegion=$(aws configure get region)
fi

if atlas clusters describe "${cluster1}" --projectId "${projectId}"; then
	echo "Cluster found"
else
	echo "Cluster1 not found, creating..."
	atlas clusters create "${cluster1}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
	atlas clusters watch "${cluster1}" --projectId "${projectId}"
	echo -e "Created Cluster \"${cluster1}\""
fi

if atlas clusters describe "${cluster2}" --projectId "${projectId}"; then
	echo "Cluster found"
else
	echo "Cluster2 not found, creating..."
	atlas clusters create "${cluster2}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --diskSizeGB 10 --output=json
	atlas clusters watch "${cluster2}" --projectId "${projectId}"
	echo -e "Created Cluster \"${cluster2}\""
fi

# shellcheck disable=SC2001
keyRegionUnderScore=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegionUnderScore=$(echo "$keyRegionUnderScore" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"

roleName="mongodb-test-data-federation-role-${projectName}-${keyRegionUnderScore}"
policyName="atlas-data-federation-bucket-role-policy-${projectName}-${keyRegionUnderScore}"

echo "roleName: ${roleName} , policyName: ${policyName}"

echo -e "--------------------------------create key and key policy document starts ----------------------------\n"

echo -e "--------------------------------create aws bucket document starts ----------------------------\n"
bucketName="cfn-data-federation-test1-${projectName}-${keyRegion}"
aws s3 rb "s3://${bucketName}" --force
aws s3 mb "s3://${bucketName}" --output json

echo -e "--------------------------------create aws bucket document  ends ----------------------------\n"

roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo -e "--------------------------------Mongo CLI Role creation ends ----------------------------\n"

echo -e "--------------------------------printing mongodb role details ----------------------------\n"
atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg NAME "${projectName}" -r '.awsIamRoles[] |select(.iamAssumedRoleArn |test( "mongodb-test-export-role$")?)'
echo -e "--------------------------------AWS Role policy creation starts ----------------------------\n"

atlasAWSAccountArn=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId == $roleID) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId == $roleID) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
	--arg atlasAWSAccountArn "$atlasAWSAccountArn" \
	'.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json
echo -e "--------------------------------AWS Role creation ends ----------------------------\n"

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

sleep 65
atlas cloudProviders accessRoles aws authorize "${roleID}" --iamAssumedRoleArn "${awsArne}" --projectId "${projectId}"

echo -e "--------------------------------authorize mongodb  Role ends ----------------------------\n"

jq -r --arg projectId "$projectId" \
	--arg role "$roleID" \
	--arg bucketName "$bucketName" \
	--arg cluster1 "$cluster1" \
	--arg cluster2 "$cluster2" \
	'.Storage.Stores[0].Name?|=$cluster1 |.Storage.Stores[0].ClusterName?|=$cluster1 |.Storage.Stores[1].Name?|=$cluster2 |.Storage.Stores[1].ClusterName?|=$cluster2 |.Storage.Databases[0].Collections[0].DataSources[0].StoreName?|=$cluster1 |.Storage.Databases[0].Collections[0].DataSources[1].StoreName?|=$cluster2 |.CloudProviderConfig.TestS3Bucket?|=$bucketName |.CloudProviderConfig.RoleId?|=$role |.ProjectId?|=$projectId' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

ls -l inputs
