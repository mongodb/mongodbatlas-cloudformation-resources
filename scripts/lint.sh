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

echo "==> Fixing lint errors..."
for FILE in ${STAGED_GO_FILES}; do
	golangci-lint run --fix --timeout 5m "${FILE}"
done

echo "==> Linting GitHub Actions..."
STAGED_ACTION_FILES=$(git diff --name-only | grep -E "\.github/workflows/.*(\.yaml|\.yml)$")
for FILE in ${STAGED_ACTION_FILES}; do
	actionlint -color -verbose "${FILE}"
done

echo "==> Done..."
