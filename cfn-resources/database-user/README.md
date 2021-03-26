# MongoDB::Atlas::DatabaseUser

## Description
The databaseUsers resource lets you retrieve, create and modify the MongoDB users in your cluster. Each user has a set of roles that provides access to the project’s databases. A user’s roles apply to all the clusters in the project: if two clusters have a products database and a user has a role granting read access on the products database, the user has that access on both clusters.

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
cd ${repo_root}/cfn-resources/database-user
./test/databaseuser.create-sample-cfn-request.sh YourProjectID SomeIAMRole > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Integration Testing w/ AWS

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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/database-user/test/databaseuser.sample-template.yaml SampleAccessList1 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID> ParameterKey=Username,ParameterValue=<USERNAME> ParameterKey=Password,ParameterValue=<PWD> 
```


