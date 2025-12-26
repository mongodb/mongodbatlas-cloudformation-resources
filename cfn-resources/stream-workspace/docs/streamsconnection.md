# MongoDB::Atlas::StreamWorkspace StreamsConnection

Settings that define a connection to an external data store.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#name" title="Name">Name</a>" : <i>String</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>,
    "<a href="#authentication" title="Authentication">Authentication</a>" : <i><a href="streamskafkaauthentication.md">StreamsKafkaAuthentication</a></i>,
    "<a href="#bootstrapservers" title="BootstrapServers">BootstrapServers</a>" : <i>String</i>,
    "<a href="#security" title="Security">Security</a>" : <i><a href="streamskafkasecurity.md">StreamsKafkaSecurity</a></i>,
    "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
    "<a href="#dbroletoexecute" title="DbRoleToExecute">DbRoleToExecute</a>" : <i><a href="dbroletoexecute.md">DBRoleToExecute</a></i>
}
</pre>

### YAML

<pre>
<a href="#name" title="Name">Name</a>: <i>String</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
<a href="#authentication" title="Authentication">Authentication</a>: <i><a href="streamskafkaauthentication.md">StreamsKafkaAuthentication</a></i>
<a href="#bootstrapservers" title="BootstrapServers">BootstrapServers</a>: <i>String</i>
<a href="#security" title="Security">Security</a>: <i><a href="streamskafkasecurity.md">StreamsKafkaSecurity</a></i>
<a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
<a href="#dbroletoexecute" title="DbRoleToExecute">DbRoleToExecute</a>: <i><a href="dbroletoexecute.md">DBRoleToExecute</a></i>
</pre>

## Properties

#### Name

Human-readable label that identifies the stream connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

Type of the connection. Can be either Cluster or Kafka.

_Required_: No

_Type_: String

_Allowed Values_: <code>Kafka</code> | <code>Cluster</code> | <code>Sample</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Authentication

User credentials required to connect to a Kafka Cluster. Includes the authentication type, as well as the parameters for that authentication mode.

_Required_: No

_Type_: <a href="streamskafkaauthentication.md">StreamsKafkaAuthentication</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BootstrapServers

Comma separated list of server addresses.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Security

Properties for the secure transport connection to Kafka. For SSL, this can include the trusted certificate to use.

_Required_: No

_Type_: <a href="streamskafkasecurity.md">StreamsKafkaSecurity</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

Name of the cluster configured for this connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DbRoleToExecute

The name of a Built in or Custom DB Role to connect to an Atlas Cluster.

_Required_: No

_Type_: <a href="dbroletoexecute.md">DBRoleToExecute</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

