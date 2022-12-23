# MongoDB::Atlas::OnlineArchive PartitionFieldView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#fieldname" title="FieldName">FieldName</a>" : <i>String</i>,
    "<a href="#fieldtype" title="FieldType">FieldType</a>" : <i>String</i>,
    "<a href="#order" title="Order">Order</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#fieldname" title="FieldName">FieldName</a>: <i>String</i>
<a href="#fieldtype" title="FieldType">FieldType</a>: <i>String</i>
<a href="#order" title="Order">Order</a>: <i>Integer</i>
</pre>

## Properties

#### FieldName

Human-readable label that identifies the parameter that MongoDB Cloud uses to partition data. To specify a nested parameter, use the dot notation.

_Required_: No

_Type_: String

_Maximum_: <code>700</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FieldType

Data type of the parameter that that MongoDB Cloud uses to partition data. Partition parameters of type [UUID](http://bsonspec.org/spec.html) must be of binary subtype 4. MongoDB Cloud skips partition parameters of type UUID with subtype 3.

_Required_: No

_Type_: String

_Allowed Values_: <code>date</code> | <code>int</code> | <code>long</code> | <code>objectId</code> | <code>string</code> | <code>uuid</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Order

Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero. The value of the **criteria.dateField** parameter defaults as the first item in the partition sequence.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

