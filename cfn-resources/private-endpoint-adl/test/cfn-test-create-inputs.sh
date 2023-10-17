#!/usr/bin/env bash
# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage: cfn-test-create-inputs.sh <projectId> <endpoint>"
	echo "Creates a new Search Index"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

if ! test -v AWS_DEFAULT_REGION; then
	region=$(aws configure get region)
else
	region=$AWS_DEFAULT_REGION
fi

#Getting Aws vpc and subnet
vpc_id=$(aws ec2 describe-vpcs --query "Vpcs[0].[VpcId]" --output text)
subnet_ids=$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$vpc_id" --output json)
subnet_id=$(echo "$subnet_ids" | jq -r '.Subnets[0].SubnetId')

#creating atlas private endpoint
output=$(atlas privateEndpoints aws list --projectId "${projectId}" --output json)
private_endpoint_id=""
# Check if the output is empty
if [ "$(echo "$output" | jq -e '. | length == 0')" = true ]; then
	echo "Empty"
	# Execute the create command if the output is empty
	create_output=$(atlas privateEndpoints aws create --region "${region}" --projectId "${projectId}" --output json)
	private_endpoint_id=$(echo "$create_output" | jq -r '.id')
	echo "Created endpoint with ID: $private_endpoint_id"

	# Poll and wait for the status to become "AVAILABLE"
	while true; do
		status=$(atlas privateEndpoints aws describe "$private_endpoint_id" --projectId "$projectId" --output json | jq -r '.status')
		if [ "$status" = "AVAILABLE" ]; then
			echo "Status: $status"
			break
		fi
		echo "Status: $status (waiting for AVAILABLE)"
		sleep 5
	done
else
	# Use jq to extract the ID of the first result
	private_endpoint_id=$(echo "$output" | jq -r '.[0].id')
	echo "ID: $private_endpoint_id"
fi

endpoint_service_id=$(atlas privateEndpoints aws describe "$private_endpoint_id" --projectId "$projectId" --output json | jq -r '.endpointServiceName')

#creating aws private endpoint
aws_private_endpoint_id=$(aws ec2 create-vpc-endpoint \
	--vpc-id "$vpc_id" \
	--service-name "$endpoint_service_id" \
	--region "$region" \
	--subnet-ids "$subnet_id" \
	--vpc-endpoint-type Interface \
	--output json | jq -r '.VpcEndpoint.VpcEndpointId')

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg proj "$projectId" \
		--arg endpoint_id "$aws_private_endpoint_id" \
		'.ProjectId?|=$proj |.EndpointId?|=$endpoint_id' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
