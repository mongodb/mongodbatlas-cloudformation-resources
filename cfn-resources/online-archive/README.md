# MongoDB::Atlas::OnlineArchive

## Description
Returns, adds, edits, or removes an online archive.

## Attributes & Parameters

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
cd ${repo_root}/cfn-resources/online-archive
./test/cfn-test-create-inputs.sh YourProjectID YourClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json 
cfn invoke resource DELETE test.request.json 
cd -
```

Both CREATE & DELETE tests must pass.

## Installation

Installation currently requires the follow 2 steps to build and then submit/register the
new MongoDB::Atlas::SearchIndex Resource Type into your AWS Region. Note, this command uses the
default AWS region.

```bash
TAGS=logging make
cfn submit --verbose --set-default
```
## CloudFormation Examples

Please see the [test/inputs_1_create.template.json](test/inputs_1_create.template.json) for example resource for example resource.

## Integration Testing w/ AWS

Once the resource is installed, you can do integrated testing from your shell to AWS.

The [../../quickstart-mongodb-atlas/scripts/launch-quickstart.sh]( ../../quickstart-mongodb-atlas/scripts/launch-quickstart.sh)  script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/online-archive/test/inputs_1_create.json SampleCluster-123 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID>
```

## For More Information
See the MongoDB Atlas API [Online Archive](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Online-Archive) documentation.