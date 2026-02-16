# MongoDB::Atlas::ProjectServiceAccountAccessListEntry

Manages IP access list entries for MongoDB Atlas Project Service Accounts.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ProjectServiceAccountAccessListEntry",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#clientid" title="ClientId">ClientId</a>" : <i>String</i>,
        "<a href="#cidrblock" title="CIDRBlock">CIDRBlock</a>" : <i>String</i>,
        "<a href="#ipaddress" title="IPAddress">IPAddress</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ProjectServiceAccountAccessListEntry
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#clientid" title="ClientId">ClientId</a>: <i>String</i>
    <a href="#cidrblock" title="CIDRBlock">CIDRBlock</a>: <i>String</i>
    <a href="#ipaddress" title="IPAddress">IPAddress</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies the project.

_Required_: Yes

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ClientId

The Client ID of the Service Account.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### CIDRBlock

Range of IP addresses in CIDR notation to be added to the access list. You can set a value for this parameter or IPAddress, but not both.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### IPAddress

Single IP address to be added to the access list. You can set a value for this parameter or CIDRBlock, but not both.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### CreatedAt

Date and time when the access list entry was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### LastUsedAddress

Network address that issued the most recent request to the API. This parameter may not be present if no requests have originated from this IP address.

#### LastUsedAt

Date and time when the API received the most recent request originating from this IP address. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### RequestCount

Total number of requests that have originated from this IP address.

