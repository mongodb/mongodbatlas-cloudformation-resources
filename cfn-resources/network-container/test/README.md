# Network Container

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Network container L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- You should be able to see the container created when running the following Atlas CLI command:
```bash
atlas networking container ls --projectId 641d818e36e2eb47d038f8c2
[
  {
    "atlasCidrBlock": "10.8.2.0/24",
    "id": "<network-container-id>",
    "providerName": "AWS",
    "provisioned": false,
    "regionName": "US_EAST_1"
  }
]
```



## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Network-Peering)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/reference/atlas-operator/ak8so-network-peering/)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/network-container
./test/networkcontainer.create-sample-cfn-request.sh <PROJECT_ID> > test.request.json
echo "Sample request:"
cat test.request.json
```
There is only 1 Network Container resource per Atlas project for AWS for a given region. So depending on your project the CREATE test may fail.

```
cfn invoke resource CREATE test.request.json 
cfn invoke resource READ test.request.json 
cfn invoke resource UPDATE test.request.json
cfn invoke resource LIST test.request.json 
cfn invoke resource DELETE test.request.json 
```

Use the `LIST` method to find the id of any existing
network container. Here is an example of the command and sample output. 

```
cfn invoke LIST test.request.json
...<output omitted>...
=== Handler response ===
{
  "message": "List Complete",
  "status": "SUCCESS",
  "resourceModel": [
    {
      "RegionName": "US_EAST_1",
      "Provisioned": "true",
      "VpcId": "vpc-ffffgggghhhhijj1232",
      "AtlasCIDRBlock": "192.168.248.0/21",
      "Id": "5f871f997cd85921961f62a5",
      "ApiKeys": {}
    }
  ],
  "bearerToken": "92f914c7-23b3-4ea5-a1e1-8215a6aa4b78",
  "resourceModels": null
}
```

You can use the `resourceModel.Id` property as the container id when creating a [Network Peering](../network-peering).

CREATE, READ, UPDATE, LIST, and DELETE tests must pass 
