# MongoDB::Atlas::AlertConfiguration IntegerThresholdView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#operator" title="Operator">Operator</a>" : <i>String</i>,
    "<a href="#threshold" title="Threshold">Threshold</a>" : <i>Double</i>,
    "<a href="#units" title="Units">Units</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#operator" title="Operator">Operator</a>: <i>String</i>
<a href="#threshold" title="Threshold">Threshold</a>: <i>Double</i>
<a href="#units" title="Units">Units</a>: <i>String</i>
</pre>

## Properties

#### Operator

Comparison operator to apply when checking the current metric value.

_Required_: No

_Type_: String

_Allowed Values_: <code>GREATER_THAN</code> | <code>LESS_THAN</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Threshold

Value of metric that, when exceeded, triggers an alert.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Units

Element used to express the quantity. This can be an element of time, storage capacity, and the like.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

