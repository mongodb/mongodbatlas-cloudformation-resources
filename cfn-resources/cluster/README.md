# MongoDB::Atlas::Cluster

## Description
The cluster resource provides access to your cluster configurations. The resource lets you create, edit and delete clusters.

Resource for managing [Clusters](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters).

## Requirements

The resource requires your Project ID.

## Attributes and Parameters

Please consult the [Resource Docs](docs/README.md)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/cluster
./test/cluster.create-sample-cfn-request.sh YourProjectID YourClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json 
cfn invoke resource DELETE test.request.json 
cd -
```

Both CREATE & DELETE tests must pass.

## Installation

Installation currently requires the follow 2 steps to build and then submit/register the 
new MongoDB::Atlas::Cluster Resource Type into your AWS Region. Note, this command uses the
default AWS region.

```bash
TAGS=logging make
cfn submit --verbose --set-default
```
## Cloudformation Examples

Please see the [CFN Template](test/cluster.sample-cfn-request.json) for example resource

## Integration Testing w/ AWS

Once the resource is installed, you can do integrated testing from your shell to AWS.

The [launch-x-quickstart.sh](../../quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the [CFN Template](../../examples/cluster/cluster.json) to create a stack using the resource.
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/cluster/test/cluster.sample-template.yaml SampleCluster-123 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID>
```

## For More Information
See the MongoDB Atlas API [Cluster Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Multi-Cloud-Clusters ) documentation.