#!/usr/bin/env bash

# cfn-deactivate-helper.sh
#
#
# This tool helps run the AWS CloudFormation cli & api's needed to
# DEACTIVATE
# the resources from the AWS CloudFormation Public Registry
#
#
trap "exit" INT TERM ERR
trap "kill 0" EXIT
set -o errexit
set -o nounset
set -o pipefail


# Default, find all the directory names with the json custom resource schema files.
resources="${1:-project database-user project-ip-access-list network-peering cluster}"
echo "cfn-deactivate-helper.sh"
echo "$(basename "$0") running for the following resources: ${resources}"

PUBLISHER_ID=bb989456c78c398a858fef18f2ca1bfc1fbba082
PUBLIC_RESOURCE_TYPE_ARN_BASE="arn:aws:cloudformation:us-east-1::type/resource/${PUBLISHER_ID}/"
aws_account_id=$(aws sts get-caller-identity --query Account --output text)
ACTIVE_PUBLIC_RESOURCE_TYPE_ARN_BASE="arn:aws:cloudformation:us-east-1:${aws_account_id}:type/resource/"
for resource in ${resources};
do
    echo "Working on resource:${resource}"
    cwd=$(pwd)
    cd "${resource}"
    jsonschema="mongodb-atlas-$(echo ${resource}| sed s/-//g).json"
    echo "jsonschema=${jsonschema}"

    type_name=$(cat ${jsonschema}| jq -r '.typeName' | sed s/::/-/g)
    echo "type_name=${type_name}"

    stack_name="mongodb-atlas-${resource}-cfn-resource-role"

    arn="${ACTIVE_PUBLIC_RESOURCE_TYPE_ARN_BASE}${type_name}"
    echo "arn=${arn}"
    #aws cloudformation deactivate-type --public-type-arn ${public_type_arn}
    aws uno deactivate-type --arn ${arn}



    aws cloudformation delete-stack --stack-name ${stack_name}
    aws cloudformation wait stack-delete-complete --stack-name ${stack_name}
    cd -
done
echo "cfn-deactivate-helper.sh:done"

