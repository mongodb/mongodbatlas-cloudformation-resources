# MongoDB::Atlas::CloudBackupSchedule ApiPolicyView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="ID">ID</a>" : <i>String</i>,
    "<a href="#policyitems" title="PolicyItems">PolicyItems</a>" : <i>[ <a href="apipolicyitemview.md">ApiPolicyItemView</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#id" title="ID">ID</a>: <i>String</i>
<a href="#policyitems" title="PolicyItems">PolicyItems</a>: <i>
      - <a href="apipolicyitemview.md">ApiPolicyItemView</a></i>
</pre>

## Properties

#### ID

Unique 24-hexadecimal digit string that identifies this backup policy. The policy id can be retrieved by running: atlas backups schedule describe "${clusterName}" --projectId "${projectId}" | jq -r '.policies[0].id'

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItems

_Required_: No

_Type_: List of <a href="apipolicyitemview.md">ApiPolicyItemView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

