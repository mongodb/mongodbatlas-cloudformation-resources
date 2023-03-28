# Datalakes


## CFN resource type used
- MongoDB::Atlas::DataLakes

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [DataLakes CFN resource](../../../../cfn-resources/datalakes/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should see your data federation in the "Data Federation" page:
![image](https://user-images.githubusercontent.com/5663078/227923171-b9aa0067-d8a0-41b3-94c5-9e5ce5ea222c.png)


## Important Links
- [API Documentation](https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/openapi-docs-test/reference/api-resources-spec/#tag/Data-Federationt)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/data-federation/config/config-adl-datasets/)