# MongoDB::Atlas::CloudBackUpRestoreJobs PrivateDownloadDeliveryUrl

Private endpoint delivery URL for manual download.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#deliveryurl" title="DeliveryUrl">DeliveryUrl</a>" : <i>String</i>,
    "<a href="#endpointid" title="EndpointId">EndpointId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#deliveryurl" title="DeliveryUrl">DeliveryUrl</a>: <i>String</i>
<a href="#endpointid" title="EndpointId">EndpointId</a>: <i>String</i>
</pre>

## Properties

#### DeliveryUrl

One URL that points to the compressed snapshot files for manual download.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EndpointId

Unique identifier of the private endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

