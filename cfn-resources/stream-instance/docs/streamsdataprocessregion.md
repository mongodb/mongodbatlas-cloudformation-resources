# MongoDB::Atlas::StreamInstance StreamsDataProcessRegion

Information about the cloud provider region in which MongoDB Cloud processes the stream.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>" : <i>String</i>,
    "<a href="#links" title="Links">Links</a>" : <i>[ <a href="link.md">Link</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>: <i>String</i>
<a href="#links" title="Links">Links</a>: <i>
      - <a href="link.md">Link</a></i>
</pre>

## Properties

#### CloudProvider

Label that identifies the cloud service provider where MongoDB Cloud performs stream processing. Currently, this parameter supports AWS only.

_Required_: No

_Type_: String

_Allowed Values_: <code>AWS</code> | <code>GCP</code> | <code>AZURE</code> | <code>TENANT</code> | <code>SERVERLESS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Links

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

_Required_: No

_Type_: List of <a href="link.md">Link</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

