# cfn-resources

## MongoDB Atlas AWS CloudFormation Custom Resource Type

This folder contains the source code for each of the AWS CloudFormation 
MongoDB Atlas Resource Types.

### Resource Status Table

| Resource                             | Status                                             | Examples                                                                                   |
|--------------------------------------|----------------------------------------------------|--------------------------------------------------------------------------------------------|
| alert-configuration                  | ![Build](https://img.shields.io/badge/Beta-yellow) | [./alert-configuration/test](./alert-configuration/test)                                   |
| auditing                             | ![Build](https://img.shields.io/badge/Beta-yellow) | [./auditing/test](./auditing/test)                                                         |
| cloud-backup-restore-jobs            | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cloud-backup-restore-jobs/test](./cloud-backup-restore-jobs/test)                       |
| cloud-backup-schedule                | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cloud-backup-schedule/test](./cloud-backup-schedule/test)                               |
| cloud-backup-snapshot                | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cloud-backup-snapshot/test](./cloud-backup-snapshot/test)                               |
| cloud-backup-snapshot-export-bucket  | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cloud-backup-snapshot-export-bucket/test](./cloud-backup-snapshot-export-bucket/test)   |
| cloud-provider-access                | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cloud-provider-access/test](./cloud-provider-access/test)                               |
| cluster                              | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cluster/test](./cluster/test)                                                           |
| custom-dns-configuration-cluster-aws | ![Build](https://img.shields.io/badge/Beta-yellow) | [./custom-db-role/test](./custom-db-role/test)                                             |
| custom-db-role                       | ![Build](https://img.shields.io/badge/Beta-yellow) | [./custom-dns-configuration-cluster-aws/test](./custom-dns-configuration-cluster-aws/test) |
| database-user                        | ![Build](https://img.shields.io/badge/Beta-yellow) | [./database-user/test](./database-user/test)                                               |
| datalakes                            | ![Build](https://img.shields.io/badge/Beta-yellow) | [./datalakes/test](./datalakes/test)                                                       |
| encryption-at-rest                   | ![Build](https://img.shields.io/badge/Beta-yellow) | [./encryption-at-rest/test](./encryption-at-rest/test)                                     |
| federated-settings-org-role-mapping  | ![Build](https://img.shields.io/badge/Beta-yellow) | [./federated-settings-org-role-mapping/test](./federated-settings-org-role-mapping/test)   |
| global-cluster-config                | ![Build](https://img.shields.io/badge/Beta-yellow) | [./global-cluster-config/test](./global-cluster-config/test)                               |
| ldap-configuration                   | ![Build](https://img.shields.io/badge/Beta-yellow) | [./ldap-configuration/test](./ldap-configuration/test)                                     |
| ldap-verify                          | ![Build](https://img.shields.io/badge/Beta-yellow) | [./ldap-verify/test](./ldap-verify/test)                                                   |
| maintenance-window                   | ![Build](https://img.shields.io/badge/Beta-yellow) | [./maintenance-window/test](./maintenance-window/test)                                     |
| network-container                    | ![Build](https://img.shields.io/badge/Beta-yellow) | [./network-container/test](./network-container/test)                                       |
| network-peering                      | ![Build](https://img.shields.io/badge/Beta-yellow) | [./network-peering/test](./network-peering/test)                                           |
| online-archive                       | ![Build](https://img.shields.io/badge/Beta-yellow) | [./online-archive/test](./online-archive/test)                                             |
| org-invitation                       | ![Build](https://img.shields.io/badge/Beta-yellow) | [./org-invitation/test](./org-invitation/test)                                             |
| private-endpoint                     | ![Build](https://img.shields.io/badge/Beta-yellow) | [./private-endpoint/test](./private-endpoint/test)                                         |
| private-endpoint-adl                 | ![Build](https://img.shields.io/badge/Beta-yellow) | [./private-endpoint-adl/test](./private-endpoint-adl/test)                                 |
| private-endpoint-regional-mode       | ![Build](https://img.shields.io/badge/Beta-yellow) | [./private-endpoint-regional-mode/test](./private-endpoint-regional-mode/test)             |
| project                              | ![Build](https://img.shields.io/badge/Beta-yellow) | [./project/test](./project/test)                                                           |
| project-invitation                   | ![Build](https://img.shields.io/badge/Beta-yellow) | [./project-invitation/test](./project-invitation/test)                                     |
| project-ip-access-list               | ![Build](https://img.shields.io/badge/Beta-yellow) | [./project-ip-access-list/test](./project-ip-access-list/test)                             |
| search-index                         | ![Build](https://img.shields.io/badge/Beta-yellow) | [./search-indexes/test](./search-indexes/test)                                             |
| serverless-instance                  | ![Build](https://img.shields.io/badge/Beta-yellow) | [./serverless-instance/test](./serverless-instance/test)                                   |
| teams                                | ![Build](https://img.shields.io/badge/Beta-yellow) | [./teams/test](./teams/test)                                                               |
| third-party-integration              | ![Build](https://img.shields.io/badge/Beta-yellow) | [./third-party-integration/test](./third-party-integration/test)                           |
| trigger                              | ![Build](https://img.shields.io/badge/Beta-yellow) | [./trigger/test](./trigger/test)                                                           |
Legend
---
| Badge | Meaning |
| --- | --- |
| ![Build](https://img.shields.io/badge/GA-green) | GA, production ready |
| ![Build](https://img.shields.io/badge/Beta-yellow) | Beta status, stable dev/testing |
| ![Build](https://img.shields.io/badge/Unstable-orange) | Not fully tested |
| ![Build](https://img.shields.io/badge/Beta-Admin-grey) | Beta status, stable for dev/testing but not only for advanced use |

## Test framework

### Requirements for local dev testing

* aws cli
* cfn cli
* python
* go
* bash
* [mongocli](https://github.com/mongodb/mongocli) (you don't *need* this but you want it)

### How we handle ApiKeys

All apikey are injected through environment variables. 
We have a helper script which can export your `mongocli` profile, so this makes it very easy to switch Atlas environments.

To use this, first download and install [mongocli](mongocli).
Next, run `mongocli config` and then;

```bash
$source <(./quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
$env | grep ATLAS
ATLAS_PUBLIC_KEY=XXXXXX
ATLAS_PRIVATE_KEY=XXXXXX
ATLAS_ORG_ID=XXXXXX
```

### Deployment.template.yaml
Note: section for project example line #145 contains an s3 bucket  configuration string please edit value to S3 bucket that you require -->  (SchemaHandlerPackage="s3://replace-placeholder-bucket-name-here/resources/mongodb-atlas-project.zip",)

### How tests are structured

Each resource has a folder called `test` with 3 items:

1. *<resource_name>*.sample-cfn-request.json
        Sample JSON template request to use for local testing with `cfn invoke`

2. *<resource_name>*.create-sample-cfn-request.sh
        The create-sample-cfn-request script injects parameters into the sample json template. You run this script, passing resource specific parameters, and the tooling will inject ApiKeys based upon the exported configuration from above. 
        _See each resource README for specific testing documentation._

3. *<resource_name>*.sample-template.yaml
        Sample real cloudformation template you can run with `aws cloudformation create-stack` or using  [../../quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh]( ../../quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh) 
        _See each resource README for specific testing documentation._




