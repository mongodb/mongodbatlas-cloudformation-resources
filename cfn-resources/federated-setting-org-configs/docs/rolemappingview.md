# MongoDB::Atlas::FederatedSettingOrgConfigs RoleMappingView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#externalgroupname" title="ExternalGroupName">ExternalGroupName</a>" : <i>String</i>,
    "<a href="#roleassignments" title="RoleAssignments">RoleAssignments</a>" : <i>[ <a href="roleassignment.md">RoleAssignment</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#externalgroupname" title="ExternalGroupName">ExternalGroupName</a>: <i>String</i>
<a href="#roleassignments" title="RoleAssignments">RoleAssignments</a>: <i>
      - <a href="roleassignment.md">RoleAssignment</a></i>
</pre>

## Properties

#### ExternalGroupName

Unique human-readable label that identifies the identity provider group to whichthis role mapping applies.

_Required_: No

_Type_: String

_Minimum Length_: <code>1</code>

_Maximum Length_: <code>200</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleAssignments

Atlas roles and the unique identifiers of the groups and organizations associated with each role.

_Required_: No

_Type_: List of <a href="roleassignment.md">RoleAssignment</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

