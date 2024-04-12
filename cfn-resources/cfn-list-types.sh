#!/usr/bin/env bash
set -o errexit

resources="${1:-project}"
regions="${2:-us-east-1 us-west-2 ca-central-1 us-east-2 us-west-1 sa-east-1 ap-southeast-1 ap-southeast-2 ap-southeast-3 ap-south-1 ap-east-1 ap-northeast-1 ap-northeast-2 ap-northeast-3 eu-west-1 eu-central-1 eu-north-1 eu-west-2 eu-west-3 eu-south-1 me-south-1 af-south-1}"

for resource in ${resources}; do
	echo "Step: Running 'list-type' on ${resource}"
	for region in ${regions}; do
		pushd "${resource}"
		# shellcheck disable=SC2001
		jsonschema="mongodb-atlas-$(echo "${resource}" | sed s/-//g).json"
		type_name=$(jq -r '.typeName' "${jsonschema}")
		type_info=$(aws cloudformation --region "${region}" --profile mdb list-types --visibility PUBLIC --output=json | jq --arg typeName "${type_name}" '.TypeSummaries[] | select(.TypeName==$typeName)')
		if [ -z "${type_info}" ]; then
			echo "*********** ${resource} type NOT found in region : ${region} *******************"
		else
			echo "${resource} found in region : ${region}"
			echo "${type_info}"
		fi
		popd
	done
done
