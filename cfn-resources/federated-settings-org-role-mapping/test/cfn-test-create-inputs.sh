#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

# NOTE: You need to set Federation Settings Id and Orgnization Id, in order to execute this resource.
#       You can get the Federation Settings Id on Atlas UI under the 'Manage Federation Settings' cosole

set -o nounset
set -o pipefail
WORDTOREMOVE="template."
function usage {
	echo "usage:$0 <project_name>"
}

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	# Atlas project lists are eventually consistent: when this script follows a project
	# creation (contract testing creates the project first), the lookup above can miss
	# the new project and the create below fails with HTTP 409. Retry the lookup with
	# a bounded delay because the list can still be stale right after the conflict.
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id') || projectId=""
	if [ -z "$projectId" ]; then
		for _ in $(seq 1 12); do
			projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id') || projectId=""
			if [ -n "$projectId" ]; then break; fi
			sleep 5
		done
	fi
	if [ -z "$projectId" ]; then
		echo "Failed to create or find project ${projectName}" >&2
		exit 1
	fi
	echo -e "Using project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg org "$MONGODB_ATLAS_ORG_ID" \
		--arg FederationSettingsId "$MONGODB_ATLAS_FEDERATION_SETTINGS_ID" \
		--arg projectId "$projectId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .RoleAssignments[0].OrgId?|=$org | .RoleAssignments[1].ProjectId?|=$projectId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..

ls -l inputs
