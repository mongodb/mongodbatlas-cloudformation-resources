# MongoDB::Atlas::Auditing

Returns and edits database auditing settings for MongoDB Cloud projects.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Auditing",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Auditing
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### AuditFilter

JSON document that specifies which events to record. Escape any characters that may prevent parsing, such as single or double quotes, using a backslash (`\`), for more information about audit filters refer to https://www.mongodb.com/docs/manual/tutorial/configure-audit-filters/.

#### ConfigurationType

Human-readable label that displays how to configure the audit filter.

#### AuditAuthorizationSuccess

Flag that indicates whether someone set auditing to track successful authentications. This only applies to the `"atype" : "authCheck"` audit filter. Setting this parameter to `true` degrades cluster performance.

