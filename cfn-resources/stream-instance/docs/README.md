# MongoDB::Atlas::StreamInstance

Returns, adds, edits, and removes Atlas Stream Processing Instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::StreamInstance",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>" : <i><a href="streamsdataprocessregion.md">StreamsDataProcessRegion</a></i>,
        "<a href="#streamconfig" title="StreamConfig">StreamConfig</a>" : <i><a href="streamconfig.md">StreamConfig</a></i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::StreamInstance
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>: <i><a href="streamsdataprocessregion.md">StreamsDataProcessRegion</a></i>
    <a href="#streamconfig" title="StreamConfig">StreamConfig</a>: <i><a href="streamconfig.md">StreamConfig</a></i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### InstanceName

Human-readable label that identifies the stream connection.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### DataProcessRegion

Information about the cloud provider region in which MongoDB Cloud processes the stream.

_Required_: Yes

_Type_: <a href="streamsdataprocessregion.md">StreamsDataProcessRegion</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StreamConfig

Configuration options for an Atlas Stream Processing Instance.

_Required_: Yes

_Type_: <a href="streamconfig.md">StreamConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal character string that identifies the project.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal character string that identifies the project.

#### Connections

Returns the <code>Connections</code> value.

#### Hostnames

List that contains the hostnames assigned to the stream instance.

