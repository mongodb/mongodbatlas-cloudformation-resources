# MongoDB::Atlas::SearchIndex ApiAtlasFTSAnalyzersViewManual

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#charfilters" title="CharFilters">CharFilters</a>" : <i>[ Map, ... ]</i>,
    "<a href="#name" title="Name">Name</a>" : <i>String</i>,
    "<a href="#tokenfilters" title="TokenFilters">TokenFilters</a>" : <i>[ Map, ... ]</i>,
    "<a href="#tokenizer" title="Tokenizer">Tokenizer</a>" : <i>Map</i>
}
</pre>

### YAML

<pre>
<a href="#charfilters" title="CharFilters">CharFilters</a>: <i>
      - Map</i>
<a href="#name" title="Name">Name</a>: <i>String</i>
<a href="#tokenfilters" title="TokenFilters">TokenFilters</a>: <i>
      - Map</i>
<a href="#tokenizer" title="Tokenizer">Tokenizer</a>: <i>Map</i>
</pre>

## Properties

#### CharFilters

Filters that examine text one character at a time and perform filtering operations.

_Required_: No

_Type_: List of Map

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings:
- `lucene.`
- `builtin.`
- `mongodb.`

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TokenFilters

Filter that performs operations such as:

- Stemming, which reduces related words, such as "talking", "talked", and "talks" to their root word "talk".

- Redaction, the removal of sensitive information from public documents.

_Required_: No

_Type_: List of Map

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tokenizer

Tokenizer that you want to use to create tokens. Tokens determine how Atlas Search splits up text into discrete chunks for indexing.

_Required_: No

_Type_: Map

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

