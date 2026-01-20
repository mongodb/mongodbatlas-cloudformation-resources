# MongoDB::Atlas::StreamWorkspace StreamsDataProcessRegion

Information about the cloud provider region in which MongoDB Cloud processes the stream.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>" : <i>String</i>,
    "<a href="#region" title="Region">Region</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>: <i>String</i>
<a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### CloudProvider

Label that identifies the cloud service provider where MongoDB Cloud performs stream processing. For CloudFormation, this is restricted to AWS only.

_Required_: Yes

_Type_: String

_Allowed Values_: <code>AWS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Name of the cloud provider region hosting Atlas Stream Processing.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

