# MongoDB::Atlas::CloudProviderAccessSetup

Registers the Atlas side of the AWS Cloud Provider Access trust handshake and returns the Atlas principal ARN, the per-registration external ID, and the role id that the customer's AWS::IAM::Role trust policy and MongoDB::Atlas::CloudProviderAccessAuthorization need.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudProviderAccessSetup",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudProviderAccessSetup
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### RoleId

Unique 24-hexadecimal digit string that identifies the Cloud Provider Access role. Atlas generates it when the role is registered.

#### AtlasAWSAccountArn

Amazon Resource Name that identifies the Atlas AWS account that assumes the customer's IAM role. Use it as the principal of the IAM role trust policy.

#### AtlasAssumedRoleExternalId

External ID that Atlas uses when it assumes the customer's IAM role. The IAM role trust policy must require it. Atlas generates a new value on each registration.

#### CreatedDate

Date and time when the role was registered, in the ISO 8601 timestamp format in UTC.

#### FeatureUsages

Atlas features that use the role. The role cannot be deleted while a feature still uses it.

