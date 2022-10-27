# MongoDB::Atlas::PrivateEndpointService

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::PrivateEndpointService",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikey.md">ApiKey</a></i>,
        "<a href="#interfaceendpointconnectionid" title="InterfaceEndpointConnectionId">InterfaceEndpointConnectionId</a>" : <i>String</i>,
        "<a href="#connectionstatus" title="ConnectionStatus">ConnectionStatus</a>" : <i>String</i>,
        "<a href="#errormessage" title="ErrorMessage">ErrorMessage</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::PrivateEndpointService
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikey.md">ApiKey</a></i>
    <a href="#interfaceendpointconnectionid" title="InterfaceEndpointConnectionId">InterfaceEndpointConnectionId</a>: <i>String</i>
    <a href="#connectionstatus" title="ConnectionStatus">ConnectionStatus</a>: <i>String</i>
    <a href="#errormessage" title="ErrorMessage">ErrorMessage</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikey.md">ApiKey</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InterfaceEndpointConnectionId

The unique identifier of the private endpoint from AWS

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ConnectionStatus

Status of the interface endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ErrorMessage

Error message pertaining to the interface endpoint. Atlas returns null if there are no errors.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the InterfaceEndpointConnectionId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ApiKeys

Returns the <code>ApiKeys</code> value.

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

#### EndpointServiceId

The unique identifier of the private endpoint.

