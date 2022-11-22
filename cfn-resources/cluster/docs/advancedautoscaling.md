# MongoDB::Atlas::Cluster advancedAutoScaling

AWS Automatic Cluster Scaling

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#diskgb" title="DiskGB">DiskGB</a>" : <i><a href="diskgb.md">diskGB</a></i>,
    "<a href="#compute" title="Compute">Compute</a>" : <i><a href="compute.md">compute</a></i>
}
</pre>

### YAML

<pre>
<a href="#diskgb" title="DiskGB">DiskGB</a>: <i><a href="diskgb.md">diskGB</a></i>
<a href="#compute" title="Compute">Compute</a>: <i><a href="compute.md">compute</a></i>
</pre>

## Properties

#### DiskGB

Automatic cluster storage settings that apply to this cluster.

_Required_: No

_Type_: <a href="diskgb.md">diskGB</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Compute

Automatic Compute Scaling

_Required_: No

_Type_: <a href="compute.md">compute</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

