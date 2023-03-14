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

import * as atlas_tpi from '@mongodbatlas-awscdk/third-party-integration';
import { App, Stack } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import * as constants from './constants';
import { DatadogIntegration, DatadogRegion } from '../src/datadogIntegration';
import { MicrosoftTeamsIntegration } from '../src/microsoftTeamsIntegration';
import { PagerDutyIntegration, PagerDutyRegion } from '../src/pagerDutyIntegration';
import * as util from '../src/util';

test('MicrosoftTeamsIntegration construct should be configured with properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new MicrosoftTeamsIntegration(stack, 'testing-stack', {
    apiKeys: {
      privateKey: constants.TEST_PRIVATE_KEY,
      publicKey: constants.TEST_PUBLIC_KEY,
    },
    projectId: constants.TEST_PROJECT_ID,
    microsoftTeamsWebhookUrl: constants.TEST_WEBHOOK_URL,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(constants.THIRD_PARTY_INTEGRATION_RESOURCE_NAME, {
    MicrosoftTeamsWebhookUrl: constants.TEST_WEBHOOK_URL,
    Type: atlas_tpi.CfnThirdPartyIntegrationPropsType.MICROSOFT_TEAMS,
  });
});


test('DatadogIntegration construct should be configured with properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new DatadogIntegration(stack, 'testing-stack', {
    apiKeys: {
      privateKey: constants.TEST_PRIVATE_KEY,
      publicKey: constants.TEST_PUBLIC_KEY,
    },
    projectId: constants.TEST_PROJECT_ID,
    apiKey: constants.TEST_KEY,
    region: DatadogRegion.US,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(constants.THIRD_PARTY_INTEGRATION_RESOURCE_NAME, {
    ApiKey: constants.TEST_KEY,
    Region: DatadogRegion.US,
    Type: atlas_tpi.CfnThirdPartyIntegrationPropsType.DATADOG,
  });
});


test('PagerDutyIntegration construct should be configured with properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new PagerDutyIntegration(stack, 'testing-stack', {
    apiKeys: {
      privateKey: constants.TEST_PRIVATE_KEY,
      publicKey: constants.TEST_PUBLIC_KEY,
    },
    projectId: constants.TEST_PROJECT_ID,
    serviceKey: constants.TEST_KEY,
    region: PagerDutyRegion.EU,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(constants.THIRD_PARTY_INTEGRATION_RESOURCE_NAME, {
    ServiceKey: constants.TEST_KEY,
    Region: PagerDutyRegion.EU,
    Type: atlas_tpi.CfnThirdPartyIntegrationPropsType.PAGER_DUTY,
  });
});

test('PagerDutyIntegration construct should throw exceptions when required params are not provided', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  expect(() => {
    new PagerDutyIntegration(stack, 'testing-stack-no-public-key', {
      apiKeys: {
        privateKey: constants.TEST_PRIVATE_KEY,
        publicKey: '',
      },
      projectId: constants.TEST_PROJECT_ID,
      serviceKey: constants.TEST_KEY,
      region: PagerDutyRegion.EU,
    });
  }).toThrow(util.getPropUndefinedMsg('publicKey'));

  expect(() => {
    new PagerDutyIntegration(stack, 'testing-stack-no-private-key', {
      apiKeys: {
        privateKey: '',
        publicKey: constants.TEST_PUBLIC_KEY,
      },
      projectId: constants.TEST_PROJECT_ID,
      serviceKey: constants.TEST_KEY,
      region: PagerDutyRegion.EU,
    });
  }).toThrow(util.getPropUndefinedMsg('privateKey'));

  expect(() => {
    new PagerDutyIntegration(stack, 'testing-stack-no-project-id', {
      apiKeys: {
        privateKey: constants.TEST_PRIVATE_KEY,
        publicKey: constants.TEST_PUBLIC_KEY,
      },
      projectId: '',
      serviceKey: constants.TEST_KEY,
      region: PagerDutyRegion.EU,
    });
  }).toThrow(util.getPropUndefinedMsg('projectId'));

  expect(() => {
    new PagerDutyIntegration(stack, 'testing-stack-no-service-key', {
      apiKeys: {
        privateKey: constants.TEST_PRIVATE_KEY,
        publicKey: constants.TEST_PUBLIC_KEY,
      },
      projectId: constants.TEST_PROJECT_ID,
      serviceKey: '',
      region: PagerDutyRegion.EU,
    });
  }).toThrow(util.getPropUndefinedMsg('serviceKey'));
});

test('MicrosoftTeamsIntegration construct should throw exceptions when required params are not provided', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  expect(() => {
    new MicrosoftTeamsIntegration(stack, 'testing-stack-no-microsoftTeamsWebhookUrl', {
      apiKeys: {
        privateKey: constants.TEST_PRIVATE_KEY,
        publicKey: constants.TEST_PUBLIC_KEY,
      },
      projectId: constants.TEST_PROJECT_ID,
      microsoftTeamsWebhookUrl: '',
    });
  }).toThrow(util.getPropUndefinedMsg('microsoftTeamsWebhookUrl'));
});

test('DatadogIntegration construct should throw exceptions when required params are not provided', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  expect(() => {
    new DatadogIntegration(stack, 'testing-stack-no-apiKey', {
      apiKeys: {
        privateKey: constants.TEST_PRIVATE_KEY,
        publicKey: constants.TEST_PUBLIC_KEY,
      },
      projectId: constants.TEST_PROJECT_ID,
      apiKey: '',
      region: DatadogRegion.EU,
    });
  }).toThrow(util.getPropUndefinedMsg('apiKey'));
});

