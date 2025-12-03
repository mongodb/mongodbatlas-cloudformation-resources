#!/usr/bin/env bash

# Run this script with the Makefile
# make create-test-resources
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

# setting projectName
projectName="ct-search-deployment-$(date +%s)"

./test/cfn-test-create-inputs.sh "$projectName"
