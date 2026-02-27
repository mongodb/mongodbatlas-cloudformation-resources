# MongoDB::Atlas::FederatedDatabaseInstance Store

Array that contains the data stores for the Atlas Data Federation.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#name" title="Name">Name</a>" : <i>String</i>,
    "<a href="#provider" title="Provider">Provider</a>" : <i>String</i>,
    "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
    "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    "<a href="#readpreference" title="ReadPreference">ReadPreference</a>" : <i><a href="readpreference.md">ReadPreference</a></i>
}
</pre>

### YAML

<pre>
<a href="#name" title="Name">Name</a>: <i>String</i>
<a href="#provider" title="Provider">Provider</a>: <i>String</i>
<a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
<a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
<a href="#readpreference" title="ReadPreference">ReadPreference</a>: <i><a href="readpreference.md">ReadPreference</a></i>
</pre>

## Properties

#### Name

Human-readable label that identifies the data store. The databases.[n].collections.[n].dataSources.[n].storeName field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the Atlas Data Federation requires an M10 or higher cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Provider

Allowed values atlas, http, online_archive, s3 and DataLakeAzureBlobStore.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

Human-readable label of the MongoDB Cloud cluster on which the store is based.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies the project.Regex ^([a-f0-9]{24})$ .

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReadPreference

_Required_: No

_Type_: <a href="readpreference.md">ReadPreference</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

