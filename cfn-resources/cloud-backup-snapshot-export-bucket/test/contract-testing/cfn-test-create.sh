#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

if [ -z "${AWS_DEFAULT_REGION+x}" ]; then
	echo "AWS_DEFAULT_REGION must be set"
	exit 1
fi

projectName="cfn-export-bucket-$(date +%s)-$RANDOM"

echo "projectName: $projectName"

./test/cfn-test-create-inputs.sh "$projectName"
