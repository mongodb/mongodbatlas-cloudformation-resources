# MongoDB::Atlas::AlertConfiguration Matcher

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#fieldname" title="FieldName">FieldName</a>" : <i>String</i>,
    "<a href="#operator" title="Operator">Operator</a>" : <i>String</i>,
    "<a href="#value" title="Value">Value</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#fieldname" title="FieldName">FieldName</a>: <i>String</i>
<a href="#operator" title="Operator">Operator</a>: <i>String</i>
<a href="#value" title="Value">Value</a>: <i>String</i>
</pre>

## Properties

#### FieldName

Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations.

_Required_: No

_Type_: String

_Allowed Values_: <code>CLUSTER_NAME</code> | <code>HOSTNAME</code> | <code>HOSTNAME_AND_PORT</code> | <code>PORT</code> | <code>REPLICA_SET_NAME</code> | <code>SHARD_NAME</code> | <code>TYPE_NAME</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Operator

Comparison operator to apply when checking the current metric value against **matcher[n].value**.

_Required_: No

_Type_: String

_Allowed Values_: <code>EQUALS</code> | <code>CONTAINS</code> | <code>STARTS_WITH</code> | <code>ENDS_WITH</code> | <code>NOT_EQUALS</code> | <code>NOT_CONTAINS</code> | <code>REGEX</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Value

Value to match or exceed using the specified **matchers.operator**.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

