# MongoDB::Atlas::PrivateEndpoint

The Private Endpoint creation flow consists of the creation of three related resources in the next order: 1. Atlas Private Endpoint Service 2. Aws VPC private Endpoint 3. Atlas Private Endpoint

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::PrivateEndpoint",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#endpointservicename" title="EndpointServiceName">EndpointServiceName</a>" : <i>String</i>,
        "<a href="#errormessage" title="ErrorMessage">ErrorMessage</a>" : <i>String</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
        "<a href="#region" title="Region">Region</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::PrivateEndpoint
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#endpointservicename" title="EndpointServiceName">EndpointServiceName</a>: <i>String</i>
    <a href="#errormessage" title="ErrorMessage">ErrorMessage</a>: <i>String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
    <a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

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

#### Status

Status of the Atlas PrivateEndpoint service connection

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

#### Region

Aws Region

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The unique identifier of the private endpoint service.

#### PrivateEndpoints

List of private endpoint associated to the service

#### InterfaceEndpointId

Returns the <code>InterfaceEndpointId</code> value.

