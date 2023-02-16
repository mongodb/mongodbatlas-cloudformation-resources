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
import * as util from './util';

export enum DatadogRegion {
  US = 'US',
  EU = 'EU',
  US3 = 'US3',
  US5 = 'US5'
}

export interface DatadogIntegrationProps extends ThirdPartyIntegrationProps {
  /**
   * Key that allows MongoDB Cloud to access your Datadog account.
   */
  readonly apiKey: string;

  /**
   * Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.
   */
  readonly region: DatadogRegion;
}

const validate = (props: DatadogIntegrationProps) => {
  util.validate(props);
  if (!props.apiKey) { throw Error(util.getPropUndefinedMsg('apiKey')); }
  if (!props.region) { throw Error(util.getPropUndefinedMsg('region')); }
};

export class DatadogIntegration extends Construct {
  readonly cfnThirdPartyIntegration: CfnThirdPartyIntegration;

  constructor(scope: Construct, id: string, props: DatadogIntegrationProps) {
    super(scope, id);
    validate(props);

    this.cfnThirdPartyIntegration = new CfnThirdPartyIntegration(this,
      'DATADOG_Integration',
      {
        ...props,
        type: CfnThirdPartyIntegrationPropsType.DATADOG,
      });
  }
}
