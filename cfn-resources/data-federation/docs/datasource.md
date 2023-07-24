# MongoDB::Atlas::DataFederation DataSource

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#allowinsecure" title="AllowInsecure">AllowInsecure</a>" : <i>Boolean</i>,
    "<a href="#collection" title="Collection">Collection</a>" : <i><a href="collection.md">Collection</a></i>,
    "<a href="#collectionregex" title="CollectionRegex">CollectionRegex</a>" : <i>String</i>,
    "<a href="#database" title="Database">Database</a>" : <i>String</i>,
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
<a href="#database" title="Database">Database</a>: <i>String</i>
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

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Collection

_Required_: No

_Type_: <a href="collection.md">Collection</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CollectionRegex

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Database

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DatabaseRegex

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DefaultFormat

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Path

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProvenanceFieldName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StoreName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Urls

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

