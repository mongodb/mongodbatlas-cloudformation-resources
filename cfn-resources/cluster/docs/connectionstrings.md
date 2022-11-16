# MongoDB::Atlas::Cluster connectionStrings

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#standard" title="Standard">Standard</a>" : <i>String</i>,
    "<a href="#standardsrv" title="StandardSrv">StandardSrv</a>" : <i>String</i>,
    "<a href="#private" title="Private">Private</a>" : <i>String</i>,
    "<a href="#privatesrv" title="PrivateSrv">PrivateSrv</a>" : <i>String</i>,
    "<a href="#privateendpoint" title="PrivateEndpoint">PrivateEndpoint</a>" : <i>[ <a href="privateendpoint.md">privateEndpoint</a>, ... ]</i>,
    "<a href="#awsprivatelinksrv" title="AwsPrivateLinkSrv">AwsPrivateLinkSrv</a>" : <i>String</i>,
    "<a href="#awsprivatelink" title="AwsPrivateLink">AwsPrivateLink</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#standard" title="Standard">Standard</a>: <i>String</i>
<a href="#standardsrv" title="StandardSrv">StandardSrv</a>: <i>String</i>
<a href="#private" title="Private">Private</a>: <i>String</i>
<a href="#privatesrv" title="PrivateSrv">PrivateSrv</a>: <i>String</i>
<a href="#privateendpoint" title="PrivateEndpoint">PrivateEndpoint</a>: <i>
      - <a href="privateendpoint.md">privateEndpoint</a></i>
<a href="#awsprivatelinksrv" title="AwsPrivateLinkSrv">AwsPrivateLinkSrv</a>: <i>String</i>
<a href="#awsprivatelink" title="AwsPrivateLink">AwsPrivateLink</a>: <i>String</i>
</pre>

## Properties

#### Standard

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StandardSrv

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Private

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateSrv

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateEndpoint

_Required_: No

_Type_: List of <a href="privateendpoint.md">privateEndpoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsPrivateLinkSrv

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsPrivateLink

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

