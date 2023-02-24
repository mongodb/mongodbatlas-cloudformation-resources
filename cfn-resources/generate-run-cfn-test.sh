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

resourceFolderName="${1}"
echo "BUILDING resource ${1}.."
cd "$1"
make build
cd ..
echo "GENERATING tests for ${1}, check <resource>/inputs folder.."
./cfn-testing-helper.sh "$1"
cd "$1"
echo "RUNNING tests for ${1} via cfn-test.."
cfn test
