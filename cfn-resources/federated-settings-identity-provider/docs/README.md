# MongoDB::Atlas::FederatedSettingsIdentityProvider

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedSettingsIdentityProvider",
    "Properties" : {
        "<a href="#ssodebugenabled" title="SsoDebugEnabled">SsoDebugEnabled</a>" : <i>Boolean</i>,
        "<a href="#associateddomains" title="AssociatedDomains">AssociatedDomains</a>" : <i>[ String, ... ]</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>,
        "<a href="#issueruri" title="IssuerUri">IssuerUri</a>" : <i>String</i>,
        "<a href="#requestbinding" title="RequestBinding">RequestBinding</a>" : <i>String</i>,
        "<a href="#responsesignaturealgorithm" title="ResponseSignatureAlgorithm">ResponseSignatureAlgorithm</a>" : <i>String</i>,
        "<a href="#ssourl" title="SsoUrl">SsoUrl</a>" : <i>String</i>,
        "<a href="#displayname" title="DisplayName">DisplayName</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FederatedSettingsIdentityProvider
Properties:
    <a href="#ssodebugenabled" title="SsoDebugEnabled">SsoDebugEnabled</a>: <i>Boolean</i>
    <a href="#associateddomains" title="AssociatedDomains">AssociatedDomains</a>: <i>
      - String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
    <a href="#issueruri" title="IssuerUri">IssuerUri</a>: <i>String</i>
    <a href="#requestbinding" title="RequestBinding">RequestBinding</a>: <i>String</i>
    <a href="#responsesignaturealgorithm" title="ResponseSignatureAlgorithm">ResponseSignatureAlgorithm</a>: <i>String</i>
    <a href="#ssourl" title="SsoUrl">SsoUrl</a>: <i>String</i>
    <a href="#displayname" title="DisplayName">DisplayName</a>: <i>String</i>
</pre>

## Properties

#### SsoDebugEnabled

Flag that indicates whether the identity provider has SSO debug enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AssociatedDomains

List that contains the domains associated with the identity provider.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

String enum that indicates whether the identity provider is active.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IssuerUri

Unique string that identifies the issuer of the SAML Assertion.

_Required_: No

_Type_: String

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

#### SsoUrl

URL that points to the receiver of the SAML authentication request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DisplayName

Human-readable label that identifies the IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies your federation.

#### OktaIdpId

Unique 20-hexadecimal digit string that identifies the identity provider.

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

