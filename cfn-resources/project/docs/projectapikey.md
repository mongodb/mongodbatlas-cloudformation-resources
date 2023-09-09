# MongoDB::Atlas::Project projectApiKey

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#key" title="Key">Key</a>" : <i>String</i>,
    "<a href="#rolenames" title="RoleNames">RoleNames</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#key" title="Key">Key</a>: <i>String</i>
<a href="#rolenames" title="RoleNames">RoleNames</a>: <i>
      - String</i>
</pre>

## Properties

#### Key

Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleNames

List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.Items Enum: "ORG_OWNER" "ORG_MEMBER" "ORG_GROUP_CREATOR" "ORG_BILLING_ADMIN" "ORG_READ_ONLY" "ORG_TEAM_MEMBERS_ADMIN" "GROUP_ATLAS_ADMIN" "GROUP_AUTOMATION_ADMIN" "GROUP_BACKUP_ADMIN" "GROUP_MONITORING_ADMIN" "GROUP_OWNER" "GROUP_READ_ONLY" "GROUP_USER_ADMIN" "GROUP_BILLING_ADMIN" "GROUP_DATA_ACCESS_ADMIN" "GROUP_DATA_ACCESS_READ_ONLY" "GROUP_DATA_ACCESS_READ_WRITE" "GROUP_CHARTS_ADMIN" "GROUP_CLUSTER_MANAGER" "GROUP_SEARCH_INDEX_EDITOR"

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

