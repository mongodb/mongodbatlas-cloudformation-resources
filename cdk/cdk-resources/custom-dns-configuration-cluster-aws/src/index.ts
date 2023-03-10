// Generated by cdk-import
import * as cdk from 'aws-cdk-lib';
import * as constructs from 'constructs';

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 *
 * @schema CfnCustomDnsConfigurationClusterAwsProps
 */
export interface CfnCustomDnsConfigurationClusterAwsProps {
  /**
   * Flag that indicates whether the project's clusters deployed to Amazon Web Services (AWS) use a custom Domain Name System (DNS)
   *
   * @schema CfnCustomDnsConfigurationClusterAwsProps#Enabled
   */
  readonly enabled?: boolean;

  /**
   * Unique 24-hexadecimal digit string that identifies your project.
   *
   * @schema CfnCustomDnsConfigurationClusterAwsProps#ProjectId
   */
  readonly projectId: string;

  /**
   * The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
   *
   * @schema CfnCustomDnsConfigurationClusterAwsProps#Profile
   */
  readonly profile?: string;

}

/**
 * Converts an object of type 'CfnCustomDnsConfigurationClusterAwsProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_CfnCustomDnsConfigurationClusterAwsProps(obj: CfnCustomDnsConfigurationClusterAwsProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'Enabled': obj.enabled,
    'ProjectId': obj.projectId,
    'Profile': obj.profile,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */


/**
 * A CloudFormation `MongoDB::Atlas::CustomDnsConfigurationClusterAws`
 *
 * @cloudformationResource MongoDB::Atlas::CustomDnsConfigurationClusterAws
 * @stability external
 */
export class CfnCustomDnsConfigurationClusterAws extends cdk.CfnResource {
  /**
  * The CloudFormation resource type name for this resource class.
  */
  public static readonly CFN_RESOURCE_TYPE_NAME = 'MongoDB::Atlas::CustomDnsConfigurationClusterAws';

  /**
   * Resource props.
   */
  public readonly props: CfnCustomDnsConfigurationClusterAwsProps;


  /**
   * Create a new `MongoDB::Atlas::CustomDnsConfigurationClusterAws`.
   *
   * @param scope - scope in which this resource is defined
   * @param id    - scoped id of the resource
   * @param props - resource properties
   */
  constructor(scope: constructs.Construct, id: string, props: CfnCustomDnsConfigurationClusterAwsProps) {
    super(scope, id, { type: CfnCustomDnsConfigurationClusterAws.CFN_RESOURCE_TYPE_NAME, properties: toJson_CfnCustomDnsConfigurationClusterAwsProps(props)! });

    this.props = props;

  }
}