// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { App, Stack } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { CfnFederatedSettingsOrgRoleMapping } from '../src/index';


const RESOURCE_NAME = 'MongoDB::Atlas::FederatedSettingsOrgRoleMapping';
const PROJECT_ID= '97498392039231';
const ORG_ID= '987498392039';
const PROFILE = 'default';
const FEDERATION_SETTINGS_ID = 'fedSettingsID';
const EXTERNAL_GROUP_NAME='customerExternalGroupName';
const ROLE = 'GROUP_DATA_ACCESS_READ_WRITE';

test('AtlasFederatedSettingsOrgRoleMapping construct should contain default properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new CfnFederatedSettingsOrgRoleMapping(stack, 'testing-stack', {
    profile: PROFILE,
    externalGroupName: EXTERNAL_GROUP_NAME,
    orgId: ORG_ID,
    federationSettingsId: FEDERATION_SETTINGS_ID,
    roleAssignments: [
      {
        role: ROLE,
        projectId: PROJECT_ID,
      },
    ],
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(RESOURCE_NAME, {
    profile: PROFILE,
    externalGroupName: EXTERNAL_GROUP_NAME,
    orgId: ORG_ID,
    federationSettingsId: FEDERATION_SETTINGS_ID,
    roleAssignments: [
      {
        role: ROLE,
        projectId: PROJECT_ID,
      },
    ],
  });
});
