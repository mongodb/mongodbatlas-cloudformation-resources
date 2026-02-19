#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <org_id>"
	echo "Generates test input files for service account secret"
	exit 0
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

orgId="${1}"
echo "OrgId: $orgId"

# Create a service account for testing using Atlas CLI
serviceAccountName="cfn-test-sa-$(date +%s)-$RANDOM"
echo "Creating service account: $serviceAccountName"

# Create service account using Atlas CLI with stdin
serviceAccountJson=$(cat <<EOF | atlas api serviceAccounts createOrgServiceAccount \
	--orgId "$orgId" \
	--version "2024-08-05" \
	--output json
{
  "name": "${serviceAccountName}",
  "description": "CFN test service account",
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

# Generate input files
WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg orgId "$orgId" \
		--arg clientId "$clientId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .OrgId?|=$orgId | .ClientId?|=$clientId' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
