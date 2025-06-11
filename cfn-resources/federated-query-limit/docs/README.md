# MongoDB::Atlas::FederatedQueryLimit

Query limit for one federated database instance.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedQueryLimit",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#tenantname" title="TenantName">TenantName</a>" : <i>String</i>,
        "<a href="#limitname" title="LimitName">LimitName</a>" : <i>String</i>,
        "<a href="#overrunpolicy" title="OverrunPolicy">OverrunPolicy</a>" : <i>String</i>,
        "<a href="#value" title="Value">Value</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FederatedQueryLimit
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#tenantname" title="TenantName">TenantName</a>: <i>String</i>
    <a href="#limitname" title="LimitName">LimitName</a>: <i>String</i>
    <a href="#overrunpolicy" title="OverrunPolicy">OverrunPolicy</a>: <i>String</i>
    <a href="#value" title="Value">Value</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### TenantName

Human-readable label that identifies the data federated database instance to which the query limit applies.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### LimitName

Human-readable label that identifies this data federation instance limit.

_Required_: Yes

_Type_: String

_Allowed Values_: <code>bytesProcessed.query</code> | <code>bytesProcessed.daily</code> | <code>bytesProcessed.weekly</code> | <code>bytesProcessed.monthly</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### OverrunPolicy

Only used for Data Federation limits. Action to take when the usage limit is exceeded. If limit span is set to QUERY, this is ignored because MongoDB Cloud stops the query when it exceeds the usage limit. "enum" : [ "BLOCK", "BLOCK_AND_KILL" ]

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Value

Amount to set the limit to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### CurrentUsage

Amount that indicates the current usage of the limit.

#### LastModifiedDate

Only used for Data Federation limits. Timestamp that indicates when this usage limit was last modified. This field uses the ISO 8601 timestamp format in UTC.

#### MaximumLimit

Maximum value of the limit.

#### DefaultLimit

Default value of the limit.

