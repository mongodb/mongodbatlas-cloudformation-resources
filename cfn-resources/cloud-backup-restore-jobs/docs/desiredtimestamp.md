# MongoDB::Atlas::CloudBackUpRestoreJobs DesiredTimestamp

BSON timestamp that indicates when the checkpoint token entry in the oplog occurred.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#date" title="Date">Date</a>" : <i>String</i>,
    "<a href="#increment" title="Increment">Increment</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#date" title="Date">Date</a>: <i>String</i>
<a href="#increment" title="Increment">Increment</a>: <i>Integer</i>
</pre>

## Properties

#### Date

Date and time when the oplog recorded this database operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Increment

Order of the database operation that the oplog recorded at specific date and time.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

