# MongoDB::Atlas::FederatedSettingOrgConfigs

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedSettingOrgConfigs",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#domainallowlist" title="DomainAllowList">DomainAllowList</a>" : <i>[ String, ... ]</i>,
        "<a href="#domainrestrictionenabled" title="DomainRestrictionEnabled">DomainRestrictionEnabled</a>" : <i>Boolean</i>,
        "<a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>" : <i>String</i>,
        "<a href="#postauthrolegrants" title="PostAuthRoleGrants">PostAuthRoleGrants</a>" : <i>[ String, ... ]</i>,
        "<a href="#rolemappings" title="RoleMappings">RoleMappings</a>" : <i>[ <a href="rolemappingview.md">RoleMappingView</a>, ... ]</i>,
        "<a href="#userconflicts" title="UserConflicts">UserConflicts</a>" : <i>[ <a href="federateduserview.md">FederatedUserView</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FederatedSettingOrgConfigs
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#domainallowlist" title="DomainAllowList">DomainAllowList</a>: <i>
      - String</i>
    <a href="#domainrestrictionenabled" title="DomainRestrictionEnabled">DomainRestrictionEnabled</a>: <i>Boolean</i>
    <a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>: <i>String</i>
    <a href="#postauthrolegrants" title="PostAuthRoleGrants">PostAuthRoleGrants</a>: <i>
      - String</i>
    <a href="#rolemappings" title="RoleMappings">RoleMappings</a>: <i>
      - <a href="rolemappingview.md">RoleMappingView</a></i>
    <a href="#userconflicts" title="UserConflicts">UserConflicts</a>: <i>
      - <a href="federateduserview.md">FederatedUserView</a></i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DomainAllowList

Approved domains that restrict users who can join the organization based on their email address.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DomainRestrictionEnabled

Value that indicates whether domain restriction is enabled for this connected org.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IdentityProviderId

Unique 20-hexadecimal digit string that identifies the identity provider that this connected org config is associated with.

_Required_: No

_Type_: String

_Minimum_: <code>20</code>

_Maximum_: <code>20</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PostAuthRoleGrants

Atlas roles that are granted to a user in this organization after authenticating.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleMappings

Role mappings that are configured in this organization.

_Required_: No

_Type_: List of <a href="rolemappingview.md">RoleMappingView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### UserConflicts

List that contains the users who have an email address that doesn't match any domain on the allowed list.

_Required_: No

_Type_: List of <a href="federateduserview.md">FederatedUserView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### OrgId

Returns the <code>OrgId</code> value.

#### OrgId

Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.

#### TestMode

Returns the <code>TestMode</code> value.

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies your federation.

#### GroupId

Returns the <code>GroupId</code> value.

#### Id

Returns the <code>Id</code> value.

#### UserId

Returns the <code>UserId</code> value.

