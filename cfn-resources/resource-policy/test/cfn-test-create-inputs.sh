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

if [ -z "${MONGODB_ATLAS_PUBLIC_API_KEY+x}" ] || [ -z "${MONGODB_ATLAS_PRIVATE_API_KEY+x}" ]; then
	echo "Error: MONGODB_ATLAS_PUBLIC_API_KEY  and MONGODB_ATLAS_PRIVATE_API_KEY environment variables must be set"
	exit 1
fi

orgId="${MONGODB_ATLAS_ORG_ID}"

WORDTOREMOVE="template."

policyName="${1}"
echo "$policyName"

cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg Name "$policyName" \
		--arg OrgId "$orgId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .Name?|=$Name
		| .OrgId?|=$OrgId ' \
		"$inputFile" >"../inputs/$outputFile"
done

cd ..

ls -l inputs
