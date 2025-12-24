# MongoDB::Atlas::StreamWorkspace DBRoleToExecute

The name of a Built in or Custom DB Role to connect to an Atlas Cluster.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#role" title="Role">Role</a>" : <i>String</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#role" title="Role">Role</a>: <i>String</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
</pre>

## Properties

#### Role

The name of the role to use. Can be a built in role or a custom role.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

Type of the DB role. Can be either BuiltIn or Custom.

_Required_: No

_Type_: String

_Allowed Values_: <code>BUILT_IN</code> | <code>CUSTOM</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

