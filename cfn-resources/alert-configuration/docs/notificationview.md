# MongoDB::Atlas::AlertConfiguration NotificationView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#apitoken" title="ApiToken">ApiToken</a>" : <i>String</i>,
    "<a href="#channelname" title="ChannelName">ChannelName</a>" : <i>String</i>,
    "<a href="#datadogapikey" title="DatadogApiKey">DatadogApiKey</a>" : <i>String</i>,
    "<a href="#datadogregion" title="DatadogRegion">DatadogRegion</a>" : <i>String</i>,
    "<a href="#delaymin" title="DelayMin">DelayMin</a>" : <i>Integer</i>,
    "<a href="#emailaddress" title="EmailAddress">EmailAddress</a>" : <i>String</i>,
    "<a href="#emailenabled" title="EmailEnabled">EmailEnabled</a>" : <i>Boolean</i>,
    "<a href="#intervalmin" title="IntervalMin">IntervalMin</a>" : <i>Double</i>,
    "<a href="#microsoftteamswebhookurl" title="MicrosoftTeamsWebhookUrl">MicrosoftTeamsWebhookUrl</a>" : <i>String</i>,
    "<a href="#mobilenumber" title="MobileNumber">MobileNumber</a>" : <i>String</i>,
    "<a href="#notificationtoken" title="NotificationToken">NotificationToken</a>" : <i>String</i>,
    "<a href="#opsgenieapikey" title="OpsGenieApiKey">OpsGenieApiKey</a>" : <i>String</i>,
    "<a href="#opsgenieregion" title="OpsGenieRegion">OpsGenieRegion</a>" : <i>String</i>,
    "<a href="#orgname" title="OrgName">OrgName</a>" : <i>String</i>,
    "<a href="#roles" title="Roles">Roles</a>" : <i>[ String, ... ]</i>,
    "<a href="#roomname" title="RoomName">RoomName</a>" : <i>String</i>,
    "<a href="#servicekey" title="ServiceKey">ServiceKey</a>" : <i>String</i>,
    "<a href="#severity" title="Severity">Severity</a>" : <i>String</i>,
    "<a href="#smsenabled" title="SmsEnabled">SmsEnabled</a>" : <i>Boolean</i>,
    "<a href="#teamid" title="TeamId">TeamId</a>" : <i>String</i>,
    "<a href="#teamname" title="TeamName">TeamName</a>" : <i>String</i>,
    "<a href="#typename" title="TypeName">TypeName</a>" : <i>String</i>,
    "<a href="#username" title="Username">Username</a>" : <i>String</i>,
    "<a href="#victoropsapikey" title="VictorOpsApiKey">VictorOpsApiKey</a>" : <i>String</i>,
    "<a href="#victoropsroutingkey" title="VictorOpsRoutingKey">VictorOpsRoutingKey</a>" : <i>String</i>,
    "<a href="#webhooksecret" title="WebhookSecret">WebhookSecret</a>" : <i>String</i>,
    "<a href="#webhookurl" title="WebhookUrl">WebhookUrl</a>" : <i>String</i>,
    "<a href="#notifierid" title="NotifierId">NotifierId</a>" : <i>String</i>,
    "<a href="#integrationid" title="IntegrationId">IntegrationId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#apitoken" title="ApiToken">ApiToken</a>: <i>String</i>
<a href="#channelname" title="ChannelName">ChannelName</a>: <i>String</i>
<a href="#datadogapikey" title="DatadogApiKey">DatadogApiKey</a>: <i>String</i>
<a href="#datadogregion" title="DatadogRegion">DatadogRegion</a>: <i>String</i>
<a href="#delaymin" title="DelayMin">DelayMin</a>: <i>Integer</i>
<a href="#emailaddress" title="EmailAddress">EmailAddress</a>: <i>String</i>
<a href="#emailenabled" title="EmailEnabled">EmailEnabled</a>: <i>Boolean</i>
<a href="#intervalmin" title="IntervalMin">IntervalMin</a>: <i>Double</i>
<a href="#microsoftteamswebhookurl" title="MicrosoftTeamsWebhookUrl">MicrosoftTeamsWebhookUrl</a>: <i>String</i>
<a href="#mobilenumber" title="MobileNumber">MobileNumber</a>: <i>String</i>
<a href="#notificationtoken" title="NotificationToken">NotificationToken</a>: <i>String</i>
<a href="#opsgenieapikey" title="OpsGenieApiKey">OpsGenieApiKey</a>: <i>String</i>
<a href="#opsgenieregion" title="OpsGenieRegion">OpsGenieRegion</a>: <i>String</i>
<a href="#orgname" title="OrgName">OrgName</a>: <i>String</i>
<a href="#roles" title="Roles">Roles</a>: <i>
      - String</i>
<a href="#roomname" title="RoomName">RoomName</a>: <i>String</i>
<a href="#servicekey" title="ServiceKey">ServiceKey</a>: <i>String</i>
<a href="#severity" title="Severity">Severity</a>: <i>String</i>
<a href="#smsenabled" title="SmsEnabled">SmsEnabled</a>: <i>Boolean</i>
<a href="#teamid" title="TeamId">TeamId</a>: <i>String</i>
<a href="#teamname" title="TeamName">TeamName</a>: <i>String</i>
<a href="#typename" title="TypeName">TypeName</a>: <i>String</i>
<a href="#username" title="Username">Username</a>: <i>String</i>
<a href="#victoropsapikey" title="VictorOpsApiKey">VictorOpsApiKey</a>: <i>String</i>
<a href="#victoropsroutingkey" title="VictorOpsRoutingKey">VictorOpsRoutingKey</a>: <i>String</i>
<a href="#webhooksecret" title="WebhookSecret">WebhookSecret</a>: <i>String</i>
<a href="#webhookurl" title="WebhookUrl">WebhookUrl</a>: <i>String</i>
<a href="#notifierid" title="NotifierId">NotifierId</a>: <i>String</i>
<a href="#integrationid" title="IntegrationId">IntegrationId</a>: <i>String</i>
</pre>

