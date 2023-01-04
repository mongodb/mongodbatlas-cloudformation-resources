# cfn-resources

## MongoDB Atlas AWS CloudFormation Resources

This folder contains the source code for each of the AWS CloudFormation 
MongoDB Atlas Resources. 
Note these are also hosted on AWS CloudFormation Public Registry under Third Party Extensions. 

### Resource Status Table

|#| Resource                             | Status                                             | Examples                                                                                   |
|---|--------------------------------------|----------------------------------------------------|--------------------------------------------------------------------------------------------|
| 1 | alert-configuration                  | ![GA](https://img.shields.io/badge/GA-green) | [./alert-configuration/test](./alert-configuration/test)                                   |
| 2 | auditing                             | ![GA](https://img.shields.io/badge/GA-green) | [./auditing/test](./auditing/test)                                                         |
| 3 | cloud-backup-restore-jobs            | ![GA](https://img.shields.io/badge/GA-green) | [./cloud-backup-restore-jobs/test](./cloud-backup-restore-jobs/test)                       |
| 4 | cloud-backup-schedule                | ![GA](https://img.shields.io/badge/GA-green) | [./cloud-backup-schedule/test](./cloud-backup-schedule/test)                               |
| 5 | cloud-backup-snapshot                | ![GA](https://img.shields.io/badge/GA-green) | [./cloud-backup-snapshot/test](./cloud-backup-snapshot/test)                               |
| 6 | cloud-backup-snapshot-export-bucket  | ![GA](https://img.shields.io/badge/GA-green) | [./cloud-backup-snapshot-export-bucket/test](./cloud-backup-snapshot-export-bucket/test)   |
| 7 | cloud-provider-access                | ![BETA](https://img.shields.io/badge/Beta-yellow) | [./cloud-provider-access/test](./cloud-provider-access/test)                               |
| 8 | cluster                              | ![GA](https://img.shields.io/badge/GA-green) | [./cluster/test](./cluster/test)                                                           |
| 9 | custom-dns-configuration-cluster-aws | ![GA](https://img.shields.io/badge/GA-green) | [./custom-db-role/test](./custom-db-role/test)                                             |
| 10 | custom-db-role                       | ![GA](https://img.shields.io/badge/GA-green) | [./custom-dns-configuration-cluster-aws/test](./custom-dns-configuration-cluster-aws/test) |
| 11 | database-user                        | ![GA](https://img.shields.io/badge/GA-green) | [./database-user/test](./database-user/test)                                               |
| 12 | datalakes                            | ![GA](https://img.shields.io/badge/GA-green) | [./datalakes/test](./datalakes/test)                                                       |
| 13 | encryption-at-rest                   | ![GA](https://img.shields.io/badge/GA-green) | [./encryption-at-rest/test](./encryption-at-rest/test)                                     |
| 14 | federated-settings-org-role-mapping  | ![GA](https://img.shields.io/badge/GA-green) | [./federated-settings-org-role-mapping/test](./federated-settings-org-role-mapping/test)   |
| 15 | global-cluster-config                | ![GA](https://img.shields.io/badge/GA-green) | [./global-cluster-config/test](./global-cluster-config/test)                               |
| 16 | ldap-configuration                   | ![GA](https://img.shields.io/badge/GA-green) | [./ldap-configuration/test](./ldap-configuration/test)                                     |
| 17 | ldap-verify                          | ![GA](https://img.shields.io/badge/GA-green) | [./ldap-verify/test](./ldap-verify/test)                                                   |
| 18 | maintenance-window                   | ![GA](https://img.shields.io/badge/GA-green) | [./maintenance-window/test](./maintenance-window/test)                                     |
| 19 | network-container                    | ![GA](https://img.shields.io/badge/GA-green) | [./network-container/test](./network-container/test)                                       |
| 20 | network-peering                      | ![GA](https://img.shields.io/badge/GA-green) | [./network-peering/test](./network-peering/test)                                           |
| 21 | online-archive                       | ![GA](https://img.shields.io/badge/GA-green) | [./online-archive/test](./online-archive/test)                                             |
| 22 | org-invitation                       | ![GA](https://img.shields.io/badge/GA-green) | [./org-invitation/test](./org-invitation/test)                                             |
| 23 | private-endpoint                     | ![GA](https://img.shields.io/badge/GA-green) | [./private-endpoint/test](./private-endpoint/test)                                         |
| 24 | private-endpoint-adl                 | ![GA](https://img.shields.io/badge/GA-green) | [./private-endpoint-adl/test](./private-endpoint-adl/test)                                 |
| 25 | private-endpoint-regional-mode       | ![GA](https://img.shields.io/badge/GA-green) | [./private-endpoint-regional-mode/test](./private-endpoint-regional-mode/test)             |
| 26 | project                              | ![GA](https://img.shields.io/badge/GA-green) | [./project/test](./project/test)                                                           |
| 27 | project-invitation                   | ![GA](https://img.shields.io/badge/GA-green) | [./project-invitation/test](./project-invitation/test)                                     |
| 28 | project-ip-access-list               | ![GA](https://img.shields.io/badge/GA-green) | [./project-ip-access-list/test](./project-ip-access-list/test)                             |
| 29 | search-index                         | ![GA](https://img.shields.io/badge/GA-green) | [./search-indexes/test](./search-indexes/test)                                             |
| 30 | serverless-instance                  | ![GA](https://img.shields.io/badge/GA-green) | [./serverless-instance/test](./serverless-instance/test)                                   |
| 31 | teams                                | ![GA](https://img.shields.io/badge/GA-green) | [./teams/test](./teams/test)                                                               |
| 32 | third-party-integration              | ![GA](https://img.shields.io/badge/GA-green) | [./third-party-integration/test](./third-party-integration/test)                           |
| 33 | trigger                              | ![GA](https://img.shields.io/badge/GA-green) | [./trigger/test](./trigger/test)                                                           

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




