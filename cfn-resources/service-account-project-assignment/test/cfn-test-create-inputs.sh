#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
	echo "Generates test input files for service account project assignment"
	exit 0
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectName="${1}"

#set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

if [ -z "${MONGODB_ATLAS_ORG_ID+x}" ]; then
	echo "MONGODB_ATLAS_ORG_ID is not set, exiting..."
	exit 1
fi

orgId="${MONGODB_ATLAS_ORG_ID}"

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

# Create a service account for testing
serviceAccountName="cfn-test-sa-$(date +%s)-$RANDOM"
echo "Creating service account: $serviceAccountName"

# Create service account using Atlas API CLI
serviceAccountResponse=$(echo "{\"name\":\"${serviceAccountName}\",\"description\":\"Test service account for CFN contract testing\",\"roles\":[\"ORG_MEMBER\"],\"secretExpiresAfterHours\":8760}" | \
	atlas api serviceAccounts createOrgServiceAccount \
	--orgId "${orgId}" \
	2>/dev/null)

clientId=$(echo "$serviceAccountResponse" | jq -r '.clientId')
echo "Service Account ClientId: $clientId"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg orgId "$orgId" \
		--arg clientId "$clientId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .OrgId?|=$orgId | .ProjectId?|=$projectId | .ClientId?|=$clientId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
