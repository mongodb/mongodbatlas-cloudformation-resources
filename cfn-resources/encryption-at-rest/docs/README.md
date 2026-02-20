# MongoDB::Atlas::EncryptionAtRest

Returns and edits the Encryption at Rest using Customer Key Management configuration.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::EncryptionAtRest",
    "Properties" : {
        "<a href="#awskmsconfig" title="AwsKmsConfig">AwsKmsConfig</a>" : <i><a href="awskmsconfig.md">AwsKmsConfig</a></i>,
        "<a href="#enabledforsearchnodes" title="EnabledForSearchNodes">EnabledForSearchNodes</a>" : <i>Boolean</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::EncryptionAtRest
Properties:
    <a href="#awskmsconfig" title="AwsKmsConfig">AwsKmsConfig</a>: <i><a href="awskmsconfig.md">AwsKmsConfig</a></i>
    <a href="#enabledforsearchnodes" title="EnabledForSearchNodes">EnabledForSearchNodes</a>: <i>Boolean</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### AwsKmsConfig

Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.

_Required_: Yes

_Type_: <a href="awskmsconfig.md">AwsKmsConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EnabledForSearchNodes

Flag that indicates whether Encryption at Rest for Dedicated Search Nodes is enabled in the specified project.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique identifier of the Atlas project to which the user belongs.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique identifier.

#### Valid

Returns the <code>Valid</code> value.

