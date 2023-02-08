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

import * as atlas_integrations from '@mongodbatlas-awscdk/third-party-integration';
import { Construct } from 'constructs';
import { ThirdPartyIntegrationProps } from './thirdPartyIntegrationBase';

export interface PagerDutyIntegrationProps extends ThirdPartyIntegrationProps{
  readonly serviceKey: string;
}

export class PagerDutyIntegration extends Construct {
  constructor(scope: Construct, id: string, props: PagerDutyIntegrationProps) {
    super(scope, id);

    new atlas_integrations.CfnThirdPartyIntegration(this,
      'PagerDutyThirdPartyIntegration',
      {
        ...props,
        type: atlas_integrations.CfnThirdPartyIntegrationPropsType.PAGER_DUTY,
      });
  }
}