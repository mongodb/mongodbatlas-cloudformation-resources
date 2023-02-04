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


set -o errexit
set -o nounset
set -o pipefail

temp=$(mktemp)
aws sts get-session-token > "${temp}"
echo "export AWS_SESSION_TOKEN=$(cat "${temp}" | jq -r '.Credentials.SessionToken')"
echo "export AWS_SECRET_ACCESS_KEY=$(cat "${temp}" | jq -r '.Credentials.SecretAccessKey')"
echo "export AWS_ACCESS_KEY_ID=$(cat "${temp}" | jq -r '.Credentials.AccessKeyId')"
rm "${temp}"
