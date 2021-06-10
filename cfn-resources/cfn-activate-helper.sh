#!/usr/bin/env bash

# cfn-activate-helper.sh
#
#
# This tool helps run the AWS CloudFormation cli & api's needed to
# activate the resources from the AWS CloudFormation Public Registry
#
#
trap "exit" INT TERM ERR
trap "kill 0" EXIT
#set -x
set -o errexit
set -o nounset
set -o pipefail

# Default, find all the directory names with the json custom resource schema files.
resources="${1:-project database-user project-ip-access-list network-peering cluster}"
echo "cfn-activate-helper.sh"
echo "$(basename "$0") running for the following resources: ${resources}"

PUBLISHER_ID=bb989456c78c398a858fef18f2ca1bfc1fbba082
PUBLIC_RESOURCE_TYPE_ARN_BASE="arn:aws:cloudformation:us-east-1::type/resource/${PUBLISHER_ID}/"
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
    aws cloudformation create-stack --capabilities CAPABILITY_IAM \
        --template-body "file://./resource-role.yaml" \
        --stack-name ${stack_name}
    aws cloudformation wait stack-create-complete \
        --stack-name ${stack_name}

    role_arn=$(aws cloudformation describe-stacks --stack-name ${stack_name} --query 'Stacks[0].Outputs[0].OutputValue' --output text)
    echo "role_arn=${role_arn}"

    public_type_arn="${PUBLIC_RESOURCE_TYPE_ARN_BASE}${type_name}"
    echo "public_type_arn=${public_type_arn}"
    #aws cloudformation activate-type --execution-role-arn ${role_arn} --public-type-arn ${public_type_arn}
    aws uno activate-type --execution-role-arn ${role_arn} --public-type-arn ${public_type_arn}
    cd -
done

