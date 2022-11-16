# MongoDB::Atlas::Cluster advancedReplicationSpec

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="ID">ID</a>" : <i>String</i>,
    "<a href="#numshards" title="NumShards">NumShards</a>" : <i>Integer</i>,
    "<a href="#advancedregionconfigs" title="AdvancedRegionConfigs">AdvancedRegionConfigs</a>" : <i>[ <a href="advancedregionconfig.md">AdvancedRegionConfig</a>, ... ]</i>,
    "<a href="#zonename" title="ZoneName">ZoneName</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#id" title="ID">ID</a>: <i>String</i>
<a href="#numshards" title="NumShards">NumShards</a>: <i>Integer</i>
<a href="#advancedregionconfigs" title="AdvancedRegionConfigs">AdvancedRegionConfigs</a>: <i>
      - <a href="advancedregionconfig.md">AdvancedRegionConfig</a></i>
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

#### AdvancedRegionConfigs

_Required_: No

_Type_: List of <a href="advancedregionconfig.md">AdvancedRegionConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ZoneName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

