# MongoDb::Atlas::X509AuthenticationDatabaseUser
## Description
Returns, edits, and removes user-managed X.509 configurations.
Also returns and generates MongoDB Cloud-managed X.509 certificates for database users.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/x509-authentication-database-user
./test/x509authenticationdatabaseuser.create-sample-cfn-request.sh YourProjectID > test.request.json 
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

The [/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh](launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the x509authenticationdatabaseuser.sample-template.yaml to create a stack using the resource.
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-quickstart.sh ${repo_root}/cfn-resources/x509-authentication-database-user/test/x509authenticationdatabaseuser.sample-template.yaml SampleProject1 ParameterKey=UserName,ParameterValue=${UserName}  ParameterKey=ProjectId,ParameterValue=${ProjectId}
```

## For More Information
See the MongoDB Atlas API [Project Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/X.509-Authentication-for-Database-Users)