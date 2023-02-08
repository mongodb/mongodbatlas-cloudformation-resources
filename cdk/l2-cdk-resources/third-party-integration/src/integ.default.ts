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

import * as cdk from 'aws-cdk-lib';
import { DatadogIntegration } from './datadogIntegration';
import { MicrosoftTeamsIntegration } from './microsoftTeamsIntegration';
import { PagerDutyIntegration } from './pagerDutyIntegration';
import { ApiKeyDefinition } from './thirdPartyIntegrationBase';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'demo-stack', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});

const apiKeys: ApiKeyDefinition = {
  privateKey: stack.node.tryGetContext('MONGODB_ATLAS_PRIVATE_KEY') || process.env.MONGO_PRIVATE_KEY,
  publicKey: stack.node.tryGetContext('MONGODB_ATLAS_PUBLIC_KEY') || process.env.MONGO_PUBLIC_KEY,
};

// const orgId = stack.node.tryGetContext('MONGO_ORG_ID') || process.env.MONGO_ORG_ID;

new MicrosoftTeamsIntegration(stack, 'MSFT_Integration', {
  apiKeys,
  projectId: 'demoProjectId',
  microsoftTeamsWebhookUrl: 'demoURL',
});

new DatadogIntegration(stack, 'DATADOG_Integration', {
  apiKeys,
  projectId: 'demoProjectId',
  apiKey: 'demoApiKey',
  region: 'US',
});

new PagerDutyIntegration(stack, 'PAGER_DUTY_Integration', {
  apiKeys,
  projectId: 'demoProjectId',
  serviceKey: 'demoKey',
});

