# MongoDB::Atlas::StreamWorkspace StreamsKafkaSecurity

Properties for the secure transport connection to Kafka. For SSL, this can include the trusted certificate to use.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#brokerpubliccertificate" title="BrokerPublicCertificate">BrokerPublicCertificate</a>" : <i>String</i>,
    "<a href="#protocol" title="Protocol">Protocol</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#brokerpubliccertificate" title="BrokerPublicCertificate">BrokerPublicCertificate</a>: <i>String</i>
<a href="#protocol" title="Protocol">Protocol</a>: <i>String</i>
</pre>

## Properties

#### BrokerPublicCertificate

A trusted, public x509 certificate for connecting to Kafka over SSL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Protocol

Describes the transport type. Can be either PLAINTEXT or SSL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

