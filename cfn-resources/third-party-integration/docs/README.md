# MongoDB::Atlas::ThirdPartyIntegration

Returns, adds, edits, and removes third-party service integration configurations. MongoDB Cloud sends alerts to each third-party service that you configure.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ThirdPartyIntegration",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#type" title="Type">Type</a>" : <i>String</i>,
        "<a href="#apikey" title="ApiKey">ApiKey</a>" : <i>String</i>,
        "<a href="#region" title="Region">Region</a>" : <i>String</i>,
        "<a href="#servicekey" title="ServiceKey">ServiceKey</a>" : <i>String</i>,
        "<a href="#apitoken" title="ApiToken">ApiToken</a>" : <i>String</i>,
        "<a href="#teamname" title="TeamName">TeamName</a>" : <i>String</i>,
        "<a href="#channelname" title="ChannelName">ChannelName</a>" : <i>String</i>,
        "<a href="#routingkey" title="RoutingKey">RoutingKey</a>" : <i>String</i>,
        "<a href="#url" title="Url">Url</a>" : <i>String</i>,
        "<a href="#secret" title="Secret">Secret</a>" : <i>String</i>,
        "<a href="#microsoftteamswebhookurl" title="MicrosoftTeamsWebhookUrl">MicrosoftTeamsWebhookUrl</a>" : <i>String</i>,
        "<a href="#username" title="UserName">UserName</a>" : <i>String</i>,
        "<a href="#password" title="Password">Password</a>" : <i>String</i>,
        "<a href="#servicediscovery" title="ServiceDiscovery">ServiceDiscovery</a>" : <i>String</i>,
        "<a href="#scheme" title="Scheme">Scheme</a>" : <i>String</i>,
        "<a href="#enabled" title="Enabled">Enabled</a>" : <i>Boolean</i>,
        "<a href="#listenaddress" title="ListenAddress">ListenAddress</a>" : <i>String</i>,
        "<a href="#tlspempath" title="TlsPemPath">TlsPemPath</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ThirdPartyIntegration
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#type" title="Type">Type</a>: <i>String</i>
    <a href="#apikey" title="ApiKey">ApiKey</a>: <i>String</i>
    <a href="#region" title="Region">Region</a>: <i>String</i>
    <a href="#servicekey" title="ServiceKey">ServiceKey</a>: <i>String</i>
    <a href="#apitoken" title="ApiToken">ApiToken</a>: <i>String</i>
    <a href="#teamname" title="TeamName">TeamName</a>: <i>String</i>
    <a href="#channelname" title="ChannelName">ChannelName</a>: <i>String</i>
    <a href="#routingkey" title="RoutingKey">RoutingKey</a>: <i>String</i>
    <a href="#url" title="Url">Url</a>: <i>String</i>
    <a href="#secret" title="Secret">Secret</a>: <i>String</i>
    <a href="#microsoftteamswebhookurl" title="MicrosoftTeamsWebhookUrl">MicrosoftTeamsWebhookUrl</a>: <i>String</i>
    <a href="#username" title="UserName">UserName</a>: <i>String</i>
    <a href="#password" title="Password">Password</a>: <i>String</i>
    <a href="#servicediscovery" title="ServiceDiscovery">ServiceDiscovery</a>: <i>String</i>
    <a href="#scheme" title="Scheme">Scheme</a>: <i>String</i>
    <a href="#enabled" title="Enabled">Enabled</a>: <i>Boolean</i>
    <a href="#listenaddress" title="ListenAddress">ListenAddress</a>: <i>String</i>
    <a href="#tlspempath" title="TlsPemPath">TlsPemPath</a>: <i>String</i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Type

Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.

_Required_: No

_Type_: String

_Allowed Values_: <code>PAGER_DUTY</code> | <code>MICROSOFT_TEAMS</code> | <code>SLACK</code> | <code>DATADOG</code> | <code>OPS_GENIE</code> | <code>VICTOR_OPS</code> | <code>WEBHOOK</code> | <code>PROMETHEUS</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ApiKey

Key that allows MongoDB Cloud to access your Opsgenie/Datadog account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Region

Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie/Datadog API.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ServiceKey

Service key associated with your PagerDuty account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiToken

Key that allows MongoDB Cloud to access your Slack account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TeamName

Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ChannelName

Name of the Slack channel to which MongoDB Cloud sends alert notifications.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoutingKey

Routing key associated with your Splunk On-Call account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Url

Endpoint web address to which MongoDB Cloud sends notifications.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Secret

Parameter returned if someone configure this webhook with a secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MicrosoftTeamsWebhookUrl

Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### UserName

Human-readable label that identifies your Prometheus incoming webhook.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Password

Password required for your integration with Prometheus

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ServiceDiscovery

Desired method to discover the Prometheus service.

_Required_: No

_Type_: String

_Allowed Values_: <code>http</code> | <code>file</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Scheme

Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.

_Required_: No

_Type_: String

_Allowed Values_: <code>http</code> | <code>https</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Enabled

Flag that indicates whether someone has activated the Prometheus integration.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ListenAddress

Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TlsPemPath

Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

