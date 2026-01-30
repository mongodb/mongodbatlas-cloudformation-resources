#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 [federation_settings_id]"
	echo "Generates test input files for federated settings identity provider"
	exit 0
}

if [[ "${1:-}" == "help" ]]; then usage; fi

rm -rf inputs
mkdir inputs

profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

if [ -n "${MONGODB_ATLAS_FEDERATION_SETTINGS_ID:-}" ]; then
	federationSettingsId="${MONGODB_ATLAS_FEDERATION_SETTINGS_ID}"
	echo "Using federation settings ID from environment variable: ${federationSettingsId}"
elif [[ "${1:-}" =~ ^[a-f0-9]{24}$ ]]; then
	federationSettingsId="${1}"
	echo "Using federation settings ID from argument: ${federationSettingsId}"
else
	echo "ERROR: MONGODB_ATLAS_FEDERATION_SETTINGS_ID must be set or a valid 24-char hex ID must be provided as argument"
	exit 1
fi
idpName="cfn-test-idp-$(date +%s)-$RANDOM"
updatedName="${idpName}-updated"
uniqueAudience="cfn-test-audience-$(date +%s)-$RANDOM"

idpId=""
if [ ${MONGODB_ATLAS_IDP_ID+x} ]; then
	echo "idp id set to ${MONGODB_ATLAS_IDP_ID}"
	idpId=${MONGODB_ATLAS_IDP_ID}
fi

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	nameValue="$idpName"
	if [[ "$inputFile" == *"update.template.json" ]]; then
		nameValue="$updatedName"
	fi
	jq --arg profile "$profile" \
		--arg federationSettingsId "$federationSettingsId" \
		--arg name "$nameValue" \
		--arg audience "$uniqueAudience" \
		--arg idpId "$idpId" \
		'.Profile?|=$profile
		| .FederationSettingsId?|=$federationSettingsId
		| .Name?|=$name
		| .Audience?|=$audience
		| (if $idpId != "" then .IdpId?=$idpId else . end)' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