## Properties

#### ApiToken

Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack. The resource requires this parameter when '"notifications.typeName" : "SLACK"'. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ChannelName

Name of the Slack channel to which MongoDB Cloud sends alert notifications. The resource requires this parameter when '"notifications.typeName" : "SLACK"'.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DatadogApiKey

Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. You can find this API key in the Datadog dashboard. The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.

_Required_: No

_Type_: String

_Pattern_: <code>^[0-9a-f]{32}$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DatadogRegion

Datadog region that indicates which API Uniform Resource Locator (URL) to use. The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.

_Required_: No

_Type_: String

_Allowed Values_: <code>EU</code> | <code>US</code>

_Minimum Length_: <code>2</code>

_Maximum Length_: <code>2</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DelayMin

Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EmailAddress

Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when '"notifications.typeName" : "EMAIL"'. You don't need to set this value to send emails to individual or groups of MongoDB Cloud users including:

- specific MongoDB Cloud users ('"notifications.typeName" : "USER"')
- MongoDB Cloud users with specific project roles ('"notifications.typeName" : "GROUP"')
- MongoDB Cloud users with specific organization roles ('"notifications.typeName" : "ORG"')
- MongoDB Cloud teams ('"notifications.typeName" : "TEAM"')

To send emails to one MongoDB Cloud user or grouping of users, set the **notifications.emailEnabled** parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EmailEnabled

Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:

- '"notifications.typeName" : "ORG"'
- '"notifications.typeName" : "GROUP"'
- '"notifications.typeName" : "USER"'

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IntervalMin

Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.

PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MicrosoftTeamsWebhookUrl

Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. The resource requires this parameter when '"notifications.typeName" : "MICROSOFT_TEAMS"'. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MobileNumber

Mobile phone number to which MongoDB Cloud sends alert notifications. The resource requires this parameter when '"notifications.typeName" : "SMS"'.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NotificationToken

HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. The resource requires this parameter when '"notifications.typeName" : "HIP_CHAT"'". If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OpsGenieApiKey

API Key that MongoDB Cloud needs to send this notification via Opsgenie. The resource requires this parameter when '"notifications.typeName" : "OPS_GENIE"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OpsGenieRegion

Opsgenie region that indicates which API Uniform Resource Locator (URL) to use.

_Required_: No

_Type_: String

_Allowed Values_: <code>EU</code> | <code>US</code>

_Minimum Length_: <code>2</code>

_Maximum Length_: <code>2</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgName

Flowdock organization name to which MongoDB Cloud sends alert notifications. This name appears after 'www.flowdock.com/app/' in the Uniform Resource Locator (URL) path. The resource requires this parameter when '"notifications.typeName" : "FLOWDOCK"'.

_Required_: No

_Type_: String

_Minimum Length_: <code>1</code>

_Maximum Length_: <code>64</code>

_Pattern_: <code>^([a-z\-]+)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Roles

List that contains the one or more organization or project roles that receive the configured alert. The resource requires this parameter when '"notifications.typeName" : "GROUP"' or '"notifications.typeName" : "ORG"'. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoomName

HipChat API room name to which MongoDB Cloud sends alert notifications. The resource requires this parameter when '"notifications.typeName" : "HIP_CHAT"'".

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ServiceKey

PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty. The resource requires this parameter when '"notifications.typeName" : "PAGER_DUTY"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Severity

Degree of seriousness given to this notification.

_Required_: No

_Type_: String

_Allowed Values_: <code>CRITICAL</code> | <code>ERROR</code> | <code>WARNING</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SmsEnabled

Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:

- '"notifications.typeName" : "ORG"'
- '"notifications.typeName" : "GROUP"'
- '"notifications.typeName" : "USER"'

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TeamId

Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team. The resource requires this parameter when '"notifications.typeName" : "TEAM"'.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TeamName

Name of the MongoDB Cloud team that receives this notification. The resource requires this parameter when '"notifications.typeName" : "TEAM"'.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TypeName

Human-readable label that displays the alert notification type.

_Required_: No

_Type_: String

_Allowed Values_: <code>DATADOG</code> | <code>EMAIL</code> | <code>FLOWDOCK</code> | <code>GROUP</code> | <code>MICROSOFT_TEAMS</code> | <code>OPS_GENIE</code> | <code>ORG</code> | <code>PAGER_DUTY</code> | <code>PROMETHEUS</code> | <code>SLACK</code> | <code>SMS</code> | <code>TEAM</code> | <code>USER</code> | <code>VICTOR_OPS</code> | <code>WEBHOOK</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Username

MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications. Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when '"notifications.typeName" : "USER"'.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### VictorOpsApiKey

API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when '"notifications.typeName" : "VICTOR_OPS"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### VictorOpsRoutingKey

Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when '"notifications.typeName" : "VICTOR_OPS"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### WebhookSecret

An optional field for your webhook secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### WebhookUrl

Your webhook URL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NotifierId

Unique 24-hexadecimal digit string that identifies the notifier to use for this alert configuration.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IntegrationId

Unique 24-hexadecimal digit string that identifies the third party integration to use for this alert configuration.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

