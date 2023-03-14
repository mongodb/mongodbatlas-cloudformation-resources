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
import { CfnTrigger } from '../src/index';


const RESOURCE_NAME = 'MongoDB::Atlas::Trigger';
const PROJECT_ID= 'testProjectId';
const APPID= 'appId';

test('Trigger construct should contain default properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new CfnTrigger(stack, 'testing-stack', {
    projectId: PROJECT_ID,
    appId: APPID,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(RESOURCE_NAME, {
    ProjectId: PROJECT_ID,
    AppId: APPID,
  });
});
