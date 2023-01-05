# MongoDB::Atlas::FederatedSettingsOrgRoleMapping RoleAssignment

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
    "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
    "<a href="#role" title="Role">Role</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
<a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
<a href="#role" title="Role">Role</a>: <i>String</i>
</pre>

## Properties

#### GroupId

List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.

This parameter returns an empty object if no custom zones exist.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.

This parameter returns an empty object if no custom zones exist.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Role

_Required_: No

_Type_: String

_Allowed Values_: <code>GLOBAL_AUTOMATION_ADMIN</code> | <code>GLOBAL_BACKUP_ADMIN</code> | <code>GLOBAL_METERING_USER</code> | <code>GLOBAL_METRICS_INTERNAL_USER</code> | <code>GLOBAL_MONITORING_ADMIN</code> | <code>GLOBAL_OWNER</code> | <code>GLOBAL_READ_ONLY</code> | <code>GLOBAL_USER_ADMIN</code> | <code>GLOBAL_USER_READ_ONLY</code> | <code>GLOBAL_ACCOUNT_SUSPENSION_ADMIN</code> | <code>GLOBAL_BILLING_ADMIN</code> | <code>GLOBAL_LEGAL_ADMIN</code> | <code>GLOBAL_FEATURE_FLAG_ADMIN</code> | <code>GLOBAL_ATLAS_TSE</code> | <code>GLOBAL_ATLAS_OPERATOR</code> | <code>GLOBAL_ATLAS_ADMIN</code> | <code>GLOBAL_STITCH_ADMIN</code> | <code>GLOBAL_CHARTS_ADMIN</code> | <code>GLOBAL_EXPERIMENT_ASSIGNMENT_USER</code> | <code>GLOBAL_STITCH_INTERNAL_ADMIN</code> | <code>GLOBAL_SECURITY_ADMIN</code> | <code>GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN</code> | <code>GLOBAL_PROACTIVE_SUPPORT_ADMIN</code> | <code>GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN</code> | <code>GLOBAL_SALESFORCE_ADMIN</code> | <code>GLOBAL_SALESFORCE_READ_ONLY</code> | <code>GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS</code> | <code>ORG_MEMBER</code> | <code>ORG_READ_ONLY</code> | <code>ORG_BILLING_ADMIN</code> | <code>ORG_GROUP_CREATOR</code> | <code>ORG_OWNER</code> | <code>GROUP_AUTOMATION_ADMIN</code> | <code>GROUP_BACKUP_ADMIN</code> | <code>GROUP_MONITORING_ADMIN</code> | <code>GROUP_OWNER</code> | <code>GROUP_READ_ONLY</code> | <code>GROUP_USER_ADMIN</code> | <code>GROUP_BILLING_ADMIN</code> | <code>GROUP_DATA_ACCESS_ADMIN</code> | <code>GROUP_DATA_ACCESS_READ_ONLY</code> | <code>GROUP_DATA_ACCESS_READ_WRITE</code> | <code>GROUP_CHARTS_ADMIN</code> | <code>GROUP_CLUSTER_MANAGER</code> | <code>GROUP_SEARCH_INDEX_EDITOR</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

