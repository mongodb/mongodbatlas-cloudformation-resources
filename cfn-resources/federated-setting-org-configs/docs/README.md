# MongoDB::Atlas::FederatedSettingOrgConfigs

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedSettingOrgConfigs",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#domainallowlist" title="DomainAllowList">DomainAllowList</a>" : <i>[ String, ... ]</i>,
        "<a href="#domainrestrictionenabled" title="DomainRestrictionEnabled">DomainRestrictionEnabled</a>" : <i>Boolean</i>,
        "<a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>" : <i>String</i>,
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
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
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#domainallowlist" title="DomainAllowList">DomainAllowList</a>: <i>
      - String</i>
    <a href="#domainrestrictionenabled" title="DomainRestrictionEnabled">DomainRestrictionEnabled</a>: <i>Boolean</i>
    <a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>: <i>String</i>
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
    <a href="#postauthrolegrants" title="PostAuthRoleGrants">PostAuthRoleGrants</a>: <i>
      - String</i>
    <a href="#rolemappings" title="RoleMappings">RoleMappings</a>: <i>
      - <a href="rolemappingview.md">RoleMappingView</a></i>
    <a href="#userconflicts" title="UserConflicts">UserConflicts</a>: <i>
      - <a href="federateduserview.md">FederatedUserView</a></i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

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

_Minimum Length_: <code>20</code>

_Maximum Length_: <code>20</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

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

#### TestMode

Returns the <code>TestMode</code> value.

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies your federation.

#### ProjectId

Returns the <code>ProjectId</code> value.

#### Id

Returns the <code>Id</code> value.

#### UserId

Returns the <code>UserId</code> value.

