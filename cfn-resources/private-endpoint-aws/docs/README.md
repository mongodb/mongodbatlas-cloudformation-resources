# MongoDB::Atlas::PrivateEndpointAWS

Creates one private endpoint connection for AWS PrivateLink. This resource links an AWS VPC endpoint to a MongoDB Atlas private endpoint service.

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

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### EndpointServiceId

Unique 24-hexadecimal digit string that identifies the Atlas private endpoint service (created with MongoDB::Atlas::PrivateEndpointService) to which you want to connect this VPC endpoint.

_Required_: Yes

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Id

Unique string that identifies the AWS VPC endpoint (interface endpoint) that you created in your VPC. Example: vpce-0d00c26273372c6ef

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### EnforceConnectionSuccess

If set to true, CloudFormation will only return success when the private endpoint connection status is AVAILABLE. If set to false, it returns success once the connection is created regardless of status.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### InterfaceEndpointId

Unique identifier of the interface endpoint.

#### PrivateEndpointConnectionName

Name of the connection for this private endpoint that Atlas generates.

#### PrivateEndpointResourceId

Unique identifier of the private endpoint resource.

#### DeleteRequested

Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.

#### ConnectionStatus

Status of the AWS PrivateLink connection. Returns one of: NONE, PENDING_ACCEPTANCE, PENDING, AVAILABLE, REJECTED, DELETING.

#### ErrorMessage

Error message pertaining to the interface endpoint. Returns null if there are no errors.

