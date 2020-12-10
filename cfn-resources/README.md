cfn-resources
====

## MongoDB Atlas AWS CloudFormation Custom Resource Type

This folder contains the source code for each of the AWS CloudFormation 
MongoDB Atlas Resource Types.

### Resource Status Table

| Resource | Status | Examples |
| --- | --- | --- |
| cloud-provider-snapshot-restore-jobs | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [](../examples/cloud-provider-snapshot-restore-jobs) | 
| cloud-provider-snapshots | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/cloud-provider-snapshots](../examples/cloud-provider-snapshots) |
| cluster | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/cluster](../examples/cluster) |
| database-user | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | (../examples/database-user)[../examples/ database-user] |
| encryption-at-rest | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/encryption-at-rest](../examples/encryption-at-rest) |
| network-container | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/network-container](../examples/network-container) |
| network-peering | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/network-peering](../examples/network-peering) |
| project | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/project](../examples/project) |
| project-ip-access-list | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../project-ip-access-list](../examples/project-ip-access-list) |
| table | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/table](../examples/table)] |
| cloud-provider-snapshot-restore-jobs | ![Build](https://img.shields.io/badge/Ipsum-Lorem-orange) | [../examples/cloud-provider-snapshot-restore-jobs](../examples/cloud-provider-snapshot-restore-jobs) | 


## Test framework


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




