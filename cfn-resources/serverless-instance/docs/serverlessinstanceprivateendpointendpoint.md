# MongoDB::Atlas::ServerlessInstance ServerlessInstancePrivateEndpointEndpoint

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#endpointid" title="EndpointId">EndpointId</a>" : <i>String</i>,
    "<a href="#providername" title="ProviderName">ProviderName</a>" : <i>String</i>,
    "<a href="#region" title="Region">Region</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#endpointid" title="EndpointId">EndpointId</a>: <i>String</i>
<a href="#providername" title="ProviderName">ProviderName</a>: <i>String</i>
<a href="#region" title="Region">Region</a>: <i>String</i>
</pre>

## Properties

#### EndpointId

Unique provider identifier of the private endpoint.


_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProviderName

Cloud provider where the private endpoint is deployed.


_Required_: No

_Type_: String

_Allowed Values_: <code>AWS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Region where the private endpoint is deployed.


_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

