# MongoDB::Atlas::AlertConfiguration CurrentValue

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#number" title="Number">Number</a>" : <i>Double</i>,
    "<a href="#units" title="Units">Units</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#number" title="Number">Number</a>: <i>Double</i>
<a href="#units" title="Units">Units</a>: <i>String</i>
</pre>

## Properties

#### Number

Amount of the **metricName** recorded at the time of the event. This value triggered the alert.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Units

Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert.

_Required_: No

_Type_: String

_Allowed Values_: <code>BITS</code> | <code>BYTES</code> | <code>DAYS</code> | <code>GIGABITS</code> | <code>GIGABYTES</code> | <code>HOURS</code> | <code>KILOBITS</code> | <code>KILOBYTES</code> | <code>MEGABITS</code> | <code>MEGABYTES</code> | <code>MILLISECONDS</code> | <code>MINUTES</code> | <code>PETABYTES</code> | <code>RAW</code> | <code>SECONDS</code> | <code>TERABYTES</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

