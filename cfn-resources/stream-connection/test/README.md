# MongoDB::Atlas::StreamConnection

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Stream Connection L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas Project
- Cluster
- Atlas Stream Instance

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. A Stream Connection should be created in the specified test project for the specified Atlas Stream instance.:

  ![image](https://github.com/mongodb/mongodbatlas-cloudformation-resources/assets/122359335/a1f1ef75-0df3-4c53-8e76-69f63f3ad86e)

2. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-streams)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/atlas-sp/overview/)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
cd ${repo_root}/cfn-resources/stream-connection
cfn invoke resource CREATE stream-connection-sample-cfn-request.json 
cfn invoke resource DELETE stream-connection-sample-cfn-request.json
cd -
```

Both CREATE & DELETE tests must pass.