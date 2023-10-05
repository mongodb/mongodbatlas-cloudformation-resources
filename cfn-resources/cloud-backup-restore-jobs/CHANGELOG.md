# Changelog

## (2023-20-05)

**(BREAKING CHANGE) ADDED SUPPORT TO SERVERLESS INSTANCE**

originally the resource was partially supporting serverless instance, it required some changes on the schema to fully support the 
creation of restore jobs for serverless instance:

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
ClusterName: !Ref 'ClusterNameSource' <----- Old identifier (deprecated)
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

he next options had been added to back up restore job:

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
  - TimeOutInSeconds: time in seconds until the resource stops witing and runturns success or fail depending on the ReturnSuccessIfTimeOut
  - CallbackDelaySeconds: time to wait until the resource checks again if the job is finished
  - ReturnSuccessIfTimeOut: if set to true, the process will return success, in the event of a timeOut

**FIXED READ AND DELETE VALIDATION FOR ALREADY EXISTS ERRORS**

Cloud Formation requires any third party resource to fail with a "NotFound" error in case the user tries to read it after being deleted

in order to complay to this behaviour, we introduce the next behavriour on the resource:

- **If the job is marked as Cencelled:** it will faild with a Not Found error, both in the READ and DELETE flows
- **if the Job is marked as Finished, Expired or Failed:** the READ will be excecuted (not returning error),
and the DELETE flow will return success being this final states
