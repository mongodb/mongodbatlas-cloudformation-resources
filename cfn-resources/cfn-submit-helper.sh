#!/usr/bin/env bash

# cfn-submit-helper.sh
#
#
# This tool works by running `cfn submit` on each resource
#

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

_BUILD_ONLY=${BUILD_ONLY:-0}

for resource in ${resources};
do
    echo "Working on resource:${resource}"
    cwd=$(pwd)
    cd "${resource}"
    echo "resource: ${resource}"
    echo "Building (Pass 1/2): with TAGS=\"logging callback\""
    TAGS="logging callback" make
    echo "Running gofmt before pass 2 `gofmt cmd/`"
    gofmt cmd/
    echo "Building (Pass 2/2): with TAGS=\"logging callback\""
    TAGS="logging callback" make
    if [[ "${_BUILD_ONLY}" == "true" ]]; then
        echo "BUILD_ONLY true, skipping submit to CloudFormation"
        cd -
        continue
    fi
    echo "Submiting to CloudFormation with flags: ${_CFN_FLAGS}"
    cfn submit ${_CFN_FLAGS}
    cd -
done
ls -l output



