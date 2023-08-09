#!/usr/bin/env bash
# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# project-api-key.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

#set profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ];then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi

if [ ${MONGODB_ATLAS_ORG_ID+x} ];then
    echo "MONGODB_ATLAS_ORG_ID must be set"
    exit 1
fi

#project_id
projectName="${1}"
if [ ${#projectName} -gt 24 ];then
  projectName=${projectName:0:23}
fi

projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

jq --arg projectId "$projectId" \
    --arg orgId "$MONGODB_ATLAS_ORG_ID" \
	  --arg profile "$profile" \
	'.desiredResourceState.ProjectId?|=$projectId | .desiredResourceState.OrgId?|=$orgId | .desiredResourceState.Profile?|=$profile' \
	"$(dirname "$0")/project-api-key.sample-cfn-request.json"
