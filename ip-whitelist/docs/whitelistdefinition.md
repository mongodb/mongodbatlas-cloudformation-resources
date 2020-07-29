# MongoDB::Atlas::ProjectIPWhitelist whitelistDefinition

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#comment" title="comment">comment</a>" : <i>String</i>,
    "<a href="#ipaddress" title="ipAddress">ipAddress</a>" : <i>String</i>,
    "<a href="#cidrblock" title="cidrBlock">cidrBlock</a>" : <i>String</i>,
    "<a href="#awssecuritygroup" title="awsSecurityGroup">awsSecurityGroup</a>" : <i>String</i>,
    "<a href="#projectid" title="projectId">projectId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#comment" title="comment">comment</a>: <i>String</i>
<a href="#ipaddress" title="ipAddress">ipAddress</a>: <i>String</i>
<a href="#cidrblock" title="cidrBlock">cidrBlock</a>: <i>String</i>
<a href="#awssecuritygroup" title="awsSecurityGroup">awsSecurityGroup</a>: <i>String</i>
<a href="#projectid" title="projectId">projectId</a>: <i>String</i>
</pre>

## Properties

#### comment

Comment associated with the whitelist entry.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ipAddress

Whitelisted IP address. Mutually exclusive with cidrBlock and awsSecurityGroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### cidrBlock

Whitelist entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with ipAddress and awsSecurityGroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### awsSecurityGroup

ID of the AWS security group to whitelist. Mutually exclusive with cidrBlock and ipAddress and cidrBlock.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### projectId

The unique identifier for the project to which you want to add one or more whitelist entries.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

