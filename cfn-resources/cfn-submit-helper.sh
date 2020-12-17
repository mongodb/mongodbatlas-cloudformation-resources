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

if [[ -d ../output ]]; then
    rm -rf ../output
fi
mkdir -p ../output
for resource in ${resources};
do
    echo "Working on resource:${resource}"
    cwd=$(pwd)
    cd "${resource}"
    echo "resource: ${resource}"
    TAGS=logging make
    cfn submit --verbose --set-default --dry-run
    cp mongodb-atlas-*.zip ../../output/
    cd -
done
ls -l output



