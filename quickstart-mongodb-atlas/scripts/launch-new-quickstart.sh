#!/usr/bin/env bash
STACK_NAME="${1:-aws-quickstart}"
aws cloudformation create-stack \
  --capabilities CAPABILITY_IAM --disable-rollback \
  --template-body file://templates/mongodb-atlas.template.yaml \
  --parameters ParameterKey=PublicKey,ParameterValue=${ATLAS_PUBLIC_KEY} \
               ParameterKey=PrivateKey,ParameterValue=${ATLAS_PRIVATE_KEY} \
               ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID} \
  --stack-name "${STACK_NAME}"
