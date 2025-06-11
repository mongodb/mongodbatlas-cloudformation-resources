# MongoDB::Atlas::Trigger

View and manage your application's triggers: https://www.mongodb.com/docs/atlas/app-services/triggers/

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Trigger",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#databasetrigger" title="DatabaseTrigger">DatabaseTrigger</a>" : <i><a href="databaseconfig.md">DatabaseConfig</a></i>,
        "<a href="#authtrigger" title="AuthTrigger">AuthTrigger</a>" : <i><a href="authconfig.md">AuthConfig</a></i>,
        "<a href="#scheduletrigger" title="ScheduleTrigger">ScheduleTrigger</a>" : <i><a href="scheduleconfig.md">ScheduleConfig</a></i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#type" title="Type">Type</a>" : <i>String</i>,
        "<a href="#disabled" title="Disabled">Disabled</a>" : <i>Boolean</i>,
        "<a href="#functionid" title="FunctionId">FunctionId</a>" : <i>String</i>,
        "<a href="#functionname" title="FunctionName">FunctionName</a>" : <i>String</i>,
        "<a href="#eventprocessors" title="EventProcessors">EventProcessors</a>" : <i><a href="event.md">Event</a></i>,
        "<a href="#appid" title="AppId">AppId</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Trigger
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#databasetrigger" title="DatabaseTrigger">DatabaseTrigger</a>: <i><a href="databaseconfig.md">DatabaseConfig</a></i>
    <a href="#authtrigger" title="AuthTrigger">AuthTrigger</a>: <i><a href="authconfig.md">AuthConfig</a></i>
    <a href="#scheduletrigger" title="ScheduleTrigger">ScheduleTrigger</a>: <i><a href="scheduleconfig.md">ScheduleConfig</a></i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#type" title="Type">Type</a>: <i>String</i>
    <a href="#disabled" title="Disabled">Disabled</a>: <i>Boolean</i>
    <a href="#functionid" title="FunctionId">FunctionId</a>: <i>String</i>
    <a href="#functionname" title="FunctionName">FunctionName</a>: <i>String</i>
    <a href="#eventprocessors" title="EventProcessors">EventProcessors</a>: <i><a href="event.md">Event</a></i>
    <a href="#appid" title="AppId">AppId</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### DatabaseTrigger

_Required_: No

_Type_: <a href="databaseconfig.md">DatabaseConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthTrigger

_Required_: No

_Type_: <a href="authconfig.md">AuthConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ScheduleTrigger

_Required_: No

_Type_: <a href="scheduleconfig.md">ScheduleConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

The trigger's name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Type

The trigger's type.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Disabled

If `true`, the trigger is disabled and does not listen for events or execute.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FunctionId

The ID of the function that the trigger calls when it fires.

This value is the same as `event_processors.FUNCTION.function_id`.
You can either define the value here or in `event_processors.FUNCTION.function_id`.
The App Services backend duplicates the value to the configuration location where you did not define it.

For example, if you define `function_id`, the backend duplicates it to `event_processors.FUNCTION.function_id`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FunctionName

The name of the function that the trigger calls when it
fires, i.e. the function described by `function_id`.

This value is the same as `event_processors.FUNCTION.function_name`.
You can either define the value here or in `event_processors.FUNCTION.function_name`.
The App Services backend duplicates the value to the configuration location where you did not define it.

For example, if you define `function_name`, the backend duplicates it to `event_processors.FUNCTION.function_name`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EventProcessors

_Required_: No

_Type_: <a href="event.md">Event</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AppId

App Services Application ID

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Project Id for application services

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The trigger's unique ID.

