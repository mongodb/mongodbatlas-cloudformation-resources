# MongoDB::Atlas::OnlineArchive ScheduleView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#type" title="Type">Type</a>" : <i>String</i>,
    "<a href="#endhour" title="EndHour">EndHour</a>" : <i>Integer</i>,
    "<a href="#endminute" title="EndMinute">EndMinute</a>" : <i>Integer</i>,
    "<a href="#starthour" title="StartHour">StartHour</a>" : <i>Integer</i>,
    "<a href="#startminute" title="StartMinute">StartMinute</a>" : <i>Integer</i>,
    "<a href="#dayofmonth" title="DayOfMonth">DayOfMonth</a>" : <i>Integer</i>,
    "<a href="#dayofweek" title="DayOfWeek">DayOfWeek</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#type" title="Type">Type</a>: <i>String</i>
<a href="#endhour" title="EndHour">EndHour</a>: <i>Integer</i>
<a href="#endminute" title="EndMinute">EndMinute</a>: <i>Integer</i>
<a href="#starthour" title="StartHour">StartHour</a>: <i>Integer</i>
<a href="#startminute" title="StartMinute">StartMinute</a>: <i>Integer</i>
<a href="#dayofmonth" title="DayOfMonth">DayOfMonth</a>: <i>Integer</i>
<a href="#dayofweek" title="DayOfWeek">DayOfWeek</a>: <i>Integer</i>
</pre>

## Properties

#### Type

_Required_: No

_Type_: String

_Allowed Values_: <code>DAILY</code> | <code>MONTHLY</code> | <code>DEFAULT</code> | <code>WEEKLY</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EndHour

Hour of the day when the scheduled window to run one online archive ends.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EndMinute

Minute of the hour when the scheduled window to run one online archive ends.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StartHour

Hour of the day when the when the scheduled window to run one online archive starts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StartMinute

Minute of the hour when the scheduled window to run one online archive starts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DayOfMonth

Day of the month when the scheduled archive starts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DayOfWeek

Day of the month when the scheduled archive starts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

