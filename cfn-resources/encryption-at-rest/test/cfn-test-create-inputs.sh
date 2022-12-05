set -x
function usage {
    echo "usage:$0 <project_name>"
    echo "Creates a new encryption key for the the project "
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
rm -rf inputs
mkdir inputs

projectName="${1}"

projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "--------------------------------get aws region starts ----------------------------"\n

keyRegion=$(aws configure get region)
keyRegion=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"
echo "--------------------------------get aws region ends ----------------------------"\n

echo "--------------------------------create key and key policy document starts ----------------------------"\n
keyARN=$(aws kms create-key | jq '.KeyMetadata|.Arn')
prefix='{ "Version": "2012-10-17", "Statement": ['
echo "--------------------------------printing key  starts ----------------------------"\n
echo "$keyARN"
cleanedkeyARN=$(echo ${keyARN} | sed 's/"//g')
echo $cleanedkeyARN
echo "--------------------------------printing key  ends ----------------------------"\n
policyContent=$(jq --arg cleanedkeyARN $cleanedkeyARN '.Statement[0]|.Resource[0]?|=$cleanedkeyARN' "$(dirname "$0")/key-policy-template.json" )
suffix=']}'
policyDocument="${prefix} ${policyContent} ${suffix}"
echo $policyDocument > $(dirname "$0")/policy.json

policyContent=$(jq '.Statement[0].Resource[0]' "$(dirname "$0")/policy.json" )
echo "$policyContent"
keyID=$(echo ${policyContent##*/})
cleanedKeyID=$(echo "${keyID}" | sed 's/"//g')
echo $cleanedKeyID

echo "--------------------------------create key and key policy document policy document ends ----------------------------"\n

echo "--------------------------------printing policy document policy document ----------------------------"\n
echo $policyDocument
echo "--------------------------------policy document finished ----------------------------"\n

echo "--------------------------------Mongo CLI Role creation starts ----------------------------"\n
roleID=$(mongocli atlas cloudProviders accessRoles  list --output json | jq --arg NAME "${projectName}" -r '.awsIamRoles[] |select(.iamAssumedRoleArn |test( "mongodb-test-role$")) |.roleId')
if [ -z "$roleID" ]; then
    roleID=$(mongocli atlas cloudProviders accessRoles aws create --output json | jq -r '.roleId')
    echo -e "Created id: ${roleID}\n"
else
    echo -e "FOUND id: ${roleID}\n"
fi
echo "--------------------------------Mongo CLI Role creation ends ----------------------------"\n

echo "--------------------------------printing mongodb role details ----------------------------"\n
mongocli atlas cloudProviders accessRoles  list --output json | jq --arg NAME "${projectName}" -r '.awsIamRoles[] |select(.iamAssumedRoleArn |test( "mongodb-test-role$"))'
echo "--------------------------------AWS Role creation starts ----------------------------"\n

atlasAWSAccountArn=$(mongocli atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAWSAccountArn')
atlasAssumedRoleExternalId=$(mongocli atlas cloudProviders accessRoles  list --output json | jq --arg roleID "${roleID}" -r '.awsIamRoles[] |select(.roleId |test( $roleID)) |.atlasAssumedRoleExternalId')
jq --arg atlasAssumedRoleExternalId "$atlasAssumedRoleExternalId" \
   --arg atlasAWSAccountArn "$atlasAWSAccountArn" \
  '.Statement[0].Principal.AWS?|=$atlasAWSAccountArn | .Statement[0].Condition.StringEquals["sts:ExternalId"]?|=$atlasAssumedRoleExternalId' "$(dirname "$0")/role-policy-template.json" >"$(dirname "$0")/add-policy.json"
echo cat add-policy.json
echo "--------------------------------AWS Role creation ends ----------------------------"\n

echo "--------------------------------AWS Role  creation starts ----------------------------"\n
awsRoleID=$(aws iam get-role --role-name mongodb-test-role | jq '.Role|select(.RoleName |test( "mongodb-test-role$")) |.RoleId')
if [ -z "$awsRoleID" ]; then
    awsRoleID=$(aws iam create-role --role-name mongodb-test-role --assume-role-policy-document file://$(dirname "$0")/add-policy.json | jq '.Role|select(.RoleName |test( "mongodb-test-role$")) |.RoleId')
    echo -e "Created id: ${awsRoleID}\n"
else
    aws iam delete-role-policy --role-name mongodb-test-role --policy-name atlas-kms-role-policy
    aws iam delete-role --role-name mongodb-test-role
	awsRoleID=$(aws iam create-role --role-name mongodb-test-role --assume-role-policy-document file://$(dirname "$0")/add-policy.json | jq '.Role|select(.RoleName |test( "mongodb-test-role$")) |.RoleId')
    echo -e "FOUND id: ${awsRoleID}\n"
fi
echo "--------------------------------AWS Role creation ends ----------------------------"\n

echo "--------------------------------printing AWS Role ----------------------------"\n
aws iam get-role --role-name mongodb-test-role
echo "--------------------------------printing AWS Role ----------------------------"\n

echo "--------------------------------attach mongodb  Role to AWS Role starts ----------------------------"\n
awsArn=$(aws iam get-role --role-name mongodb-test-role | jq '.Role|select(.RoleName |test( "mongodb-test-role$")) |.Arn')
mongocli atlas cloudProviders accessRoles  list --output json
aws iam put-role-policy   --role-name mongodb-test-role   --policy-name atlas-kms-role-policy   --policy-document file://$(dirname "$0")/policy.json
echo "--------------------------------attach mongodb  Role to AWS Role ends ----------------------------"\n

echo "--------------------------------authorize mongodb  Role starts ----------------------------"\n

echo "--------------------------------Role Id ----------------------------"\n"${roleID}"
awsArne=$(echo "${awsArn}" | sed 's/"//g')
# shellcheck disable=SC2086
echo "--------------------------------Role Id ----------------------------"\n${awsArne}
mongocli atlas cloudProviders accessRoles aws authorize ${roleID} --iamAssumedRoleArn ${awsArne}
echo "--------------------------------authorize mongodb  Role ends ----------------------------"\n



jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg KMS_KEY "$cleanedKeyID" \
   --arg KMS_ROLE "${roleID}" \
   --arg region "$keyRegion" \
   '.AwsKms.CustomerMasterKeyID?|=$KMS_KEY | .AwsKms.RoleID?|=$KMS_ROLE | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ProjectId?|=$projectId | .AwsKms.Region?|=$region ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg KMS_KEY "$cleanedKeyID" \
   --arg KMS_ROLE "${roleID}" \
   --arg projectId "$projectId" \
   --arg region "$keyRegion" \
    '.AwsKms.CustomerMasterKeyID?|=$KMS_KEY | .AwsKms.RoleID?|=$KMS_ROLE | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ProjectId?|=$projectId | .AwsKms.Region?|=$region' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg KMS_KEY "$cleanedKeyID" \
   --arg KMS_ROLE "${roleID}" \
   --arg region "$keyRegion" \
   '.AwsKms.CustomerMasterKeyID?|=$KMS_KEY | .AwsKms.RoleID?|=$KMS_ROLE | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ProjectId?|=$projectId | .AwsKms.Region?|=$region ' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"
ls -l inputs
#mongocli iam projects delete "${projectId}" --force



#mongocli atlas cloudProviders accessRoles aws authorize 63721b924ad9a46eeef105ae --iamAssumedRoleArn "arn:aws:iam::816546967292:role/mongodb-test-role"