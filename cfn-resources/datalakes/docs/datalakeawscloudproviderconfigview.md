# MongoDB::Atlas::DataLakes DataLakeAWSCloudProviderConfigView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#externalid" title="ExternalId">ExternalId</a>" : <i>String</i>,
    "<a href="#iamassumedrolearn" title="IamAssumedRoleARN">IamAssumedRoleARN</a>" : <i>String</i>,
    "<a href="#iamuserarn" title="IamUserARN">IamUserARN</a>" : <i>String</i>,
    "<a href="#roleid" title="RoleId">RoleId</a>" : <i>String</i>,
    "<a href="#tests3bucket" title="TestS3Bucket">TestS3Bucket</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#externalid" title="ExternalId">ExternalId</a>: <i>String</i>
<a href="#iamassumedrolearn" title="IamAssumedRoleARN">IamAssumedRoleARN</a>: <i>String</i>
<a href="#iamuserarn" title="IamUserARN">IamUserARN</a>: <i>String</i>
<a href="#roleid" title="RoleId">RoleId</a>: <i>String</i>
<a href="#tests3bucket" title="TestS3Bucket">TestS3Bucket</a>: <i>String</i>
</pre>

## Properties

#### ExternalId

Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IamAssumedRoleARN

Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores.

_Required_: No

_Type_: String

_Minimum Length_: <code>20</code>

_Maximum Length_: <code>2048</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IamUserARN

Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RoleId

Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TestS3Bucket

Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

