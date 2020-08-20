#!/usr/bin/env bash

# Generate a random name for this push
echo "Loading secrets and env from ../.env"
echo $(dirname "$0")
echo ----
source ../.env
NAME=$(curl -s https://random-word-api.herokuapp.com/word\?number\=3\&swear\=0 | jq -r '. | join("-")')
NAME_TRUNC=$( echo ${NAME} | cut -d- -f1 )
TEMPLATE=${1}
REGION=${2}
INSTANCE_SIZE=${3}

env | grep ATLAS
# create the secret for the apikey if env
if [[ ! -z ${ATLAS_PUBLIC_KEY+x} ]]; then
  echo "creating secret"
  aws secretsmanager create-secret --name "mongodbatlas/atlas-cfn-quickstart/${NAME}" \
  --secret-string '{\"AtlasMongoDBPrivateKey\": \"${ATLAS_PRIVATE_KEY}\", \"AtlasMongoDBOrgID\": \"${ATLAS_ORG_ID}\", \"AtlasMongoDBPublicKey\" : \"${ATLAS_PUBLIC_KEY\"}' \
  --region us-east-2

fi

echo "Creating ${NAME} (${INSTANCE_SIZE}) from ${TEMPLATE} in ${REGION}"
# {
#    "ParameterKey": "VPC",
#    "ParameterValue": "vpc-03330ad23571b081a"
# },
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

