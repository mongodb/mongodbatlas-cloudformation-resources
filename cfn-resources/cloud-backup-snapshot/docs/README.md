# MongoDB::Atlas::CloudBackupSnapshot

Returns, takes, and removes Cloud Backup snapshots.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudBackupSnapshot",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
        "<a href="#retentionindays" title="RetentionInDays">RetentionInDays</a>" : <i>Integer</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudBackupSnapshot
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
    <a href="#retentionindays" title="RetentionInDays">RetentionInDays</a>: <i>Integer</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

Human-readable label that identifies the cluster.

_Required_: No

_Type_: String

_Minimum_: <code>1</code>

_Maximum_: <code>64</code>

_Pattern_: <code>^([a-zA-Z0-9]([a-zA-Z0-9-]){0,21}(?<!-)([\w]{0,42}))$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InstanceName

Human-readable label that identifies the serverless instance.

_Required_: No

_Type_: String

_Minimum_: <code>1</code>

_Maximum_: <code>64</code>

_Pattern_: <code>^([a-zA-Z0-9]([a-zA-Z0-9-]){0,21}(?<!-)([\w]{0,42}))$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when `"status": "onDemand"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionInDays

Number of days that MongoDB Cloud should retain the on-demand snapshot. Must be at least **1**

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the SnapshotId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### SnapshotId

Unique 24-hexadecimal digit string that identifies the desired snapshot.

#### SnapshotIds

List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster. The resource returns this parameter when `"type": "SHARDED_CLUSTER"`. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name.

#### MasterKeyUUID

Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when `"encryptionEnabled" : true`.

#### Results

List of returned documents that MongoDB Cloud provides when completing this request.

#### ItemsPerPage

Number of items that the response returns per page.

#### IncludeCount

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

#### PageNum

Number of the page that displays the current set of the total objects that the response returns.

#### Type

Human-readable label that categorizes the cluster as a replica set or sharded cluster.

#### SnapshotType

Human-readable label that identifies when this snapshot triggers.

#### TotalCount

Number of documents returned in this response.

#### Members

List that includes the snapshots and the cloud provider that stores the snapshots. The resource returns this parameter when `"type" : "SHARDED_CLUSTER"`.

#### ExpiresAt

Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### StorageSizeBytes

Number of bytes taken to store the backup snapshot.

#### PolicyItems

List that contains unique identifiers for the policy items.

#### Id

Unique 24-hexadecimal digit string that identifies the snapshot.

#### CreatedAt

Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### Links

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

#### CloudProvider

Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when `"type": "replicaSet".`

#### MongodVersion

Version of the MongoDB host that this snapshot backs up.

#### FrequencyType

Human-readable label that identifies how often this snapshot triggers.

#### ReplicaSetName

Human-readable label that identifies the replica set from which MongoDB Cloud took this snapshot. The resource returns this parameter when `"type": "replicaSet"`

#### Status

Human-readable label that indicates the stage of the backup process for this snapshot.

#### Type

Returns the <code>Type</code> value.

#### Id

Returns the <code>Id</code> value.

#### CloudProvider

Returns the <code>CloudProvider</code> value.

#### Description

Returns the <code>Description</code> value.

#### CreatedAt

Returns the <code>CreatedAt</code> value.

#### ExpiresAt

Returns the <code>ExpiresAt</code> value.

#### MongodVersion

Returns the <code>MongodVersion</code> value.

#### StorageSizeBytes

Returns the <code>StorageSizeBytes</code> value.

#### ReplicaSetName

Returns the <code>ReplicaSetName</code> value.

#### Status

Returns the <code>Status</code> value.

#### FrequencyType

Returns the <code>FrequencyType</code> value.

#### Links

Returns the <code>Links</code> value.

#### PolicyItems

Returns the <code>PolicyItems</code> value.

#### SnapshotType

Returns the <code>SnapshotType</code> value.

#### MasterKeyUUID

Returns the <code>MasterKeyUUID</code> value.

#### CloudProvider

Returns the <code>CloudProvider</code> value.

#### Id

Returns the <code>Id</code> value.

#### ReplicaSetName

Returns the <code>ReplicaSetName</code> value.

#### Status

Returns the <code>Status</code> value.

#### SnapshotType

Returns the <code>SnapshotType</code> value.

#### Type

Returns the <code>Type</code> value.

#### PolicyItems

Returns the <code>PolicyItems</code> value.

#### CreatedAt

Returns the <code>CreatedAt</code> value.

#### ExpiresAt

Returns the <code>ExpiresAt</code> value.

#### Id

Returns the <code>Id</code> value.

#### Members

Returns the <code>Members</code> value.

#### MongodVersion

Returns the <code>MongodVersion</code> value.

#### SnapshotIds

Returns the <code>SnapshotIds</code> value.

#### FrequencyType

Returns the <code>FrequencyType</code> value.

#### StorageSizeBytes

Returns the <code>StorageSizeBytes</code> value.

#### MasterKeyUUID

Returns the <code>MasterKeyUUID</code> value.

#### Description

Returns the <code>Description</code> value.

#### Links

Returns the <code>Links</code> value.

