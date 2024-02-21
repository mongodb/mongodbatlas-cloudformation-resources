# MongoDB::Atlas::StreamInstance StreamConfig

Configuration options for an Atlas Stream Processing Instance.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#links" title="Links">Links</a>" : <i>[ <a href="link.md">Link</a>, ... ]</i>,
    "<a href="#tier" title="Tier">Tier</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#links" title="Links">Links</a>: <i>
      - <a href="link.md">Link</a></i>
<a href="#tier" title="Tier">Tier</a>: <i>String</i>
</pre>

## Properties

#### Links

_Required_: No

_Type_: List of <a href="link.md">Link</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tier

Selected tier for the Stream Instance. Configures Memory / VCPU allowances.

_Required_: No

_Type_: String

_Allowed Values_: <code>SP30</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

