# MongoDB::Atlas::Project projectTeam

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#teamid" title="TeamId">TeamId</a>" : <i>String</i>,
    "<a href="#rolenames" title="RoleNames">RoleNames</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#teamid" title="TeamId">TeamId</a>: <i>String</i>
<a href="#rolenames" title="RoleNames">RoleNames</a>: <i>
      - String</i>
</pre>

## Properties

#### TeamId

Unique 24-hexadecimal character string that identifies the team. string = 24 characters ^([a-f0-9]{24})$

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleNames

One or more organization- or project-level roles to assign to the MongoDB Cloud user. tems Enum: "GROUP_CLUSTER_MANAGER" "GROUP_DATA_ACCESS_ADMIN" "GROUP_DATA_ACCESS_READ_ONLY" "GROUP_DATA_ACCESS_READ_WRITE" "GROUP_OWNER" "GROUP_READ_ONLY"

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

