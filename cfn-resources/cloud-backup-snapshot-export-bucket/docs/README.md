# MongoDB::Atlas::CloudBackupSnapshotExportBucket

The exportBuckets resource allows you to grant Atlas access to the specified bucket for exporting backup snapshots.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudBackupSnapshotExportBucket",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#bucketname" title="BucketName">BucketName</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#iamroleid" title="IamRoleID">IamRoleID</a>" : <i>String</i>,
        "<a href="#requireprivatenetworking" title="RequirePrivateNetworking">RequirePrivateNetworking</a>" : <i>Boolean</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudBackupSnapshotExportBucket
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#bucketname" title="BucketName">BucketName</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#iamroleid" title="IamRoleID">IamRoleID</a>: <i>String</i>
    <a href="#requireprivatenetworking" title="RequirePrivateNetworking">RequirePrivateNetworking</a>: <i>Boolean</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### BucketName

Human-readable label that identifies the AWS bucket that the role is authorized to access.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### IamRoleID

Unique 24-hexadecimal character string that identifies the AWS IAM role that MongoDB Cloud uses to access the AWS S3 bucket.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RequirePrivateNetworking

Indicates whether to do exports over PrivateLink as opposed to public IPs. Defaults to false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal character string that identifies the Amazon Web Services (AWS) Simple Storage Service (S3) export bucket.

