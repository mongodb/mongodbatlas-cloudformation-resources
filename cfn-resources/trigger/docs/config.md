# MongoDB::Atlas::Trigger Config

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#operationtype" title="OperationType">OperationType</a>" : <i>String</i>,
    "<a href="#providers" title="Providers">Providers</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#operationtype" title="OperationType">OperationType</a>: <i>String</i>
<a href="#providers" title="Providers">Providers</a>: <i>
      - String</i>
</pre>

## Properties

#### OperationType

The type of authentication event that the trigger listens for.

_Required_: Yes

_Type_: String

_Allowed Values_: <code>LOGIN</code> | <code>CREATE</code> | <code>DELETE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Providers

The type(s) of authentication provider that the trigger listens to.

_Required_: Yes

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

