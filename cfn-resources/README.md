# cfn-resources

## MongoDB Atlas AWS CloudFormation Custom Resource Type

This folder contains the source code for each of the AWS CloudFormation 
MongoDB Atlas Resource Types.

### Resource Status Table

| Resource | Status | Examples |
| --- | --- | --- |
| project | ![Build](https://img.shields.io/badge/Beta-yellow) | [./project/test](./project/test) |
| project-ip-access-list | ![Build](https://img.shields.io/badge/Beta-yellow) | [./project-ip-access-list/test](./project-ip-access-list/test) |
| cluster | ![Build](https://img.shields.io/badge/Beta-yellow) | [./cluster/test](./cluster/test) |
| database-user | ![Build](https://img.shields.io/badge/Beta-yellow) | [./database-user/test](./database-user/test) |
| network-peering | ![Build](https://img.shields.io/badge/Beta-yellow) | [./network-peering/test](./network-peering/test) |
| encryption-at-rest | ![Build](https://img.shields.io/badge/Unstable-orange) | [./encryption-at-rest/test](./encryption-at-rest/test) |
| cloud-provider-snapshots | ![Build](https://img.shields.io/badge/Unstable-orange) | [../cloud-provider-snapshots/test](./cloud-provider-snapshots/test) |
| cloud-provider-snapshot-restore-jobs | ![Build](https://img.shields.io/badge/Unstable-orange) | [./cloud-provider-snapshot-restore-jobs/test](./cloud-provider-snapshot-restore-jobs/test) | 
| network-container | ![Build](https://img.shields.io/badge/Beta-Admin-grey) | [./network-container/test](./network-container/test) |


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




