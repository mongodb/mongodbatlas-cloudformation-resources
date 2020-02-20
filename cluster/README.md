# MongoDB::Atlas::Cluster

## Description
The cluster resource provides access to your cluster configurations. The resource lets you create, edit and delete clusters. The resource requires your Project ID.

## Attributes
`ID` : Unique identifier of the cluster.<br>
`MongoURI` : Base connection string for the cluster.<br>
`MongoURIUpdated` : Timestamp in ISO 8601 date and time format in UTC when the connection string was last updated. The connection string changes if you update any of the other values.<br>
`MongoURIWithOptions` : Connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.<br>
`Paused` : Flag that indicates whether the cluster is paused or not.<br>
`SrvAddress` : Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS. The mongoURI parameter lists additional options.<br>
`StateName` : Current state of the cluster.<br>

## Parameters
`AutoScaling` *(optional)* : Configure your cluster to automatically scale its storage and cluster tier.<br>
`BackupEnabled` *(optional)* : Applicable only for M10+ clusters. Set to true to enable Atlas continuous backups for the cluster. Set to false to disable continuous backups for the cluster. Atlas deletes any stored snapshots. See the continuous backup Snapshot Schedule for more information. You cannot enable continuous backups if you have an existing cluster in the project with Cloud Provider Snapshots enabled. The default value is false.<br>
`BiConnector` *(optional)* : Specifies BI Connector for Atlas configuration on this cluster.<br>
`ClusterType` *(optional)* : Type of the cluster that you want to create.<br>
`DiskSizeGB` *(optional)* : Capacity, in gigabytes, of the host’s root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive integer.<br>
`EncryptionAtRestProvider` *(optional)* : Set the Encryption at Rest parameter.<br>
`ProjectID` *(required)* : Unique identifier of the project the cluster belongs to.<br>
`Labels` *(optional)* : Array containing key-value pairs that tag and categorize the cluster.<br>
`MongoDBVersion` *(optional)* : Major version of the cluster to deploy.<br>
`MongoDBMajorVersion` *(optional)* : Major version of the cluster to deploy.<br>
`Name` *(required)* : Name of the cluster. Once the cluster is created, its name cannot be changed.<br>
`NumShards` *(optional)* : Positive integer that specifies the number of shards to deploy for a sharded cluster.<br>
`PitEnabled` *(optional)* : Flag that indicates if the cluster uses Point-in-Time backups. If set to true, providerBackupEnabled must also be set to true.<br>
`ProviderSettings` *(required)* : Configuration for the provisioned servers on which MongoDB runs.<br>
`ProviderBackupEnabled` *(optional)* : Applicable only for M10+ clusters. Set to true to enable Atlas Cloud Provider Snapshots backups for the cluster. Set to false to disable Cloud Provider Snapshots backups for the cluster. You cannot enable Cloud Provider Snapshots if you have an existing cluster in the project with continuous backups enabled. Note that you must set this value to true for NVMe clusters. The default value is false.<br>
`ReplicationFactor` *(optional)* : ReplicationFactor is deprecated. Use replicationSpecs.<br>
`ReplicationSpecs` *(optional)* : Configuration for cluster regions.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...