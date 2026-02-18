# MongoDB::Atlas::CloudBackUpRestoreJobs Component

Information on the restore job for each replica set in the sharded cluster.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#downloadurl" title="DownloadUrl">DownloadUrl</a>" : <i>String</i>,
    "<a href="#replicasetname" title="ReplicaSetName">ReplicaSetName</a>" : <i>String</i>,
    "<a href="#privatedownloaddeliveryurls" title="PrivateDownloadDeliveryUrls">PrivateDownloadDeliveryUrls</a>" : <i>[ <a href="privatedownloaddeliveryurl.md">PrivateDownloadDeliveryUrl</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#downloadurl" title="DownloadUrl">DownloadUrl</a>: <i>String</i>
<a href="#replicasetname" title="ReplicaSetName">ReplicaSetName</a>: <i>String</i>
<a href="#privatedownloaddeliveryurls" title="PrivateDownloadDeliveryUrls">PrivateDownloadDeliveryUrls</a>: <i>
      - <a href="privatedownloaddeliveryurl.md">PrivateDownloadDeliveryUrl</a></i>
</pre>

## Properties

#### DownloadUrl

One URL that points to the compressed snapshot files for this replica set.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicaSetName

Human-readable label that identifies the replica set.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateDownloadDeliveryUrls

_Required_: No

_Type_: List of <a href="privatedownloaddeliveryurl.md">PrivateDownloadDeliveryUrl</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

