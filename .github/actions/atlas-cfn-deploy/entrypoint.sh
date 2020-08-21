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

env

AWS_PROFILE="default"

if [[ "${INPUT_AWS_ACCESS_KEY_ID:-}" ]];then
    AWS_ACCESS_KEY_ID=$INPUT_AWS_ACCESS_KEY_ID
fi
if [[ -z "$AWS_ACCESS_KEY_ID" ]];then
    echo "AWS_ACCESS_KEY_ID is not SET!"; exit 1
fi

if [[ "${INPUT_AWS_SECRET_ACCESS_KEY:-}" ]];then
    AWS_SECRET_ACCESS_KEY=$INPUT_AWS_SECRET_ACCESS_KEY
fi
if [[ -z "$AWS_SECRET_ACCESS_KEY" ]];then
    echo "AWS_SECRET_ACCESS_KEY is not SET!"; exit 2
fi

if [[ "${INPUT_AWS_REGION_SECRET:-}" ]];then
    AWS_REGION=$INPUT_AWS_REGION_SECRET
fi
if [[ "${INPUT_AWS_REGION_INPUT:-}" ]];then
    AWS_REGION=$INPUT_AWS_REGION_INPUT
fi

if [[ -z "$AWS_REGION" ]];then
    echo "AWS_REGION is not SET!"; exit 3
fi

aws configure --profile ${AWS_PROFILE} set aws_access_key_id ${AWS_ACCESS_KEY_ID}
aws configure --profile ${AWS_PROFILE} set aws_secret_access_key ${AWS_SECRET_ACCESS_KEY}
aws configure --profile ${AWS_PROFILE} set region ${AWS_REGION}

ls -al
ls -l cfn-resources
echo "Cleaning up any 'mongodb-atlas-*-role-stack's' in region: ${AWS_REGION}"
./atlas-cfn-stack-cleaner.sh

echo "Deploying all MongoDB Atlas CFN resources to ${AWS_REGION}"
./atlas-cfn-deploy.py --region=${AWS_REGION} all+

echo "Deployment complete. Be calm and data on."
