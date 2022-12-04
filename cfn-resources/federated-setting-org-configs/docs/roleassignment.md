# MongoDB::Atlas::FederatedSettingOrgConfigs RoleAssignment

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#role" title="Role">Role</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#role" title="Role">Role</a>: <i>String</i>
</pre>

## Properties

#### Role

_Required_: No

_Type_: String

_Allowed Values_: <code>GLOBAL_AUTOMATION_ADMIN</code> | <code>GLOBAL_BACKUP_ADMIN</code> | <code>GLOBAL_METERING_USER</code> | <code>GLOBAL_METRICS_INTERNAL_USER</code> | <code>GLOBAL_MONITORING_ADMIN</code> | <code>GLOBAL_OWNER</code> | <code>GLOBAL_READ_ONLY</code> | <code>ORG_OWNER</code> | <code>ORG_MEMBER</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

