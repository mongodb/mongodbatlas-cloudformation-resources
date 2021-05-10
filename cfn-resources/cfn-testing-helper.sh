#!/usr/bin/env bash

# cfn-testing-helper.sh
#
#
# This tool works by running `cfn submit` on each resource
# By default it will build and submit every resource found
# in this directory.
# There are some options.
#
# BUILD_ONLY=true|false
# SUBMIT_ONLY=true|false
# LOG_LEVEL=logrus valid string loglevel
#
# Example with DEBUG logging enabled by default for set of resources:
# LOG_LEVEL=debug ./cfn-testing-helper.sh project database-user project-ip-access-list cluster network-peering
#
trap "exit" INT TERM ERR
trap "kill 0" EXIT
set -x
set -o errexit
set -o nounset
set -o pipefail

# Default, find all the directory names with the json custom resource schema files.
resources="${@: 1}"
if [ $# -eq 0 ]
  then
    echo "No arguments supplied, will submit all resource"
    resources=$(ls -F **/mongodb-atlas-*.json | cut -d/ -f1)
fi
echo "Submitting the following resources: ${resources}"

_CFN_FLAGS=${CFN_FLAGS:---verbose}

_SKIP_BUILD=${SKIP_BUILD:-false}
_BUILD_ONLY=${BUILD_ONLY:-false}
_DEFAULT_LOG_LEVEL=${LOG_LEVEL:-info}

echo "Step 1/2: Building"
for resource in ${resources};
do
    if [[ "${_SKIP_BUILD}" == "true" ]]; then
        continue
    fi
    echo "Working on resource:${resource}"
    cwd=$(pwd)
    cd "${resource}"
    echo "resource: ${resource}"
    if [[ "${_DEFAULT_LOG_LEVEL}" == "debug" ]]; then
        make debug
    else
        make
    fi
    cd -
done
if [[ "${_BUILD_ONLY}" == "true" ]]; then
    echo "BUILD_ONLY true, skipping testing with the CloudFormation CLI"
    exit 0
fi



echo "Step 2/3: Generating 'cfn test' 'inputs/' folder from each 'test/cfn-test-create-inputs.sh'"
#if [ ! -d "./inputs" ]; then
#fi

# Start full pass generating inputs based off a starting project
# We need a project to test all the other resources first,
# each resource will create it's own project off base name
# project - for that we create new project with the test
# cluster - test creates it's own project
# ditto dbuser
#, so there will a few projects total for whole test run:
# base: "${CFN_TEST_NEW_PROJECT_NAME}"
# project: "${CFN_TEST_NEW_PROJECT_NAME}-project"
# cluster: "${CFN_TEST_NEW_PROJECT_NAME}-cluster"
# etc...
#
. ./cfn-testing-helper.config
env | grep CFN_TEST_


PROJECT_NAME="${CFN_TEST_NEW_PROJECT_NAME}"
echo "PROJECT_NAME:${PROJECT_NAME}"

#if false; then

for res in ${resources};
do
    cd "${res}"
    ./test/cfn-test-create-inputs.sh "${PROJECT_NAME}-${res}" && echo "resource:${res} inputs created OK" || echo "resource:${res} input create FAILED"
    echo "Generated inputs for: ${res}"
    echo "----------------------------"
    ls -l ./inputs
    cd -
    echo ""
done


#fi

# TODO - network peering
# find vpc to use using awscli
#echo "usage:$0 <project_name> <aws_account_id> <vpc_id>"
# ./test/cfn-test-create-inputs.sh PeeringList-CFNTest-2 466197078724 vpc-fa3d7680
#res="network-peering"
#cd "${res}"
#./${res}/test/cfn-test-create-inputs.sh "${PROJECT_NAME}-2" && echo "resource:${res} inputs created OK" || echo "resource:${res} input create FAILED"




echo "Step 3/3: Running 'cfn test' on resource type"
SAM_LOG=$(mktemp)
for resource in ${resources};
do
    echo "Working on resource:${resource}"
    cwd=$(pwd)
    cd "${resource}"
    sam_log="${SAM_LOG}.${resource}"
    echo "starting resource handler lambda in background - capture output to: ${sam_log}"
    sam local start-lambda &> "${sam_log}" &
    sam_pid=$!
    echo "Started 'sam local start-lamda' with PID:${sam_pid}, wait 3 seconds to startup..." && sleep 3
    ps -ef | grep ${sam_pid}
    echo "resource: ${resource}, running 'cfn test' with flags: ${_CFN_FLAGS}"
    cfn test ${_CFN_FLAGS}
    echo "killing sam_pid:${sam_pid}"
    kill ${sam_pid}
    echo "sam_log: ${sam_log}"
    cat "${sam_log}"
    cd -
done



echo "Clean up afterwards"
for resource in ${resources};
do
    echo "res=${res}"
    p_id=$(mongocli iam project list --output=json | jq --arg name "${PROJECT_NAME}-${res}" -r '.results[] | select(.name==$name) | .id')
    p_name=$(mongocli iam project list --output=json | jq --arg name "${PROJECT_NAME}-${res}" -r '.results[] | select(.name==$name) | .name')
    echo "Cleaning up for resource:${res}, project:${p_name} id:${p_id}"
    mongocli iam project delete ${p_id} --force && echo "Clean up project:${p_name} id:${p_id}" || (echo "Failed cleaning up project:${p_id}" && exit 1)
done



