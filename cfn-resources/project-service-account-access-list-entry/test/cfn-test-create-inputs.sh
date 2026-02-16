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
	echo "Generates test input files for project service account access list entry"
	exit 0
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

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

serviceAccountName="cfn-test-sa-accesslist-$(date +%s)-$RANDOM"
echo "Creating project service account: $serviceAccountName"

serviceAccountJson=$(cat <<EOF | atlas api serviceAccounts createGroupServiceAccount \
	--groupId "$projectId" \
	--version "2024-08-05" \
	--output json
{
  "name": "${serviceAccountName}",
  "description": "CFN test project service account for access list",
  "roles": ["GROUP_READ_ONLY"],
  "secretExpiresAfterHours": 8760
}
EOF
)

clientId=$(echo "$serviceAccountJson" | jq -r '.clientId')

if [ -z "$clientId" ] || [ "$clientId" = "null" ]; then
	echo "Error creating service account:"
	echo "$serviceAccountJson"
	exit 1
fi

echo "ClientId: $clientId"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*.template.json; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg clientId "$clientId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .ProjectId?|=$projectId | .ClientId?|=$clientId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
