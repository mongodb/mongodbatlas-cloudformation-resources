# MongoDB::Atlas::NetworkContainer

## Description
This resource allows you to list and delete network containers, ONLY. 
Network containers are required for network peering. With each Atlas Project having exactly 1 container per AWS region that will be peered with. The containers are provisioned dyanmically through the [Network Peering](../network-peering) resource. Direct use of this Resource Type is not expected, however it is included for completeness and supportability.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Unit Testing Locally

The Network Container resource and it's companion [Network Peering](../network-peering) should be unit tested together.

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
There is only 1 Network Container resource per Atlas project for AWS. So depending on your project the CREATE test may fail.

```
cfn invoke LIST test.request.json 
cfn invoke DELETE test.request.json 
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

LIST, & DELETE tests must pass 

## Installation

Installation currently requires the follow 2 steps to build and then submit/register the 
new MongoDB::Atlas::networkcontainer Resource Type into your AWS Region. Note, this command uses the
default AWS region.

```bash
TAGS=logging make
cfn submit --verbose --set-default
```

## Integration Testing w/ AWS

Once the resource is installed, you can do integrated testing from your shell to AWS.

The [/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh](launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the project.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-project-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/network-container/test/networkcontainer.sample-template.yaml SampleNetworkContainer-123 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID>
 
 
```

