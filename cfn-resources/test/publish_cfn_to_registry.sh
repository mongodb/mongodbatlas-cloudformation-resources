#!/bin/bash
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


# This script performs the following actions:
#
# 1) Generates the resource typename to use for the e2e test. Note: we cannot use the default typeName as it will affect other cfn stacks/tests
# using that resource
# 2) Updates .rpdk-config to use the typename generated at step 1
# 3) Updates the resource schema to use the typename generated at the step 1
# 4) Cleanings: Updates the files changed in the previous steps to the correct typename

set -eu

#t.Setenv("RESOURCE_TYPE_NAME", rctx.ResourceTypeName)
#//	t.Setenv("RESOURCE_TYPE_NAME_FOR_E2E", rctx.ResourceTypeNameForE2e)
#//	t.Setenv("E2E_RAND_SUFFIX", rctx.E2eRandSuffix)
#//	t.Setenv("RESOURCE_DIRECTORY_NAME", rctx.ResourceDirectory)

export RESOURCE_TYPE_NAME=${1}
export RESOURCE_DIRECTORY_NAME=${2}
#
##RANDOM_NUMBER=$((1 + RANDOM % 10000))
##RESOURCE_TYPE_NAME_FOR_E2E="${RESOURCE_TYPE_NAME}${RANDOM_NUMBER}"
#
export E2E_RAND_SUFFIX=$((1 + RANDOM % 10000))
export RESOURCE_TYPE_NAME_FOR_E2E="${RESOURCE_TYPE_NAME}${E2E_RAND_SUFFIX}"

resource_directory=$RESOURCE_DIRECTORY_NAME


echo "Updating .rpdk-config with the E2E resource type $RESOURCE_TYPE_NAME_FOR_E2E"
echo "...... $PWD"
rpdk_file="../$RESOURCE_DIRECTORY_NAME/.rpdk-config"
tmp_rpdk_file="../$RESOURCE_DIRECTORY_NAME/.rpdk-config$E2E_RAND_SUFFIX"
jq --arg type_name "$RESOURCE_TYPE_NAME_FOR_E2E" \
	'.typeName?|=$type_name' \
	"${rpdk_file}" >"${tmp_rpdk_file}"
rm "${rpdk_file}"
mv "${tmp_rpdk_file}" "${rpdk_file}"


echo "Creating a new resource schema"
schema_file_name="${resource_directory//-}"
echo "schema name file........ : ${schema_file_name}"
resource_schema_file="../$RESOURCE_DIRECTORY_NAME/mongodb-atlas-${schema_file_name}.json"
tmp_resource_schema_file="../$RESOURCE_DIRECTORY_NAME/mongodb-atlas-${schema_file_name}$E2E_RAND_SUFFIX.json"
jq --arg type_name "$RESOURCE_TYPE_NAME_FOR_E2E" \
	'.typeName?|=$type_name' \
	"${resource_schema_file}" >"${tmp_resource_schema_file}"


echo "Releasing the resource to private registry $RESOURCE_TYPE_NAME_FOR_E2E"
cd ../"$RESOURCE_DIRECTORY_NAME"
make build && cfn submit --set-default


echo "Reverting .rpdk-config with the original resource typeName $RESOURCE_TYPE_NAME"
rpdk_file=".rpdk-config"
tmp_rpdk_file=".rpdk-config$E2E_RAND_SUFFIX"
jq --arg type_name "$RESOURCE_TYPE_NAME" \
	'.typeName?|=$type_name' \
	"${rpdk_file}" >"${tmp_rpdk_file}"
rm "${rpdk_file}"
mv "${tmp_rpdk_file}" "${rpdk_file}"


echo "Deleting resource JSON schema used for the E2E test"
rm "${tmp_resource_schema_file}"
