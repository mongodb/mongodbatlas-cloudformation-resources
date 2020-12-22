# MongoDB::Atlas::NetworkPeering

This resource allows to create, read, update and delete a network peering

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::NetworkPeering",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#containerid" title="ContainerId">ContainerId</a>" : <i>String</i>,
        "<a href="#accepterregionname" title="AccepterRegionName">AccepterRegionName</a>" : <i>String</i>,
        "<a href="#awsaccountid" title="AwsAccountId">AwsAccountId</a>" : <i>String</i>,
        "<a href="#providername" title="ProviderName">ProviderName</a>" : <i>String</i>,
        "<a href="#routetablecidrblock" title="RouteTableCIDRBlock">RouteTableCIDRBlock</a>" : <i>String</i>,
        "<a href="#vpcid" title="VpcId">VpcId</a>" : <i>String</i>,
        "<a href="#connectionid" title="ConnectionId">ConnectionId</a>" : <i>String</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::NetworkPeering
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#containerid" title="ContainerId">ContainerId</a>: <i>String</i>
    <a href="#accepterregionname" title="AccepterRegionName">AccepterRegionName</a>: <i>String</i>
    <a href="#awsaccountid" title="AwsAccountId">AwsAccountId</a>: <i>String</i>
    <a href="#providername" title="ProviderName">ProviderName</a>: <i>String</i>
    <a href="#routetablecidrblock" title="RouteTableCIDRBlock">RouteTableCIDRBlock</a>: <i>String</i>
    <a href="#vpcid" title="VpcId">VpcId</a>: <i>String</i>
    <a href="#connectionid" title="ConnectionId">ConnectionId</a>: <i>String</i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
</pre>

## Properties

#### ProjectId

The unique identifier of the project.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ContainerId

Unique identifier of the Atlas VPC container for the AWS region.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AccepterRegionName

AWS region where the peer VPC resides. Returns null if the region is the same region in which the Atlas VPC resides.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsAccountId

AWS account ID of the owner of the peer VPC.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProviderName

The name of the provider

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RouteTableCIDRBlock

Peer VPC CIDR block or subnet.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### VpcId

Unique identifier of the peer VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ConnectionId

Unique identifier for the peering connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique identifier of the Network Peer.

#### StatusName

The VPC peering connection status

#### ErrorStateName

Error state, if any.

