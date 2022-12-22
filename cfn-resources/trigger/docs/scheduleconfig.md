# MongoDB::Atlas::Trigger ScheduleConfig

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#schedule" title="Schedule">Schedule</a>" : <i>String</i>,
    "<a href="#skipcatchupevents" title="SkipcatchupEvents">SkipcatchupEvents</a>" : <i>Boolean</i>
}
</pre>

### YAML

<pre>
<a href="#schedule" title="Schedule">Schedule</a>: <i>String</i>
<a href="#skipcatchupevents" title="SkipcatchupEvents">SkipcatchupEvents</a>: <i>Boolean</i>
</pre>

## Properties

#### Schedule

A [cron expression](https://www.mongodb.com/docs/atlas/app-services/triggers/scheduled-triggers/#cron-expressions) that specifies when the trigger executes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SkipcatchupEvents

If `true`, enabling the trigger after it was disabled
will not invoke events that occurred while the trigger
was disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

