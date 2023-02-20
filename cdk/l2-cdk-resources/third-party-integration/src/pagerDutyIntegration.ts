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

export enum PagerDutyRegion {
  US = 'US',
  EU = 'EU',
}

export interface PagerDutyIntegrationProps extends ThirdPartyIntegrationProps {
  /**
   * Service key associated with your PagerDuty account.
   */
  readonly serviceKey: string;

  /**
   * PagerDuty region that indicates the API Uniform Resource Locator (URL) to use.
   */
  readonly region: PagerDutyRegion;
}

const validate = (props: PagerDutyIntegrationProps) => {
  util.validate(props);
  if (!props.serviceKey) { throw Error(util.getPropUndefinedMsg('serviceKey')); }
  if (!props.region) { throw Error(util.getPropUndefinedMsg('region')); }
};

export class PagerDutyIntegration extends Construct {
  readonly cfnThirdPartyIntegration: CfnThirdPartyIntegration;

  constructor(scope: Construct, id: string, props: PagerDutyIntegrationProps) {
    super(scope, id);
    validate(props);

    this.cfnThirdPartyIntegration = new CfnThirdPartyIntegration(this,
      'PAGER_DUTY_Integration',
      {
        ...props,
        type: CfnThirdPartyIntegrationPropsType.PAGER_DUTY,
      });
  }
}
