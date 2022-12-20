# MongoDB::Atlas::DataLakes DataLakeStorageView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#databases" title="Databases">Databases</a>" : <i>[ <a href="datalakedatabaseview.md">DataLakeDatabaseView</a>, ... ]</i>,
    "<a href="#stores" title="Stores">Stores</a>" : <i>[ <a href="storedetail.md">StoreDetail</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#databases" title="Databases">Databases</a>: <i>
      - <a href="datalakedatabaseview.md">DataLakeDatabaseView</a></i>
<a href="#stores" title="Stores">Stores</a>: <i>
      - <a href="storedetail.md">StoreDetail</a></i>
</pre>

## Properties

#### Databases

Array that contains the queryable databases and collections for this data lake.

_Required_: No

_Type_: List of <a href="datalakedatabaseview.md">DataLakeDatabaseView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Stores

Array that contains the data stores for the data lake.

_Required_: No

_Type_: List of <a href="storedetail.md">StoreDetail</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

