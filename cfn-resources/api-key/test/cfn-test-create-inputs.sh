#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "Creates a template for org apikey creation"
	exit 1
}

if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

# shellcheck source=/dev/null
source "$(dirname "$0")/../../tag-policy.sh"
require_tag_policy_vars

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

if [ -z "${MONGODB_ATLAS_ORG_ID}" ]; then
	echo "MONGODB_ATLAS_ORG_ID must be set"
	exit 1
fi

orgId="${MONGODB_ATLAS_ORG_ID}"

# create ProjectId
projectName="${1}"
if [ ${#projectName} -gt 22 ]; then
	projectName=${projectName:0:21}
fi

# create aws secret key
awsSecretName="mongodb/atlas/apikey/${projectName}"
aws secretsmanager delete-secret --secret-id "${awsSecretName}" --force-delete-without-recovery 2>/dev/null || true
if aws secretsmanager create-secret --name "${awsSecretName}" --secret-string "atlas api-keys goes here" --tags Key=mongodb-owner,Value="${MONGODB_TAG_OWNER}" Key=mongodb-env,Value="${MONGODB_TAG_ENV}"; then
	echo "aws secret created with name : ${awsSecretName}"
else
	echo "aws secret create failed with name : ${awsSecretName}"
	exit 1
fi

project1="${projectName}-1"
project2="${projectName}-2"
project3="${projectName}-3"

# shellcheck source=/dev/null
source "$(dirname "$0")/../../project-resolution.sh"

projectId1=$(resolve_project_id "${project1}")
if [ -z "$projectId1" ]; then
	echo "Failed to create or find project ${project1}" >&2
	exit 1
fi

projectId2=$(resolve_project_id "${project2}")
if [ -z "$projectId2" ]; then
	echo "Failed to create or find project ${project2}" >&2
	exit 1
fi

projectId3=$(resolve_project_id "${project3}")
if [ -z "$projectId3" ]; then
	echo "Failed to create or find project ${project3}" >&2
	exit 1
fi

# Create assigns 2 projects
jq --arg orgId "$orgId" \
	--arg profile "$profile" \
	--arg awsSecretName "$awsSecretName" \
	--arg projectId1 "$projectId1" \
	--arg projectId2 "$projectId2" \
	'.OrgId?|=$orgId | .Profile?|=$profile | .AwsSecretName?|=$awsSecretName |
	 .ProjectAssignments[0].ProjectId?|=$projectId1 |
	 .ProjectAssignments[1].ProjectId?|=$projectId2' \
	"$(dirname "$0")/inputs_1_create.json" >"inputs/inputs_1_create.json"

# Update with un-assign 1, update 1 and assign 1 Project.
jq --arg orgId "$orgId" \
	--arg profile "$profile" \
	--arg projectId2 "$projectId2" \
	--arg projectId3 "$projectId3" \
	'.OrgId?|=$orgId | .Profile?|=$profile |
  	 .ProjectAssignments[0].ProjectId?|=$projectId2 |
  	 .ProjectAssignments[1].ProjectId?|=$projectId3' \
	"$(dirname "$0")/inputs_1_update.json" >"inputs/inputs_1_update.json"

ls -l inputs
