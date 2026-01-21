# MongoDB::Atlas::StreamConnection

Returns, adds, edits, and removes one connection for a stream workspace in the specified project. To use this resource, the requesting API Key must have the Project Owner roles.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::StreamConnection",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#connectionname" title="ConnectionName">ConnectionName</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#workspacename" title="WorkspaceName">WorkspaceName</a>" : <i>String</i>,
        "<a href="#type" title="Type">Type</a>" : <i>String</i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#clusterprojectid" title="ClusterProjectId">ClusterProjectId</a>" : <i>String</i>,
        "<a href="#dbroletoexecute" title="DbRoleToExecute">DbRoleToExecute</a>" : <i><a href="dbroletoexecute.md">DBRoleToExecute</a></i>,
        "<a href="#authentication" title="Authentication">Authentication</a>" : <i><a href="streamskafkaauthentication.md">StreamsKafkaAuthentication</a></i>,
        "<a href="#bootstrapservers" title="BootstrapServers">BootstrapServers</a>" : <i>String</i>,
        "<a href="#security" title="Security">Security</a>" : <i><a href="streamskafkasecurity.md">StreamsKafkaSecurity</a></i>,
        "<a href="#config" title="Config">Config</a>" : <i><a href="config.md">Config</a></i>,
        "<a href="#networking" title="Networking">Networking</a>" : <i><a href="networking.md">Networking</a></i>,
        "<a href="#aws" title="Aws">Aws</a>" : <i><a href="aws.md">Aws</a></i>,
        "<a href="#url" title="Url">Url</a>" : <i>String</i>,
        "<a href="#headers" title="Headers">Headers</a>" : <i><a href="headers.md">Headers</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::StreamConnection
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#connectionname" title="ConnectionName">ConnectionName</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#workspacename" title="WorkspaceName">WorkspaceName</a>: <i>String</i>
    <a href="#type" title="Type">Type</a>: <i>String</i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#clusterprojectid" title="ClusterProjectId">ClusterProjectId</a>: <i>String</i>
    <a href="#dbroletoexecute" title="DbRoleToExecute">DbRoleToExecute</a>: <i><a href="dbroletoexecute.md">DBRoleToExecute</a></i>
    <a href="#authentication" title="Authentication">Authentication</a>: <i><a href="streamskafkaauthentication.md">StreamsKafkaAuthentication</a></i>
    <a href="#bootstrapservers" title="BootstrapServers">BootstrapServers</a>: <i>String</i>
    <a href="#security" title="Security">Security</a>: <i><a href="streamskafkasecurity.md">StreamsKafkaSecurity</a></i>
    <a href="#config" title="Config">Config</a>: <i><a href="config.md">Config</a></i>
    <a href="#networking" title="Networking">Networking</a>: <i><a href="networking.md">Networking</a></i>
    <a href="#aws" title="Aws">Aws</a>: <i><a href="aws.md">Aws</a></i>
    <a href="#url" title="Url">Url</a>: <i>String</i>
    <a href="#headers" title="Headers">Headers</a>: <i><a href="headers.md">Headers</a></i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ConnectionName

Human-readable label that identifies the stream connection. In the case of the Sample type, this is the name of the sample source.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### InstanceName

Human-readable label that identifies the stream instance. Deprecated: Use WorkspaceName instead.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### WorkspaceName

Human-readable label that identifies the stream workspace.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Type

Type of the connection. Can be Cluster, Kafka, Sample, AWSLambda, or Https.

_Required_: Yes

_Type_: String

_Allowed Values_: <code>Kafka</code> | <code>Cluster</code> | <code>Sample</code> | <code>AWSLambda</code> | <code>Https</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ClusterName

Name of the cluster configured for this connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterProjectId

Unique 24-hexadecimal digit string that identifies the project containing the cluster for cross-project cluster connections.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DbRoleToExecute

The name of a Built in or Custom DB Role to connect to an Atlas Cluster.

_Required_: No

_Type_: <a href="dbroletoexecute.md">DBRoleToExecute</a>

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

#### Config

A map of Kafka key-value pairs for optional configuration. This is a flat object, and keys can have '.' characters.

_Required_: No

_Type_: <a href="config.md">Config</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Networking

Networking configuration for AWS PrivateLink connections.

_Required_: No

_Type_: <a href="networking.md">Networking</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Aws

AWS Lambda connection configuration.

_Required_: No

_Type_: <a href="aws.md">Aws</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Url

URL endpoint for HTTPS type connections.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Headers

HTTP headers for HTTPS type connections.

_Required_: No

_Type_: <a href="headers.md">Headers</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

