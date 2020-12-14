# MongoDB::Atlas::DatabaseUser scopeDefinition

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#name" title="name">name</a>" : <i>String</i>,
    "<a href="#type" title="type">type</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#name" title="name">name</a>: <i>String</i>
<a href="#type" title="type">type</a>: <i>String</i>
</pre>

## Properties

#### name

_Required_: No

_Type_: String

_Minimum_: <code>1</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### type

_Required_: No

_Type_: String

_Allowed Values_: <code>CLUSTER</code> | <code>DATA_LAKE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

