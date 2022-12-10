# MongoDB::Atlas::CloudBackUpRestoreJobs

Returns, starts, and cancels Cloud Backup restore jobs.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudBackUpRestoreJobs",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#deliverytype" title="DeliveryType">DeliveryType</a>" : <i>String</i>,
        "<a href="#snapshotid" title="SnapshotId">SnapshotId</a>" : <i>String</i>,
        "<a href="#oplogts" title="OpLogTs">OpLogTs</a>" : <i>String</i>,
        "<a href="#oploginc" title="OpLogInc">OpLogInc</a>" : <i>String</i>,
        "<a href="#pointintimeutcseconds" title="PointInTimeUtcSeconds">PointInTimeUtcSeconds</a>" : <i>Integer</i>,
        "<a href="#targetprojectid" title="TargetProjectId">TargetProjectId</a>" : <i>String</i>,
        "<a href="#targetclustername" title="TargetClusterName">TargetClusterName</a>" : <i>String</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudBackUpRestoreJobs
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#deliverytype" title="DeliveryType">DeliveryType</a>: <i>String</i>
    <a href="#snapshotid" title="SnapshotId">SnapshotId</a>: <i>String</i>
    <a href="#oplogts" title="OpLogTs">OpLogTs</a>: <i>String</i>
    <a href="#oploginc" title="OpLogInc">OpLogInc</a>: <i>String</i>
    <a href="#pointintimeutcseconds" title="PointInTimeUtcSeconds">PointInTimeUtcSeconds</a>: <i>Integer</i>
    <a href="#targetprojectid" title="TargetProjectId">TargetProjectId</a>: <i>String</i>
    <a href="#targetclustername" title="TargetClusterName">TargetClusterName</a>: <i>String</i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
</pre>

## Properties

#### ProjectId

The unique identifier of the project for the Atlas cluster.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InstanceName

The instance name of the Serverless cluster whose snapshot you want to restore or you want to retrieve restore jobs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DeliveryType

Type of restore job to create.The value can be any one of download,automated or point_in_time 

_Required_: No

_Type_: String

_Allowed Values_: <code>download</code> | <code>automated</code> | <code>point_in_time</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SnapshotId

Unique identifier of the source snapshot ID of the restore job.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OpLogTs

Timestamp in the number of seconds that have elapsed since the UNIX epoch from which to you want to restore this snapshot. This is the first part of an Oplog timestamp.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OpLogInc

Oplog operation number from which to you want to restore this snapshot. This is the second part of an Oplog timestamp.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PointInTimeUtcSeconds

If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TargetProjectId

Name of the target Atlas project of the restore job. Only visible if deliveryType is automated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TargetClusterName

Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

 The unique identifier of the restore job.

#### DeliveryUrl

One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.

#### Cancelled

Indicates whether the restore job was canceled.

#### CreatedAt

UTC ISO 8601 formatted point in time when Atlas created the restore job.

#### Expired

Indicates whether the restore job expired.

#### ExpiresAt

UTC ISO 8601 formatted point in time when the restore job expires.

#### FinishedAt

UTC ISO 8601 formatted point in time when the restore job completed.

#### Timestamp

Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.

#### Links

One or more links to sub-resources and/or related resources.

