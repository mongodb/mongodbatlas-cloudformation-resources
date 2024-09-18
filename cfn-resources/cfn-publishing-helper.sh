#!/usr/bin/env bash

# cfn-publishing-helper.sh
#
#
# This tool helps run the AWS CloudFormation cli & api's needed to
# both test and publish the resources for the CloudFormation Public Registry
# You should only run this tool with appropriate AWS account which is linked to the
# Marketplace
#
# There are some options.
#
# TEST_ONLY=true|false
# PUBLISH_ONLY=true|false
# LOG_LEVEL=logrus valid string loglevel
#
# Example with DEBUG logging enabled by default for set of resources:
# LOG_LEVEL=debug ./cfn-publishing-helper.sh project database-user project-ip-access-list cluster network-peering
#
#trap "exit" INT TERM ERR
set -e
#set -o errexit
#set -o nounset
#set -o pipefail

# shellcheck source=/dev/null
. ./cfn-publishing-helper.config
env | grep CFN_PUBLISH_
echo $MONGODB_ATLAS_PROFILE

_DRY_RUN=${DRY_RUN:-false}
# shellcheck disable=SC2034
_CFN_FLAGS=${CFN_FLAGS:---verbose}
_TEST_ONLY=${TEST_ONLY:-false}
_PUBLISH_ONLY=${PUBLISH_ONLY:-false}
# shellcheck disable=SC2034
_DEFAULT_LOG_LEVEL=${LOG_LEVEL:-info}
_CFN_TEST_LOG_BUCKET=${CFN_TEST_LOG_BUCKET:-mongodb-cfn-testing}
version="${2:-00000001}"
profile=$MONGODB_ATLAS_PROFILE

#echo " ******************** version : ${version}"
[[ "${_DRY_RUN}" == "true" ]] && echo "*************** DRY_RUN mode enabled **************"

# Default, find all the directory names with the json custom resource schema files.
resources="${1:-project}"
echo "$(basename "$0") running for the following resources: ${resources}"

echo "Step 1/2: cfn test in the cloud...."
if aws s3api head-bucket --bucket "${_CFN_TEST_LOG_BUCKET}"; then
	echo "found bucket with ${_CFN_TEST_LOG_BUCKET}"
else
	aws s3 mb "s3://${_CFN_TEST_LOG_BUCKET}"
fi

for resource in ${resources}; do
	echo "Working on resource:${resource}"
	[[ "${_DRY_RUN}" == "true" ]] && echo "[dry-run] would have run make on:${resource}" && continue
	if [[ "${_PUBLISH_ONLY}" == "true" ]]; then
		echo "_PUBLISH_ONLY was true, not running 'cfn test' in cloud"
		continue
	fi

	cd "${resource}"
	echo "resource: ${resource}"
	# shellcheck disable=SC2001
	jsonschema="mongodb-atlas-$(echo "${resource}" | sed s/-//g).json"
	# shellcheck disable=SC2002
	res_type=$(cat "${jsonschema}" | jq -r '.typeName')
	echo "res_type=${res_type}"
	version=$(aws cloudformation list-types --output=json | jq --arg typeName "${res_type}" '.TypeSummaries[] | select(.TypeName==$typeName)' | jq -r '.DefaultVersionId')
	echo "version from cfn-publishing-helper=${version}"
	arn=$(aws cloudformation test-type --profile "${profile}" --type RESOURCE --type-name "${res_type}" --log-delivery-bucket "${_CFN_TEST_LOG_BUCKET}" --version-id "${version}" --profile "${profile}" | jq -r '.TypeVersionArn')
	echo "arn from cfn-publishing-helper=${arn}"

	echo "********** Initiated test-type command ***********"
	sleep 10
	echo "Found arn:${arn}"
	# sit and watch the test----
	dt=$(aws cloudformation describe-type --arn "${arn}")
	echo "dt=${dt}"
	# sometime the status is not_tested after triggering the test, so keeping delay
	status=$(echo "${dt}" | jq -r '.TypeTestsStatus')
	if [[ "$status" == "NOT_TESTED" ]]; then
		test_type_resp=$(aws cloudformation test-type --profile "${profile}" --type RESOURCE --type-name "${res_type}" --log-delivery-bucket "${_CFN_TEST_LOG_BUCKET}" --version-id "${version}")
		arn=$(echo "${test_type_resp}" | jq -r '.TypeVersionArn')
		sleep 60
	fi

	while [[ "$status" == "IN_PROGRESS" ]]; do
		sleep 15
		dt=$(aws cloudformation describe-type --arn "${arn}")
		status=$(echo "${dt}" | jq -r '.TypeTestsStatus')
		echo "status=${status}"
	done
	if [[ "${status}" == "FAILED" || "${status}" == "NOT_TESTED" ]]; then
		echo "Test_type STATUS is ${status}"
		exit 1
	fi
	# Fetch the resource type
	cd -
done
if [[ "${_TEST_ONLY}" == "true" ]]; then
	echo "TEST_ONLY true, skipping testing with the CloudFormation CLI"
	exit 0
fi

echo "Step: Running 'publish-type' on ${resources}"
for resource in ${resources}; do
	cd "${resource}"
	[[ "${_DRY_RUN}" == "true" ]] && echo "[dry-run] would have run 'publish-type' for:${resource}" && continue
	jsonschema="mongodb-atlas-${resource//-/}.json"
	type_name=$(jq <"${jsonschema}" -r '.typeName')
	type_info=$(aws cloudformation list-types --output=json | jq --arg typeName "${type_name}" '.TypeSummaries[] | select(.TypeName==$typeName)')
	type_arn=$(echo "${type_info}" | jq -r '.TypeArn')
	echo "resource:${resource}, jsonschema=${jsonschema}, type_name=${type_name}, version=${version}, type_arn=${type_arn}"

	if [ -n "${RESOURCE_VERSION_PUBLISHING}" ]; then
		version_param="--public-version-number ${RESOURCE_VERSION_PUBLISHING}"
	fi
	command="aws cloudformation publish-type --type RESOURCE --arn ${type_arn} ${version_param}"
	echo "${command}"
	${command}

	echo "Deleting role stack as it is not needeed anymore"
	roleStack="mongodb-atlas-${resource//-/}-role-stack"
	command="aws cloudformation update-termination-protection --no-enable-termination-protection --stack-name ${roleStack}"
	echo "${command}"
	${command}
	command="aws cloudformation delete-stack --stack-name ${roleStack}"
	echo "${command}"
	${command}

	cd -
done

echo "Clean up afterwards"
for resource in ${resources}; do
	[[ "${_DRY_RUN}" == "true" ]] && echo "[dry-run] would have run clean-up step for:${resource}" && continue
	echo "running clean-up step for ${resource}"
done
