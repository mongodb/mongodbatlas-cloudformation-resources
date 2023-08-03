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
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#project_id
projectName="${1}"
if [ ${#projectName} -gt 24 ];then
  projectName=${projectName:0:23}
fi

region="${AWS_DEFAULT_REGION}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi


id=$(atlas privateEndpoints aws create --region "${region}" --projectId "${projectId}" --output json | jq -r '.id' )

# WAIT UNTIL CREATES
atlas privateEndpoints aws watch "${id}" --projectId "${projectId}"

#Read the service name once created.
endpointServiceName=$(atlas privateEndpoints aws describe "${id}" --projectId "${projectId}" | jq -r '.endpointServiceName')
echo "endpointServiceName : ${endpointServiceName}"

#Transforming endpointServiceName. eg: com.amazonaws.vpce.us-east-1.vpce-svc-00e311695874992b4 to vpce-00e311695874992b4
endpointId="vpce-${endpointServiceName: (-17)}"
echo "endpointID: ${endpointId}"

jq --arg projectId "$projectId" \
	--arg endpointId "$endpointId" \
	'.ProjectId?|=$projectId | .EndpointId?|=$endpointId' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg projectId "$projectId" \
	--arg endpointId "(*ksadfks)" \
	'.ProjectId?|=$projectId | .EndpointId?|=$endpointId' \
	"$(dirname "$0")/inputs_1_invalid.template.json" >"inputs/inputs_1_invalid.json"

jq --arg projectId "$projectId" \
	--arg endpointId "$endpointId" \
	'.ProjectId?|=$projectId | .EndpointId?|=$endpointId' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
