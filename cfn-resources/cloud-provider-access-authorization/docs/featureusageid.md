# MongoDB::Atlas::CloudProviderAccessAuthorization FeatureUsageId

Identifying characteristics about the AWS Simple Storage Service (S3) export bucket linked to this AWS IAM role.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
    "<a href="#bucketname" title="BucketName">BucketName</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
<a href="#bucketname" title="BucketName">BucketName</a>: <i>String</i>
</pre>

## Properties

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BucketName

Name of the AWS S3 bucket to which your logs are exported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

