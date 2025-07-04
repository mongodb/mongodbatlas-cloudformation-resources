# MongoDB::Atlas::SearchIndex

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Search index L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas project
- Atlas Cluster with sample data
- Database and collection (using sample_airbnb)



## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md). Make sure to use all the templates in [example/search-index](../../../examples/search-index/).


### Success criteria when testing the resource
1. Search Index should be created in `ACTIVE` state for the database:

![image](https://user-images.githubusercontent.com/122359335/227660157-b51c16cd-7a87-40b6-bdd9-9bbf44efeeec.png)

2. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-atlas-search)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/atlas-search/create-index/)
