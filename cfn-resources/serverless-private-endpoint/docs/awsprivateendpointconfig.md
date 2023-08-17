# MongoDB::Atlas::ServerlessPrivateEndpoint awsPrivateEndpointConfig

AWS Private endpoint configuration properties

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#vpcid" title="VpcId">VpcId</a>" : <i>String</i>,
    "<a href="#subnetids" title="SubnetIds">SubnetIds</a>" : <i>[ String, ... ]</i>,
    "<a href="#interfaceendpointid" title="InterfaceEndpointId">InterfaceEndpointId</a>" : <i>String</i>,
    "<a href="#awsprivateendpointstatus" title="AWSPrivateEndpointStatus">AWSPrivateEndpointStatus</a>" : <i>String</i>,
    "<a href="#region" title="Region">Region</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#vpcid" title="VpcId">VpcId</a>: <i>String</i>
<a href="#subnetids" title="SubnetIds">SubnetIds</a>: <i>
      - String</i>
<a href="#interfaceendpointid" title="InterfaceEndpointId">InterfaceEndpointId</a>: <i>String</i>
<a href="#awsprivateendpointstatus" title="AWSPrivateEndpointStatus">AWSPrivateEndpointStatus</a>: <i>String</i>
<a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### VpcId

String Representing the AWS VPC ID (like: vpc-xxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint)

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SubnetIds

List of string representing the AWS VPC Subnet ID (like: subnet-xxxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint)

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InterfaceEndpointId

Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AWSPrivateEndpointStatus

Status of the AWS PrivateEndpoint connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Aws Region

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

