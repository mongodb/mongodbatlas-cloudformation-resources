# MongoDB::Atlas::PrivateEndpointAWS

Creates one private endpoint for the specified cloud service provider. At this current version only AWS is supported

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::PrivateEndpointAWS",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#endpointserviceid" title="EndpointServiceId">EndpointServiceId</a>" : <i>String</i>,
        "<a href="#id" title="Id">Id</a>" : <i>String</i>,
        "<a href="#enforceconnectionsuccess" title="EnforceConnectionSuccess">EnforceConnectionSuccess</a>" : <i>Boolean</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::PrivateEndpointAWS
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#endpointserviceid" title="EndpointServiceId">EndpointServiceId</a>: <i>String</i>
    <a href="#id" title="Id">Id</a>: <i>String</i>
    <a href="#enforceconnectionsuccess" title="EnforceConnectionSuccess">EnforceConnectionSuccess</a>: <i>Boolean</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### EndpointServiceId

Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Id

Unique string that identifies the private endpoint. for AWS is the VPC endpoint ID, example: vpce-xxxxxxxx

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### EnforceConnectionSuccess

If this proper is set to TRUE, the cloud formation resource will return success Only if the private connection is Succeeded

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### InterfaceEndpointId

Unique identifier of the AWS VPC interface endpoint (e.g., vpce-0d00c26273372c6ef).

#### DeleteRequested

Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.

#### ConnectionStatus

State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request.

#### ErrorMessage

Error message returned when requesting private connection resource. The resource returns null if the request succeeded.

