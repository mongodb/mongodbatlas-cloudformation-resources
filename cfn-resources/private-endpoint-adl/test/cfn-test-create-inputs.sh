#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

function usage {
    echo "usage:$2 <projectId> $3<endpoint>"
    echo "Creates a new Search Index"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectId="${2:-625454459c4e6108393d650d}"
endpoint="${3:-vpce-0cd1f01dbd6a3047d}"
WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*;
do
  outputFile=${inputFile//$WORDTOREMOVE/};
 jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
    --arg pvtkey "$ATLAS_PRIVATE_KEY" \
    --arg org  "$projectId" \
    --arg endpoint_id  "$endpoint" \
    '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |.GroupId?|=$org |.EndpointId?|=$endpoint_id' \
     "$inputFile" > "../inputs/$outputFile"
done
cd ..
ls -l inputs
