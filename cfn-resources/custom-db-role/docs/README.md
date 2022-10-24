# MongoDB::Atlas::CustomDBRole

Returns, adds, edits, and removes custom database user privilege roles.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CustomDBRole",
    "Properties" : {
        "<a href="#actions" title="Actions">Actions</a>" : <i>[ <a href="action.md">Action</a>, ... ]</i>,
        "<a href="#inheritedroles" title="InheritedRoles">InheritedRoles</a>" : <i>[ <a href="inheritedrole.md">InheritedRole</a>, ... ]</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikey.md">ApiKey</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CustomDBRole
Properties:
    <a href="#actions" title="Actions">Actions</a>: <i>
      - <a href="action.md">Action</a></i>
    <a href="#inheritedroles" title="InheritedRoles">InheritedRoles</a>: <i>
      - <a href="inheritedrole.md">InheritedRole</a></i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikey.md">ApiKey</a></i>
</pre>

## Properties

#### Actions

List of the individual privilege actions that the role grants.

_Required_: No

_Type_: List of <a href="action.md">Action</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InheritedRoles

List of the built-in roles that this custom role inherits.

_Required_: No

_Type_: List of <a href="inheritedrole.md">InheritedRole</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: No

_Type_: <a href="apikey.md">ApiKey</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the RoleName.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### RoleName

Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.

