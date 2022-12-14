# MongoDB::Atlas::DataLakes DataLakeDatabaseDataSourceView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#allowinsecure" title="AllowInsecure">AllowInsecure</a>" : <i>Boolean</i>,
    "<a href="#collection" title="Collection">Collection</a>" : <i>String</i>,
    "<a href="#collectionregex" title="CollectionRegex">CollectionRegex</a>" : <i>String</i>,
    "<a href="#database" title="Database">Database</a>" : <i>String</i>,
    "<a href="#defaultformat" title="DefaultFormat">DefaultFormat</a>" : <i>String</i>,
    "<a href="#path" title="Path">Path</a>" : <i>String</i>,
    "<a href="#storename" title="StoreName">StoreName</a>" : <i>String</i>,
    "<a href="#urls" title="Urls">Urls</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#allowinsecure" title="AllowInsecure">AllowInsecure</a>: <i>Boolean</i>
<a href="#collection" title="Collection">Collection</a>: <i>String</i>
<a href="#collectionregex" title="CollectionRegex">CollectionRegex</a>: <i>String</i>
<a href="#database" title="Database">Database</a>: <i>String</i>
<a href="#defaultformat" title="DefaultFormat">DefaultFormat</a>: <i>String</i>
<a href="#path" title="Path">Path</a>: <i>String</i>
<a href="#storename" title="StoreName">StoreName</a>: <i>String</i>
<a href="#urls" title="Urls">Urls</a>: <i>
      - String</i>
</pre>

## Properties

#### AllowInsecure

Flag that validates the scheme in the specified URLs. If `true`, allows insecure `HTTP` scheme, doesn't verify the server's certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If `false`, allows secure `HTTPS` scheme only.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Collection

Human-readable label that identifies the collection in the database. For creating a wildcard (`*`) collection, you must omit this parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CollectionRegex

Regex pattern to use for creating the wildcard (*) collection. To learn more about the regex syntax, see [Go programming language](https://pkg.go.dev/regexp).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Database

Human-readable label that identifies the database, which contains the collection in the cluster. You must omit this parameter to generate wildcard (`*`) collections for dynamically generated databases.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DefaultFormat

File format that MongoDB Cloud uses if it encounters a file without a file extension while searching **storeName**.

_Required_: No

_Type_: String

_Allowed Values_: <code>.avro</code> | <code>.avro.gz</code> | <code>.bson</code> | <code>.bson.gz</code> | <code>.csv</code> | <code>.json</code> | <code>.json.gz</code> | <code>.orc</code> | <code>.tsv</code> | <code>.tsv.gz</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Path

File path that controls how MongoDB Cloud searches for and parses files in the **storeName** before mapping them to a collection.Specify ``/`` to capture all files and folders from the ``prefix`` path.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StoreName

Human-readable label that identifies the data store that MongoDB Cloud maps to the collection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Urls

URLs of the publicly accessible data files. You can't specify URLs that require authentication. Atlas Data Lake creates a partition for each URL. If empty or omitted, Data Lake uses the URLs from the store specified in the **dataSources.storeName** parameter.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

