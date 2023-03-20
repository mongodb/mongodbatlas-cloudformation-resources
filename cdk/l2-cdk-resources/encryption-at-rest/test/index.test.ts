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
import { AtlasEncryptionAtRest } from '../src/index';


const RESOURCE_NAME = 'MongoDB::Atlas::EncryptionAtRest';
const PROFILE = 'testProfile';
const PROJECT_ID= 'testProjectId';
const ROLE_ID = 'roleId';
const REGION = 'region';
const CUSTOMER_MASTER_KEY_ID='customerMasterKeyId';

test('AtlasEncryptionAtRest construct should contain default properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new AtlasEncryptionAtRest(stack, 'testing-stack', {
    projectId: PROJECT_ID,
    roleId: ROLE_ID,
    customerMasterKeyId: CUSTOMER_MASTER_KEY_ID,
    profile: PROFILE,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(RESOURCE_NAME, {
    ProjectId: PROJECT_ID,
    Profile: PROFILE,
    AwsKms: {
      RoleID: ROLE_ID,
      CustomerMasterKeyID: CUSTOMER_MASTER_KEY_ID,
      Enabled: true,
      Region: 'US_EAST_1',
    },
  });
});

test('AtlasEncryptionAtRest construct should contain all the properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new AtlasEncryptionAtRest(stack, 'testing-stack', {
    profile: PROFILE,
    projectId: PROJECT_ID,
    roleId: ROLE_ID,
    customerMasterKeyId: CUSTOMER_MASTER_KEY_ID,
    region: REGION,
    enabled: false,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(RESOURCE_NAME, {
    ProjectId: PROJECT_ID,
    Profile: PROFILE,
    AwsKms: {
      RoleID: ROLE_ID,
      CustomerMasterKeyID: CUSTOMER_MASTER_KEY_ID,
      Enabled: false,
      Region: REGION,
    },
  });
});

test('AtlasEncryptionAtRest construct should thorow exceptions when required params are not provided', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  expect(() => {
    new AtlasEncryptionAtRest(stack, 'testing-stack-no-project-id', {
      profile: PROFILE,
      projectId: '',
      roleId: ROLE_ID,
      customerMasterKeyId: CUSTOMER_MASTER_KEY_ID,
      region: REGION,
      enabled: false,
    });
  }).toThrow('Validation error: projectId is not defined');

  expect(() => {
    new AtlasEncryptionAtRest(stack, 'testing-stack-no-role-id', {
      profile: PROFILE,
      projectId: PROJECT_ID,
      roleId: '',
      customerMasterKeyId: CUSTOMER_MASTER_KEY_ID,
      region: REGION,
      enabled: false,
    });
  }).toThrow('Validation error: roleId is not defined');

  expect(() => {
    new AtlasEncryptionAtRest(stack, 'testing-stack-no-customer-master-key-id', {
      profile: PROFILE,
      projectId: PROJECT_ID,
      roleId: ROLE_ID,
      customerMasterKeyId: '',
      region: REGION,
      enabled: false,
    });
  }).toThrow('Validation error: customerMasterKeyId is not defined');
});
