# MongoDB::Atlas::FederatedSettingsIdentityProvider

## Description
Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

The resource is in status UNSTABLE: 
is not possible to deploy the resource before completing the Pendings

Pending:
- Functionlity:
  - CREATE: the resource requires a CREATE function
  - DELETE: the resource requires a DELETE function
- Testing:
  - the testing files provided (test file) are only representative, and might need to be edited once create and delte are implemented
- Example file:
  - An example file must be provided

the functionality here has been replicated from the terraform api

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Example

TODO: add example file once the resource is finished

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project
./test/project.create-sample-cfn-request.sh YourProjectName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Usage

The [launch-x-quickstart.sh](../../quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh) script
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/project/test/project.sample-template.yaml SampleProject1 ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID}
```

## For More Information
See the MongoDB Atlas API [Federated-Authentication](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Federated-Authentication) Documentation.
