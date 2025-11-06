# MongoDB::Atlas::FederatedDatabaseInstance CloudProviderConfig

Cloud provider linked to this Atlas Data Federation.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#roleid" title="RoleId">RoleId</a>" : <i>String</i>,
    "<a href="#tests3bucket" title="TestS3Bucket">TestS3Bucket</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#roleid" title="RoleId">RoleId</a>: <i>String</i>
<a href="#tests3bucket" title="TestS3Bucket">TestS3Bucket</a>: <i>String</i>
</pre>

## Properties

#### RoleId

Unique identifier of the role that the Atlas Data Federation can use to access the data stores.Required if specifying cloudProviderConfig.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TestS3Bucket

Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

