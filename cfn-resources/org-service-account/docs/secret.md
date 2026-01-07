# MongoDB::Atlas::OrgServiceAccount Secret

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="Id">Id</a>" : <i>String</i>,
    "<a href="#createdat" title="CreatedAt">CreatedAt</a>" : <i>String</i>,
    "<a href="#expiresat" title="ExpiresAt">ExpiresAt</a>" : <i>String</i>,
    "<a href="#lastusedat" title="LastUsedAt">LastUsedAt</a>" : <i>String</i>,
    "<a href="#maskedsecretvalue" title="MaskedSecretValue">MaskedSecretValue</a>" : <i>String</i>,
    "<a href="#secret" title="Secret">Secret</a>" : <i><a href="secret.md">Secret</a></i>
}
</pre>

### YAML

<pre>
<a href="#id" title="Id">Id</a>: <i>String</i>
<a href="#createdat" title="CreatedAt">CreatedAt</a>: <i>String</i>
<a href="#expiresat" title="ExpiresAt">ExpiresAt</a>: <i>String</i>
<a href="#lastusedat" title="LastUsedAt">LastUsedAt</a>: <i>String</i>
<a href="#maskedsecretvalue" title="MaskedSecretValue">MaskedSecretValue</a>: <i>String</i>
<a href="#secret" title="Secret">Secret</a>: <i><a href="secret.md">Secret</a></i>
</pre>

## Properties

#### Id

Unique identifier of the secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CreatedAt

Date and time that the secret was created on.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ExpiresAt

Date and time when the secret expires.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### LastUsedAt

Date and time when the secret was last used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaskedSecretValue

Masked value of the secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Secret

_Required_: No

_Type_: <a href="secret.md">Secret</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

