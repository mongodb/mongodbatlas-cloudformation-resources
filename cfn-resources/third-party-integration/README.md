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
Examples aws cloudformation template is available here [example template.](../../examples/thirdpartyintegrations)


```bash
#Configure you AWS Credentials to create Cloudformation Stack
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_REGION=""
export AWS_DEFAULT_REGION=""

#Command to deploy the sample thirdpartyintegration stack (Before this step "cfn submit" should have been executed successfully)
./examples/thirdpartyintegration/deploy.sh
```

| Integrations    | Status | Reference links                                                                                                                                                                                                                                                 |
|-----------------|--------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| NEW_RELIC       | EOL    | [From MongoDB](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/?_ga=2.141767858.1639178218.1667927805-1433452924.1667927805), [From NewRelic](https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267) |
| FLOWDOCK        | EOL    | [From MongoDB](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/?_ga=2.141767858.1639178218.1667927805-1433452924.1667927805)                                                                                                       |
| DATA_DOG        | GA     | [Datadog template.](../../examples/thirdpartyintegrations/datadog.json)                                                                                                                                                                                         |
| PROMETHEUS      | GA     | [Prometheus template.](../../examples/thirdpartyintegrations/prometheus.json)                                                                                                                                                                                   |
| MICROSOFT_TEAMS | GA     | [Microsoft Teams CFN template.](../../examples/thirdpartyintegrations/microsoftteams.json)                                                                                                                                                                      |
| OPS_GENIE       | GA     | [OpsGenie CFN Template.](../../examples/thirdpartyintegrations/opsgenie.json)                                                                                                                                                                                   |
| PAGERDUTY       | GA     | [Pagerduty CFN template.](../../examples/thirdpartyintegrations/pagerduty.json)                                                                                                                                                                                 |
| WEBHOOK         | GA     | [Webhook CFN template.](../../examples/thirdpartyintegrations/webhook.json)                                                                                                                                                                                     |



For more information see: MongoDB Atlas API [ThirdPartyIntegrations Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Service-Integrations) Documentation.



