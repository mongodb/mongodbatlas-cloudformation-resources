# Project IP Access List 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- Atlas Basic L3 CDK constructor
- Encryption at rest L3 CDK constructor
- Atlas Quickstart


## CFN resource type used
- MongoDB::Atlas::ProjectIpAccessList

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Project Access List CFN resource](../../../../cfn-resources/project-ip-access-list/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- A new entry should be added to the "Network Access" page:
![image](https://user-images.githubusercontent.com/5663078/227484402-9189af3d-a3f0-4bde-a288-9ee847e6eeab.png)
## Important Links
- [API Documentation](https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/openapi-docs-test/reference/api-resources-spec/#operation/createProjectIpAccessList)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/ip-access-list/)


## Important Links
- [API Documentation](https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/openapi-docs-test/reference/api-resources-spec/#operation/createProjectIpAccessList)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/ip-access-list/)