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

Hardware specification for the search node instance sizes. The [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Atlas-Search/operation/createAtlasSearchDeployment) describes the valid values. More details can also be found in the [Search Node Documentation](https://www.mongodb.com/docs/atlas/cluster-config/multi-cloud-distribution/#search-tier).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NodeCount

Number of search nodes in the cluster.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

