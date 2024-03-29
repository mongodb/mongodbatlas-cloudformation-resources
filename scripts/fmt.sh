#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

STAGED_GO_FILES=$(git diff --name-only | grep ".go$")

echo "==> Formatting changed go files..."
for FILE in ${STAGED_GO_FILES}; do
	gofmt -w -s "${FILE}"
	goimports -w "${FILE}"
done

echo "==> Formatting changed sh files..."
STAGED_SH_FILES=$(git diff --name-only | grep ".sh$")
for FILE in ${STAGED_SH_FILES}; do
	shfmt -l -w "${FILE}"
done

STAGED_JSON_FILES=$(git diff --name-only | grep ".json$")
for FILE in ${STAGED_JSON_FILES}; do
	prettyFile=$(jq . "${FILE}")
	echo "${prettyFile}" >"${FILE}"
	git add "${FILE}"
done

echo "==> Done..."
