# MongoDB::Atlas::OnlineArchive PartitionFieldView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#fieldname" title="FieldName">FieldName</a>" : <i>String</i>,
    "<a href="#order" title="Order">Order</a>" : <i>Double</i>
}
</pre>

### YAML

<pre>
<a href="#fieldname" title="FieldName">FieldName</a>: <i>String</i>
<a href="#order" title="Order">Order</a>: <i>Double</i>
</pre>

## Properties

#### FieldName

Human-readable label that identifies the parameter that MongoDB Cloud uses to partition data. To specify a nested parameter, use the dot notation.

_Required_: No

_Type_: String

_Maximum Length_: <code>700</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Order

Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero. The value of the **criteria.dateField** parameter defaults as the first item in the partition sequence.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

