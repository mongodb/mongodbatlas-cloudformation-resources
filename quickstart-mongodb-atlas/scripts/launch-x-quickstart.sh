#!/usr/bin/env bash
set -ex
TEMPLATE="${1:-templates/mongodb-atlas.template.yaml}"
STACK_NAME="${2:-aws-quickstart}"
EXTRA_PARAMS="${@: 3}"
echo "STACK_NAME=${STACK_NAME}, TEMPLATE=${TEMPLATE}"
echo "EXTRA_PARAMS=${EXTRA_PARAMS}"
aws cloudformation create-stack \
  --capabilities CAPABILITY_IAM --disable-rollback \
  --template-body "file://${TEMPLATE}" \
  --parameters ParameterKey=PublicKey,ParameterValue=${ATLAS_PUBLIC_KEY} \
               ParameterKey=PrivateKey,ParameterValue=${ATLAS_PRIVATE_KEY} \
               ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID} \
               ${EXTRA_PARAMS} \
  --stack-name "${STACK_NAME}"
