# MongoDB::Atlas::AlertConfiguration

Returns and edits the conditions that trigger alerts and how MongoDB Cloud notifies users. This collection remains under revision and may change. Refer to the legacy documentation for this collection in the following link.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::AlertConfiguration",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#eventtypename" title="EventTypeName">EventTypeName</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#matchers" title="Matchers">Matchers</a>" : <i>[ <a href="matcher.md">Matcher</a>, ... ]</i>,
        "<a href="#metricthreshold" title="MetricThreshold">MetricThreshold</a>" : <i><a href="metricthresholdview.md">MetricThresholdView</a></i>,
        "<a href="#notifications" title="Notifications">Notifications</a>" : <i>[ <a href="notificationview.md">NotificationView</a>, ... ]</i>,
        "<a href="#threshold" title="Threshold">Threshold</a>" : <i><a href="integerthresholdview.md">IntegerThresholdView</a></i>,
        "<a href="#typename" title="TypeName">TypeName</a>" : <i>String</i>,
        "<a href="#severityoverride" title="SeverityOverride">SeverityOverride</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::AlertConfiguration
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#eventtypename" title="EventTypeName">EventTypeName</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#matchers" title="Matchers">Matchers</a>: <i>
      - <a href="matcher.md">Matcher</a></i>
    <a href="#metricthreshold" title="MetricThreshold">MetricThreshold</a>: <i><a href="metricthresholdview.md">MetricThresholdView</a></i>
    <a href="#notifications" title="Notifications">Notifications</a>: <i>
      - <a href="notificationview.md">NotificationView</a></i>
    <a href="#threshold" title="Threshold">Threshold</a>: <i><a href="integerthresholdview.md">IntegerThresholdView</a></i>
    <a href="#typename" title="TypeName">TypeName</a>: <i>String</i>
    <a href="#severityoverride" title="SeverityOverride">SeverityOverride</a>: <i>String</i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### EventTypeName

Event type that triggers an alert.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Matchers

List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster.

_Required_: No

_Type_: List of <a href="matcher.md">Matcher</a>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### MetricThreshold

_Required_: No

_Type_: <a href="metricthresholdview.md">MetricThresholdView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Notifications

List that contains the targets that MongoDB Cloud sends notifications.

_Required_: No

_Type_: List of <a href="notificationview.md">NotificationView</a>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Threshold

_Required_: No

_Type_: <a href="integerthresholdview.md">IntegerThresholdView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TypeName

Human-readable label that displays the alert type.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### SeverityOverride

Degree of seriousness given to this alert. This value overrides the default severity level for the alert.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal digit string that identifies the alert configuration.

#### Enabled

Flag that indicates whether someone enabled this alert configuration for the specified project.

#### Updated

Date and time when someone last updated this alert configuration. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### Created

Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

