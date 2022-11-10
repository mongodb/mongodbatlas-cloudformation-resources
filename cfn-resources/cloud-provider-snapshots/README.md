# MongoDB::Atlas::CloudProviderSnapshots

## Description
This resource allows you to take one on-demand snapshot, get one or all cloud provider snapshot and delete one cloud provider snapshot.

## Attributes
`Id` : Unique identifier of the snapshot.<br>
`RetentionInDays` : The number of days that Atlas should retain the on-demand snapshot. <br>
`CreatedAt` : UTC ISO 8601, formatted point in time when Atlas took the snapshot.<br>
`MasterKeyUuid` : Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot.<br>
`MongoVersion` : Version of the MongoDB server.<br>
`SnapshotType` : Specified the type of snapshot.<br>
`Status` : Current status of the snapshot.<br>
`StorageSizeBytes` : Specifies the size of the snapshot in bytes.<br>
`Type` : Specifies the type of cluster.<br>

## Parameters
`ProjectId` *(required)* : The unique identifier of the project for the Atlas cluster.<br>
`ClusterName` *(required)* : The name of the Atlas cluster that contains the snapshots you want to retrieve.<br>
`Description` *(required)* : Description of the on-demand snapshot.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...