Cluster 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- AtlasBasic L3 CDK constructor
- Encryption at Rest L3 CDK constructor
- Atlas Quickstart
- Atlas Quickstart Fargate

## Uses CFN resource type
- MongoDB::Atlas::Cluster

This CFN resource must be active in your AWS account while using this constructor.


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
All resources are created as part of `cfn-testing-helper.sh`


## Manual QA
- Follow prerequisite steps for testing a CDK construct.
- Follow prerequisite steps for corresponding [Cluster CFN resource](../../../../cfn-resources/cluster/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).



### Success criteria when testing the resource
- A new Cluster should be added to the "Database Deployments" page:
![image](https://user-images.githubusercontent.com/5663078/227485960-fab8e1c9-b4df-41bb-8fbb-4895e37da2f1.png)

