# MongoDB::Atlas::ProjectIpAccessList accessListDefinition

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#deleteafterdate" title="DeleteAfterDate">DeleteAfterDate</a>" : <i>String</i>,
    "<a href="#awssecuritygroup" title="AwsSecurityGroup">AwsSecurityGroup</a>" : <i>String</i>,
    "<a href="#cidrblock" title="CIDRBlock">CIDRBlock</a>" : <i>String</i>,
    "<a href="#comment" title="Comment">Comment</a>" : <i>String</i>,
    "<a href="#ipaddress" title="IPAddress">IPAddress</a>" : <i>String</i>,
    "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#deleteafterdate" title="DeleteAfterDate">DeleteAfterDate</a>: <i>String</i>
<a href="#awssecuritygroup" title="AwsSecurityGroup">AwsSecurityGroup</a>: <i>String</i>
<a href="#cidrblock" title="CIDRBlock">CIDRBlock</a>: <i>String</i>
<a href="#comment" title="Comment">Comment</a>: <i>String</i>
<a href="#ipaddress" title="IPAddress">IPAddress</a>: <i>String</i>
<a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### DeleteAfterDate

Date and time after which MongoDB Cloud deletes the temporary access list entry. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. The date must be later than the current date but no later than one week after you submit this request. The resource returns this parameter if you specified an expiration date when creating this IP access list entry.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsSecurityGroup

Unique string of the Amazon Web Services (AWS) security group that you want to add to the project's IP access list. Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. You must configure Virtual Private Connection (VPC) peering for your project before you can add an AWS security group to an IP access list. You cannot set AWS security groups as temporary access list entries. Don't set this parameter if you set cidrBlock or ipAddress.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CIDRBlock

Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project's IP access list. Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. Don't set this parameter if you set awsSecurityGroup or ipAddress

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Comment

Remark that explains the purpose or scope of this IP access list entry.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IPAddress

IP address that you want to add to the project's IP access list. Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. Don't set this parameter if you set awsSecurityGroup or cidrBlock.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

