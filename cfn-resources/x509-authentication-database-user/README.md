# MongoDb::Atlas::X509AuthenticationDatabaseUser
## Description

Returns, edits, and removes user-managed X.509 configurations.
Also returns and generates MongoDB Atlas-managed X.509 certificates for database users.

## Attributes and Parameters

See the [Resource Docs](docs/README.md).

## Cloudformation Examples

See the [CFN Template](/examples/x509-authentication-db-user/x509-authentication-db-user.json) for example resource.

## Installation

From the repository root directory, run the following commands:

```
TAGS=logging make
cfn submit --verbose --set-default
```

## Usage

The [`launch-x-quickstart.sh`](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/main/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Atlas ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the [x509authenticationdatabaseuser.sample-template.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/x509-authentication-db-user/x509-authentication-db-user.json) to create a stack using the resource.
Similar to [Local Testing](#local-testing), you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session, run:

```
aws logs tail mongodb-atlas-project-logs --follow
```

Then create the stack with a helper script. The script inserts the API keys for you.

```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-quickstart.sh ${repo_root}/cfn-resources/x509-authentication-database-user/test/x509authenticationdatabaseuser.sample-template.yaml SampleProject1 ParameterKey=UserName,ParameterValue=${UserName}  ParameterKey=ProjectId,ParameterValue=${ProjectId}
```

## Testing

For information on running the tests, see [here](./test/README.md).

## For More Information

See the MongoDB Atlas API [X.509 Authentication for Database Users Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/X.509-Authentication-for-Database-Users)
