#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

orgName="${1}"
echo "$orgName"
orgId=$(atlas organization list --output json | jq --arg NAME "${organizationName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$orgId" ]; then
	orgId=$(atlas organization create "${orgName}" --output=json | jq -r '.id')

	echo -e "Created organization \"${orgName}\" with id: ${orgId}\n"
else
	echo -e "FOUND project \"${orgName}\" with id: ${orgId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${orgId} --force\n====="

WORDTOREMOVE="template."

policyName="${2}"
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg Name "$policyName" \
		--arg OrgId "$orgId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .Name?|=$policyName
		| .OrgId?|=$orgId ' \
		"$inputFile" >"../inputs/$outputFile"
done

cd ..

ls -l inputs
