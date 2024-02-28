# MongoDB::Atlas::SearchDeployment ApiSearchDeploymentSpec

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#instancesize" title="InstanceSize">InstanceSize</a>" : <i>String</i>,
    "<a href="#nodecount" title="NodeCount">NodeCount</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#instancesize" title="InstanceSize">InstanceSize</a>: <i>String</i>
<a href="#nodecount" title="NodeCount">NodeCount</a>: <i>Integer</i>
</pre>

## Properties

#### InstanceSize

Hardware specification for the search node instance sizes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NodeCount

Number of search nodes in the cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

