# MongoDB::Atlas::ServerlessInstance ServerlessInstanceProviderSettings

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#providername" title="ProviderName">ProviderName</a>" : <i>String</i>,
    "<a href="#regionname" title="RegionName">RegionName</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#providername" title="ProviderName">ProviderName</a>: <i>String</i>
<a href="#regionname" title="RegionName">RegionName</a>: <i>String</i>
</pre>

## Properties

#### ProviderName

Human-readable label that identifies the cloud service provider.

_Required_: No

_Type_: String

_Allowed Values_: <code>SERVERLESS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RegionName

Human-readable label that identifies the geographic location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

