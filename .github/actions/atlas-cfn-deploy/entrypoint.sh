#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit  # same as -e
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch if the pipe fucntion fails
set -o pipefail
set -x


echo "GITHUB_REF=${GITHUB_REF}"

echo "Setting up deploy tool dependencies"
python3 -m pip install -r util/atlas-cfn-deploy/requirements.txt

AWS_PROFILE="default"

#Check AWS credetials are defined in Gitlab Secrets
if [[ -z "$INPUT_AWS_ACCESS_KEY_ID" ]];then
    echo "AWS_ACCESS_KEY_ID is not SET!"; exit 1
fi

if [[ -z "$INPUT_AWS_SECRET_ACCESS_KEY" ]];then
    echo "AWS_SECRET_ACCESS_KEY is not SET!"; exit 2
fi

if [[ -z "$INPUT_AWS_REGION" ]];then
echo "AWS_REGION is not SET!"; exit 3
fi

env | grep "INPUT_"
AWS_ACCESS_KEY_ID="${INPUT_AWS_ACCESS_KEY_ID}"
AWS_SECRET_ACCESS_KEY="${INPUT_AWS_SECRET_ACCESS_KEY}"
AWS_REGION="${INPUT_AWS_REGION}"
ATLAS_PUBLIC_KEY="${INPUT_ATLAS_PUBLIC_KEY}"
ATLAS_PRIVATE_KEY="${INPUT_ATLAS_PRIVATE_KEY}"
ATLAS_ORG_ID="${INPUT_ATLAS_ORG_ID}"
LAUNCH_TEST_CLUSTER="${INPUT_LAUNCH_TEST_CLUSTER}"
env | grep AWS
env | grep ATLAS

aws configure --profile ${AWS_PROFILE} set aws_access_key_id ${AWS_ACCESS_KEY_ID}
aws configure --profile ${AWS_PROFILE} set aws_secret_access_key ${AWS_SECRET_ACCESS_KEY}
aws configure --profile ${AWS_PROFILE} set region ${AWS_REGION}

echo "Cleaning up any 'mongodb-atlas-*-role-stack's' in region: ${AWS_REGION}"
./util/atlas-cfn-stack-cleaner.sh

echo "Deploying all MongoDB Atlas CFN resources to ${AWS_REGION}"
./util/atlas-cfn-deploy/atlas-cfn-deploy.py --region=${AWS_REGION} all+


echo "LAUNCH_TEST_CLUSTER:${LAUNCH_TEST_CLUSTER}"
echo "Deployment complete. Be calm and data on."
