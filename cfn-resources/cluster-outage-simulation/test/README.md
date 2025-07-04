# Cluster outage simulation

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Cloud backup snapshot L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Cluster

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- Cluster Outage Simulation is created successfully:
![image](TODO: add image)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cluster-outage-simulation)
- [Resource Usage Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cluster-outage-simulation)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:
```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/cluster-outage-simulation
./test/cluster-outage-simulation.create-sample-cfn-request.sh YourProjectID ClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.
