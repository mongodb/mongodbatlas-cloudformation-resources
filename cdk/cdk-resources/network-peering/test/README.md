# Network peering

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- Quickstart VPC peering


## CFN resource type used
- MongoDB::Atlas::NetworkPeering

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Network peering CFN resource](../../../../cfn-resources/network-peering/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should be able to see the network peering in the "Network Access" page:
![image](https://user-images.githubusercontent.com/5663078/227514067-123c7343-1066-4ba7-802a-03a73a810c78.png)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Network-Peering)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/reference/atlas-operator/ak8so-network-peering/)