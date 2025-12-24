# MongoDB::Atlas::StreamWorkspace StreamConfig

Configuration options for an Atlas Stream Processing Workspace.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#tier" title="Tier">Tier</a>" : <i>String</i>,
    "<a href="#maxtiersize" title="MaxTierSize">MaxTierSize</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#tier" title="Tier">Tier</a>: <i>String</i>
<a href="#maxtiersize" title="MaxTierSize">MaxTierSize</a>: <i>String</i>
</pre>

## Properties

#### Tier

Selected tier for the Stream Workspace. Configures Memory / VCPU allowances.

_Required_: No

_Type_: String

_Allowed Values_: <code>SP2</code> | <code>SP5</code> | <code>SP10</code> | <code>SP30</code> | <code>SP50</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaxTierSize

Max tier size for the Stream Workspace. Configures Memory / VCPU allowances.

_Required_: No

_Type_: String

_Allowed Values_: <code>SP2</code> | <code>SP5</code> | <code>SP10</code> | <code>SP30</code> | <code>SP50</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

