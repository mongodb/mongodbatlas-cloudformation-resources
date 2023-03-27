## MongoDB::Atlas::ServerlessInstance

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Serverless L1 CDK constructor


### Resources (and parameters for local tests) needed to manually QA:
- Atlas Organization (ATLAS_ORG_ID)
- Atlas Project (created as part of `cfn-testing-helper.sh`)

## Manual QA:

### Prerequisite steps:
1. Create an Atlas Organization if you don’t already have one and note the OrgId from URL (https://cloud.mongodb.com/v2#/org/<ATLAS_ORG_ID>/projects)
2. Export ATLAS_ORG_ID environment variable with OrgId from #1.

### Steps to test:
1. Ensure prerequisites above for this resource and general [prerequisites](../../../TESTING.md.md#prerequisites) are complete.
2. Follow [general steps](../../../TESTING.md.md#steps) to test a CFN resource.


### Success criteria when testing the resource
1. A serverless instance should show correctly configured for the database:

![image](https://user-images.githubusercontent.com/122359335/227501805-7eee80cc-12a0-4a80-8400-09a283655187.png)

2. General [CFN resource success criteria](../../../TESTING.md.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Serverless-Instances)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/create-serverless-instance/)