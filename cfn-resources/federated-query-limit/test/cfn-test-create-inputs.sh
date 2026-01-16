#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "Creates project and tenant for the query limit test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

#LimitName is an enum from this list. i.e. "bytesProcessed.query" "bytesProcessed.daily" "bytesProcessed.weekly" "bytesProcessed.monthly"

#create federated tenant
tenantName="${projectName}-tenant"
if atlas dataFederation describe "${tenantName}" --projectId "${projectId}"; then
	echo "tenant already exists with name ${tenantName}"
else
	atlas dataFederation create "${tenantName}" --projectId "${projectId}"
	echo "tenant created with name ${tenantName}"
fi

jq --arg projectId "${projectId}" \
	--arg tenantName "${tenantName}" \
	--arg profile "$profile" \
	'.Profile?|=$profile |.ProjectId?|=$projectId | .TenantName?|=$tenantName' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg projectId "${projectId}" \
	--arg tenantName "${tenantName}" \
	--arg profile "$profile" \
	'.Profile?|=$profile |.ProjectId?|=$projectId | .TenantName?|=$tenantName' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

jq --arg projectId "${projectId}" \
	--arg tenantName "${tenantName}" \
	--arg profile "$profile" \
	'.Profile?|=$profile |.ProjectId?|=$projectId | .TenantName?|=$tenantName' \
	"$(dirname "$0")/inputs_2_create.template.json" >"inputs/inputs_2_create.json"

jq --arg projectId "${projectId}" \
	--arg tenantName "${tenantName}" \
	--arg profile "$profile" \
	'.Profile?|=$profile |.ProjectId?|=$projectId | .TenantName?|=$tenantName' \
	"$(dirname "$0")/inputs_2_update.template.json" >"inputs/inputs_2_update.json"
