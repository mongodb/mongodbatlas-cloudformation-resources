# MongoDB::Atlas::StreamProcessor

Returns, adds, edits, and removes Atlas Stream Processors.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::StreamProcessor",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#workspacename" title="WorkspaceName">WorkspaceName</a>" : <i>String</i>,
        "<a href="#processorname" title="ProcessorName">ProcessorName</a>" : <i>String</i>,
        "<a href="#pipeline" title="Pipeline">Pipeline</a>" : <i>String</i>,
        "<a href="#state" title="State">State</a>" : <i>String</i>,
        "<a href="#options" title="Options">Options</a>" : <i><a href="streamsoptions.md">StreamsOptions</a></i>,
        "<a href="#timeouts" title="Timeouts">Timeouts</a>" : <i><a href="timeouts.md">Timeouts</a></i>,
        "<a href="#deleteoncreatetimeout" title="DeleteOnCreateTimeout">DeleteOnCreateTimeout</a>" : <i>Boolean</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::StreamProcessor
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#workspacename" title="WorkspaceName">WorkspaceName</a>: <i>String</i>
    <a href="#processorname" title="ProcessorName">ProcessorName</a>: <i>String</i>
    <a href="#pipeline" title="Pipeline">Pipeline</a>: <i>String</i>
    <a href="#state" title="State">State</a>: <i>String</i>
    <a href="#options" title="Options">Options</a>: <i><a href="streamsoptions.md">StreamsOptions</a></i>
    <a href="#timeouts" title="Timeouts">Timeouts</a>: <i><a href="timeouts.md">Timeouts</a></i>
    <a href="#deleteoncreatetimeout" title="DeleteOnCreateTimeout">DeleteOnCreateTimeout</a>: <i>Boolean</i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project. 

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### InstanceName

Label that identifies the stream processing workspace. This field is deprecated in favor of WorkspaceName. Exactly one of InstanceName or WorkspaceName must be provided.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### WorkspaceName

Label that identifies the stream processing workspace. This is the preferred field name. Exactly one of InstanceName or WorkspaceName must be provided.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProcessorName

Label that identifies the stream processor.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Pipeline

Stream aggregation pipeline you want to apply to your streaming data. This should be a JSON-encoded array of pipeline stages. Refer to MongoDB Atlas Docs for more information on stream aggregation pipelines.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### State

The state of the stream processor. Commonly occurring states are 'CREATED', 'STARTED', 'STOPPED' and 'FAILED'. Used to start or stop the Stream Processor. Valid values are CREATED, STARTED or STOPPED. When a Stream Processor is created without specifying the state, it will default to CREATED state. When a Stream Processor is updated without specifying the state, it will default to the Previous state.

**NOTE** When a Stream Processor is updated without specifying the state, it is stopped and then restored to previous state upon update completion.

_Required_: No

_Type_: String

_Allowed Values_: <code>CREATED</code> | <code>STARTED</code> | <code>STOPPED</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Options

Optional configuration for the stream processor.

_Required_: No

_Type_: <a href="streamsoptions.md">StreamsOptions</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Timeouts

Configurable timeouts for stream processor operations.

_Required_: No

_Type_: <a href="timeouts.md">Timeouts</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DeleteOnCreateTimeout

Indicates whether to delete the resource being created if a timeout is reached when waiting for completion. When set to `true` and timeout occurs, it triggers the deletion and returns immediately without waiting for deletion to complete. When set to `false`, the timeout will not trigger resource deletion. If you suspect a transient error when the value is `true`, wait before retrying to allow resource deletion to finish. Default is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal character string that identifies the stream processor.

#### Stats

The stats associated with the stream processor as a JSON string. Refer to the MongoDB Atlas Docs for more information.

