#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage: $0"
	echo "  Requires MONGODB_ATLAS_ORG_ID environment variable to be set."
	echo "Generates test input files for service account access list entry"
	exit 0
}

if [[ "${1:-}" == help ]]; then usage; fi

if [ -z "${MONGODB_ATLAS_ORG_ID:-}" ]; then
	echo "Error: MONGODB_ATLAS_ORG_ID environment variable is not set"
	exit 1
fi

orgId="${MONGODB_ATLAS_ORG_ID}"
echo "OrgId: $orgId"

rm -rf inputs
mkdir inputs

#set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

serviceAccountName="cfn-test-sa-accesslist-$(date +%s)-$RANDOM"
echo "Creating service account: $serviceAccountName"

serviceAccountJson=$(cat <<EOF | atlas api serviceAccounts createOrgServiceAccount \
	--orgId "$orgId" \
	--version "2024-08-05" \
	--output json
{
  "name": "${serviceAccountName}",
  "description": "CFN test service account for access list",
  "roles": ["ORG_MEMBER"],
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
	jq --arg orgId "$orgId" \
		--arg clientId "$clientId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .OrgId?|=$orgId | .ClientId?|=$clientId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
