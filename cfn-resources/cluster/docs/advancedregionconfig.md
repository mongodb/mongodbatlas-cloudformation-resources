# MongoDB::Atlas::Cluster advancedRegionConfig

Hardware specifications for nodes set for a given region. Each regionConfigs object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each regionConfigs object must have either an analyticsSpecs object, electableSpecs object, or readOnlySpecs object. Tenant clusters only require electableSpecs. Dedicated clusters can specify any of these specifications, but must have at least one electableSpecs object within a replicationSpec. Every hardware specification must use the same instanceSize.

Example:

If you set "replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize" : "M30", set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30"if you have electable nodes and"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize" : "M30" if you have read-only nodes.",

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#analyticsautoscaling" title="AnalyticsAutoScaling">AnalyticsAutoScaling</a>" : <i><a href="advancedautoscaling.md">advancedAutoScaling</a></i>,
    "<a href="#autoscaling" title="AutoScaling">AutoScaling</a>" : <i><a href="advancedautoscaling.md">advancedAutoScaling</a></i>,
    "<a href="#regionname" title="RegionName">RegionName</a>" : <i>String</i>,
    "<a href="#backingprovidername" title="BackingProviderName">BackingProviderName</a>" : <i>String</i>,
    "<a href="#providername" title="ProviderName">ProviderName</a>" : <i>String</i>,
    "<a href="#analyticsspecs" title="AnalyticsSpecs">AnalyticsSpecs</a>" : <i><a href="specs.md">specs</a></i>,
    "<a href="#electablespecs" title="ElectableSpecs">ElectableSpecs</a>" : <i><a href="specs.md">specs</a></i>,
    "<a href="#priority" title="Priority">Priority</a>" : <i>Integer</i>,
    "<a href="#readonlyspecs" title="ReadOnlySpecs">ReadOnlySpecs</a>" : <i><a href="specs.md">specs</a></i>
}
</pre>

### YAML

<pre>
<a href="#analyticsautoscaling" title="AnalyticsAutoScaling">AnalyticsAutoScaling</a>: <i><a href="advancedautoscaling.md">advancedAutoScaling</a></i>
<a href="#autoscaling" title="AutoScaling">AutoScaling</a>: <i><a href="advancedautoscaling.md">advancedAutoScaling</a></i>
<a href="#regionname" title="RegionName">RegionName</a>: <i>String</i>
<a href="#backingprovidername" title="BackingProviderName">BackingProviderName</a>: <i>String</i>
<a href="#providername" title="ProviderName">ProviderName</a>: <i>String</i>
<a href="#analyticsspecs" title="AnalyticsSpecs">AnalyticsSpecs</a>: <i><a href="specs.md">specs</a></i>
<a href="#electablespecs" title="ElectableSpecs">ElectableSpecs</a>: <i><a href="specs.md">specs</a></i>
<a href="#priority" title="Priority">Priority</a>: <i>Integer</i>
<a href="#readonlyspecs" title="ReadOnlySpecs">ReadOnlySpecs</a>: <i><a href="specs.md">specs</a></i>
</pre>

## Properties

#### AnalyticsAutoScaling

AWS Automatic Cluster Scaling

_Required_: No

_Type_: <a href="advancedautoscaling.md">advancedAutoScaling</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AutoScaling

_Required_: No

_Type_: <a href="advancedautoscaling.md">advancedAutoScaling</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RegionName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BackingProviderName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProviderName

_Required_: No

_Type_: String

_Allowed Values_: <code>AWS</code> | <code>GCP</code> | <code>AZURE</code> | <code>TENANT</code> | <code>FLEX</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AnalyticsSpecs

_Required_: No

_Type_: <a href="specs.md">specs</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ElectableSpecs

_Required_: No

_Type_: <a href="specs.md">specs</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Priority

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReadOnlySpecs

_Required_: No

_Type_: <a href="specs.md">specs</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

