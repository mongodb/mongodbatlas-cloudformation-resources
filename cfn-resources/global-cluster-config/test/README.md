# MongoDB::Atlas::GlobalClusterConfig

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Global cluster configuration  L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
These resources are created as part of `cfn-testing-helper.sh`
- Atlas Project
- Atlas Cluster (at least M30)

## Manual QA
### Prerequisites Steps:
1. In your Atlas Cluster, ensure `Global Cluster Configuration` is enabled (found under `Edit Configuration` setting for the cluster)

### Steps to test:
Please follow the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- Custom Zone Mappings and ManagedNamespaces should be configured for the global cluster as specified in the template. 

This can be validated via GET API call at URL:
  `https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/<ATLAS_PROJECT_ID>/clusters/<ATLAS_CLUSTER_NAME>/globalWrites`

![image](https://user-images.githubusercontent.com/122359335/229160264-92715616-656e-4e7c-bd33-b6241041f9ae.png)

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Global-Clusters)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/global-clusters/)
