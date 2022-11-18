# MongoDB::Atlas::PrivateEndpoint

The Private Endpoint creation flow consists of the creation of three related resources in the next order.

1. Atlas Private Endpoint Service
2. Aws VPC private Endpoint
3. Atlas Private Endpoint

### CREATION FLOW

The creation process waits until all the resources gets created the next order

1. Creates the Private Endpoint Service and waits until is completed (Available Status)
2. Once the Private Endpoint Service gets completed, creates the Aws VPC Private endpoint and the Atlas Private Endpoint
3. Once the two last resources gets completed (Available Status) the flow returns Success

### SIMPLIFIED CREATION FLOW

![](../docs_files/simplified_creation_flow.png)

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::PrivateEndpoint",
    "Properties" : {
        "<a href="#endpointservicename" title="EndpointServiceName">EndpointServiceName</a>" : <i>String</i>,
        "<a href="#errormessage" title="ErrorMessage">ErrorMessage</a>" : <i>String</i>,
        "<a href="#interfaceendpoints" title="InterfaceEndpoints">InterfaceEndpoints</a>" : <i>[ String, ... ]</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
        "<a href="#vpcid" title="VpcId">VpcId</a>" : <i>String</i>,
        "<a href="#subnetid" title="SubnetId">SubnetId</a>" : <i>String</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikey.md">ApiKey</a></i>,
        "<a href="#region" title="Region">Region</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::PrivateEndpoint
Properties:
    <a href="#endpointservicename" title="EndpointServiceName">EndpointServiceName</a>: <i>String</i>
    <a href="#errormessage" title="ErrorMessage">ErrorMessage</a>: <i>String</i>
    <a href="#interfaceendpoints" title="InterfaceEndpoints">InterfaceEndpoints</a>: <i>
      - String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
    <a href="#vpcid" title="VpcId">VpcId</a>: <i>String</i>
    <a href="#subnetid" title="SubnetId">SubnetId</a>: <i>String</i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikey.md">ApiKey</a></i>
    <a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### EndpointServiceName

Name of the AWS PrivateLink endpoint service. Atlas returns null while it is creating the endpoint service.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ErrorMessage

Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InterfaceEndpoints

Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

Status of the AWS PrivateLink connection. Atlas returns one of the following

_Required_: No

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### VpcId

String Representing the AWS VPC ID (like: vpc-xxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint)

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SubnetId

String Representing the AWS VPC Subnet ID (like: subnet-xxxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint)

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikey.md">ApiKey</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Aws Region

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The unique identifier of the private endpoint service.

