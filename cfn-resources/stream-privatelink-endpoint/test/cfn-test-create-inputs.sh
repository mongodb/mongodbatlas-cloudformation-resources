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
	echo "Generates test input files for stream private link endpoint"
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
# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

projectId=$(resolve_project_id "${projectName}")
if [ -z "$projectId" ]; then
	echo "Failed to create or find project ${projectName}" >&2
	exit 1
fi
echo -e "Using project \"${projectName}\" with id: ${projectId}\n"

region="${AWS_DEFAULT_REGION:-}"
if [ -z "$region" ]; then
	region=$(aws configure get region 2>/dev/null || echo "")
fi
if [ -z "$region" ]; then
	region="${AWS_REGION:-eu-west-1}"
fi
echo "Using region: ${region}"

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}
	jq --arg projectId "$projectId" \
		--arg profile "$profile" \
		--arg region "$region" \
		'.Profile?|=$profile | .ProjectId?|=$projectId | .Region?|=$region | .ServiceEndpointId?|="com.amazonaws." + $region + ".s3"' \
		"$inputFile" >"../inputs/$outputFile"
done
cd ..
ls -l inputs
