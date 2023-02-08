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

import {
  CfnThirdPartyIntegration,
  CfnThirdPartyIntegrationPropsType,
} from '@mongodbatlas-awscdk/third-party-integration';
import { Construct } from 'constructs';
import { ThirdPartyIntegrationProps } from './thirdPartyIntegrationBase';

export interface MicrosoftTeamsIntegrationProps extends ThirdPartyIntegrationProps {
  readonly microsoftTeamsWebhookUrl: string;
}

export class MicrosoftTeamsIntegration extends Construct {
  constructor(scope: Construct, id: string, props: MicrosoftTeamsIntegrationProps) {
    super(scope, id);

    new CfnThirdPartyIntegration(this,
      'MicrosoftTeamsIntegration',
      {
        ...props,
        type: CfnThirdPartyIntegrationPropsType.MICROSOFT_TEAMS,
      });
  }
}