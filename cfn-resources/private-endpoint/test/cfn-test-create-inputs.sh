#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
	echo "Creates a new private endpoint role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

region=$AWS_DEFAULT_REGION
if [ -z "$region" ]; then
	region=$(aws configure get region)
fi

projectName="${1}"
vpcId="${2}"
subnetId="${3}"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"

jq --arg groupId "$projectId" \
	--arg region "$region" \
	--arg vpcId "$vpcId" \
	--arg subnetId "$subnetId" \
	'.GroupId?|=$groupId | .Region?|=$region | .PrivateEndpoints[0].VpcId?|=$vpcId | .PrivateEndpoints[0].SubnetIds[0]?|=$subnetId' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

echo "mongocli iam projects delete ${projectId} --force"
