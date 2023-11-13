# MongoDB::Atlas::SearchIndex ApiAtlasFTSAnalyzersTokenizer

Tokenizer that you want to use to create tokens. Tokens determine how Atlas Search splits up text into discrete chunks for indexing.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#maxgram" title="MaxGram">MaxGram</a>" : <i>Integer</i>,
    "<a href="#mingram" title="MinGram">MinGram</a>" : <i>Integer</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>,
    "<a href="#group" title="Group">Group</a>" : <i>Integer</i>,
    "<a href="#pattern" title="Pattern">Pattern</a>" : <i>String</i>,
    "<a href="#maxtokenlength" title="MaxTokenLength">MaxTokenLength</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#maxgram" title="MaxGram">MaxGram</a>: <i>Integer</i>
<a href="#mingram" title="MinGram">MinGram</a>: <i>Integer</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
<a href="#group" title="Group">Group</a>: <i>Integer</i>
<a href="#pattern" title="Pattern">Pattern</a>: <i>String</i>
<a href="#maxtokenlength" title="MaxTokenLength">MaxTokenLength</a>: <i>Integer</i>
</pre>

## Properties

#### MaxGram

Characters to include in the longest token that Atlas Search creates.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MinGram

Characters to include in the shortest token that Atlas Search creates.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

Human-readable label that identifies this tokenizer type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Group

Index of the character group within the matching expression to extract into tokens. Use `0` to extract all character groups.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Pattern

Regular expression to match against.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaxTokenLength

Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

