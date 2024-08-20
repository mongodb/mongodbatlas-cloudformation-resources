#!/usr/bin/env bash

set -e
set -o nounset

resources="${1:-project}"
otherParams="${2:-}"

if [ -n "${otherParams}" ]; then
	paramKeys=$(echo "$otherParams" | jq -c -r 'keys[]' | tr '\n' ' ')
	echo "Exporting the following keys: ${paramKeys}"
	for param in ${paramKeys}; do
		paramKey="${param}="
		paramValue=$(echo "$otherParams" | jq -c -r --arg key "$param" '.[$key]')
		exportString="$paramKey$paramValue"
		export "${exportString?}"
	done
fi

cloud_publish=${3:-true}

export CLOUD_PUBLISH="${cloud_publish}"

echo "CLOUD_PUBLISH : ${CLOUD_PUBLISH}"

for resource in ${resources}; do
	# shellcheck source=/dev/null
	. ./cfn-testing-helper.config
	env | grep CFN_TEST_

	PROJECT_NAME="${CFN_TEST_NEW_PROJECT_NAME}"
	echo "PROJECT_NAME:${PROJECT_NAME}"

	echo "Started Publishing ${resource} resource"
	echo "Step 1: cfn test"

	if ! ./cfn-testing-helper.sh "${resource}" "${PROJECT_NAME}"; then
		echo "Error in Testing phase"
		exit 1
	fi

	echo "Step 2: cfn submit for ${resource}"

	if ! ./cfn-submit-helper.sh "${resource}"; then
		echo "Error in Submit phase"
		exit 1
	fi

	echo "Step 3: update default version for ${resource}"

	cd "${resource}"
	pwd
	# shellcheck disable=SC2001
	jsonschema="mongodb-atlas-$(echo "${resource}" | sed s/-//g).json"
	res_type=$(jq -r '.typeName' "${jsonschema}")
	echo "${res_type}"
	cd -

	# keeping 5s sleep to get the updated submit version (may not be required)
	sleep 5

	# get latest submit version of the resource
	versionResp=$(aws cloudformation list-type-versions --type RESOURCE --type-name "${res_type}")
	nextPageToken=$(echo "${versionResp}" | jq -r '.NextToken')
	while [ "${nextPageToken}" != "null" ]; do
		versionResp=$(aws cloudformation list-type-versions --type RESOURCE --type-name "${res_type}" --next-token "${nextPageToken}")
		nextPageToken=$(echo "${versionResp}" | jq -r '.NextToken')
	done
	latestVersion=$(echo "${versionResp}" | jq -r '.TypeVersionSummaries[-1].VersionId')

	echo "Setting default version to ${latestVersion} "
	aws cloudformation set-type-default-version --type RESOURCE --type-name "${res_type}" --version-id "${latestVersion}"

	echo "Step 4:  Publishing  ${resource}"

	if ! ./cfn-publishing-helper.sh "${resource}" "${latestVersion}"; then
		echo "Error in Publishing phase"
		exit 1
	fi

	#delete the input params
	cd "${resource}"
	./test/cfn-test-delete-inputs.sh "${PROJECT_NAME}-${resource}" && echo "resource:${resource} inputs delete OK" || echo "resource:${resource} input delete FAILED"

	#    Deleting the projects
	echo "Looking up Atlas project id for resource:${resource} project name:${PROJECT_NAME}-${resource}"
	p_id=$(atlas project list --output=json | jq --arg name "${PROJECT_NAME}-${resource}" -r '.results[] | select(.name==$name) | .id')
	[ -z "$p_id" ] && echo "No project found" && continue
	p_name=$(atlas project list --output=json | jq --arg name "${PROJECT_NAME}-${resource}" -r '.results[] | select(.name==$name) | .name')
	echo "Cleaning up for resource:${resource}, project:${p_name} id:${p_id}"

	if atlas project delete "${p_id}" --force; then
		echo "Cleaned up project:${p_name} id:${p_id}"
	else
		echo "Failed cleaning up project:${p_id}"
		exit 1
	fi

	echo "******** Successfully published ${resource} *************"
done
