# MongoDB::Atlas::Cluster AdvancedRegionConfig

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#analyticsautoscaling" title="AnalyticsAutoScaling">AnalyticsAutoScaling</a>" : <i><a href="advancedautoscaling.md">advancedAutoScaling</a></i>,
    "<a href="#autoscaling" title="AutoScaling">AutoScaling</a>" : <i><a href="advancedautoscaling.md">advancedAutoScaling</a></i>,
    "<a href="#regionname" title="RegionName">RegionName</a>" : <i>String</i>,
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
<a href="#analyticsspecs" title="AnalyticsSpecs">AnalyticsSpecs</a>: <i><a href="specs.md">specs</a></i>
<a href="#electablespecs" title="ElectableSpecs">ElectableSpecs</a>: <i><a href="specs.md">specs</a></i>
<a href="#priority" title="Priority">Priority</a>: <i>Integer</i>
<a href="#readonlyspecs" title="ReadOnlySpecs">ReadOnlySpecs</a>: <i><a href="specs.md">specs</a></i>
</pre>

## Properties

#### AnalyticsAutoScaling

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

