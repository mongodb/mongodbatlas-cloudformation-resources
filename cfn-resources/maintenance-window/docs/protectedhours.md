# MongoDB::Atlas::MaintenanceWindow ProtectedHours

Protected hours during which MongoDB Cloud cannot start maintenance.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#starthourofday" title="StartHourOfDay">StartHourOfDay</a>" : <i>Integer</i>,
    "<a href="#endhourofday" title="EndHourOfDay">EndHourOfDay</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#starthourofday" title="StartHourOfDay">StartHourOfDay</a>: <i>Integer</i>
<a href="#endhourofday" title="EndHourOfDay">EndHourOfDay</a>: <i>Integer</i>
</pre>

## Properties

#### StartHourOfDay

Hour of the day when protected hours start (0-23).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EndHourOfDay

Hour of the day when protected hours end (0-23).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

