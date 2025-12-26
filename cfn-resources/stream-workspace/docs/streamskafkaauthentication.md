# MongoDB::Atlas::StreamWorkspace StreamsKafkaAuthentication

User credentials required to connect to a Kafka Cluster. Includes the authentication type, as well as the parameters for that authentication mode.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#mechanism" title="Mechanism">Mechanism</a>" : <i>String</i>,
    "<a href="#username" title="Username">Username</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#mechanism" title="Mechanism">Mechanism</a>: <i>String</i>
<a href="#username" title="Username">Username</a>: <i>String</i>
</pre>

## Properties

#### Mechanism

Style of authentication. Can be one of PLAIN, SCRAM-256, or SCRAM-512.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Username

Username of the account to connect to the Kafka cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

