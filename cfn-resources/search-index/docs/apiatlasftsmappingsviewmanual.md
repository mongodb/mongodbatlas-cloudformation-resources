# MongoDB::Atlas::SearchIndex ApiAtlasFTSMappingsViewManual

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#dynamic" title="Dynamic">Dynamic</a>" : <i>Boolean</i>,
    "<a href="#fields" title="Fields">Fields</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#dynamic" title="Dynamic">Dynamic</a>: <i>Boolean</i>
<a href="#fields" title="Fields">Fields</a>: <i>
      - String</i>
</pre>

## Properties

#### Dynamic

Flag that indicates whether the index uses dynamic or static mappings. Required if **mappings.fields** is omitted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Fields

One or more field specifications for the Atlas Search index. The element of the array must have the format fieldName:{Stringify json, containing a list of types with all its properties}. Required if **mappings.dynamic** is omitted or set to **false**.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

