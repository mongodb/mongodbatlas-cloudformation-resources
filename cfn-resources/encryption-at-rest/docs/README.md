# MongoDB::Atlas::EncryptionAtRest

This resource allows administrators to enable, disable, configure, and retrieve the configuration for Encryption at Rest.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::EncryptionAtRest",
    "Properties" : {
        "<a href="#awskms" title="AwsKms">AwsKms</a>" : <i><a href="awskms.md">AwsKms</a></i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::EncryptionAtRest
Properties:
    <a href="#awskms" title="AwsKms">AwsKms</a>: <i><a href="awskms.md">AwsKms</a></i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
</pre>

## Properties

#### AwsKms

Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.

_Required_: Yes

_Type_: <a href="awskms.md">AwsKms</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ProjectId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ProjectId

Unique identifier of the Atlas project to which the user belongs.

