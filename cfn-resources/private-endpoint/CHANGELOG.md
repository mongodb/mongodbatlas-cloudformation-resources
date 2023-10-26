# Changelog

## (2023-10-20) **(BREAKING CHANGE) Deprecation of MongoDB::Atlas::PrivateEndpoint**

### Why is this resource deprecated?

The MongoDB::Atlas::PrivateEndpoint resource, which was responsible for several tasks, is now marked as DEPRECATED, and will no longer receive support after 1 January 2024.
Users are encouraged to use the following new resources instead:

1. **MongoDB::Atlas::PrivateEndpointService**: This resource is responsible for creating an unconfigured Atlas Private Endpoint Service. 
Subsequent Private Endpoint configuration with AWS or any other provider can be done using this service.
2. **MongoDB::Atlas::PrivateEndpointAWS**: The existing resource has been modified to focus solely on adding a Private Endpoint to the Service with an AWS provider.

### Drawbacks of the Previous Approach
The deprecation of the MongoDB::Atlas::PrivateEndpoint resource is due to the following drawbacks in the previous approach:

- Complexity: Managing three different functions within a single resource made it challenging to maintain.
- Limitations: Users found it limiting to manage an AWS Private Endpoint within an Atlas resource, restricting certain configurations.
- Progress Visibility: When a resource was created, there was no visibility into the progress of each component.

### New Resource Structure
With the introduction of the new resources, users can now configure a Private Endpoint by defining these resources separately:

**Stack:**

- MongoDB::Atlas::PrivateEndpointService: A new resource responsible for creating an unconfigured Private Endpoint Service, which can be used to configure a Private Endpoint with AWS or any other provider.
- AWS::EC2::VPCEndpoint: The existing AWS resource.
- MongoDB::Atlas::PrivateEndpointAWS: The current resource, which has been modified to add a Private Endpoint to the Service.

For detailed upgrade instructions, please refer to the [V2 Upgrade Guide](upgradeguidev2/V2-UpgradeGuide.md).