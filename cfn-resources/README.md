# cfn-resources

## MongoDB Atlas AWS CloudFormation Custom Resource Type

This folder contains the source code for each of the AWS CloudFormation 
MongoDB Atlas Resources. 

Note these are also hosted on AWS CloudFormation Public Registry under Third Party Extensions. 

### Resource Status Table

| #   | Resource                             | Status                                                 | Examples                                                                                   |
|-----|--------------------------------------|--------------------------------------------------------|--------------------------------------------------------------------------------------------|
| 1   | alert-configuration                  | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./alert-configuration/test](./alert-configuration/test)                                   |
| 2   | auditing                             | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./auditing/test](./auditing/test)                                                         |
| 3   | cloud-backup-restore-jobs            | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./cloud-backup-restore-jobs/test](./cloud-backup-restore-jobs/test)                       |
| 4   | cloud-backup-schedule                | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./cloud-backup-schedule/test](./cloud-backup-schedule/test)                               |
| 5   | cloud-backup-snapshot                | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./cloud-backup-snapshot/test](./cloud-backup-snapshot/test)                               |
| 6   | cloud-backup-snapshot-export-bucket  | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./cloud-backup-snapshot-export-bucket/test](./cloud-backup-snapshot-export-bucket/test)   |
| 7   | cloud-backup-snapshot-export-job     | ![Build](https://img.shields.io/badge/Unstable-orange) |                                                                                            |
| 8   | cloud-provider-access                | ![Build](https://img.shields.io/badge/Unstable-orange) | [./cloud-provider-access/test](./cloud-provider-access/test)                               |
| 9   | cluster                              | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./cluster/test](./cluster/test)                                                           |
| 10  | custom-dns-configuration-cluster-aws | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./custom-db-role/test](./custom-db-role/test)                                             |
| 11  | custom-db-role                       | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./custom-dns-configuration-cluster-aws/test](./custom-dns-configuration-cluster-aws/test) |
| 12  | database-user                        | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./database-user/test](./database-user/test)                                               |
| 13  | datalakes                            | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./datalakes/test](./datalakes/test)                                                       |
| 14  | encryption-at-rest                   | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./encryption-at-rest/test](./encryption-at-rest/test)                                     |
| 15  | federated-settings-identity-provider | ![Build](https://img.shields.io/badge/Unstable-orange) |                                                                                            |
| 16  | federated-settings-org-configs       | ![Build](https://img.shields.io/badge/Unstable-orange) |                                                                                            |
| 17  | federated-settings-org-role-mapping  | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./federated-settings-org-role-mapping/test](./federated-settings-org-role-mapping/test)   |
| 18  | global-cluster-config                | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./global-cluster-config/test](./global-cluster-config/test)                               |
| 19  | ldap-configuration                   | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./ldap-configuration/test](./ldap-configuration/test)                                     |
| 20  | ldap-verify                          | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./ldap-verify/test](./ldap-verify/test)                                                   |
| 21  | maintenance-window                   | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./maintenance-window/test](./maintenance-window/test)                                     |
| 22  | network-container                    | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./network-container/test](./network-container/test)                                       |
| 23  | network-peering                      | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./network-peering/test](./network-peering/test)                                           |
| 24  | online-archive                       | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./online-archive/test](./online-archive/test)                                             |
| 25  | org-invitation                       | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./org-invitation/test](./org-invitation/test)                                             |
| 26  | private-endpoint                     | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./private-endpoint/test](./private-endpoint/test)                                         |
| 27  | private-endpoint-adl                 | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./private-endpoint-adl/test](./private-endpoint-adl/test)                                 |
| 28  | private-endpoint-regional-mode       | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./private-endpoint-regional-mode/test](./private-endpoint-regional-mode/test)             |
| 29  | project                              | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./project/test](./project/test)                                                           |
| 30  | project-invitation                   | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./project-invitation/test](./project-invitation/test)                                     |
| 31  | project-ip-access-list               | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./project-ip-access-list/test](./project-ip-access-list/test)                             |
| 32  | search-index                         | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./search-indexes/test](./search-indexes/test)                                             |
| 33  | serverless-instance                  | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./serverless-instance/test](./serverless-instance/test)                                   |
| 34  | teams                                | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./teams/test](./teams/test)                                                               |
| 35  | third-party-integration              | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./third-party-integration/test](./third-party-integration/test)                           |
| 36  | trigger                              | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./trigger/test](./trigger/test)                                                           |
| 37  | X509AuthenticationDatabaseUser       | ![Build](https://img.shields.io/badge/Beta-yellow)     | [./x509-authentication-database-user/test](./x509-authentication-database-user/test)       |

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




