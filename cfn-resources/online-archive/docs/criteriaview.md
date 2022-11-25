# MongoDB::Atlas::onlinearchives CriteriaView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#type" title="Type">Type</a>" : <i>String</i>,
    "<a href="#datefield" title="DateField">DateField</a>" : <i>String</i>,
    "<a href="#dateformat" title="DateFormat">DateFormat</a>" : <i>String</i>,
    "<a href="#expireafterdays" title="ExpireAfterDays">ExpireAfterDays</a>" : <i>Integer</i>,
    "<a href="#query" title="Query">Query</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#type" title="Type">Type</a>: <i>String</i>
<a href="#datefield" title="DateField">DateField</a>: <i>String</i>
<a href="#dateformat" title="DateFormat">DateFormat</a>: <i>String</i>
<a href="#expireafterdays" title="ExpireAfterDays">ExpireAfterDays</a>: <i>Integer</i>
<a href="#query" title="Query">Query</a>: <i>String</i>
</pre>

## Properties

#### Type

Means by which MongoDB Cloud selects data to archive. Data can be chosen using the age of the data or a MongoDB query.
**DATE** selects documents to archive based on a date.
**CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `"collectionType": "TIMESERIES"`.

_Required_: No

_Type_: String

_Allowed Values_: <code>DATE</code> | <code>CUSTOM</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DateField

Indexed database parameter that stores the date that determines when data moves to the online archive. MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the expireAfterDays parameter. Set this parameter when you set "criteria.type" : "DATE".

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DateFormat

Syntax used to write the date after which data moves to the online archive. Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when "criteria.type" : "DATE". You must set "criteria.type" : "DATE" if "collectionType": "TIMESERIES".

_Required_: No

_Type_: String

_Allowed Values_: <code>ISODATE</code> | <code>EPOCH_SECONDS</code> | <code>EPOCH_MILLIS</code> | <code>EPOCH_NANOSECONDS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ExpireAfterDays

Number of days after the value in the criteria.dateField when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set "criteria.type" : "DATE".

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Query

MongoDB find query that selects documents to archive. The specified query follows the syntax of the db.collection.find(query) command. This query can't use the empty document ({}) to return all documents. Set this parameter when "criteria.type" : "CUSTOM".

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

