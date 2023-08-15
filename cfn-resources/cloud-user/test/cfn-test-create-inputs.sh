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
    echo "Creates project and tenant for the query limit test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi

# Check MONGODB_ATLAS_ORG_ID is set
if [ -z "${MONGODB_ATLAS_ORG_ID+x}" ];then
  echo "MONGODB_ATLAS_ORG_ID must be set"
  exit 1
fi
orgId="${MONGODB_ATLAS_ORG_ID}"

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

username="govardhan.pagidi@mongodb.com"
mobileNumber="1234567890"
password="CfnTestingPwd"
jq --arg projectId "${projectId}" \
   --arg orgId "${orgId}" \
   --arg profile "$profile" \
   --arg username "$username" \
   --arg mobileNumber "$mobileNumber" \
   --arg password  "${password}"\
   '.Password?|=$password |.Profile?|=$profile | .Roles[1].ProjectId?|=$projectId |.Roles[0].OrgId?|=$orgId | .Username?|=$username | .EmailAddress?|=$username | .MobileNumber?|=$mobileNumber' \
   "$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

