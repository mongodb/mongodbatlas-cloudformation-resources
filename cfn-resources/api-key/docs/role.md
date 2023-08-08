# MongoDB::Atlas::APIKey Role

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
    "<a href="#rolename" title="RoleName">RoleName</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
<a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
<a href="#rolename" title="RoleName">RoleName</a>: <i>String</i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies the project to which this role belongs. You can set a value for this parameter or orgId but not both in the same request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization to which this role belongs. You can set a value for this parameter or groupId but not both in the same request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleName

Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include organization- and project-level roles.

_Required_: No

_Type_: String

_Allowed Values_: <code>ORG_OWNER</code> | <code>ORG_MEMBER</code> | <code>ORG_GROUP_CREATOR</code> | <code>ORG_BILLING_ADMIN</code> | <code>ORG_READ_ONLY</code> | <code>ORG_TEAM_MEMBERS_ADMIN</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

