# cfn-resources

## MongoDB Atlas AWS CloudFormation Custom Resource Type

This folder contains the source code for each of the AWS CloudFormation 
MongoDB Atlas Resources. 

Note these are also hosted on AWS CloudFormation Public Registry under Third Party Extensions. 

### Resource Status Table

| #   | Resource                                                    | Status                                          | Examples                                                                                                                                            | Local Testing Scripts                                                                                                                    |
|-----|-------------------------------------------------------------|-------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| 1   | alert-configuration                                         | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/alert-configuration/alert-configuration.json)                                                                                 | [./alert-configuration/test](./alert-configuration/test)                                                                                 |
| 2   | auditing                                                    | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/auditing/auditing.json)                                                                                                       | [./auditing/test](./auditing/test)                                                                                                       |
| 3   | cloud-backup-restore-jobs                                   | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/cloud-backup-restore-jobs/restore.json)                                                                                       | [./cloud-backup-restore-jobs/test](./cloud-backup-restore-jobs/test)                                                                     |
| 4   | cloud-backup-schedule                                       | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/cloud-backup-schedule/cloudBackupSchedule.json)                                                                               | [./cloud-backup-schedule/test](./cloud-backup-schedule/test)                                                                             |
| 5   | cloud-backup-snapshot                                       | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/cloud-backup-snapshot/snapshot.json)                                                                                          | [./cloud-backup-snapshot/test](./cloud-backup-snapshot/test)                                                                             |
| 6   | cloud-backup-snapshot-export-bucket                         | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/cloud-backup-snapshot-export-bucket/CloudBackupSnapshotExportBucket.json)                                                     | [./cloud-backup-snapshot-export-bucket/test](./cloud-backup-snapshot-export-bucket/test)                                                 |
| 7   | cluster                                                     | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/cluster/cluster.json)                                                                                                         | [./cluster/test](./cluster/test)                                                                                                         |
| 8   | custom-dns-configuration-cluster-aws                        | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/custom-dns-configuration-cluster-aws/CustomDnsConfigurationClusterAws.json)                                                   | [./custom-db-role/test](./custom-db-role/test)                                                                                           |
| 9   | custom-db-role                                              | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/custom-db-role/custom-db-role.json)                                                                                           | [./custom-dns-configuration-cluster-aws/test](./custom-dns-configuration-cluster-aws/test)                                               |
| 10  | database-user                                               | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/database-user/user.json)                                                                                                      | [./database-user/test](./database-user/test)                                                                                             |
| 11  | datalakes                                                   | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/data-lakes/datalake.json)                                                                                                     | [./datalakes/test](./datalakes/test)                                                                                                     |
| 12  | encryption-at-rest                                          | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/encryption-at-rest/encryption-at-rest.json)                                                                                   | [./encryption-at-rest/test](./encryption-at-rest/test)                                                                                   |
| 13  | federated-settings-org-role-mapping                         | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/federated-settings-org-role-mapping/federatedSettingsOrgRoleMapping.json)                                                     | [./federated-settings-org-role-mapping/test](./federated-settings-org-role-mapping/test)                                                 |
| 14  | global-cluster-config                                       | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/global-cluster-config/global-cluster-config.json)                                                                             | [./global-cluster-config/test](./global-cluster-config/test)                                                                             |
| 15  | ldap-configuration                                          | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/ldap-configuration/LDAPConfiguration.json)                                                                                    | [./ldap-configuration/test](./ldap-configuration/test)                                                                                   |
| 16  | ldap-verify                                                 | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/LDAPVerify/LDAPVerify.json)                                                                                                   | [./ldap-verify/test](./ldap-verify/test)                                                                                                 |
| 17  | maintenance-window                                          | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/maintenance-window/maintenance-window.json)                                                                                   | [./maintenance-window/test](./maintenance-window/test)                                                                                   |
| 18  | network-container                                           | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/network-container/container.json)                                                                                             | [./network-container/test](./network-container/test)                                                                                     |
| 19  | network-peering                                             | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/network-peering/peering.json)                                                                                                 | [./network-peering/test](./network-peering/test)                                                                                         |
| 20  | online-archive                                              | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/online-archive/online-archive.json)                                                                                           | [./online-archive/test](./online-archive/test)                                                                                           |
| 21  | org-invitation                                              | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/org-invitation/org-invitation.json)                                                                                           | [./org-invitation/test](./org-invitation/test)                                                                                           |
| 22  | private-endpoint                                            | ![Build](https://img.shields.io/badge/Deprecated-red) | [example](../examples/private-endpoint/privateEndpoint.json)                                                                                        | [./private-endpoint/test](./private-endpoint/test)                                                                                       |
| 23  | private-endpoint-adl                                        | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/private-endpoint-adl/endpoint-adl.json)                                                                                       | [./private-endpoint-adl/test](./private-endpoint-adl/test)                                                                               |
| 24  | private-endpoint-regional-mode                              | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/private-endpoint-regional-mode/privateEndpointRegionalMode.json)                                                              | [./private-endpoint-regional-mode/test](./private-endpoint-regional-mode/test)                                                           |
| 25  | project                                                     | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/project/project.json)                                                                                                         | [./project/test](./project/test)                                                                                                         |
| 26  | project-invitation                                          | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/project-invitation/project-invitation.json)                                                                                   | [./project-invitation/test](./project-invitation/test)                                                                                   |
| 27  | project-ip-access-list                                      | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/project-ip-access-list/ip-access-list.yaml)                                                                                   | [./project-ip-access-list/test](./project-ip-access-list/test)                                                                           |
| 28  | search-index                                                | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/search-index/searchIndex.json)                                                                                                | [./search-indexes/test](./search-indexes/test)                                                                                           |
| 29  | serverless-instance                                         | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/serverless-instance/serverless-instance.json)                                                                                 | [./serverless-instance/test](./serverless-instance/test)                                                                                 |
| 30  | teams                                                       | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/teams/teams.json)                                                                                                             | [./teams/test](./teams/test)                                                                                                             |
| 31  | third-party-integration                                     | ![Build](https://img.shields.io/badge/GA-green) | [example files](../examples/thirdpartyintegrations)                                                                                                 | [./third-party-integration/test](./third-party-integration/test)                                                                         |
| 32  | trigger                                                     | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/trigger/trigger.json)                                                                                                         | [./trigger/test](./trigger/test)                                                                                                         |
| 33  | X509AuthenticationDatabaseUser                              | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/x509-authentication-db-user/x509-authentication-db-user.json)                                                                 | [./x509-authentication-database-user/test](./x509-authentication-database-user/test)                                                     |
| 34  | federated-database-instance                                 | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/federated-database-instance/federatedDatabaseInstance.json)                                                                   | [./federated-database-instance/test](./federated-database-instance/test)                                                                 |
| 35  | privatelink-endpoint-service-data-federation-online-archive | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/privatelink-endpoint-service-data-federation-online-archive/privatelink-endpoint-service-data-federation-online-archive.json) | [./privatelink-endpoint-service-data-federation-online-archive/test](./privatelink-endpoint-service-data-federation-online-archive/test) |
| 36  | federated-query-limit                                       | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/federated-query-limit/federatedQueryLimit.json)                                                                               | [./federated-query-limit/test](./federated-query-limit/test)                                                                             |
| 37  | serverless-private-endpoint                                 | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/serverless-private-endpoint/serverless-private-endpoint-with-aws-private-endpoint.json)                                       | [./serverless-private-endpoint/test](./serverless-private-endpoint/test)                                                                 |
| 38  | api-key                                                     | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/api-key/api-key.json)                                                                                                         | [./api-key/test](./api-key/test)                                                                                                         |
| 39  | access-list-api-key                                         | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/access-list-api-key/access-list-api-key.json)                                                                                 | [./access-list-api-key/test](./access-list-api-key/test)                                                                                 |
| 40  | data-lake-pipeline                                          | ![Build](https://img.shields.io/badge/Deprecated-red) | [example](../examples/data-lake-pipeline/datalakepipeline.json)                                                                                     | [./data-lake-pipeline/test](./data-lake-pipeline/test)                                                                                   |
| 41  | organization                                                | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/organization/organization.json)                                                                                               | [./organization/test](./organization/test)                                                                                               |
| 42  | cloud-outage-simulation                                     | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/cluster-outage-simulation/cluster-outage-simulation.json)                                                                     | [./cloud-outage-simulation/test](./cl-outage-simulation/test)                                                                            |
| 43  | private-endpoint-service                                    | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/private-endpoint/privateEndpointV2.json)                                                                                      | [./private-endpoint-service/test](./private-endpoint-service/test)                                                                       |
| 44  | private-endpoint-aws                                        | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/private-endpoint/privateEndpointV2.json)                                                                                      | [./private-endpoint-aws/test](./private-endpoint-aws/test)   
| 45  | search-deployment                                        | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/search-deployment/search-deployment.json)                                                                                      | [./search-deployment/test](./search-deployment/test)  
| 46  | stream-instance                                        | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/atlas-streams/stream-instance/stream-instance.json)                                                                                      | [./stream-instance/test](./stream-instance/test)  
| 47  | stream-connection                                        | ![Build](https://img.shields.io/badge/GA-green) | [example](../examples/atlas-streams/stream-connection/stream-connection.json)                                                                                      | [./stream-connection/test](./stream-connection/test)  
|47  | resource-policy                                        | ![Build](https://img.shields.io/badge/Beta-yellow) | [example](../examples/resource-policy/resource-policy.json)                                                                                      | [./resource-policy/test](./resource-policy/test)  

## Resource Import Operations

All MongoDB Atlas AWS CloudFormation resources support the import operation, allowing you to bring existing Atlas resources under CloudFormation management. When importing resources, please consider the following:

### Import Requirements and Considerations

1. **DeletionPolicy Attribute**: Your resource template must include the [DeletionPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) attribute in order to specify what happens to the resource when it is removed from the stack.

2. **UPDATE Operation Required**: The import operation executes the UPDATE operation behind the scenes. Therefore, UPDATE functionality must be properly implemented for the resource type to support import. For more datails on when the update is run, please see [Create a stack from existing resources using the AWS Management Console](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-new-stack.html#resource-import-new-stack-console) or [Import an existing resource into a stack using the AWS Management Console](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-existing-stack.html#resource-import-existing-stack-console) or 

3. **Outputs Restrictions**: You cannot modify or add [Outputs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) to the template during the import operation. However, you can add outputs after the import is complete.

For more details, refer to the [official AWS Resource Import documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import.html).

Legend
---
| Badge | Meaning |
| --- | --- |
| ![Build](https://img.shields.io/badge/GA-green) | GA, production ready |
| ![Build](https://img.shields.io/badge/Beta-yellow) | Beta status, stable dev/testing |
| ![Build](https://img.shields.io/badge/Unstable-orange) | Not fully tested |
| ![Build](https://img.shields.io/badge/Beta-Admin-grey) | Beta status, stable for dev/testing but not only for advanced use |
| ![Build](https://img.shields.io/badge/Deprecated-red) | Deprecated |

## Test framework

### Requirements for local dev testing

* aws cli
* cfn cli
* python
* go
* bash
* [atlascli](https://github.com/mongodb/mongodb-atlas-cli) (you don't *need* this but will make testing easier)

### How we handle ApiKeys

All apikey are injected through environment variables. 
We have a helper script which can export your `mongocli` profile, so this makes it very easy to switch Atlas environments.

To use this, first download and install [mongocli](mongocli).
Next, run `mongocli config` and then;

```bash
$source <(./quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
$env | grep ATLAS
MONGODB_ATLAS_PUBLIC_KEY=XXXXXX
MONGODB_ATLAS_PRIVATE_KEY=XXXXXX
MONGODB_ATLAS_ORG_ID=XXXXXX
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




