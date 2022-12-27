# MongoDB::Atlas::FederatedSettingsOrgRoleMapping

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedSettingsOrgRoleMapping",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#externalgroupname" title="ExternalGroupName">ExternalGroupName</a>" : <i>String</i>,
        "<a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>" : <i>String</i>,
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
        "<a href="#roleassignments" title="RoleAssignments">RoleAssignments</a>" : <i>[ <a href="roleassignment.md">RoleAssignment</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FederatedSettingsOrgRoleMapping
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#externalgroupname" title="ExternalGroupName">ExternalGroupName</a>: <i>String</i>
    <a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>: <i>String</i>
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
    <a href="#roleassignments" title="RoleAssignments">RoleAssignments</a>: <i>
      - <a href="roleassignment.md">RoleAssignment</a></i>
</pre>

## Properties

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ExternalGroupName

Unique human-readable label that identifies the identity provider group to whichthis role mapping applies.

_Required_: No

_Type_: String

_Minimum_: <code>1</code>

_Maximum_: <code>200</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies your federation.

_Required_: Yes

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization that contains your projects.

_Required_: Yes

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleAssignments

Atlas roles and the unique identifiers of the groups and organizations associated with each role.

_Required_: No

_Type_: List of <a href="roleassignment.md">RoleAssignment</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### GroupId

Returns the <code>GroupId</code> value.

#### OrgId

Returns the <code>OrgId</code> value.

#### Id

Unique 24-hexadecimal digit string that identifies the role mapping.

