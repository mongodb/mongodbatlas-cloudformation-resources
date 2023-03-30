# @mongodbatlas-awscdk/atlas-basic-private-endpoint

## CFN resource type used
- MongoDB::Atlas::Project
- MongoDB::Atlas::Cluster
- MongoDB::Atlas::DatabaseUser
- MongoDB::Atlas::ProjectIpAccessList
- MongoDB::Atlas::PrivateEndpoint

These CFN resources must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Ensure all underlying resources are configured correctly as specified (Atlas Project, Cluster, Network IPAccessList, DatabaseUser, PrivateEndpoint):

![image](https://user-images.githubusercontent.com/122359335/228276480-0cadd908-e674-4d26-9e17-c9203d1e5072.png)

![image](https://user-images.githubusercontent.com/122359335/228276519-0fccfa24-84c0-4ff4-aebf-508009cc6db8.png)

![image](https://user-images.githubusercontent.com/122359335/228276560-3d2104e1-bed8-49de-bc4d-290e6ee4a5da.png)

![image](https://user-images.githubusercontent.com/122359335/228276604-116e6251-ac5d-4f91-aa6e-87a221ee0f15.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.
