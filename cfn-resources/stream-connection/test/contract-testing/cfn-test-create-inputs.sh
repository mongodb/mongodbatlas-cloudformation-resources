#!/usr/bin/env bash

# Run this script with the Makefile
# make create-test-resources
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -euo pipefail
set -x

if [ -z "${AWS_DEFAULT_REGION+x}" ]; then
	echo "AWS_DEFAULT_REGION must be set"
	exit 1
fi

# setting projectName
projectName="cfn-stream-conn-$((1 + RANDOM % 10000))"

./test/cfn-test-create-inputs.sh $projectName
