#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#
# Run this script with the Makefile
# make create-test-file
#
set -eu

project_id=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)

if [ -z "${MONGODB_ATLAS_PUBLIC_API_KEY+x}" ] || [ -z "${MONGODB_ATLAS_PRIVATE_API_KEY+x}" ] || [ -z "${MONGODB_ATLAS_ORG_ID+x}" ]; then
	echo "Error: MONGODB_ATLAS_PUBLIC_API_KEY, MONGODB_ATLAS_PRIVATE_API_KEY and MONGODB_ATLAS_ORG_ID environment variables must be set"
	exit 1
fi

#delete project
atlas projects delete "${projectId}" --force
