# MongoDB::Atlas::APIKey ProjectAssignment

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#roles" title="Roles">Roles</a>" : <i>[ String, ... ]</i>,
    "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#roles" title="Roles">Roles</a>: <i>
      - String</i>
<a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### Roles

List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization.

_Required_: No

_Type_: List of String

_Allowed Values_: <code>GROUP_ATLAS_ADMIN</code> | <code>GROUP_AUTOMATION_ADMIN</code> | <code>GROUP_BACKUP_ADMIN</code> | <code>GROUP_MONITORING_ADMIN</code> | <code>GROUP_OWNER</code> | <code>GROUP_READ_ONLY</code> | <code>GROUP_USER_ADMIN</code> | <code>GROUP_BILLING_ADMIN</code> | <code>GROUP_DATA_ACCESS_ADMIN</code> | <code>GROUP_DATA_ACCESS_READ_ONLY</code> | <code>GROUP_DATA_ACCESS_READ_WRITE</code> | <code>GROUP_CHARTS_ADMIN</code> | <code>GROUP_CLUSTER_MANAGER</code> | <code>GROUP_SEARCH_INDEX_EDITOR</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies the project in an organization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

