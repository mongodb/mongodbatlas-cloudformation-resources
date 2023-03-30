# @mongodbatlas-awscdk/atlas-basic

## CFN resource type used
- MongoDB::Atlas::Project
- MongoDB::Atlas::Cluster
- MongoDB::Atlas::DatabaseUser
- MongoDB::Atlas::ProjectIpAccessList

These CFN resources must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Ensure all underlying resources are configured correctly as specified (Atlas Project, Cluster, Network IPAccessList and DatabaseUser):

![image](https://user-images.githubusercontent.com/122359335/228263898-9d9c3a8a-ddc5-4cf6-9f7e-256b7c976b54.png)

![image](https://user-images.githubusercontent.com/122359335/228263913-8fbad8e7-7a60-4eae-aac5-9bccc2dc9242.png)

![image](https://user-images.githubusercontent.com/122359335/228263944-f4b35480-7cc3-4b6d-afbb-cd8975ecab19.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.
