# MongoDB::Atlas::FederatedDatabaseInstance Collection

Array of collections and data sources that map to a stores data store.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#datasources" title="DataSources">DataSources</a>" : <i>[ <a href="datasource.md">DataSource</a>, ... ]</i>,
    "<a href="#name" title="Name">Name</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#datasources" title="DataSources">DataSources</a>: <i>
      - <a href="datasource.md">DataSource</a></i>
<a href="#name" title="Name">Name</a>: <i>String</i>
</pre>

## Properties

#### DataSources

Array that contains the data stores that map to a collection for this Atlas Data Federation.

_Required_: No

_Type_: List of <a href="datasource.md">DataSource</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

