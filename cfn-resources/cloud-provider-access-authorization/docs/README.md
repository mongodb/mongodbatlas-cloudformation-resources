# MongoDB::Atlas::CloudProviderAccessAuthorization

Authorizes an AWS IAM role against a Cloud Provider Access role registration so Atlas can assume it. Use MongoDB::Atlas::CloudProviderAccessSetup to register the role first, then use this resource to authorize it. Deleting this resource removes the whole Cloud Provider Access role registration, not only the authorization.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudProviderAccessAuthorization",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#roleid" title="RoleId">RoleId</a>" : <i>String</i>,
        "<a href="#iamassumedrolearn" title="IamAssumedRoleArn">IamAssumedRoleArn</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudProviderAccessAuthorization
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#roleid" title="RoleId">RoleId</a>: <i>String</i>
    <a href="#iamassumedrolearn" title="IamAssumedRoleArn">IamAssumedRoleArn</a>: <i>String</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### RoleId

Unique 24-hexadecimal digit string that identifies the Cloud Provider Access role registration to authorize. This is the RoleId returned by MongoDB::Atlas::CloudProviderAccessSetup.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### IamAssumedRoleArn

Amazon Resource Name (ARN) that identifies the AWS IAM role that MongoDB Cloud assumes when it accesses resources in your AWS account. Changing this value re-authorizes the same Cloud Provider Access role in place.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>20</code>

_Maximum Length_: <code>2048</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### AuthorizedDate

Date and time when the role was authorized for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### FeatureUsages

List of MongoDB Cloud features that use the authorized AWS IAM role.

