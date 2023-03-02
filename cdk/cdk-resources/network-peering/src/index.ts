// Generated by cdk-import
import * as cdk from 'aws-cdk-lib';
import * as constructs from 'constructs';

/**
 * Returns, adds, edits, and removes network peering containers and peering connections.
 *
 * @schema CfnNetworkPeeringProps
 */
export interface CfnNetworkPeeringProps {
  /**
   * Unique 24-hexadecimal digit string that identifies your project.
   *
   * @schema CfnNetworkPeeringProps#ProjectId
   */
  readonly projectId: string;

  /**
   * Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
   *
   * @schema CfnNetworkPeeringProps#ContainerId
   */
  readonly containerId: string;

  /**
   * Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns null if your VPC and the MongoDB Cloud VPC reside in the same region.
   *
   * @schema CfnNetworkPeeringProps#AccepterRegionName
   */
  readonly accepterRegionName?: string;

  /**
   * Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC.
   *
   * @schema CfnNetworkPeeringProps#AwsAccountId
   */
  readonly awsAccountId?: string;

  /**
   * Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC's subnet that you want to peer with the MongoDB Cloud VPC.
   *
   * @schema CfnNetworkPeeringProps#RouteTableCIDRBlock
   */
  readonly routeTableCidrBlock?: string;

  /**
   * Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC.
   *
   * @schema CfnNetworkPeeringProps#VpcId
   */
  readonly vpcId: string;

  /**
   * The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
   *
   * @schema CfnNetworkPeeringProps#Profile
   */
  readonly profile?: string;

}

/**
 * Converts an object of type 'CfnNetworkPeeringProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_CfnNetworkPeeringProps(obj: CfnNetworkPeeringProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'ProjectId': obj.projectId,
    'ContainerId': obj.containerId,
    'AccepterRegionName': obj.accepterRegionName,
    'AwsAccountId': obj.awsAccountId,
    'RouteTableCIDRBlock': obj.routeTableCidrBlock,
    'VpcId': obj.vpcId,
    'Profile': obj.profile,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */


/**
 * A CloudFormation `MongoDB::Atlas::NetworkPeering`
 *
 * @cloudformationResource MongoDB::Atlas::NetworkPeering
 * @stability external
 */
export class CfnNetworkPeering extends cdk.CfnResource {
  /**
  * The CloudFormation resource type name for this resource class.
  */
  public static readonly CFN_RESOURCE_TYPE_NAME = 'MongoDB::Atlas::NetworkPeering';

  /**
   * Resource props.
   */
  public readonly props: CfnNetworkPeeringProps;

  /**
   * Attribute `MongoDB::Atlas::NetworkPeering.Id`
   */
  public readonly attrId: string;
  /**
   * Attribute `MongoDB::Atlas::NetworkPeering.StatusName`
   */
  public readonly attrStatusName: string;
  /**
   * Attribute `MongoDB::Atlas::NetworkPeering.ErrorStateName`
   */
  public readonly attrErrorStateName: string;
  /**
   * Attribute `MongoDB::Atlas::NetworkPeering.ConnectionId`
   */
  public readonly attrConnectionId: string;

  /**
   * Create a new `MongoDB::Atlas::NetworkPeering`.
   *
   * @param scope - scope in which this resource is defined
   * @param id    - scoped id of the resource
   * @param props - resource properties
   */
  constructor(scope: constructs.Construct, id: string, props: CfnNetworkPeeringProps) {
    super(scope, id, { type: CfnNetworkPeering.CFN_RESOURCE_TYPE_NAME, properties: toJson_CfnNetworkPeeringProps(props)! });

    this.props = props;

    this.attrId = cdk.Token.asString(this.getAtt('Id'));
    this.attrStatusName = cdk.Token.asString(this.getAtt('StatusName'));
    this.attrErrorStateName = cdk.Token.asString(this.getAtt('ErrorStateName'));
    this.attrConnectionId = cdk.Token.asString(this.getAtt('ConnectionId'));
  }
}