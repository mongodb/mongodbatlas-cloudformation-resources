#!/usr/bin/env bash
# Generate a random name for this push
ENV=${4}
echo "Loading secrets/environment from ${ENV}"
echo ----
source "${ENV}"
NAME=$(curl -s https://random-word-api.herokuapp.com/word\?number\=3\&swear\=0 | jq -r '. | join("-")')
NAME_TRUNC=$( echo ${NAME} | cut -d- -f1 )
TEMPLATE=${1}
REGION=${2}
INSTANCE_SIZE=${3}

env | grep ATLAS

secret=$(mktemp)
cat << EOF > "${secret}"
{
    "AtlasMongoDBPrivateKey": "${ATLAS_PRIVATE_KEY}",
    "AtlasMongoDBOrgID": "${ATLAS_ORG_ID}",
    "AtlasMongoDBPublicKey" : "${ATLAS_PUBLIC_KEY}"
}
EOF

echo "${secret}"
cat "${secret}"
echo "standby, baking a fresh yummy secret for your MongoDB Atlas API Key..."
echo "Provisioning AWS Secret for MongoDB Atlas API Key called 'mongodbatlas/atlas-cfn-quickstart/$NAME"
secret_info=$(aws secretsmanager create-secret \
 --name="mongodbatlas/atlas-cfn-quickstart/$NAME" \
 --region="${REGION}" \
 --secret-string="file://${secret}" )
echo "secret_info=$secret_info"

echo "Creating ${NAME} (${INSTANCE_SIZE}) from ${TEMPLATE} in ${REGION}"

PARAMS=$(cat <<PARAMETERS
[
 {
    "ParameterKey": "MongoDBAtlasClusterName",
    "ParameterValue": "${NAME}"
 },
 {
    "ParameterKey": "MongoDBAtlasUsername",
    "ParameterValue": "${NAME_TRUNC}"
 },
 {
    "ParameterKey": "MongoDBAtlasPassword",
    "ParameterValue": "MongoDB12345@"
 },
 {
    "ParameterKey": "MongoDBAtlasAPIKeySecretName",
    "ParameterValue": "mongodbatlas/atlas-cfn-quickstart/${NAME}"
 },
 {
    "ParameterKey": "MongoDBAtlasInstanceSize",
    "ParameterValue": "${INSTANCE_SIZE}"
 }
]
PARAMETERS
)

tp=$(mktemp)
echo "${PARAMS}" > "${tp}"
echo "Using ${tp} with ${PARAMS}"

aws cloudformation create-stack \
 --disable-rollback \
 --stack-name="${NAME}" \
 --region "${REGION}" \
 --template-body="file://${TEMPLATE}" \
 --parameters="file://${tp}"

