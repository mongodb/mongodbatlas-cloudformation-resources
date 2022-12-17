# MongoDB::Atlas::SearchIndex ApiAtlasFTSSynonymMappingDefinitionView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#analyzer" title="Analyzer">Analyzer</a>" : <i>String</i>,
    "<a href="#name" title="Name">Name</a>" : <i>String</i>,
    "<a href="#source" title="Source">Source</a>" : <i><a href="synonymsource.md">SynonymSource</a></i>
}
</pre>

### YAML

<pre>
<a href="#analyzer" title="Analyzer">Analyzer</a>: <i>String</i>
<a href="#name" title="Name">Name</a>: <i>String</i>
<a href="#source" title="Source">Source</a>: <i><a href="synonymsource.md">SynonymSource</a></i>
</pre>

## Properties

#### Analyzer

Specific pre-defined method chosen to apply to the synonyms to be searched.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Human-readable label that identifies the synonym definition. Each **synonym.name** must be unique within the same index definition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Source

_Required_: No

_Type_: <a href="synonymsource.md">SynonymSource</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

