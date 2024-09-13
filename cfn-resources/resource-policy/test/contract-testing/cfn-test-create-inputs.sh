#!/usr/bin/env bash

# Run this script with the Makefile
# make create-test-resources
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

if [ -z "${AWS_DEFAULT_REGION+x}" ]; then
	echo "AWS_DEFAULT_REGION must be set"
	exit 1
fi

# setting projectName
orgName="ct-resource-policy-$((1 + RANDOM % 10000))"

./test/cfn-test-create-inputs.sh $orgName
