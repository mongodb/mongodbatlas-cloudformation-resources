# Network Peering

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Network peering L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Atlas Container

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should be able to see the network peering in the "Network Access" page:
![image](https://user-images.githubusercontent.com/5663078/227514067-123c7343-1066-4ba7-802a-03a73a810c78.png)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-network-peering)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/reference/atlas-operator/ak8so-network-peering/)

## Unit Testing Locally

Suggested to use a new Project to test the Network Peering resource.
(Shortcut: `mongocli iam projects create Network-Peering-Test-1` can grab the id)

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/network-peering
 ./test/networkpeering.create-sample-cfn-request.sh <PROJECT_ID> "10.0.0.0/16" <YOUR_VPC_ID>  > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
cd -
```

Both CREATE & DELETE tests must pass.