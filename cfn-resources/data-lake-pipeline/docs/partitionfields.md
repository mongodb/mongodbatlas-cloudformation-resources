# MongoDB::Atlas::DataLakePipeline partitionFields

Ordered fields used to physically organize data in the destination.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#fieldname" title="FieldName">FieldName</a>" : <i>String</i>,
    "<a href="#order" title="Order">Order</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#fieldname" title="FieldName">FieldName</a>: <i>String</i>
<a href="#order" title="Order">Order</a>: <i>Integer</i>
</pre>

## Properties

#### FieldName

Human-readable label that identifies the field name used to partition data.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Order

Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

