# MongoDB::Atlas::EncryptionAtRest AwsKmsConfig

Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#roleid" title="RoleID">RoleID</a>" : <i>String</i>,
    "<a href="#customermasterkeyid" title="CustomerMasterKeyID">CustomerMasterKeyID</a>" : <i>String</i>,
    "<a href="#enabled" title="Enabled">Enabled</a>" : <i>Boolean</i>,
    "<a href="#region" title="Region">Region</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#roleid" title="RoleID">RoleID</a>: <i>String</i>
<a href="#customermasterkeyid" title="CustomerMasterKeyID">CustomerMasterKeyID</a>: <i>String</i>
<a href="#enabled" title="Enabled">Enabled</a>: <i>Boolean</i>
<a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### RoleID

ID of an AWS IAM role authorized to manage an AWS customer master key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomerMasterKeyID

The AWS customer master key used to encrypt and decrypt the MongoDB master keys.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Enabled

Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

The AWS region in which the AWS customer master key exists.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

