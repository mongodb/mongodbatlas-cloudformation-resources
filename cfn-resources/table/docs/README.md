# MongoDB::Atlas::Table

The MongoDB::Atlas::Table resource creates a MongoDB collection in a new or existing MongoDB Atlas cluster and a database user for accessing the connection. For more information, see Table in the MongoDB API Reference..

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Table",
    "Properties" : {
        "<a href="#publicapikey" title="PublicApiKey">PublicApiKey</a>" : <i>String</i>,
        "<a href="#privateapikey" title="PrivateApiKey">PrivateApiKey</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#username" title="Username">Username</a>" : <i>String</i>,
        "<a href="#labels" title="Labels">Labels</a>" : <i>[ <a href="labeldefinition.md">labelDefinition</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Table
Properties:
    <a href="#publicapikey" title="PublicApiKey">PublicApiKey</a>: <i>String</i>
    <a href="#privateapikey" title="PrivateApiKey">PrivateApiKey</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#username" title="Username">Username</a>: <i>String</i>
    <a href="#labels" title="Labels">Labels</a>: <i>
      - <a href="labeldefinition.md">labelDefinition</a></i>
</pre>

## Properties

#### PublicApiKey

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateApiKey

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique identifier of the Atlas project to which this Table belongs.

_Required_: Yes

_Type_: String

_Minimum_: <code>2</code>

_Maximum_: <code>250</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Username

Username for authenticating to MongoDB, this is optional. TODO: Support IAM

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Labels

Array containing key-value pairs that tag and categorize the database user.

_Required_: No

_Type_: List of <a href="labeldefinition.md">labelDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the TableCNFIdentifier.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### TableName

A name for the table. If you don't specify a name, MongoDB Atlas will generate a unique ID and uses that ID for the table name. This is synomous with a MongoDB Collection

#### ClusterName

Name of the the MongoDB Atlas Cluster.

#### DatabaseName

Name of the the MongoDB Atlas Database.

#### RegionName

Returns the <code>RegionName</code> value.

#### ConnectionStringsStandard

Connection strings for your MongoDB Table.

#### ConnectionStringsStandardSrv

Connection strings for your MongoDB Table.

#### TableCNFIdentifier

A unique identifier comprised of the Atlas Project ID, TableName, and Username

