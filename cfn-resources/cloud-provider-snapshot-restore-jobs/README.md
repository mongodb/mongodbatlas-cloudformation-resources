# MongoDB::Atlas::CloudProviderSnapshotRestoreJobs

## Description
This resource allows you to create, cancel, get one or list all cloud provider snapshot restore jobs.

## Attributes
`Id` :  The unique identifier of the restore job.<br>
`DeliveryUrl` : One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is `download`.<br>
`Cancelled` : Indicates whether the restore job was canceled.<br>
`CreatedAt` : UTC ISO 8601 formatted point in time when Atlas created the restore job.<br>
`Expired` : Indicates whether the restore job expired.<br>
`ExpiresAt` : UTC ISO 8601 formatted point in time when the restore job expires.<br>
`FinishedAt` : UTC ISO 8601 formatted point in time when the restore job completed.<br>
`Timestamp` : Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.<br>
`Links` : One or more links to sub-resources and/or related resources.<br>
`OpLogTs` : If you performed a Point-in-Time restores at a time specified by a timestamp from the oplog, oplogTs indicates the timestamp used.<br>
`PointInTimeUtcSeconds` : If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used.<br>
`TargetProjectId` : Name of the target Atlas project of the restore job. Only visible if deliveryType is automated.<br>
`TargetClusterName` : Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.<br>

## Parameters
`ProjectId` *(required)* : The unique identifier of the project for the Atlas cluster.<br>
`ClusterName` *(required)* : The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs.<br>
`DeliveryType` *(required)* : Type of restore job to create. <br>
`SnapshotId` *(required)* : Unique identifier of the source snapshot ID of the restore job.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit