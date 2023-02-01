// Generated by cdk-import
import * as cdk from 'aws-cdk-lib';
import * as constructs from 'constructs';

/**
 * The Private Endpoint creation flow consists of the creation of three related resources in the next order: 1. Atlas Private Endpoint Service 2. Aws VPC private Endpoint 3. Atlas Private Endpoint
 *
 * @schema CfnPrivateEndpointProps
 */
export interface CfnPrivateEndpointProps {
  /**
   * Name of the AWS PrivateLink endpoint service. Atlas returns null while it is creating the endpoint service.
   *
   * @schema CfnPrivateEndpointProps#EndpointServiceName
   */
  readonly endpointServiceName?: string;

  /**
   * Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.
   *
   * @schema CfnPrivateEndpointProps#ErrorMessage
   */
  readonly errorMessage?: string;

  /**
   * Status of the Atlas PrivateEndpoint service connection
   *
   * @schema CfnPrivateEndpointProps#Status
   */
  readonly status?: string;

  /**
   * Unique 24-hexadecimal digit string that identifies your project.
   *
   * @schema CfnPrivateEndpointProps#GroupId
   */
  readonly groupId: string;

  /**
   * @schema CfnPrivateEndpointProps#ApiKeys
   */
  readonly apiKeys: ApiKey;

  /**
   * Aws Region
   *
   * @schema CfnPrivateEndpointProps#Region
   */
  readonly region: string;

  /**
   * List of private endpoint associated to the service
   *
   * @schema CfnPrivateEndpointProps#PrivateEndpoints
   */
  readonly privateEndpoints?: PrivateEndpoint[];

}

/**
 * Converts an object of type 'CfnPrivateEndpointProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_CfnPrivateEndpointProps(obj: CfnPrivateEndpointProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'EndpointServiceName': obj.endpointServiceName,
    'ErrorMessage': obj.errorMessage,
    'Status': obj.status,
    'GroupId': obj.groupId,
    'ApiKeys': toJson_ApiKey(obj.apiKeys),
    'Region': obj.region,
    'PrivateEndpoints': obj.privateEndpoints?.map(y => toJson_PrivateEndpoint(y)),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema ApiKey
 */
export interface ApiKey {
  /**
   * @schema ApiKey#PublicKey
   */
  readonly publicKey?: string;

  /**
   * @schema ApiKey#PrivateKey
   */
  readonly privateKey?: string;

}

/**
 * Converts an object of type 'ApiKey' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_ApiKey(obj: ApiKey | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'PublicKey': obj.publicKey,
    'PrivateKey': obj.privateKey,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema PrivateEndpoint
 */
export interface PrivateEndpoint {
  /**
   * String Representing the AWS VPC ID (like: vpc-xxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint)
   *
   * @schema PrivateEndpoint#VpcId
   */
  readonly vpcId?: string;

  /**
   * String Representing the AWS VPC Subnet ID (like: subnet-xxxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint)
   *
   * @schema PrivateEndpoint#SubnetId
   */
  readonly subnetId?: string;

  /**
   * Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.
   *
   * @schema PrivateEndpoint#InterfaceEndpointId
   */
  readonly interfaceEndpointId?: string;

  /**
   * Status of the AWS PrivateEndpoint connection.
   *
   * @schema PrivateEndpoint#AWSPrivateEndpointStatus
   */
  readonly awsPrivateEndpointStatus?: string;

  /**
   * Status of the Atlas PrivateEndpoint connection.
   *
   * @schema PrivateEndpoint#AtlasPrivateEndpointStatus
   */
  readonly atlasPrivateEndpointStatus?: string;

}

/**
 * Converts an object of type 'PrivateEndpoint' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_PrivateEndpoint(obj: PrivateEndpoint | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'VpcId': obj.vpcId,
    'SubnetId': obj.subnetId,
    'InterfaceEndpointId': obj.interfaceEndpointId,
    'AWSPrivateEndpointStatus': obj.awsPrivateEndpointStatus,
    'AtlasPrivateEndpointStatus': obj.atlasPrivateEndpointStatus,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */


/**
 * A CloudFormation `MongoDB::Atlas::PrivateEndpoint`
 *
 * @cloudformationResource MongoDB::Atlas::PrivateEndpoint
 * @stability external
 */
export class CfnPrivateEndpoint extends cdk.CfnResource {
  /**
  * The CloudFormation resource type name for this resource class.
  */
  public static readonly CFN_RESOURCE_TYPE_NAME = 'MongoDB::Atlas::PrivateEndpoint';

  /**
   * Resource props.
   */
  public readonly props: CfnPrivateEndpointProps;

  /**
   * Attribute `MongoDB::Atlas::PrivateEndpoint.Id`
   */
  public readonly attrId: string;

  /**
   * Create a new `MongoDB::Atlas::PrivateEndpoint`.
   *
   * @param scope - scope in which this resource is defined
   * @param id    - scoped id of the resource
   * @param props - resource properties
   */
  constructor(scope: constructs.Construct, id: string, props: CfnPrivateEndpointProps) {
    super(scope, id, { type: CfnPrivateEndpoint.CFN_RESOURCE_TYPE_NAME, properties: toJson_CfnPrivateEndpointProps(props)! });

    this.props = props;

    this.attrId = cdk.Token.asString(this.getAtt('Id'));
  }
}