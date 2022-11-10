# MongoDB::Atlas::BackupSnapshot ApiAtlasDiskBackupShardedClusterSnapshotView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
    "<a href="#createdat" title="CreatedAt">CreatedAt</a>" : <i>String</i>,
    "<a href="#description" title="Description">Description</a>" : <i>String</i>,
    "<a href="#expiresat" title="ExpiresAt">ExpiresAt</a>" : <i>String</i>,
    "<a href="#frequencytype" title="FrequencyType">FrequencyType</a>" : <i>String</i>,
    "<a href="#id" title="Id">Id</a>" : <i>String</i>,
    "<a href="#links" title="Links">Links</a>" : <i>[ <a href="link.md">Link</a>, ... ]</i>,
    "<a href="#masterkeyuuid" title="MasterKeyUUID">MasterKeyUUID</a>" : <i>String</i>,
    "<a href="#members" title="Members">Members</a>" : <i>[ <a href="apiatlasdiskbackupshardedclustersnapshotmemberview.md">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>, ... ]</i>,
    "<a href="#mongodversion" title="MongodVersion">MongodVersion</a>" : <i>String</i>,
    "<a href="#policyitems" title="PolicyItems">PolicyItems</a>" : <i>[ String, ... ]</i>,
    "<a href="#snapshotids" title="SnapshotIds">SnapshotIds</a>" : <i>[ String, ... ]</i>,
    "<a href="#snapshottype" title="SnapshotType">SnapshotType</a>" : <i>String</i>,
    "<a href="#status" title="Status">Status</a>" : <i>String</i>,
    "<a href="#storagesizebytes" title="StorageSizeBytes">StorageSizeBytes</a>" : <i>Integer</i>,
    "<a href="#type" title="Type">Type</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
<a href="#createdat" title="CreatedAt">CreatedAt</a>: <i>String</i>
<a href="#description" title="Description">Description</a>: <i>String</i>
<a href="#expiresat" title="ExpiresAt">ExpiresAt</a>: <i>String</i>
<a href="#frequencytype" title="FrequencyType">FrequencyType</a>: <i>String</i>
<a href="#id" title="Id">Id</a>: <i>String</i>
<a href="#links" title="Links">Links</a>: <i>
      - <a href="link.md">Link</a></i>
<a href="#masterkeyuuid" title="MasterKeyUUID">MasterKeyUUID</a>: <i>String</i>
<a href="#members" title="Members">Members</a>: <i>
      - <a href="apiatlasdiskbackupshardedclustersnapshotmemberview.md">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a></i>
<a href="#mongodversion" title="MongodVersion">MongodVersion</a>: <i>String</i>
<a href="#policyitems" title="PolicyItems">PolicyItems</a>: <i>
      - String</i>
<a href="#snapshotids" title="SnapshotIds">SnapshotIds</a>: <i>
      - String</i>
<a href="#snapshottype" title="SnapshotType">SnapshotType</a>: <i>String</i>
<a href="#status" title="Status">Status</a>: <i>String</i>
<a href="#storagesizebytes" title="StorageSizeBytes">StorageSizeBytes</a>: <i>Integer</i>
<a href="#type" title="Type">Type</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CreatedAt

Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when `"status": "onDemand"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ExpiresAt

Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyType

Human-readable label that identifies how often this snapshot triggers.

_Required_: No

_Type_: String

_Allowed Values_: <code>hourly</code> | <code>daily</code> | <code>weekly</code> | <code>monthly</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Id

Unique 24-hexadecimal digit string that identifies the snapshot.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Links

_Required_: No

_Type_: List of <a href="link.md">Link</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MasterKeyUUID

Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when `"encryptionEnabled" : true`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Members

_Required_: No

_Type_: List of <a href="apiatlasdiskbackupshardedclustersnapshotmemberview.md">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MongodVersion

Version of the MongoDB host that this snapshot backs up.

_Required_: No

_Type_: String

_Pattern_: <code>([\d]+\.[\d]+\.[\d]+)</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItems

List that contains unique identifiers for the policy items.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SnapshotIds

List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster. The resource returns this parameter when `"type": "SHARDED_CLUSTER"`. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SnapshotType

Human-readable label that identifies when this snapshot triggers.

_Required_: No

_Type_: String

_Allowed Values_: <code>onDemand</code> | <code>scheduled</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

Human-readable label that indicates the stage of the backup process for this snapshot.

_Required_: No

_Type_: String

_Allowed Values_: <code>queued</code> | <code>inProgress</code> | <code>completed</code> | <code>failed</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StorageSizeBytes

Number of bytes taken to store the backup snapshot.

_Required_: No

_Type_: Integer

_Pattern_: <code>^([0-9]+)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

Human-readable label that categorizes the cluster as a replica set or sharded cluster.

_Required_: No

_Type_: String

_Allowed Values_: <code>REPLICA_SET</code> | <code>SHARDED_CLUSTER</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

