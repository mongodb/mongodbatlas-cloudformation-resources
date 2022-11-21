# MongoDB::Atlas::ThirdPartyIntegration

## Description
Returns, adds, edits, and removes third-party service integration configurations. MongoDB Cloud sends alerts to each third-party service that you configure.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

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

cd ${repo_root}/cfn-resources/thirdpartyintegration
./test/thirdpartyintegration.create-sample-cfn-request.sh > test.request.json 
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

You can use the project.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-project-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
#Configure you AWS Credentials to create Cloudformation Stack
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_REGION=""
export AWS_DEFAULT_REGION=""

#Command to deploy the sample thirdpartyintegration stack (Before this step "cfn submit" should have been executed successfully)
aws cloudformation deploy --stack-name atlas-thirdpartyintegration-test --template-file ./test/thirdpartyintegration.sample-template.yaml --no-fail-on-empty-changeset --parameter-overrides PublicKey=$MCLI_PUBLIC_API_KEY Privatekey=$MCLI_PRIVATE_API_KEY ProjectId=$MCLI_PROJECT_ID
```

| Integrations           | Status                                             | Reference links                                                                                                                                                                                                                                                 |
|------------------------|----------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| NEW_RELIC              | EOL                                                | [From MongoDB](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/?_ga=2.141767858.1639178218.1667927805-1433452924.1667927805), [From NewRelic](https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267) |
| FLOWDOCK | EOL                                                | [From MongoDB](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/?_ga=2.141767858.1639178218.1667927805-1433452924.1667927805)                                                                                                       |
| VICTOR_OPS                | ![Build](https://img.shields.io/badge/Beta-yellow) | [Jira Ticket](https://jira.mongodb.org/browse/HELP-39527)                                                                                                                                                                                                       |


For more information see: MongoDB Atlas API ThirdPartyIntegrations [Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Service-Integrations) Documentation.



