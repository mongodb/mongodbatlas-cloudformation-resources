# MongoDB::Atlas::DataLakePipeline

Data Lake is deprecated. As of September 2024, Data Lake is deprecated and will reach end-of-life. It will be removed on September 30, 2025. If you use Data Lake, you should migrate to alternative solutions before the service is removed. To learn more, see <https://dochub.mongodb.org/core/data-lake-deprecation>. Returns, adds, edits, and removes data lake pipelines.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::DataLakePipeline",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#state" title="State">State</a>" : <i>String</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#sink" title="Sink">Sink</a>" : <i><a href="sink.md">sink</a></i>,
        "<a href="#source" title="Source">Source</a>" : <i><a href="source.md">source</a></i>,
        "<a href="#transformations" title="Transformations">Transformations</a>" : <i>[ <a href="transformations.md">transformations</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::DataLakePipeline
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#state" title="State">State</a>: <i>String</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#sink" title="Sink">Sink</a>: <i><a href="sink.md">sink</a></i>
    <a href="#source" title="Source">Source</a>: <i><a href="source.md">source</a></i>
    <a href="#transformations" title="Transformations">Transformations</a>: <i>
      - <a href="transformations.md">transformations</a></i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### State

State of the Data Lake Pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Name of this Data Lake Pipeline.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Sink

Ingestion destination of a Data Lake Pipeline.

_Required_: Yes

_Type_: <a href="sink.md">sink</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Source

Ingestion destination of a Data Lake Pipeline.

_Required_: No

_Type_: <a href="source.md">source</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Transformations

Ingestion destination of a Data Lake Pipeline.

_Required_: Yes

_Type_: List of <a href="transformations.md">transformations</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### CreatedDate

Timestamp that indicates when the Data Lake Pipeline was created.

#### Id

Unique 24-hexadecimal digit string that identifies your project.

#### LastUpdatedDate

Timestamp that indicates the last time that the Data Lake Pipeline was updated.

