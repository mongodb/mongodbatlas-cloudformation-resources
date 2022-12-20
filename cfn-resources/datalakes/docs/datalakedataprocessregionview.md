# MongoDB::Atlas::DataLakes DataLakeDataProcessRegionView

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

Name of the cloud service that hosts the data lake's data stores.

_Required_: No

_Type_: String

_Allowed Values_: <code>AWS</code> | <code>GCP</code> | <code>AZURE</code> | <code>TENANT</code> | <code>SERVERLESS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Name of the region to which the data lake routes client connections.

_Required_: No

_Type_: String

_Allowed Values_: <code>DUBLIN_IRL</code> | <code>FRANKFURT_DEU</code> | <code>LONDON_GBR</code> | <code>MUMBAI_IND</code> | <code>OREGON_USA</code> | <code>SYDNEY_AUS</code> | <code>VIRGINIA_USA</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

