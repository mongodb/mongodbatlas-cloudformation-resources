# MongoDB::Atlas::FederatedDatabaseInstance DataSource

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#allowinsecure" title="AllowInsecure">AllowInsecure</a>" : <i>Boolean</i>,
    "<a href="#collection" title="Collection">Collection</a>" : <i><a href="collection.md">Collection</a></i>,
    "<a href="#collectionregex" title="CollectionRegex">CollectionRegex</a>" : <i>String</i>,
    "<a href="#database" title="Database">Database</a>" : <i><a href="database.md">Database</a></i>,
    "<a href="#databaseregex" title="DatabaseRegex">DatabaseRegex</a>" : <i>String</i>,
    "<a href="#defaultformat" title="DefaultFormat">DefaultFormat</a>" : <i>String</i>,
    "<a href="#path" title="Path">Path</a>" : <i>String</i>,
    "<a href="#provenancefieldname" title="ProvenanceFieldName">ProvenanceFieldName</a>" : <i>String</i>,
    "<a href="#storename" title="StoreName">StoreName</a>" : <i>String</i>,
    "<a href="#urls" title="Urls">Urls</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#allowinsecure" title="AllowInsecure">AllowInsecure</a>: <i>Boolean</i>
<a href="#collection" title="Collection">Collection</a>: <i><a href="collection.md">Collection</a></i>
<a href="#collectionregex" title="CollectionRegex">CollectionRegex</a>: <i>String</i>
<a href="#database" title="Database">Database</a>: <i><a href="database.md">Database</a></i>
<a href="#databaseregex" title="DatabaseRegex">DatabaseRegex</a>: <i>String</i>
<a href="#defaultformat" title="DefaultFormat">DefaultFormat</a>: <i>String</i>
<a href="#path" title="Path">Path</a>: <i>String</i>
<a href="#provenancefieldname" title="ProvenanceFieldName">ProvenanceFieldName</a>: <i>String</i>
<a href="#storename" title="StoreName">StoreName</a>: <i>String</i>
<a href="#urls" title="Urls">Urls</a>: <i>
      - String</i>
</pre>

## Properties

#### AllowInsecure

Flag that validates the scheme in the specified URLs. If true, allows insecure HTTP scheme, doesn't verify the server's certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If false, allows secure HTTPS scheme only.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Collection

_Required_: No

_Type_: <a href="collection.md">Collection</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CollectionRegex

Regex pattern to use for creating the wildcard (*) collection. To learn more about the regex syntax, see Go programming language.( https://pkg.go.dev/regexp ).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Database

_Required_: No

_Type_: <a href="database.md">Database</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DatabaseRegex

Regex pattern to use for creating the wildcard (*) collection. To learn more about the regex syntax, see Go programming language.( https://pkg.go.dev/regexp ).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DefaultFormat

File format that MongoDB Cloud uses if it encounters a file without a file extension while searching storeName.Enum: ".avro" ".avro.bz2" ".avro.gz" ".bson" ".bson.bz2" ".bson.gz" ".bsonx" ".csv" ".csv.bz2" ".csv.gz" ".json" ".json.bz2" ".json.gz" ".orc" ".parquet" ".tsv" ".tsv.bz2" ".tsv.gz"

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Path

File path that controls how MongoDB Cloud searches for and parses files in the storeName before mapping them to a collection.Specify / to capture all files and folders from the prefix path.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProvenanceFieldName

Name for the field that includes the provenance of the documents in the results. MongoDB Cloud returns different fields in the results for each supported provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StoreName

Human-readable label that identifies the data store that MongoDB Cloud maps to the collection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Urls

URLs of the publicly accessible data files. You can't specify URLs that require authentication. Atlas Data Federation creates a partition for each URL. If empty or omitted, Data Federation uses the URLs from the store specified in the dataSources.storeName parameter.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

