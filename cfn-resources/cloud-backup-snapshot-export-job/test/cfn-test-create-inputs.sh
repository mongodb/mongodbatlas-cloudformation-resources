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


projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"
ClusterName="${projectName}"
clusterId=$(atlas clusters list --output json  | jq --arg NAME ${ClusterName} -r '.results[] | select(.name==$NAME) | .name')
if [ -z "$clusterId" ]; then
    clusterId=$(atlas cluster create ${ClusterName} --projectId ${projectId} --provider AWS --region US_EAST_1 --members 3 --backup --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json | jq -r '.name')
    sleep 1200
    echo -e "Created Cluster \"${ClusterName}\" with id: ${clusterId}\n"
else
    echo -e "FOUND Cluster \"${ClusterName}\" with id: ${clusterId}\n"
fi

SnapshotId=$( atlas backup snapshots list ${ClusterName} --output json  | jq --arg ID "6396e836d8a0ce6c6e99c3ef" -r '.results[] | select(.id==$ID) | .id')
if [ -z "$SnapshotId" ]; then
    SnapshotId=$( atlas backup snapshots create ${ClusterName} --desc "cfn unit test" --retention 3 --output=json | jq -r '.id')
    sleep 120
    echo -e "Created snapshot \"${SnapshotId}\""
else
    echo -e "FOUND snapshot  with id: ${SnapshotId}\n"
fi


echo "--------------------------------get aws region starts ----------------------------"\n

keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
keyRegion=$(aws configure get region)
fi
keyRegion=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"

roleName="mongodb-test-export-role-${keyRegion}"
policyName="atlas-bucket-role-policy-${keyRegion}"
bucketName="cfntest-demo-test-${AWS_DEFAULT_REGION}"
echo "roleName: ${roleName} , policyName: ${policyName}"
echo "--------------------------------get aws region ends ----------------------------"\n

echo "--------------------------------create aws bucket document starts ----------------------------"\n

aws s3 rb s3://${bucketName} --force
aws s3 mb s3://${bucketName} --output json

echo "--------------------------------create aws bucket document  ends ----------------------------"\n



roleID=$(atlas cloudProviders accessRoles aws create --output json | jq -r '.roleId')
echo "--------------------------------Mongo CLI Role creation ends ----------------------------"\n


atlasAWSAccountArn=$(atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
  '.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json
echo "--------------------------------AWS Role creation ends ----------------------------"\n


awsRoleID=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
if [ -z "$awsRoleID" ]; then
    awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://$(dirname "$0")/add-policy.json | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
    echo -e "No role found, hence creating the role. Created id: ${awsRoleID}\n"
else
    aws iam delete-role-policy --role-name "${roleName}" --policy-name "${policyName}"
    aws iam delete-role --role-name "${roleName}"
 awsRoleID=$(aws iam create-role --role-name "${roleName}" --assume-role-policy-document file://$(dirname "$0")/add-policy.json | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.RoleId')
    echo -e "FOUND id: ${awsRoleID}\n"
fi
echo "--------------------------------AWS Role creation ends ----------------------------"\n


awsArn=$(aws iam get-role --role-name "${roleName}" | jq --arg roleName "${roleName}" -r '.Role | select(.RoleName==$roleName) |.Arn')

aws iam put-role-policy   --role-name "${roleName}"   --policy-name "${policyName}"   --policy-document file://$(dirname "$0")/policy.json
echo "--------------------------------attach mongodb  Role to AWS Role ends ----------------------------"\n

awsArne=$(echo "${awsArn}" | sed 's/"//g')
# shellcheck disable=SC2086
#TODO Needs change to while loop using get operation
sleep 65

atlas cloudProviders accessRoles aws authorize ${roleID} --iamAssumedRoleArn ${awsArne}
echo "--------------------------------authorize mongodb  Role ends ----------------------------"\n

sleep 15
ExportBucketId=$(atlas backup export buckets create ${bucketName} --cloudProvider AWS --iamRoleId ${roleID} --output json | jq '._id')
echo "--------------------------------authorize mongodb  Role ends ----------------------------"\n
echo ${ExportBucketId}

ExportBucketId=$(echo ${ExportBucketId} | sed 's/"//g')
echo $ExportBucketId


jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg ClusterName "$ClusterName" \
   --arg group_id "$projectId" \
   --arg SnapshotId "$SnapshotId" \
    --arg ExportBucketId "$ExportBucketId" \
   '.ExportBucketId?|=$ExportBucketId |.SnapshotId?|=$SnapshotId | .GroupId?|=$group_id | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ClusterName?|=$ClusterName' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg region "${region}- more B@d chars !@(!(@====*** ;;::" \
   --arg group_id "$projectId" \
   --arg ClusterName "$ClusterName" \
   --arg SnapshotId "$SnapshotId" \
   --arg ExportBucketId "$ExportBucketId" \
   '.ExportBucketId?|=$ExportBucketId |.SnapshotId?|=$SnapshotId |.GroupId?|=$group_id | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ClusterName?|=$ClusterName' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"


#echo "mongocli iam projects delete ${projectId} --force"

ls -l  inputs
