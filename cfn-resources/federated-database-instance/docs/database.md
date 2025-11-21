# MongoDB::Atlas::FederatedDatabaseInstance Database

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#collections" title="Collections">Collections</a>" : <i>[ <a href="collection.md">Collection</a>, ... ]</i>,
    "<a href="#maxwildcardcollections" title="MaxWildcardCollections">MaxWildcardCollections</a>" : <i>String</i>,
    "<a href="#name" title="Name">Name</a>" : <i>String</i>,
    "<a href="#views" title="Views">Views</a>" : <i>[ <a href="view.md">View</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#collections" title="Collections">Collections</a>: <i>
      - <a href="collection.md">Collection</a></i>
<a href="#maxwildcardcollections" title="MaxWildcardCollections">MaxWildcardCollections</a>: <i>String</i>
<a href="#name" title="Name">Name</a>: <i>String</i>
<a href="#views" title="Views">Views</a>: <i>
      - <a href="view.md">View</a></i>
</pre>

## Properties

#### Collections

Array of collections and data sources that map to a stores data store.

_Required_: No

_Type_: List of <a href="collection.md">Collection</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaxWildcardCollections

Maximum number of wildcard collections in the database. This only applies to S3 data sources.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Human-readable label that identifies the database to which the Atlas Data Federation maps data.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Views

Array of aggregation pipelines that apply to the collection. This only applies to S3 data sources.

_Required_: No

_Type_: List of <a href="view.md">View</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

