#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o nounset
set -o pipefail
set -x

function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
region="us-east-1"

#project_id

projectName="${1}"
projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"
ClusterName="${projectName}"
clusterId=$(mongocli atlas clusters list --output json  | jq --arg NAME ${ClusterName} -r '.results[] | select(.name==$NAME) | .name')
if [ -z "$clusterId" ]; then
    clusterId=$(mongocli atlas cluster create ${ClusterName} --projectId ${projectId} --provider AWS --region US_EAST_1 --members 3 --backup --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json | jq -r '.name')
    sleep 20m
    echo -e "Created Cluster \"${ClusterName}\" with id: ${clusterId}\n"
else
    echo -e "FOUND Cluster \"${ClusterName}\" with id: ${clusterId}\n"
fi


echo "--------------------------------get aws region starts ----------------------------"\n

keyRegion=$(aws configure get region)
keyRegion=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"
echo "--------------------------------get aws region ends ----------------------------"\n

echo "--------------------------------create aws bucket document starts ----------------------------"\n
bucketName="cfntest-demo-test-123"
aws s3 rb s3://${bucketName} --force
aws s3 mb s3://${bucketName} --output json

echo "--------------------------------create aws bucket document  ends ----------------------------"\n


echo "--------------------------------Mongo CLI Role creation starts ----------------------------"\n
roleID=$(mongocli atlas cloudProviders accessRoles  list --output json | jq --arg NAME "${projectName}" -r '.awsIamRoles[] |select(.iamAssumedRoleArn |test( "mongodb-test-export-role$")?) |.roleId')
if [ -z "$roleID" ]; then
    roleID=$(mongocli atlas cloudProviders accessRoles aws create --output json | jq -r '.roleId')
    echo -e "Created id: ${roleID}\n"
else
    echo -e "FOUND id: ${roleID}\n"
fi
echo "--------------------------------Mongo CLI Role creation ends ----------------------------"\n

echo "--------------------------------printing mongodb role details ----------------------------"\n
mongocli atlas cloudProviders accessRoles  list --output json | jq --arg NAME "${projectName}" -r '.awsIamRoles[] |select(.iamAssumedRoleArn |test( "mongodb-test-export-role$")?)'
echo "--------------------------------AWS Role policy creation starts ----------------------------"\n

atlasAWSAccountArn=$(mongocli atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)?) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(mongocli atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)?) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
  '.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json
echo "--------------------------------AWS Role policy creation ends ----------------------------"\n

echo "--------------------------------AWS Role  creation starts ----------------------------"\n
awsRoleID=$(aws iam get-role --role-name mongodb-test-export-role | jq '.Role|select(.RoleName |test( "mongodb-test-export-role$")?) |.RoleId')
if [ -z "$awsRoleID" ]; then
    awsRoleID=$(aws iam create-role --role-name mongodb-test-export-role --assume-role-policy-document file://$(dirname "$0")/add-policy.json | jq '.Role|select(.RoleName |test( "mongodb-test-export-role$")) |.RoleId')
    echo -e "Created id: ${awsRoleID}\n"
else
    aws iam delete-role-policy --role-name mongodb-test-export-role --policy-name atlas-role-policy
    aws iam delete-role --role-name mongodb-test-export-role
	awsRoleID=$(aws iam create-role --role-name mongodb-test-export-role --assume-role-policy-document file://$(dirname "$0")/add-policy.json | jq '.Role|select(.RoleName |test( "mongodb-test-export-role$")) |.RoleId')
    echo -e "FOUND id: ${awsRoleID}\n"
fi
echo "--------------------------------AWS Role creation ends ----------------------------"\n

echo "--------------------------------printing AWS Role ----------------------------"\n
aws iam get-role --role-name mongodb-test-export-role
echo "--------------------------------printing AWS Role ----------------------------"\n

echo "--------------------------------attach mongodb  Role to AWS Role starts ----------------------------"\n
awsArn=$(aws iam get-role --role-name mongodb-test-export-role | jq '.Role|select(.RoleName |test( "mongodb-test-export-role$")?) |.Arn')
mongocli atlas cloudProviders accessRoles  list --output json
aws iam put-role-policy   --role-name mongodb-test-export-role   --policy-name atlas-role-policy   --policy-document file://$(dirname "$0")/policy.json
echo "--------------------------------attach mongodb  Role to AWS Role ends ----------------------------"\n

echo "--------------------------------authorize mongodb  Role starts ----------------------------"\n

echo "--------------------------------Role Id ----------------------------"\n"${roleID}"
awsArne=$(echo "${awsArn}" | sed 's/"//g')
# shellcheck disable=SC2086
echo "--------------------------------Role Id ----------------------------"\n${awsArne}
mongocli atlas cloudProviders accessRoles aws authorize ${roleID} --iamAssumedRoleArn ${awsArne}
sleep 65


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg group_id "$projectId" \
   --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
   --arg AWSAssumedArn "$awsArne" \
   --arg role "$roleID" \
   --arg bucketName "$bucketName" \
   '.TenantName?|=$bucketName |.CloudProviderConfig.Aws.TestS3Bucket?|=$bucketName |.CloudProviderConfig.Aws.RoleId?|=$role |.CloudProviderConfig.Aws.IamUserARN?|=$atlasAWSAccountArn |.CloudProviderConfig.Aws.ExternalId?|=$atlasAssumedRoleExternalId | .CloudProviderConfig.Aws.IamAssumedRoleARN?|=$AWSAssumedArn | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$group_id' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg group_id "$projectId" \
   --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
   --arg AWSAssumedArn "$awsArne" \
   --arg role "$roleID" \
   --arg bucketName "$bucketName" \
   '.TenantName?|=$bucketName |.CloudProviderConfig.Aws.TestS3Bucket?|=$bucketName |.CloudProviderConfig.Aws.RoleId?|=$role |.CloudProviderConfig.Aws.IamUserARN?|=$atlasAWSAccountArn |.CloudProviderConfig.Aws.ExternalId?|=$atlasAssumedRoleExternalId | .CloudProviderConfig.Aws.IamAssumedRoleARN?|=$AWSAssumedArn | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$group_id' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg group_id "$projectId" \
   --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
   --arg AWSAssumedArn "$awsArne" \
   --arg role "$roleID" \
   --arg bucketName "$bucketName" \
   '.TenantName?|=$bucketName |.CloudProviderConfig.Aws.TestS3Bucket?|=$bucketName |.CloudProviderConfig.Aws.RoleId?|=$role |.CloudProviderConfig.Aws.IamUserARN?|=$atlasAWSAccountArn |.CloudProviderConfig.Aws.ExternalId?|=$atlasAssumedRoleExternalId | .CloudProviderConfig.Aws.IamAssumedRoleARN?|=$AWSAssumedArn | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$group_id' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"
#echo "mongocli iam projects delete ${projectId} --force"

ls -l inputs
