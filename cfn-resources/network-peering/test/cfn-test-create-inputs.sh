#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o nounset
set -o pipefail
WORDTOREMOVE="template."
function usage {
	echo "usage:$0 <project_name> <AwsAccountId> <vpc-id>"
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

echo "Check if a project is created $projectId"

echo "Creating network container"

region=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
	region=$(aws configure get region)
fi

region_name_1=$(echo "$region" | tr '[:lower:]' '[:upper:]')
region_name=$(echo "$region_name_1" | sed 's/-/_/g')
public_key=$ATLAS_PUBLIC_KEY
private_key=$ATLAS_PRIVATE_KEY

# Generate a random IPv4 address
IP=10.$(($RANDOM % 256)).$(($RANDOM % 256)).0
CIDR=24

# Generate the CIDR block
cidr_block="$IP/$CIDR"

response=$(curl --user "$public_key:$private_key" --digest --request POST \
  --url https://cloud.mongodb.com/api/atlas/v1.0/groups/$projectId/containers \
  --header "Content-Type: application/json" \
  --data '{
    "providerName": "AWS",
    "atlasCidrBlock": "'"$cidr_block"'",
    "regionName": "'"$region_name"'"
  }')

http_status=$(echo "$response" | jq -r '.error')
http_message=$(echo "$response" | jq -r '.errorCode')

echo "result: $response"
echo "result: $http_status"

if [[ "$http_status" != "null" ]]; then
	echo "FAIL TO CREATE Container HTTP error: $http_status , message: $http_message"
  exit 1
fi

container_id=$(echo $response | jq -r '.id')

echo "Container ID: $container_id"

awsId="${2}"
vpcId="${3}"
nwkConId="${container_id}"
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg awsId "$awsId" \
		--arg vpcId "$vpcId" \
		--arg nwkConId "$nwkConId" \
		'.ProjectId?|=$projectId | .AwsAccountId?|=$awsId | .VpcId|=$vpcId | .ContainerId|=$nwkConId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..

ls -l inputs
#mongocli iam projects delete "${projectId}" --force
