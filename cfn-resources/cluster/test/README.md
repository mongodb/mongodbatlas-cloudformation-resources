# Cluster 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Cluster L1 CDK constructor
- AtlasBasic L3 CDK constructor
- Encryption at Rest L3 CDK constructor
- Atlas Quickstart
- Atlas Quickstart Fargate




## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- A new Cluster should be added to the "Database Deployments" page:
![image](https://user-images.githubusercontent.com/5663078/227485960-fab8e1c9-b4df-41bb-8fbb-4895e37da2f1.png)
## Important Links
- [API Documentation](https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/openapi-docs-test/reference/api-resources-spec/#tag/Global-Clusters)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/manage-clusters/)
