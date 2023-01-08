# MongoDB::Atlas::federatedsettingsidentityprovider

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::federatedsettingsidentityprovider",
    "Properties" : {
        "<a href="#acsurl" title="AcsUrl">AcsUrl</a>" : <i>String</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>,
        "<a href="#associateddomains" title="AssociatedDomains">AssociatedDomains</a>" : <i>[ String, ... ]</i>,
        "<a href="#associatedorgs" title="AssociatedOrgs">AssociatedOrgs</a>" : <i>[ <a href="connectedorgconfigview.md">ConnectedOrgConfigView</a>, ... ]</i>,
        "<a href="#audienceuri" title="AudienceUri">AudienceUri</a>" : <i>String</i>,
        "<a href="#displayname" title="DisplayName">DisplayName</a>" : <i>String</i>,
        "<a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>" : <i>String</i>,
        "<a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>" : <i>String</i>,
        "<a href="#issueruri" title="IssuerUri">IssuerUri</a>" : <i>String</i>,
        "<a href="#oktaidpid" title="OktaIdpId">OktaIdpId</a>" : <i>String</i>,
        "<a href="#pemfileinfo" title="PemFileInfo">PemFileInfo</a>" : <i>[ Map, ... ]</i>,
        "<a href="#requestbinding" title="RequestBinding">RequestBinding</a>" : <i>String</i>,
        "<a href="#responsesignaturealgorithm" title="ResponseSignatureAlgorithm">ResponseSignatureAlgorithm</a>" : <i>String</i>,
        "<a href="#ssodebugenabled" title="SsoDebugEnabled">SsoDebugEnabled</a>" : <i>Boolean</i>,
        "<a href="#ssourl" title="SsoUrl">SsoUrl</a>" : <i>String</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::federatedsettingsidentityprovider
Properties:
    <a href="#acsurl" title="AcsUrl">AcsUrl</a>: <i>String</i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>
    <a href="#associateddomains" title="AssociatedDomains">AssociatedDomains</a>: <i>
      - String</i>
    <a href="#associatedorgs" title="AssociatedOrgs">AssociatedOrgs</a>: <i>
      - <a href="connectedorgconfigview.md">ConnectedOrgConfigView</a></i>
    <a href="#audienceuri" title="AudienceUri">AudienceUri</a>: <i>String</i>
    <a href="#displayname" title="DisplayName">DisplayName</a>: <i>String</i>
    <a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>: <i>String</i>
    <a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>: <i>String</i>
    <a href="#issueruri" title="IssuerUri">IssuerUri</a>: <i>String</i>
    <a href="#oktaidpid" title="OktaIdpId">OktaIdpId</a>: <i>String</i>
    <a href="#pemfileinfo" title="PemFileInfo">PemFileInfo</a>: <i>
      - Map</i>
    <a href="#requestbinding" title="RequestBinding">RequestBinding</a>: <i>String</i>
    <a href="#responsesignaturealgorithm" title="ResponseSignatureAlgorithm">ResponseSignatureAlgorithm</a>: <i>String</i>
    <a href="#ssodebugenabled" title="SsoDebugEnabled">SsoDebugEnabled</a>: <i>Boolean</i>
    <a href="#ssourl" title="SsoUrl">SsoUrl</a>: <i>String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
</pre>

## Properties

#### AcsUrl

URL that points to where to send the SAML response.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">ApiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AssociatedDomains

List that contains the domains associated with the identity provider.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AssociatedOrgs

List that contains the connected organization configurations associated with the identity provider.

_Required_: No

_Type_: List of <a href="connectedorgconfigview.md">ConnectedOrgConfigView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AudienceUri

Unique string that identifies the intended audience of the SAML assertion.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DisplayName

Human-readable label that identifies the identity provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies your federation.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IdentityProviderId

Unique 20-hexadecimal digit string that identifies the identity provider.

_Required_: No

_Type_: String

_Minimum Length_: <code>20</code>

_Maximum Length_: <code>20</code>

_Pattern_: <code>^([a-f0-9]{20})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IssuerUri

Unique string that identifies the issuer of the SAML Assertion.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OktaIdpId

Unique 20-hexadecimal digit string that identifies the identity provider.

_Required_: No

_Type_: String

_Minimum Length_: <code>20</code>

_Maximum Length_: <code>20</code>

_Pattern_: <code>^([a-f0-9]{20})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PemFileInfo

List that contains the PEM file information for the identity provider's current certificates.

_Required_: No

_Type_: List of Map

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RequestBinding

SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ResponseSignatureAlgorithm

Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SsoDebugEnabled

Flag that indicates whether the identity provider has SSO debug enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SsoUrl

URL that points to the receiver of the SAML authentication request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

String enum that indicates whether the identity provider is active.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### GroupId

Returns the <code>GroupId</code> value.

#### OrgId

Returns the <code>OrgId</code> value.

#### Id

Returns the <code>Id</code> value.

#### UserId

Returns the <code>UserId</code> value.

#### OrgId

Returns the <code>OrgId</code> value.

