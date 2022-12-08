# MongoDB::Atlas::PrivateEndPointRegionalMode

## Description
Returns, adds and removes regional mode setting for private endpoints.

## Attributes & Parameters
Please consult the [Resource Docs](https://github.com/PeerIslands/mongodbatlas-cloudformation-resources/blob/feature-private-endpoint-regional-mode/cfn-resources/private-endpoint-regional-mode/docs/README.md)

## Local Testing
The local tests are integrated with the AWS sam local and cfn invoke tooling features:

```
sam local start-lambda --skip-pull-image
```

then in another shell:
```bash
#https://www.mongodb.com/docs/mongocli/stable/configure/environment-variables/
#Set the public API key for commands that interact with your MongoDB service.
export MCLI_PUBLIC_API_KEY = ""
#Set the private API key for commands that interact with your MongoDB service.
export MCLI_PRIVATE_API_KEY=""
#Sets the project ID for commands that require the --projectId option.
export MCLI_PROJECT_ID = ""

cd ${repo_root}/cfn-resources/private-endpoint-regional-mode
./test/private-endpoint-regional-mode.create-sample-cfn-request.sh > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json
cfn invoke DELETE test.request.json
```

Both CREATE & DELETE tests must pass.

## Installation
```
TAGS=logging make
cfn submit --verbose --set-default
```

## Usage
Examples aws cloudformation template is available here example template.

```bash
#Configure you AWS Credentials to create Cloudformation Stack
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_REGION=""
export AWS_DEFAULT_REGION=""

#Command to deploy the sample PrivateEndPointRegionalMode stack (Before this step "cfn submit" should have been executed successfully)
./examples/private-endpoint-regional-mode/Deploy.sh
```

| Operation | Flag       | Reference links                                                                                                                                                                                                                                                 |
|-----------|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| PATCH     | true/false | [From MongoDB](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Private-Endpoint-Services/operation/toggleRegionalizedPrivateEndpointStatus) |
| READ      | true/false | [From MongoDB](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Private-Endpoint-Services/operation/returnRegionalizedPrivateEndpointStatus) |                                   

For more information see: MongoDB Atlas API [PrivateEndpointRegionalizedMode Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Private-Endpoint-Services) Documentation.