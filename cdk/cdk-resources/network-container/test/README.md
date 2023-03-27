# Network container

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- Quickstart VPC peering


## CFN resource type used
- MongoDB::Atlas::NetworkContainer

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Project Access List CFN resource](../../../../cfn-resources/network-peering/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


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