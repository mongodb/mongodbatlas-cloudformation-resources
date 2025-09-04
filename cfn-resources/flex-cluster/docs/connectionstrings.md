# MongoDB::Atlas::FlexCluster ConnectionStrings

Collection of Uniform Resource Locators that point to the MongoDB database.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#standard" title="Standard">Standard</a>" : <i>String</i>,
    "<a href="#standardsrv" title="StandardSrv">StandardSrv</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#standard" title="Standard">Standard</a>: <i>String</i>
<a href="#standardsrv" title="StandardSrv">StandardSrv</a>: <i>String</i>
</pre>

## Properties

#### Standard

Public connection string that you can use to connect to this cluster. This connection string uses the mongodb:// protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StandardSrv

Public connection string that you can use to connect to this flex cluster. This connection string uses the mongodb+srv:// protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

