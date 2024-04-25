# MongoDB::Atlas::CloudBackupSchedule ApiPolicyItemView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="ID">ID</a>" : <i>String</i>,
    "<a href="#frequencytype" title="FrequencyType">FrequencyType</a>" : <i>String</i>,
    "<a href="#frequencyinterval" title="FrequencyInterval">FrequencyInterval</a>" : <i>Integer</i>,
    "<a href="#retentionvalue" title="RetentionValue">RetentionValue</a>" : <i>Integer</i>,
    "<a href="#retentionunit" title="RetentionUnit">RetentionUnit</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#id" title="ID">ID</a>: <i>String</i>
<a href="#frequencytype" title="FrequencyType">FrequencyType</a>: <i>String</i>
<a href="#frequencyinterval" title="FrequencyInterval">FrequencyInterval</a>: <i>Integer</i>
<a href="#retentionvalue" title="RetentionValue">RetentionValue</a>: <i>Integer</i>
<a href="#retentionunit" title="RetentionUnit">RetentionUnit</a>: <i>String</i>
</pre>

## Properties

#### ID

Unique identifier of the backup policy item.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyType

Frequency associated with the backup policy item. One of the following values: hourly, daily, weekly, monthly or yearly.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyInterval

Desired frequency of the new backup policy item specified by frequencyType.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionValue

Duration for which the backup is kept. Associated with retentionUnit.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionUnit

Metric of duration of the backup policy item: days, weeks, months or years.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

