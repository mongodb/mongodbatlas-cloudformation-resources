# MongoDB::Atlas::BackupCompliancePolicy ScheduledPolicyItem

Scheduled backup policy item configuration (hourly, daily, weekly, monthly, or yearly). When provided, FrequencyInterval, RetentionUnit, and RetentionValue are required.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="Id">Id</a>" : <i>String</i>,
    "<a href="#frequencytype" title="FrequencyType">FrequencyType</a>" : <i>String</i>,
    "<a href="#frequencyinterval" title="FrequencyInterval">FrequencyInterval</a>" : <i>Integer</i>,
    "<a href="#retentionunit" title="RetentionUnit">RetentionUnit</a>" : <i>String</i>,
    "<a href="#retentionvalue" title="RetentionValue">RetentionValue</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#id" title="Id">Id</a>: <i>String</i>
<a href="#frequencytype" title="FrequencyType">FrequencyType</a>: <i>String</i>
<a href="#frequencyinterval" title="FrequencyInterval">FrequencyInterval</a>: <i>Integer</i>
<a href="#retentionunit" title="RetentionUnit">RetentionUnit</a>: <i>String</i>
<a href="#retentionvalue" title="RetentionValue">RetentionValue</a>: <i>Integer</i>
</pre>

## Properties

#### Id

Unique identifier of the backup policy item.

_Required_: No

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyType

Frequency associated with the backup policy item. One of the following values: hourly, daily, weekly, monthly, or yearly. This is a read-only value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyInterval

Desired frequency of the new backup policy item specified by frequencyType. Required when the policy item is provided.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionUnit

Unit of time in which MongoDB Cloud measures snapshot retention. Required when the policy item is provided.

_Required_: No

_Type_: String

_Allowed Values_: <code>days</code> | <code>weeks</code> | <code>months</code> | <code>years</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionValue

Duration in days, weeks, months, or years that MongoDB Cloud retains the snapshot. Required when the policy item is provided. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

