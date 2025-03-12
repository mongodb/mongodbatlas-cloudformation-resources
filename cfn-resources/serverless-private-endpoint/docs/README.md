# MongoDB::Atlas::ServerlessPrivateEndpoint

**WARNING:** This resource is deprecated and will be removed in May 2025. If you try to create a ServerlessPrivateEndpoint for a newly created ServerlessInstance, it will fail because the newly created ServerlessInstance are now Flex clusters. For more details, see [Migrate your programmatic tools from M2, M5, or Serverless Instances to Flex Clusters](https://www.mongodb.com/docs/atlas/flex-migration/). Returns, adds, edits, and removes private endpoints for serverless instances. To learn more, see the Atlas Administration API tab on the following tutorial.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ServerlessPrivateEndpoint",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#comment" title="Comment">Comment</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#cloudproviderendpointid" title="CloudProviderEndpointId">CloudProviderEndpointId</a>" : <i>String</i>,
        "<a href="#privateendpointipaddress" title="PrivateEndpointIpAddress">PrivateEndpointIpAddress</a>" : <i>String</i>,
        "<a href="#createandassignawsprivateendpoint" title="CreateAndAssignAWSPrivateEndpoint">CreateAndAssignAWSPrivateEndpoint</a>" : <i>Boolean</i>,
        "<a href="#awsprivateendpointconfigurationproperties" title="AwsPrivateEndpointConfigurationProperties">AwsPrivateEndpointConfigurationProperties</a>" : <i><a href="awsprivateendpointconfig.md">awsPrivateEndpointConfig</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ServerlessPrivateEndpoint
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#comment" title="Comment">Comment</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#cloudproviderendpointid" title="CloudProviderEndpointId">CloudProviderEndpointId</a>: <i>String</i>
    <a href="#privateendpointipaddress" title="PrivateEndpointIpAddress">PrivateEndpointIpAddress</a>: <i>String</i>
    <a href="#createandassignawsprivateendpoint" title="CreateAndAssignAWSPrivateEndpoint">CreateAndAssignAWSPrivateEndpoint</a>: <i>Boolean</i>
    <a href="#awsprivateendpointconfigurationproperties" title="AwsPrivateEndpointConfigurationProperties">AwsPrivateEndpointConfigurationProperties</a>: <i><a href="awsprivateendpointconfig.md">awsPrivateEndpointConfig</a></i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### InstanceName

Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Comment

Human-readable comment associated with the private endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### CloudProviderEndpointId

Unique string that identifies the private endpoint's network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateEndpointIpAddress

IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CreateAndAssignAWSPrivateEndpoint

If true the resource will create the aws private endpoint and assign the Endpoint ID

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsPrivateEndpointConfigurationProperties

AWS Private endpoint configuration properties

_Required_: No

_Type_: <a href="awsprivateendpointconfig.md">awsPrivateEndpointConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal digit string that identifies the private endpoint.

#### EndpointServiceName

Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.

#### ErrorMessage

Human-readable error message that indicates error condition associated with establishing the private endpoint connection.

#### ProviderName

Human-readable error message that indicates error condition associated with establishing the private endpoint connection.

#### Status

Human-readable error message that indicates error condition associated with establishing the private endpoint connection.

#### AwsPrivateEndpointMetaData

Metadata used to track information about the aws private endpoint

