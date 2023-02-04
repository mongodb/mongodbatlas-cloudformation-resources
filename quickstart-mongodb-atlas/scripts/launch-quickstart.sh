#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


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
               ${EXTRA_PARAMS} \
  --stack-name "${STACK_NAME}"
