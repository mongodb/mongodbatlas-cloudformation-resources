# MongoDB::Atlas::ServerlessInstance ServerlessInstancePrivateEndpoint

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#endpoints" title="Endpoints">Endpoints</a>" : <i>[ <a href="serverlessinstanceprivateendpointendpoint.md">ServerlessInstancePrivateEndpointEndpoint</a>, ... ]</i>,
    "<a href="#srvconnectionstring" title="SrvConnectionString">SrvConnectionString</a>" : <i>String</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#endpoints" title="Endpoints">Endpoints</a>: <i>
      - <a href="serverlessinstanceprivateendpointendpoint.md">ServerlessInstancePrivateEndpointEndpoint</a></i>
<a href="#srvconnectionstring" title="SrvConnectionString">SrvConnectionString</a>: <i>String</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
</pre>

## Properties

#### Endpoints

List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**.

_Required_: No

_Type_: List of <a href="serverlessinstanceprivateendpointendpoint.md">ServerlessInstancePrivateEndpointEndpoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SrvConnectionString

Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to MongoDB Cloud through a private endpoint. The `mongodb+srv` protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

MongoDB process type to which your application connects.


_Required_: No

_Type_: String

_Allowed Values_: <code>MONGOS</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

