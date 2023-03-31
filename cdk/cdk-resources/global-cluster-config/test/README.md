# @mongodbatlas-awscdk/global-cluster-config

## CFN resource type used
- MongoDB::Atlas::GlobalClusterConfig

This CFN resource must be active in your AWS account while using this constructor.

## Prerequisites 
### Resources needed to run the manual QA
This resource is created as part of `cfn-testing-helper.sh`
- Atlas Project
- Atlas Cluster (at least M30)

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [GlobalClusterConfig CFN resource](../../../../cfn-resources/global-cluster-config/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Custom Zone Mappings and ManagedNamespaces should be configured for the global cluster as specified in the template. This can be validated via API call:
![image](https://user-images.githubusercontent.com/122359335/229160264-92715616-656e-4e7c-bd33-b6241041f9ae.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Global-Clusters)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/global-clusters/)
