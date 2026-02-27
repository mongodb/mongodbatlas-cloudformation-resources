# Changelog

## (2025-02-26)

**(BREAKING CHANGE) REMOVED SUPPORT FOR SERVERLESS INSTANCES**

Serverless instances are no longer supported. The `InstanceType` field now only accepts `"cluster"`. Serverless API calls and code paths have been removed.

## (2023-10-17)

**(BREAKING CHANGE) FIELD CHANGES**

Removed fields:
- CreatedAt (string). This property was not being populated and therefore removed.

Added fields:
- Failed (bool). This property is being sent by the server so it is being made available.


## (2023-10-05)

**(BREAKING CHANGE) ADDED SUPPORT TO SERVERLESS INSTANCE**

Originally, the resource partially supported serverless instances, but it required some schema changes to fully support
the creation of restore jobs for serverless instances.

- Deprecated fields:
    - ClusterName
- New fields:
    - InstanceName: string
    - InstanceType: string ["serverless", "cluster"]

example:

Original:
``` yaml
RestoreJob:
Type: MongoDB::Atlas::CloudBackUpRestoreJobs
Properties:
ProjectId: !Ref 'ProjectSource'
ClusterName: !Ref 'ClusterNameSource' <----- Old identifier (removed)
DeliveryType: 'automated'
SnapshotId: !Ref 'Snapshot'
TargetProjectId: !Ref 'ProjectDest'
TargetClusterName: !Ref 'ClusterNameDest'
Profile: !Ref 'Profile'
```
New:
``` yaml
RestoreJob:
Type: MongoDB::Atlas::CloudBackUpRestoreJobs
Properties:
ProjectId: !Ref 'ProjectSource'
InstanceName: !Ref 'ClusterNameSource'  <----- New identifier
InstanceType: 'cluster'                 <----- New identifier
DeliveryType: 'automated'
SnapshotId: !Ref 'Snapshot'
TargetProjectId: !Ref 'ProjectDest'
TargetClusterName: !Ref 'ClusterNameDest'
Profile: !Ref 'Profile'
```

**ADDED SUPPORT TO SYNCHRONOUS CREATE FLOW**

We have introduced the capability to have a CloudFormation resource wait until the job is completed before returning a success status.

The following options have been added to the backup restore job:

``` json
"EnableSynchronousCreation" : "true",
        "SynchronousCreationOptions" : {
          "TimeOutInSeconds" : 900,
          "CallbackDelaySeconds" : 30,
          "ReturnSuccessIfTimeOut" : true
        }
```

- EnableSynchronousCreation: if set to TRUE, the resource will wait until the job is finished
- SynchronousCreationOptions:
  - TimeOutInSeconds: Time in seconds until the resource stops waiting and returns success or fails, depending on ReturnSuccessIfTimeOut.
  - CallbackDelaySeconds: Time to wait until the resource checks again if the job is finished.
  - ReturnSuccessIfTimeOut: If set to true, the resource will return success in the event of a timeOut.

**FIXED READ AND DELETE VALIDATION FOR ALREADY EXISTS ERRORS**

CloudFormation requires any third-party resource to fail with a "NotFound" error in case the user tries to read it after being deleted.

In order to comply with this behavior, we introduce the following behavior on the resource:

- **If the job is marked as Cancelled:** It will fail with a Not Found error, both in the READ and DELETE flows.
- **if the Job is marked as Finished, Expired or Failed:** The READ will be executed (not returning an error), 
and the DELETE flow will return success, being these final states.
