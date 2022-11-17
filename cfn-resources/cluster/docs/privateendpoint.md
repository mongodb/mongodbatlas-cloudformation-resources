# MongoDB::Atlas::Cluster privateEndpoint

List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint. This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#connectionstring" title="ConnectionString">ConnectionString</a>" : <i>String</i>,
    "<a href="#endpoints" title="Endpoints">Endpoints</a>" : <i>[ <a href="endpoint.md">endpoint</a>, ... ]</i>,
    "<a href="#srvconnectionstring" title="SRVConnectionString">SRVConnectionString</a>" : <i>String</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#connectionstring" title="ConnectionString">ConnectionString</a>: <i>String</i>
<a href="#endpoints" title="Endpoints">Endpoints</a>: <i>
      - <a href="endpoint.md">endpoint</a></i>
<a href="#srvconnectionstring" title="SRVConnectionString">SRVConnectionString</a>: <i>String</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
</pre>

## Properties

#### ConnectionString

Private endpoint-aware connection string that uses the mongodb:// protocol to connect to MongoDB Cloud through a private endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Endpoints

List that contains the private endpoints through which you connect to MongoDB Cloud when you use connectionStrings.privateEndpoint[n].connectionString or connectionStrings.privateEndpoint[n].srvConnectionString.

_Required_: No

_Type_: List of <a href="endpoint.md">endpoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SRVConnectionString

Private endpoint-aware connection string that uses the mongodb+srv:// protocol to connect to MongoDB Cloud through a private endpoint. The mongodb+srv protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your application supports it. If it doesn't, use connectionStrings.privateEndpoint[n].connectionString.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

Enum: "MONGOD" "MONGOS"
MongoDB process type to which your application connects. Use MONGOD for replica sets and MONGOS for sharded clusters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

