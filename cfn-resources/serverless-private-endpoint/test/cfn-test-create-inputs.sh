#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
    echo "usage:$0 <project/cluster_name>"
    echo "Creates a new project and cluster by that name for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi

rm -rf inputs
mkdir inputs

if ! test -v AWS_DEFAULT_REGION; then
    region=$(aws configure get region)
else
  region=$AWS_DEFAULT_REGION
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

# Check if the instance already exists
if atlas serverless describe "${projectName}" --projectId "${projectId}"; then
  echo "Serverless found"
else
  echo "Serverless not found, creating..."
  atlas serverless create "$projectName" --provider AWS --region US_EAST_1 --projectId "$projectId"
  atlas serverless watch "${projectName}" --projectId "${projectId}"
  echo -e "Created Serverless \"${projectName}\""
fi

jq --arg region "$region" \
   --arg instanceName "$projectName" \
   --arg projectId "$projectId" \
   --arg profile "$profile" \
   '.Profile?|=$profile | .InstanceName?|=$instanceName | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg region "$region" \
   --arg instanceName "$projectName" \
   --arg projectId "$projectId" \
   --arg profile "$profile" \
   '.Profile?|=$profile | .InstanceName?|=$instanceName | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"
