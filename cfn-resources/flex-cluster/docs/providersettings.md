# MongoDB::Atlas::FlexCluster ProviderSettings

Group of cloud provider settings that configure the provisioned MongoDB flex cluster.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#backingprovidername" title="BackingProviderName">BackingProviderName</a>" : <i>String</i>,
    "<a href="#regionname" title="RegionName">RegionName</a>" : <i>String</i>,
}
</pre>

### YAML

<pre>
<a href="#backingprovidername" title="BackingProviderName">BackingProviderName</a>: <i>String</i>
<a href="#regionname" title="RegionName">RegionName</a>: <i>String</i>
</pre>

## Properties

#### BackingProviderName

Cloud service provider on which MongoDB Cloud provisioned the flex cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### RegionName

Human-readable label that identifies the geographic location of your MongoDB flex cluster. The region you choose can affect network latency for clients accessing your databases.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

