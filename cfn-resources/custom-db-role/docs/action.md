# MongoDB::Atlas::CustomDBRole Action

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#action" title="Action">Action</a>" : <i><a href="action.md">Action</a></i>,
    "<a href="#resources" title="Resources">Resources</a>" : <i>[ <a href="resource.md">Resource</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#action" title="Action">Action</a>: <i><a href="action.md">Action</a></i>
<a href="#resources" title="Resources">Resources</a>: <i>
      - <a href="resource.md">Resource</a></i>
</pre>

## Properties

#### Action

_Required_: No

_Type_: <a href="action.md">Action</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Resources

Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.

_Required_: No

_Type_: List of <a href="resource.md">Resource</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

