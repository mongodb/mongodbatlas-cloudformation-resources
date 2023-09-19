# MongoDB::Atlas::Organization APIKey

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#description" title="Description">Description</a>" : <i>String</i>,
    "<a href="#roles" title="Roles">Roles</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#description" title="Description">Description</a>: <i>String</i>
<a href="#roles" title="Roles">Roles</a>: <i>
      - String</i>
</pre>

## Properties

#### Description

Purpose or explanation provided when someone created this organization API key. 1 to 250 characters

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Roles

List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization.

_Required_: No

_Type_: List of String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

