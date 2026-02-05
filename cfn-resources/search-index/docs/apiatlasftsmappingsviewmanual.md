# MongoDB::Atlas::SearchIndex ApiAtlasFTSMappingsViewManual

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#dynamic" title="Dynamic">Dynamic</a>" : <i>Boolean</i>,
    "<a href="#dynamicconfig" title="DynamicConfig">DynamicConfig</a>" : <i>String</i>,
    "<a href="#fields" title="Fields">Fields</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#dynamic" title="Dynamic">Dynamic</a>: <i>Boolean</i>
<a href="#dynamicconfig" title="DynamicConfig">DynamicConfig</a>: <i>String</i>
<a href="#fields" title="Fields">Fields</a>: <i>String</i>
</pre>

## Properties

#### Dynamic

Flag that indicates whether the index uses dynamic or static mappings. If DynamicConfig is specified, this field is ignored (DynamicConfig takes precedence). Required for search indexes if **mappings.fields** is omitted and **mappings.dynamicConfig** is not specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DynamicConfig

Stringify json representation of dynamic mapping configuration object. This allows for more complex dynamic mapping configurations beyond a simple boolean. If both Dynamic and DynamicConfig are specified, DynamicConfig takes precedence.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Fields

One or more field specifications for the Atlas Search index. Stringify json representation of field with types and properties. Required for search indexes if **mappings.dynamic** and **mappings.dynamicConfig** are omitted or if **mappings.dynamic** is set to **false**.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

