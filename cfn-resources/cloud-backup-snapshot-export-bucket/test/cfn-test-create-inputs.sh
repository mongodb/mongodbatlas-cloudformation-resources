#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
#set -o errexit
#set -o nounset
#set -o pipefail

#set -x

function usage {
    echo "Creates a new cloud backup export bucket role for the test"
}

region=$AWS_DEFAULT_REGION
awsRegion=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
region=$(aws configure get region)
fi

# shellcheck disable=SC2001
region=$(echo "$region" | sed -e "s/-/_/g")
region=$(echo "$region" | tr '[:lower:]' '[:upper:]')
echo "$region"

roleName="mongodb-test-cloud-backup-export-bucket-role-${region}"
policyName="atlas-cloud-backup-export-bucket-S3-role-policy-${region}"

echo "roleName: ${roleName} , policyName: ${policyName}"

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

#------------ CREATING AtlAS ROLE -------------------
roleID=$(atlas cloudProviders accessRoles aws create --projectId "${projectId}" --output json | jq -r '.roleId')
echo -e "--------------------------------Mongo CLI Role creation ends ----------------------------\n"

#------------ Get role information-------------------
atlasAWSAccountArn=$(atlas cloudProviders accessRoles  list --projectId "${projectId}" --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles --projectId "${projectId}"  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
  '.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json

#------------ Create aws Iam role-------------------

awsRoleID=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
if [ -z "$awsRoleID" ]; then
    echo -e "No role found, hence creating the role. Created id: ${awsRoleID}\n"
else
    aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}"
    aws iam delete-role --role-name "${roleName}"
 awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document "file://$(dirname "$0")/add-policy.json" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
    echo -e "FOUND id: ${awsRoleID}\n"
fi
echo -e "--------------------------------AWS Role creation ends ----------------------------\n"

# Role Arn not returning immediately
sleep 30

#------------ get Role arn-------------------
awsArn=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.Arn')

aws iam put-role-policy   --role-name "${roleName}"   --policy-name "${policyName}"   --policy-document "file://$(dirname "$0")/policy.json"
echo -e "--------------------------------attach mongodb  Role to AWS Role ends ----------------------------\n"

# shellcheck disable=SC2001
awsArne=$(echo "${awsArn}" | sed 's/"//g')
# shellcheck disable=SC2086
# TODO Needs change to while loop using get operation
sleep 30

atlas cloudProviders accessRoles aws authorize "${roleID}" --projectId "${projectId}" --iamAssumedRoleArn "${awsArne}"
echo -e "--------------------------------authorize mongodb  Role ends ----------------------------\n"

#create the s3 bucket

bucketName="cloud-backup-snapshot-${CFN_TEST_TAG}-${awsRegion}"

aws s3 rb "s3://${bucketName}" --force
aws s3 mb "s3://${bucketName}" --output json

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

jq --arg projectId "$projectId" \
   --arg iamRoleID "$roleID" \
   --arg bucketName "$bucketName" \
   '.ProjectId?|=$projectId | .IamRoleID?|=$iamRoleID | .BucketName?|=$bucketName ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"