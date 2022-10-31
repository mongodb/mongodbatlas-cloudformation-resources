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

ID of the AWS security group to allow access. Mutually exclusive with CIDRBlock and IPAddress.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CIDRBlock

Accessable entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with ipAddress and AwsSecurityGroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Comment

Comment associated with the ip access list entry.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IPAddress

Accessable IP address. Mutually exclusive with CIDRBlock and AwsSecurityGroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

The unique identifier for the project to which you want to add one or more ip access list entries.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

