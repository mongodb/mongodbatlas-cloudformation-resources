# MongoDB::Atlas::FederatedSettingsIdentityProvider

Resource for managing MongoDB Atlas Federated Settings Identity Providers (SAML and OIDC) within an Atlas federation.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedSettingsIdentityProvider",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>" : <i>String</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#issueruri" title="IssuerUri">IssuerUri</a>" : <i>String</i>,
        "<a href="#requestbinding" title="RequestBinding">RequestBinding</a>" : <i>String</i>,
        "<a href="#responsesignaturealgorithm" title="ResponseSignatureAlgorithm">ResponseSignatureAlgorithm</a>" : <i>String</i>,
        "<a href="#associateddomains" title="AssociatedDomains">AssociatedDomains</a>" : <i>[ String, ... ]</i>,
        "<a href="#ssodebugenabled" title="SsoDebugEnabled">SsoDebugEnabled</a>" : <i>Boolean</i>,
        "<a href="#ssourl" title="SsoUrl">SsoUrl</a>" : <i>String</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>,
        "<a href="#protocol" title="Protocol">Protocol</a>" : <i>String</i>,
        "<a href="#audience" title="Audience">Audience</a>" : <i>String</i>,
        "<a href="#clientid" title="ClientId">ClientId</a>" : <i>String</i>,
        "<a href="#groupsclaim" title="GroupsClaim">GroupsClaim</a>" : <i>String</i>,
        "<a href="#requestedscopes" title="RequestedScopes">RequestedScopes</a>" : <i>[ String, ... ]</i>,
        "<a href="#userclaim" title="UserClaim">UserClaim</a>" : <i>String</i>,
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#authorizationtype" title="AuthorizationType">AuthorizationType</a>" : <i>String</i>,
        "<a href="#idptype" title="IdpType">IdpType</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FederatedSettingsIdentityProvider
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>: <i>String</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#issueruri" title="IssuerUri">IssuerUri</a>: <i>String</i>
    <a href="#requestbinding" title="RequestBinding">RequestBinding</a>: <i>String</i>
    <a href="#responsesignaturealgorithm" title="ResponseSignatureAlgorithm">ResponseSignatureAlgorithm</a>: <i>String</i>
    <a href="#associateddomains" title="AssociatedDomains">AssociatedDomains</a>: <i>
      - String</i>
    <a href="#ssodebugenabled" title="SsoDebugEnabled">SsoDebugEnabled</a>: <i>Boolean</i>
    <a href="#ssourl" title="SsoUrl">SsoUrl</a>: <i>String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
    <a href="#protocol" title="Protocol">Protocol</a>: <i>String</i>
    <a href="#audience" title="Audience">Audience</a>: <i>String</i>
    <a href="#clientid" title="ClientId">ClientId</a>: <i>String</i>
    <a href="#groupsclaim" title="GroupsClaim">GroupsClaim</a>: <i>String</i>
    <a href="#requestedscopes" title="RequestedScopes">RequestedScopes</a>: <i>
      - String</i>
    <a href="#userclaim" title="UserClaim">UserClaim</a>: <i>String</i>
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#authorizationtype" title="AuthorizationType">AuthorizationType</a>: <i>String</i>
    <a href="#idptype" title="IdpType">IdpType</a>: <i>String</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secrets Manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies your federation.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Name

Human-readable name (display name) of the identity provider.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IssuerUri

Issuer URI of the identity provider.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RequestBinding

SAML request binding.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ResponseSignatureAlgorithm

SAML response signature algorithm.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AssociatedDomains

List of associated domains for this identity provider.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SsoDebugEnabled

Flag that indicates whether to enable SSO debug.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SsoUrl

SSO URL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

Identity provider status.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Protocol

Identity provider protocol.

_Required_: No

_Type_: String

_Allowed Values_: <code>SAML</code> | <code>OIDC</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Audience

OIDC audience.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClientId

OIDC client ID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupsClaim

OIDC groups claim.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RequestedScopes

OIDC requested scopes.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### UserClaim

OIDC user claim.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

Description of the identity provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthorizationType

OIDC authorization type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IdpType

Identity provider type (for OIDC).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### OktaIdpId

Legacy identity provider identifier (Okta IdP ID). Returned by the API.

#### IdpId

Unique identifier of the identity provider. Returned by the API.

