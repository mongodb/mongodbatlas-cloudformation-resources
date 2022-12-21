# MongoDB::Atlas::ServerlessInstance ServerlessInstanceConnectionStrings

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#privateendpoint" title="PrivateEndpoint">PrivateEndpoint</a>" : <i>[ <a href="serverlessinstanceprivateendpoint.md">ServerlessInstancePrivateEndpoint</a>, ... ]</i>,
    "<a href="#standardsrv" title="StandardSrv">StandardSrv</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#privateendpoint" title="PrivateEndpoint">PrivateEndpoint</a>: <i>
      - <a href="serverlessinstanceprivateendpoint.md">ServerlessInstancePrivateEndpoint</a></i>
<a href="#standardsrv" title="StandardSrv">StandardSrv</a>: <i>String</i>
</pre>

## Properties

#### PrivateEndpoint

List of private endpoint connection strings that you can use to connect to this serverless instance through a private endpoint. This parameter returns only if you created a private endpoint for this serverless instance and it is AVAILABLE.

_Required_: No

_Type_: List of <a href="serverlessinstanceprivateendpoint.md">ServerlessInstancePrivateEndpoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StandardSrv

Public connection string that you can use to connect to this serverless instance. This connection string uses the `mongodb+srv://` protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

