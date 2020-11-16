# MongoDB::Atlas::Cluster ReplicationSpec

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="ID">ID</a>" : <i>String</i>,
    "<a href="#numshards" title="NumShards">NumShards</a>" : <i>Integer</i>,
    "<a href="#regionsconfig" title="RegionsConfig">RegionsConfig</a>" : <i>[ [ <a href="regionsconfig.md">RegionsConfig</a>, ... ], ... ]</i>,
    "<a href="#zonename" title="ZoneName">ZoneName</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#id" title="ID">ID</a>: <i>String</i>
<a href="#numshards" title="NumShards">NumShards</a>: <i>Integer</i>
<a href="#regionsconfig" title="RegionsConfig">RegionsConfig</a>: <i>
      - 
      - <a href="regionsconfig.md">RegionsConfig</a></i>
<a href="#zonename" title="ZoneName">ZoneName</a>: <i>String</i>
</pre>

## Properties

#### ID

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NumShards

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RegionsConfig

_Required_: No

_Type_: List of List of <a href="regionsconfig.md">RegionsConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ZoneName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

