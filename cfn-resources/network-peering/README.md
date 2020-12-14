# MongoDB::Atlas::NetworkPeering

## Description
This resource allows you to create, read, update and delete a network peering.
The NetworkPeering resource will automatically manage the corresponding MongoDB::Atlas::NetworkContainer for the given Project and AWS Region specified. 

The peering resource will either return an existing Atlas network container for a given (ProjectID, AWS Region) pair, or create a new one on the fly and use that for the new peering resource. This means that users of VPC Peering should not normally need to ever directly use network container resource.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

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

## Installation

Installation currently requires the follow 2 steps to build and then submit/register the 
new MongoDB::Atlas::networkpeering Resource Type into your AWS Region. Note, this command uses the
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/network-peering/test/networkpeering.sample-template.yaml Sample-NetworkPeering-1 ParameterKey=ProjectId,ParameterValue=<PROJECT_ID> ParameterKey=RouteTableCIDRBlock,ParameterValue=192.168.0.0/24 ParameterKey=VPC,ParameterValue=<YOUR_VPC_ID>  
 
```

