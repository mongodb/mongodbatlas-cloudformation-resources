#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "Creates test inputs for Backup Compliance Policy"
	echo "Usage: $0 <project_name> [authorized_email]"
	echo "Creates a new project (cluster is not required - policy can be created without a cluster)"
}

if [ "$#" -lt 1 ]; then usage; exit 1; fi
if [[ "$*" == help ]]; then usage; exit 0; fi

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

# Get the current user email for authorized email
# Second argument is optional - if not provided, use default test email
authorizedEmail="${2:-test@example.com}"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg authorizedEmail "$authorizedEmail" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ProjectId?|=$projectId |.AuthorizedEmail?|=$authorizedEmail' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..

ls -l inputs
