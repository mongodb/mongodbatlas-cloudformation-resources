#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o nounset
set -o pipefail
WORDTOREMOVE="template."
function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

cd "$(dirname "$0")" || exit
for inputFile in inputs_*;
do
  outputFile=${inputFile//$WORDTOREMOVE/};
  jq --arg pubkey "$MCLI_PUBLIC_API_KEY" \
     --arg pvtkey "$MCLI_PRIVATE_API_KEY" \
	   --arg org "$ATLAS_ORG_ID" \
     --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
     '.FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey ' \
     "$inputFile" > "../inputs/$outputFile"
done
cd ..

ls -l inputs