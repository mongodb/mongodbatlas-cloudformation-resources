#!/usr/bin/env bash

# Cloud Tag Policy: values are required when the test scripts create AWS resources.
require_tag_policy_vars() {
	if [ -z "${MONGODB_TAG_OWNER:-}" ] || [ -z "${MONGODB_TAG_ENV:-}" ]; then
		echo "MONGODB_TAG_OWNER and MONGODB_TAG_ENV must be set (Cloud Tag Policy)" >&2
		exit 1
	fi
}

# Creates an S3 bucket with the Cloud Tag Policy tags in the create request, which
# creation-time enforcement requires. LocationConstraint must be omitted for us-east-1.
create_tagged_bucket() {
	local bucket="$1" region="$2"
	local tags="Tags=[{Key=mongodb-owner,Value=${MONGODB_TAG_OWNER}},{Key=mongodb-env,Value=${MONGODB_TAG_ENV}}]"
	if [ "${region}" = "us-east-1" ]; then
		aws s3api create-bucket --bucket "${bucket}" --region "${region}" --create-bucket-configuration "${tags}"
	else
		aws s3api create-bucket --bucket "${bucket}" --region "${region}" --create-bucket-configuration "LocationConstraint=${region},${tags}"
	fi
}
