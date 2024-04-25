# MongoDB::Atlas::CloudBackupSchedule Export

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#exportbucketid" title="ExportBucketId">ExportBucketId</a>" : <i>String</i>,
    "<a href="#frequencytype" title="FrequencyType">FrequencyType</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#exportbucketid" title="ExportBucketId">ExportBucketId</a>: <i>String</i>
<a href="#frequencytype" title="FrequencyType">FrequencyType</a>: <i>String</i>
</pre>

## Properties

#### ExportBucketId

Unique identifier of the AWS bucket to export the cloud backup snapshot to

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyType

Frequency associated with the export policy. Value can be daily, weekly, monthly or yearly.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

