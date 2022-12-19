# MongoDB::Atlas::ldapconfiguration

Returns, edits, verifies, and removes LDAP configurations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ldapconfiguration",
    "Properties" : {
        "<a href="#bindusername" title="BindUsername">BindUsername</a>" : <i>String</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>,
        "<a href="#hostname" title="Hostname">Hostname</a>" : <i>String</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>,
        "<a href="#authenticationenabled" title="AuthenticationEnabled">AuthenticationEnabled</a>" : <i>Boolean</i>,
        "<a href="#authorizationenabled" title="AuthorizationEnabled">AuthorizationEnabled</a>" : <i>Boolean</i>,
        "<a href="#cacertificate" title="CaCertificate">CaCertificate</a>" : <i>String</i>,
        "<a href="#authzquerytemplate" title="AuthzQueryTemplate">AuthzQueryTemplate</a>" : <i>String</i>,
        "<a href="#bindpassword" title="BindPassword">BindPassword</a>" : <i>String</i>,
        "<a href="#customerx509" title="CustomerX509">CustomerX509</a>" : <i><a href="apiatlascustomerx509view.md">ApiAtlasCustomerX509View</a></i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
        "<a href="#port" title="Port">Port</a>" : <i>Integer</i>,
        "<a href="#usertodnmapping" title="UserToDNMapping">UserToDNMapping</a>" : <i>[ <a href="apiatlasndsusertodnmappingview.md">ApiAtlasNDSUserToDNMappingView</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ldapconfiguration
Properties:
    <a href="#bindusername" title="BindUsername">BindUsername</a>: <i>String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
    <a href="#hostname" title="Hostname">Hostname</a>: <i>String</i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>
    <a href="#authenticationenabled" title="AuthenticationEnabled">AuthenticationEnabled</a>: <i>Boolean</i>
    <a href="#authorizationenabled" title="AuthorizationEnabled">AuthorizationEnabled</a>: <i>Boolean</i>
    <a href="#cacertificate" title="CaCertificate">CaCertificate</a>: <i>String</i>
    <a href="#authzquerytemplate" title="AuthzQueryTemplate">AuthzQueryTemplate</a>: <i>String</i>
    <a href="#bindpassword" title="BindPassword">BindPassword</a>: <i>String</i>
    <a href="#customerx509" title="CustomerX509">CustomerX509</a>: <i><a href="apiatlascustomerx509view.md">ApiAtlasCustomerX509View</a></i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
    <a href="#port" title="Port">Port</a>: <i>Integer</i>
    <a href="#usertodnmapping" title="UserToDNMapping">UserToDNMapping</a>: <i>
      - <a href="apiatlasndsusertodnmappingview.md">ApiAtlasNDSUserToDNMappingView</a></i>
</pre>

## Properties

#### BindUsername

Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. LDAP distinguished names must be formatted according to RFC 2253.

_Required_: Yes

_Type_: String

_Pattern_: <code>^(?:(?<cn>CN=(?<name>[^,]*)),)?(?:(?<path>(?:(?:CN|OU)=[^,]+,?)+),)?(?<domain>(?:DC=[^,]+,?)+)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

The current status of the LDAP over TLS/SSL configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Hostname

Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster.

_Required_: No

_Type_: String

_Pattern_: <code>^([0-9]{1,3}\.){3}[0-9]{1,3}|([0-9a-f]{1,4}\:){7}([0-9a-f]{1,4})|(([a-z0-9]+\.){1,10}[a-z]+)?$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeydefinition.md">ApiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthenticationEnabled

Flag that indicates whether users can authenticate using an Lightweight Directory Access Protocol (LDAP) host.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthorizationEnabled

Flag that indicates whether users can authorize access to MongoDB Cloud resources using an Lightweight Directory Access Protocol (LDAP) host.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CaCertificate

Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: `"caCertificate": ""`

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthzQueryTemplate

Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user. MongoDB Cloud uses this parameter only for user authorization. Use the `{USER}` placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query according to [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BindPassword

Password that MongoDB Cloud uses to authenticate the **bindUsername**.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomerX509

_Required_: No

_Type_: <a href="apiatlascustomerx509view.md">ApiAtlasCustomerX509View</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Port

Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### UserToDNMapping

User-to-Distinguished Name (DN) map that MongoDB Cloud uses to transform a Lightweight Directory Access Protocol (LDAP) username into an LDAP DN.

_Required_: No

_Type_: List of <a href="apiatlasndsusertodnmappingview.md">ApiAtlasNDSUserToDNMappingView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the GroupId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Links

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

#### Links

Returns the <code>Links</code> value.

