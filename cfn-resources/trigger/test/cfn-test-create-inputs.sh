#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage: cfn-test-create-inputs.sh <projectId> <endpoint>"
	echo "Creates a new Search Index"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

project_Id="${2:-625454459c4e6108393d650d}"
db_name="${3:-store}"
coll_name="${4:-sales}"
trigger_name="salestrigger-${RANDOM}"
func_name="${5:-cfn_func}"
func_id="${6:-63862553ac0702272aa701ba}"
service_id="${7:-6387aee08659af5254b0a51e}"
app_id="${7:-638624a5167f5659feb75971}"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
		--arg pvtkey "$ATLAS_PRIVATE_KEY" \
		--arg db "$db_name" \
		--arg coll "$coll_name" \
		--arg funcName "$func_name" \
		--arg funcId "$func_id" \
		--arg appId "$app_id" \
		--arg serviceId "$service_id" \
		--arg projectId "$project_Id" \
		--arg triggerName "$trigger_name" \
		'.AppId?|=$appId |.DatabaseTrigger.ServiceId?|=$serviceId | .ProjectId?|=$projectId |
 .EventProcessors.FUNCTION.FuncConfig.FunctionName?|=$funcName |.EventProcessors.FUNCTION.FuncConfig.FunctionId?|=$funcId|
 .DatabaseTrigger.Collection?|=$coll |.DatabaseTrigger.Database?|=$db | .RealmConfig.PublicKey?|=$pubkey |
.RealmConfig.PrivateKey?|=$pvtkey | .Name?|=$triggerName' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
