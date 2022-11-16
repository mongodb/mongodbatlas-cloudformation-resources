# MongoDB::Atlas::Cluster privateEndpoint

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#connectionstring" title="ConnectionString">ConnectionString</a>" : <i>String</i>,
    "<a href="#endpoints" title="Endpoints">Endpoints</a>" : <i>[ <a href="endpoint.md">endpoint</a>, ... ]</i>,
    "<a href="#srvconnectionstring" title="SRVConnectionString">SRVConnectionString</a>" : <i>String</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#connectionstring" title="ConnectionString">ConnectionString</a>: <i>String</i>
<a href="#endpoints" title="Endpoints">Endpoints</a>: <i>
      - <a href="endpoint.md">endpoint</a></i>
<a href="#srvconnectionstring" title="SRVConnectionString">SRVConnectionString</a>: <i>String</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
</pre>

## Properties

#### ConnectionString

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Endpoints

_Required_: No

_Type_: List of <a href="endpoint.md">endpoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SRVConnectionString

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

