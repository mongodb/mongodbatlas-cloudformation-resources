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

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies the project in an organization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

