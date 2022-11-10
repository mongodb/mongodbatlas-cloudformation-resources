# MongoDB::Atlas::CloudBackupSchedule

## Description
Returns, adds,  and removes cloud backup schedule.
This resource allows you to get, add or remove cloud backup schedule for a cluster within the specified project

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```


Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default



For more information see: MongoDB Atlas API Cloud BackUp Schedule [Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backup-Schedule) Documentation.