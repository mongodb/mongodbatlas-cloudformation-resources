# MongoDB::Atlas::FederatedDatabaseInstance DataProcessRegion

Information about the cloud provider region to which the Atlas Data Federation routes client connections. MongoDB Cloud supports AWS only.

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

Name of the cloud service that hosts the Atlas Data Federation data stores.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Name of the region to which the Atlas Data Federation routes client connections.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

