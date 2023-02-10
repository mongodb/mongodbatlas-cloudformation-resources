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
  CfnEncryptionAtRest,
} from '@mongodbatlas-awscdk/encryption-at-rest';
import { Construct } from 'constructs';

const US_EAST_1='US_EAST_1';

export interface AtlasEncryptionAtRestProps {
  /**
     * ID of an AWS IAM role authorized to manage an AWS customer master key.
     *
     * @schema AwsKms#RoleID
     */
  readonly roleId: string;
  /**
     * The AWS customer master key used to encrypt and decrypt the MongoDB master keys.
     *
     * @schema AwsKms#CustomerMasterKeyID
     */
  readonly customerMasterKeyId: string;
  /**
     * Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
     *
     * @schema AwsKms#Enabled
     */
  readonly enabled?: boolean;
  /**
     * The AWS region in which the AWS customer master key exists.
     *
     * @schema AwsKms#Region
     */
  readonly region?: string;

  /**
     * Unique identifier of the Atlas project to which the user belongs.
     *
     * @schema CfnEncryptionAtRestProps#ProjectId
     */
  readonly projectId: string;

  /**
     * @schema apiKeyDefinition#PublicKey
     */
  readonly publicKey: string;

  /**
      * @schema apiKeyDefinition#PrivateKey
    */
  readonly privateKey: string;
}


/**
 * It throws an error if any of the required properties are not defined
 * @param {AtlasEncryptionAtRestProps} props - AtlasEncryptionAtRestProps
 */
const validate = (props: AtlasEncryptionAtRestProps) => {
  if (!props.projectId) {throw Error('Validation error: projectId is not defined');}

  if (!props.publicKey) {throw Error('Validation error: publicKey is not defined');}

  if (!props.privateKey) {throw Error('Validation error: privateKey is not defined');}

  if (!props.customerMasterKeyId) {throw Error('Validation error: customerMasterKeyId is not defined');}

  if (!props.roleId) {throw Error('Validation error: roleId is not defined');}
};

/**
 *
 *
 * @export
 * @class AtlasEncryptionAtRest
 * @extends {Construct}
 */
export class AtlasEncryptionAtRest extends Construct {

  /**
   *
   *
   * @type {CfnEncryptionAtRest}
   * @memberof AtlasEncryptionAtRest
   */
  readonly cfnEncryptionAtRest: CfnEncryptionAtRest;

  constructor(scope: Construct, id: string, props: AtlasEncryptionAtRestProps) {
    super(scope, id);

    validate(props);
    this.cfnEncryptionAtRest = new CfnEncryptionAtRest(this,
      'AtlasEncryptionAtRest',
      {
        apiKeys: {
          publicKey: props.publicKey,
          privateKey: props.privateKey,
        },

        awsKms: {
          enabled: props.enabled == undefined ? true: props.enabled,
          region: !props.region ? US_EAST_1 : props.region,
          roleId: props.roleId,
          customerMasterKeyId: props.customerMasterKeyId,
        },

        projectId: props.projectId,
      });
  }
}


