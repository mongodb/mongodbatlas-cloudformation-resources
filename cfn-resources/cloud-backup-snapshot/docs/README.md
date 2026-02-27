# MongoDB::Atlas::CloudBackupSnapshot

Returns, takes, and removes Cloud Backup snapshots.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudBackupSnapshot",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#instancetype" title="InstanceType">InstanceType</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#frequencytype" title="FrequencyType">FrequencyType</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>,
        "<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>" : <i>Integer</i>,
        "<a href="#members" title="Members">Members</a>" : <i>[ <a href="apiatlasdiskbackupshardedclustersnapshotmemberview.md">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>, ... ]</i>,
        "<a href="#pagenum" title="PageNum">PageNum</a>" : <i>Integer</i>,
        "<a href="#policyitems" title="PolicyItems">PolicyItems</a>" : <i>[ String, ... ]</i>,
        "<a href="#results" title="Results">Results</a>" : <i>[ <a href="apiatlasdiskbackupshardedclustersnapshotview.md">ApiAtlasDiskBackupShardedClusterSnapshotView</a>, ... ]</i>,
        "<a href="#retentionindays" title="RetentionInDays">RetentionInDays</a>" : <i>Integer</i>,
        "<a href="#snapshottype" title="SnapshotType">SnapshotType</a>" : <i>String</i>,
        "<a href="#totalcount" title="TotalCount">TotalCount</a>" : <i>Double</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudBackupSnapshot
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#instancetype" title="InstanceType">InstanceType</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#frequencytype" title="FrequencyType">FrequencyType</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
    <a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>: <i>Integer</i>
    <a href="#members" title="Members">Members</a>: <i>
      - <a href="apiatlasdiskbackupshardedclustersnapshotmemberview.md">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a></i>
    <a href="#pagenum" title="PageNum">PageNum</a>: <i>Integer</i>
    <a href="#policyitems" title="PolicyItems">PolicyItems</a>: <i>
      - String</i>
    <a href="#results" title="Results">Results</a>: <i>
      - <a href="apiatlasdiskbackupshardedclustersnapshotview.md">ApiAtlasDiskBackupShardedClusterSnapshotView</a></i>
    <a href="#retentionindays" title="RetentionInDays">RetentionInDays</a>: <i>Integer</i>
    <a href="#snapshottype" title="SnapshotType">SnapshotType</a>: <i>String</i>
    <a href="#totalcount" title="TotalCount">TotalCount</a>: <i>Double</i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### InstanceType

Type of instance specified on the Instance Name.

_Required_: Yes

_Type_: String

_Allowed Values_: <code>cluster</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### InstanceName

Human-readable label that identifies the cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Description

Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when `"status": "onDemand"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FrequencyType

Human-readable label that identifies how often this snapshot triggers.

_Required_: No

_Type_: String

_Allowed Values_: <code>hourly</code> | <code>daily</code> | <code>weekly</code> | <code>monthly</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### IncludeCount

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ItemsPerPage

Number of items that the response returns per page.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Members

List that includes the snapshots and the cloud provider that stores the snapshots. The resource returns this parameter when `"type" : "SHARDED_CLUSTER"`.

_Required_: No

_Type_: List of <a href="apiatlasdiskbackupshardedclustersnapshotmemberview.md">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PageNum

Number of the page that displays the current set of the total objects that the response returns.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItems

List that contains unique identifiers for the policy items.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Results

List of returned documents that MongoDB Cloud provides when completing this request.

_Required_: No

_Type_: List of <a href="apiatlasdiskbackupshardedclustersnapshotview.md">ApiAtlasDiskBackupShardedClusterSnapshotView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionInDays

Number of days that MongoDB Cloud should retain the on-demand snapshot. Must be at least **1**

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SnapshotType

Human-readable label that identifies when this snapshot triggers.

_Required_: No

_Type_: String

_Allowed Values_: <code>onDemand</code> | <code>scheduled</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TotalCount

Number of documents returned in this response.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### SnapshotId

Unique 24-hexadecimal digit string that identifies the desired snapshot.

#### SnapshotIds

List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster. The resource returns this parameter when `"type": "SHARDED_CLUSTER"`. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name.

#### MasterKeyUUID

Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when `"encryptionEnabled" : true`.

#### Type

Human-readable label that categorizes the cluster as a replica set or sharded cluster.

#### ExpiresAt

Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### StorageSizeBytes

Number of bytes taken to store the backup snapshot.

#### Id

Unique 24-hexadecimal digit string that identifies the snapshot.

#### CreatedAt

Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### CloudProvider

Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when `"type": "replicaSet".`

#### MongodVersion

Version of the MongoDB host that this snapshot backs up.

#### ReplicaSetName

Human-readable label that identifies the replica set from which MongoDB Cloud took this snapshot. The resource returns this parameter when `"type": "replicaSet"`

#### Status

Human-readable label that indicates the stage of the backup process for this snapshot.

