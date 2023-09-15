# MongoDB::Atlas::DataLakePipeline sink

Ingestion destination of a Data Lake Pipeline.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#type" title="Type">Type</a>" : <i>String</i>,
    "<a href="#metadataprovider" title="MetadataProvider">MetadataProvider</a>" : <i>String</i>,
    "<a href="#metadataregion" title="MetadataRegion">MetadataRegion</a>" : <i>String</i>,
    "<a href="#partitionfields" title="PartitionFields">PartitionFields</a>" : <i>[ <a href="partitionfields.md">partitionFields</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#type" title="Type">Type</a>: <i>String</i>
<a href="#metadataprovider" title="MetadataProvider">MetadataProvider</a>: <i>String</i>
<a href="#metadataregion" title="MetadataRegion">MetadataRegion</a>: <i>String</i>
<a href="#partitionfields" title="PartitionFields">PartitionFields</a>: <i>
      - <a href="partitionfields.md">partitionFields</a></i>
</pre>

## Properties

#### Type

Type of ingestion destination of this Data Lake Pipeline.

_Required_: No

_Type_: String

_Allowed Values_: <code>DLS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MetadataProvider

Target cloud provider for this Data Lake Pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MetadataRegion

Target cloud provider region for this Data Lake Pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PartitionFields

Ordered fields used to physically organize data in the destination.

_Required_: No

_Type_: List of <a href="partitionfields.md">partitionFields</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

