# Search Index

## CFN resource type used
- MongoDB::Atlas::SearchIndex

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [SearchIndex CFN resource](../../../../cfn-resources/search-index/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Search Index should be created in `ACTIVE` state for the database:

![image](https://user-images.githubusercontent.com/122359335/227660157-b51c16cd-7a87-40b6-bdd9-9bbf44efeeec.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Atlas-Search)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/atlas-search/create-index/)