# MongoDB::Atlas::StreamInstance

Returns, adds, edits, and removes stream instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::StreamInstance",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>" : <i><a href="streamsdataprocessregion.md">StreamsDataProcessRegion</a></i>,
        "<a href="#streamconfig" title="StreamConfig">StreamConfig</a>" : <i><a href="streamconfig.md">StreamConfig</a></i>,
        "<a href="#connections" title="Connections">Connections</a>" : <i><a href="streamsconnection.md">StreamsConnection</a></i>,
        "<a href="#links" title="Links">Links</a>" : <i>[ <a href="link.md">Link</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::StreamInstance
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>: <i><a href="streamsdataprocessregion.md">StreamsDataProcessRegion</a></i>
    <a href="#streamconfig" title="StreamConfig">StreamConfig</a>: <i><a href="streamconfig.md">StreamConfig</a></i>
    <a href="#connections" title="Connections">Connections</a>: <i><a href="streamsconnection.md">StreamsConnection</a></i>
    <a href="#links" title="Links">Links</a>: <i>
      - <a href="link.md">Link</a></i>
</pre>

## Properties

#### Name

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

#### Connections

Settings that define a connection to an external data store.

_Required_: No

_Type_: <a href="streamsconnection.md">StreamsConnection</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Links

_Required_: No

_Type_: List of <a href="link.md">Link</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal character string that identifies the project.

#### Connections

Settings that define a connection to an external data store.

#### GroupId

Unique 24-hexadecimal character string that identifies the project.

#### Hostnames

List that contains the hostnames assigned to the stream instance.

#### Links

Returns the <code>Links</code> value.

