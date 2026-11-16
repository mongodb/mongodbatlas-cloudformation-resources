# MongoDB::Atlas::CloudProviderAccessAuthorization FeatureId

Identifying characteristics of the feature that uses the role, when Atlas returns them, for example the S3 bucket name for a push-based log export. Not set for features that Atlas does not identify, for example ENCRYPTION_AT_REST.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    "<a href="#bucketname" title="BucketName">BucketName</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
<a href="#bucketname" title="BucketName">BucketName</a>: <i>String</i>
</pre>

## Properties

#### ProjectId

Unique 24-hexadecimal digit string that identifies the project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BucketName

Name of the AWS S3 bucket that the push-based log export writes to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

