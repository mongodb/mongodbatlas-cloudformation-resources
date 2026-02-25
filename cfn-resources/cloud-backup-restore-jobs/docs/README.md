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
        "<a href="#instancetype" title="InstanceType">InstanceType</a>" : <i>String</i>,
        "<a href="#instancename" title="InstanceName">InstanceName</a>" : <i>String</i>,
        "<a href="#deliverytype" title="DeliveryType">DeliveryType</a>" : <i>String</i>,
        "<a href="#snapshotid" title="SnapshotId">SnapshotId</a>" : <i>String</i>,
        "<a href="#oplogts" title="OpLogTs">OpLogTs</a>" : <i>String</i>,
        "<a href="#oploginc" title="OpLogInc">OpLogInc</a>" : <i>String</i>,
        "<a href="#pointintimeutcseconds" title="PointInTimeUtcSeconds">PointInTimeUtcSeconds</a>" : <i>Integer</i>,
        "<a href="#targetprojectid" title="TargetProjectId">TargetProjectId</a>" : <i>String</i>,
        "<a href="#targetclustername" title="TargetClusterName">TargetClusterName</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#enablesynchronouscreation" title="EnableSynchronousCreation">EnableSynchronousCreation</a>" : <i>Boolean</i>,
        "<a href="#synchronouscreationoptions" title="SynchronousCreationOptions">SynchronousCreationOptions</a>" : <i><a href="synchronouscreationoptions.md">SynchronousCreationOptions</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudBackUpRestoreJobs
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#instancetype" title="InstanceType">InstanceType</a>: <i>String</i>
    <a href="#instancename" title="InstanceName">InstanceName</a>: <i>String</i>
    <a href="#deliverytype" title="DeliveryType">DeliveryType</a>: <i>String</i>
    <a href="#snapshotid" title="SnapshotId">SnapshotId</a>: <i>String</i>
    <a href="#oplogts" title="OpLogTs">OpLogTs</a>: <i>String</i>
    <a href="#oploginc" title="OpLogInc">OpLogInc</a>: <i>String</i>
    <a href="#pointintimeutcseconds" title="PointInTimeUtcSeconds">PointInTimeUtcSeconds</a>: <i>Integer</i>
    <a href="#targetprojectid" title="TargetProjectId">TargetProjectId</a>: <i>String</i>
    <a href="#targetclustername" title="TargetClusterName">TargetClusterName</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#enablesynchronouscreation" title="EnableSynchronousCreation">EnableSynchronousCreation</a>: <i>Boolean</i>
    <a href="#synchronouscreationoptions" title="SynchronousCreationOptions">SynchronousCreationOptions</a>: <i><a href="synchronouscreationoptions.md">SynchronousCreationOptions</a></i>
</pre>

## Properties

#### ProjectId

The unique identifier of the project for the Atlas cluster.

_Required_: Yes

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

#### DeliveryType

Type of restore job to create.The value can be any one of download,automated or point_in_time 

_Required_: Yes

_Type_: String

_Allowed Values_: <code>download</code> | <code>automated</code> | <code>pointInTime</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SnapshotId

Unique identifier of the source snapshot ID of the restore job.

_Required_: Yes

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

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### EnableSynchronousCreation

If set to true, the CloudFormation resource will wait until the job is completed, WARNING: if the snapshot has a big load of data, the cloud formation resource might take a long time to finish leading to high costs

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SynchronousCreationOptions

Options that needs to be set to control the synchronous creation flow, this options need to be set if EnableSynchronousCreation is se to TRUE

_Required_: No

_Type_: <a href="synchronouscreationoptions.md">SynchronousCreationOptions</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

 The unique identifier of the restore job.

#### DeliveryUrl

One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.

#### Cancelled

Indicates whether the restore job was canceled.

#### Failed

Indicates whether the restore job failed.

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

