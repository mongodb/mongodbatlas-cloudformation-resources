#!/bin/bash

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

# Run this script with the Makefile
# make e2e-test
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -eu

RANDOM_NUMBER=$((1 + RANDOM % 10000))
RESOURCE_TYPE_NAME="MongoDB::Atlas::Project"
RESOURCE_TYPE_NAME_FOR_E2E="${RESOURCE_TYPE_NAME}${RANDOM_NUMBER}"

echo "Update .rpdk-config with the E2E resource type ${RESOURCE_TYPE_NAME_FOR_E2E}"
jq --arg type_name "${RESOURCE_TYPE_NAME_FOR_E2E}" \
	'.typeName?|=$type_name' \
	".rpdk-config" >".rpdk-config${RESOURCE_TYPE_NAME_FOR_E2E}"
rm ".rpdk-config"
mv ".rpdk-config${RESOURCE_TYPE_NAME_FOR_E2E}" ".rpdk-config"

echo "Create a new resource JSON schema"
jq --arg type_name "${RESOURCE_TYPE_NAME_FOR_E2E}" \
	'.typeName?|=$type_name' \
	"mongodb-atlas-project.json" >"mongodb-atlas-project${RANDOM_NUMBER}.json"

echo "Update the template to be use for the E2E test"
jq --arg type_name "${RESOURCE_TYPE_NAME_FOR_E2E}" \
	'.Resources.Project.Type?|=$type_name' \
	"test/e2e/template/project.json" >"test/e2e/template/project${RESOURCE_TYPE_NAME_FOR_E2E}.json"
rm "test/e2e/template/project.json"
mv "test/e2e/template/project${RESOURCE_TYPE_NAME_FOR_E2E}.json" "test/e2e/template/project.json"

echo "Release the resource ${RESOURCE_TYPE_NAME_FOR_E2E}"
make release

# Run e2e test
# go test ...

echo "Cleaning..."

echo "Revert changes to the template"
jq --arg type_name "${RESOURCE_TYPE_NAME}" \
	'.Resources.Project.Type?|=$type_name' \
	"test/e2e/template/project.json" >"test/e2e/template/project${RESOURCE_TYPE_NAME}.json"
rm "test/e2e/template/project.json"
mv "test/e2e/template/project${RESOURCE_TYPE_NAME}.json" "test/e2e/template/project.json"

echo "Update .rpdk-config with the original resource typeName ${RESOURCE_TYPE_NAME}"
jq --arg type_name "${RESOURCE_TYPE_NAME}" \
	'.typeName?|=$type_name' \
	".rpdk-config" >".rpdk-config${RESOURCE_TYPE_NAME}"
rm ".rpdk-config"
mv ".rpdk-config${RESOURCE_TYPE_NAME}" ".rpdk-config"

echo "Delete resource JSON schema used for the E2E test"
rm "mongodb-atlas-project${RANDOM_NUMBER}.json"

echo "Deactivate the CFN resource ${RESOURCE_TYPE_NAME_FOR_E2E}"
# Deactivate the CFN resource (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-register.html)
aws cloudformation deregister-type --type-name "${RESOURCE_TYPE_NAME_FOR_E2E}" --type RESOURCE
