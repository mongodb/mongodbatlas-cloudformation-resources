# Network Container

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Network container L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should be able to see the container created when running the following Atlas CLI command:
```bash
atlas networking container ls --projectId 641d818e36e2eb47d038f8c2
[
  {
    "atlasCidrBlock": "10.8.2.0/24",
    "id": "<network-container-id>",
    "providerName": "AWS",
    "provisioned": false,
    "regionName": "US_EAST_1"
  }
]
```



## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-network-peering)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/reference/atlas-operator/ak8so-network-peering/)