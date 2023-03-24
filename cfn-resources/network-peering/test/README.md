# Network Peering

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Network peering L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Atlas Container

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- You should be able to see the network peering in the "Network Access" page:
![image](https://user-images.githubusercontent.com/5663078/227514067-123c7343-1066-4ba7-802a-03a73a810c78.png)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/reference/atlas-operator/ak8so-network-peering/)