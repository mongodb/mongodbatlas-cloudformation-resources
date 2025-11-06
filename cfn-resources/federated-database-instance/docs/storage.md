# MongoDB::Atlas::FederatedDatabaseInstance Storage

Configuration information for each data store and its mapping to MongoDB Cloud databases.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#databases" title="Databases">Databases</a>" : <i>[ <a href="database.md">Database</a>, ... ]</i>,
    "<a href="#stores" title="Stores">Stores</a>" : <i>[ <a href="store.md">Store</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#databases" title="Databases">Databases</a>: <i>
      - <a href="database.md">Database</a></i>
<a href="#stores" title="Stores">Stores</a>: <i>
      - <a href="store.md">Store</a></i>
</pre>

## Properties

#### Databases

Array that contains the queryable databases and collections for this Atlas Data Federation.

_Required_: No

_Type_: List of <a href="database.md">Database</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Stores

Array that contains the data stores for the Atlas Data Federation.

_Required_: No

_Type_: List of <a href="store.md">Store</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

