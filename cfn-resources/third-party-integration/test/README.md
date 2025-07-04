# MongoDB::Atlas::ThirdPartyIntegration

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Third-party integration L1 CDK constructor
 - Third-party integration L2 CDK constructor



## Prerequisites 
### Resources needed to manually QA
- Atlas Project
All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Atlas Project should show correctly configured integration with chosen third-party service, for example Datadog:

![image](https://user-images.githubusercontent.com/122359335/227501805-7eee80cc-12a0-4a80-8400-09a283655187.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createthirdpartyintegration)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/third-party-service-integrations/)

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