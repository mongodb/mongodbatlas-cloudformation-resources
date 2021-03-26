#!/usr/bin/env bash

# cfn-submit-helper.sh
#
#
# This tool works by running `cfn submit` on each resource
#
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

_CFN_FLAGS=${CFN_FLAGS:---verbose --set-default}

_BUILD_ONLY=${BUILD_ONLY:-false}
_SUBMIT_ONLY=${SUBMIT_ONLY:-false}

echo "Step 1/2: Building"
if [[ "${_SUBMIT_ONLY}" == "true" ]]; then
    echo "SUBMIT_ONLY, skipping build."
else
    for resource in ${resources};
    do
        echo "Working on resource:${resource}"
        cwd=$(pwd)
        cd "${resource}"
        echo "resource: ${resource}"
        make
        cd -
    done
fi
if [[ "${_BUILD_ONLY}" == "true" ]]; then
    echo "BUILD_ONLY true, skipping submit to CloudFormation"
    exit 0
fi
echo "Step 2/2: Submit resource type to CloudFormation Private Registry"
for resource in ${resources};
do
    echo "Working on resource:${resource}"
    cwd=$(pwd)
    cd "${resource}"
    echo "resource: ${resource}"
    echo "Submiting to CloudFormation with flags: ${_CFN_FLAGS}"
    cfn submit ${_CFN_FLAGS}
    cd -
done



