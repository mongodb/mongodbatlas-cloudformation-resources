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

# This script 'de-registers' the test CFN resource type published to AWS private registry

resource_type_name_for_e2e="${RESOURCE_TYPE_NAME_FOR_E2E}"

echo "Cleaning up..."

# Deactivate the CFN resource (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-register.html)
echo "Deactivating the CFN resource ${resource_type_name_for_e2e}"
aws cloudformation deregister-type --type-name "${resource_type_name_for_e2e}" --type RESOURCE
