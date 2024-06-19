# MongoDB::Atlas::PrivateEndpointService

Creates one private endpoint service for the specified cloud service provider. This cloud service provider manages the private endpoint service for the project. When you create a private endpoint service, MongoDB Cloud creates a network container in the project for the cloud provider for which you create the private endpoint service if one doesn't already exist.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::PrivateEndpointService",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#region" title="Region">Region</a>" : <i>String</i>,
        "<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::PrivateEndpointService
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#region" title="Region">Region</a>: <i>String</i>
    <a href="#cloudprovider" title="CloudProvider">CloudProvider</a>: <i>String</i>
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

#### Region

Aws Region

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### CloudProvider

Cloud service provider that manages this private endpoint, default : AWS

_Required_: Yes

_Type_: String

_Allowed Values_: <code>AWS</code> | <code>AZURE</code> | <code>GCP</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The unique identifier of the private endpoint service.

#### EndpointServiceName

Name of the AWS PrivateLink endpoint service. Atlas returns null while it is creating the endpoint service.

#### ErrorMessage

Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.

#### Status

Status of the Atlas PrivateEndpoint service connection

#### InterfaceEndpoints

List of interface endpoint ids associated to the service

