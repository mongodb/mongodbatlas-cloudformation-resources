# MongoDB::Atlas::MaintenanceWindow

The maintenanceWindow resource provides access to retrieve or update the current Atlas project maintenance window.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::MaintenanceWindow",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#autodeferonceenabled" title="AutoDeferOnceEnabled">AutoDeferOnceEnabled</a>" : <i>Boolean</i>,
        "<a href="#dayofweek" title="DayOfWeek">DayOfWeek</a>" : <i>Integer</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#hourofday" title="HourOfDay">HourOfDay</a>" : <i>Integer</i>,
        "<a href="#startasap" title="StartASAP">StartASAP</a>" : <i>Boolean</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::MaintenanceWindow
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#autodeferonceenabled" title="AutoDeferOnceEnabled">AutoDeferOnceEnabled</a>: <i>Boolean</i>
    <a href="#dayofweek" title="DayOfWeek">DayOfWeek</a>: <i>Integer</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#hourofday" title="HourOfDay">HourOfDay</a>: <i>Integer</i>
    <a href="#startasap" title="StartASAP">StartASAP</a>: <i>Boolean</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### AutoDeferOnceEnabled

Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DayOfWeek

One-based integer that represents the day of the week that the maintenance window starts.

| Value | Day of Week |
|---|---|
| `1` | Sunday |
| `2` | Monday |
| `3` | Tuesday |
| `4` | Wednesday |
| `5` | Thursday |
| `6` | Friday |
| `7` | Saturday |


_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### HourOfDay

Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock. Use `0` for midnight and `12` for noon.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StartASAP

Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request. To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to `false` after MongoDB Cloud completes maintenance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

