# MongoDB::Atlas::SearchIndex

Returns, adds, edits, and removes Atlas Search indexes. Also returns and updates user-defined analyzers.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::SearchIndex",
    "Properties" : {
        "<a href="#analyzer" title="Analyzer">Analyzer</a>" : <i>String</i>,
        "<a href="#analyzers" title="Analyzers">Analyzers</a>" : <i>[ <a href="apiatlasftsanalyzersviewmanual.md">ApiAtlasFTSAnalyzersViewManual</a>, ... ]</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#collectionname" title="CollectionName">CollectionName</a>" : <i>String</i>,
        "<a href="#database" title="Database">Database</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#mappings" title="Mappings">Mappings</a>" : <i><a href="apiatlasftsmappingsviewmanual.md">ApiAtlasFTSMappingsViewManual</a></i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#type" title="Type">Type</a>" : <i>String</i>,
        "<a href="#searchanalyzer" title="SearchAnalyzer">SearchAnalyzer</a>" : <i>String</i>,
        "<a href="#synonyms" title="Synonyms">Synonyms</a>" : <i>[ <a href="apiatlasftssynonymmappingdefinitionview.md">ApiAtlasFTSSynonymMappingDefinitionView</a>, ... ]</i>,
        "<a href="#fields" title="Fields">Fields</a>" : <i>String</i>,
        "<a href="#storedsource" title="StoredSource">StoredSource</a>" : <i>String</i>,
        "<a href="#typesets" title="TypeSets">TypeSets</a>" : <i>[ <a href="typeset.md">TypeSet</a>, ... ]</i>,
        "<a href="#numpartitions" title="NumPartitions">NumPartitions</a>" : <i>Integer</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::SearchIndex
Properties:
    <a href="#analyzer" title="Analyzer">Analyzer</a>: <i>String</i>
    <a href="#analyzers" title="Analyzers">Analyzers</a>: <i>
      - <a href="apiatlasftsanalyzersviewmanual.md">ApiAtlasFTSAnalyzersViewManual</a></i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#collectionname" title="CollectionName">CollectionName</a>: <i>String</i>
    <a href="#database" title="Database">Database</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#mappings" title="Mappings">Mappings</a>: <i><a href="apiatlasftsmappingsviewmanual.md">ApiAtlasFTSMappingsViewManual</a></i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#type" title="Type">Type</a>: <i>String</i>
    <a href="#searchanalyzer" title="SearchAnalyzer">SearchAnalyzer</a>: <i>String</i>
    <a href="#synonyms" title="Synonyms">Synonyms</a>: <i>
      - <a href="apiatlasftssynonymmappingdefinitionview.md">ApiAtlasFTSSynonymMappingDefinitionView</a></i>
    <a href="#fields" title="Fields">Fields</a>: <i>String</i>
    <a href="#storedsource" title="StoredSource">StoredSource</a>: <i>String</i>
    <a href="#typesets" title="TypeSets">TypeSets</a>: <i>
      - <a href="typeset.md">TypeSet</a></i>
    <a href="#numpartitions" title="NumPartitions">NumPartitions</a>: <i>Integer</i>
</pre>

## Properties

#### Analyzer

Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves a variety of changes made to the text in fields:

- extracting words
- removing punctuation
- removing accents
- changing to lowercase
- removing common words
- reducing words to their root form (stemming)
- changing words to their base form (lemmatization)
 MongoDB Cloud uses the selected process to build the Atlas Search index.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Analyzers

List of user-defined methods to convert database field text into searchable words.

_Required_: No

_Type_: List of <a href="apiatlasftsanalyzersviewmanual.md">ApiAtlasFTSAnalyzersViewManual</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ClusterName

Name of the cluster that contains the database and collection with one or more Application Search indexes.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>1</code>

_Maximum Length_: <code>64</code>

_Pattern_: <code>^[a-zA-Z0-9][a-zA-Z0-9-]*$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### CollectionName

Human-readable label that identifies the collection that contains one or more Atlas Search indexes.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Database

Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Mappings

_Required_: No

_Type_: <a href="apiatlasftsmappingsviewmanual.md">ApiAtlasFTSMappingsViewManual</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Human-readable label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Type

Type of index: **search** or **vectorSearch**. Default type is **search**.

_Required_: No

_Type_: String

_Allowed Values_: <code>search</code> | <code>vectorSearch</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### SearchAnalyzer

Method applied to identify words when searching this index.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Synonyms

Rule sets that map words to their synonyms in this index.

_Required_: No

_Type_: List of <a href="apiatlasftssynonymmappingdefinitionview.md">ApiAtlasFTSSynonymMappingDefinitionView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Fields

Array of [Fields](https://www.mongodb.com/docs/atlas/atlas-search/field-types/knn-vector/#std-label-fts-data-types-knn-vector) to configure this vectorSearch index. Stringify json representation of field with types and properties. Required for vector indexes. It must contain at least one **vector** type field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StoredSource

Flag that indicates whether to store the original document in the index. Can be a boolean ("true" or "false") or a stringified JSON object specifying which fields to include/exclude. When stored, this allows the index to return the original document for queries.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TypeSets

Array of type sets that define alternate types for fields in the index. Each type set allows you to group related fields under a common name.

_Required_: No

_Type_: List of <a href="typeset.md">TypeSet</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NumPartitions

Number of partitions for the index. This is used to improve search performance for large datasets by distributing the index across multiple partitions.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### IndexId

Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.

#### Status

Condition of the search index when you made this request.

| Status | Index Condition |
 |---|---|
 | IN_PROGRESS | Atlas is building or re-building the index after an edit. |
 | STEADY | You can use this search index. |
 | FAILED | Atlas could not build the index. |
 | MIGRATING | Atlas is upgrading the underlying cluster tier and migrating indexes. |


