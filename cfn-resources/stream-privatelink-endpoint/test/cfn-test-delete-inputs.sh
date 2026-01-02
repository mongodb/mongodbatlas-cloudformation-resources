#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the AWS resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail

# Accept project_name parameter for consistency with cfn-testing-helper.sh
# (not used, but accepted to avoid errors)
projectName="${1:-}"

# Get region from inputs file or use default
region=""
if [ -f "./inputs/inputs_1_create.json" ]; then
	region=$(jq -r '.Region // empty' ./inputs/inputs_1_create.json)
fi

if [ -z "$region" ]; then
	region="${AWS_DEFAULT_REGION:-}"
	if [ -z "$region" ]; then
		region=$(aws configure get region 2>/dev/null || echo "")
	fi
	if [ -z "$region" ]; then
		region="${AWS_REGION:-eu-west-1}"
	fi
fi

echo "Using region: ${region}"

# Delete S3 bucket (static name)
echo -e "--------------------------------delete aws s3 bucket starts ----------------------------\n"
bucketName="mongodb-atlas-stream-test-${region}"

echo "Deleting S3 bucket: ${bucketName}"
aws s3 rb "s3://${bucketName}" --force 2>/dev/null || echo "⚠️  Warning: Failed to delete bucket ${bucketName} (may not exist or already deleted)"

echo -e "--------------------------------delete aws s3 bucket ends ----------------------------\n"
