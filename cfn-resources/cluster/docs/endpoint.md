# MongoDB::Atlas::Cluster endpoint

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#endpointid" title="EndpointID">EndpointID</a>" : <i>String</i>,
    "<a href="#providername" title="ProviderName">ProviderName</a>" : <i>String</i>,
    "<a href="#region" title="Region">Region</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#endpointid" title="EndpointID">EndpointID</a>: <i>String</i>
<a href="#providername" title="ProviderName">ProviderName</a>: <i>String</i>
<a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### EndpointID

Unique string that the cloud provider uses to identify the private endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProviderName

Cloud provider in which MongoDB Cloud deploys the private endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Region in which MongoDB Cloud deploys the private endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

