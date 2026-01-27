#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <service_account_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

serviceAccountName="${1}"

#set profile
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

WORDTOREMOVE="template."

cd "$(dirname "$0")" || exit
for inputFile in inputs_*.template.json; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg Name "$serviceAccountName" \
		--arg OrgId "$orgId" \
		--arg profile "$profile" \
		'.Profile?|=$profile | .Name?|=$Name
		| .OrgId?|=$OrgId ' \
		"$inputFile" >"../inputs/$outputFile"
done

cd ..

ls -l inputs
