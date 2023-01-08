# MongoDB::Atlas::federatedsettingsidentityprovider ConnectedOrgConfigView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>,
    "<a href="#domainallowlist" title="DomainAllowList">DomainAllowList</a>" : <i>[ String, ... ]</i>,
    "<a href="#domainrestrictionenabled" title="DomainRestrictionEnabled">DomainRestrictionEnabled</a>" : <i>Boolean</i>,
    "<a href="#identityproviderid" title="IdentityProviderId">IdentityProviderId</a>" : <i>String</i>,
    "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
    "<a href="#postauthrolegrants" title="PostAuthRoleGrants">PostAuthRoleGrants</a>" : <i>[ String, ... ]</i>,
    "<a href="#rolemappings" title="RoleMappings">RoleMappings</a>" : <i>[ <a href="rolemappingview.md">RoleMappingView</a>, ... ]</i>,
    "<a href="#userconflicts" title="UserConflicts">UserConflicts</a>" : <i>[ <a href="federateduserview.md">FederatedUserView</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>
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

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">ApiKeyDefinition</a>

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

_Minimum Length_: <code>20</code>

_Maximum Length_: <code>20</code>

_Pattern_: <code>^([a-f0-9]{20})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the connected organization configuration.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

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

