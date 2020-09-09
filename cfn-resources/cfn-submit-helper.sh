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
resources="${1:-$(ls -F **/mongodb-atlas-*.json | cut -d/ -f1)}"
echo "Submitting the following resources: ${resources}"

for resource in ${resources};
do
    cwd=$(pwd)
    cd "${resource}"
    echo "resource: ${resource}"
    TAGS=logging make
    cfn submit --verbose --set-default
    cd -
done
