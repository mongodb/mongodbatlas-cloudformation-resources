# MongoDB::Atlas::Cluster compute

Automatic Compute Scaling

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#enabled" title="Enabled">Enabled</a>" : <i>Boolean</i>,
    "<a href="#scaledownenabled" title="ScaleDownEnabled">ScaleDownEnabled</a>" : <i>Boolean</i>,
    "<a href="#mininstancesize" title="MinInstanceSize">MinInstanceSize</a>" : <i>String</i>,
    "<a href="#maxinstancesize" title="MaxInstanceSize">MaxInstanceSize</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#enabled" title="Enabled">Enabled</a>: <i>Boolean</i>
<a href="#scaledownenabled" title="ScaleDownEnabled">ScaleDownEnabled</a>: <i>Boolean</i>
<a href="#mininstancesize" title="MinInstanceSize">MinInstanceSize</a>: <i>String</i>
<a href="#maxinstancesize" title="MaxInstanceSize">MaxInstanceSize</a>: <i>String</i>
</pre>

## Properties

#### Enabled

Flag that indicates whether someone enabled instance size auto-scaling.

Set to true to enable instance size auto-scaling. If enabled, you must specify a value for replicationSpecs[n].regionConfigs[m].autoScaling.compute.maxInstanceSize.
Set to false to disable instance size automatic scaling.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ScaleDownEnabled

Flag that indicates whether the instance size may scale down. MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true. If you enable this option, specify a value for replicationSpecs[n].regionConfigs[m].autoScaling.compute.minInstanceSize.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MinInstanceSize

Minimum instance size to which your cluster can automatically scale. MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaxInstanceSize

Maximum instance size to which your cluster can automatically scale. MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

