# Changelog

## (2023-10-20)

**(BREAKING CHANGE) Resource Structure modification**

We have made significant improvements to our Private Endpoint resource structure to enhance its usability and maintainability.

Previously, the MongoDB::Atlas::PrivateEndpoint resource was responsible for handling three distinct tasks:

Creating and deleting an Atlas Private Endpoint service.
Adding and removing an Atlas Private Endpoint to/from a Private Endpoint Service.
Creating and deleting an AWS private endpoint.
This approach had several drawbacks:

The complexity of having a single resource manage three different functions made it challenging to maintain.
Users found it limiting to manage an AWS Private Endpoint within an Atlas resource, which restricted certain configurations.
When a resource was created, there was no visibility into the progress of each component.
To address these issues, we have introduced a new resource, MongoDB::Atlas::PrivateEndpointService. This resource is specifically designed to create an Atlas Private Endpoint Service without any preconfiguration, allowing for subsequent Private Endpoint configuration with AWS or any other provider.

Additionally, we have defined a new MongoDB::Atlas::PrivateEndpointAWS resource to focus solely on adding a Private Endpoint to the Service.

	- MongoDB::Atlas::PrivateEndpointService: new, resource responsible for creating a Private Endpoint Service, unconfigured, and then using it to configure a Private Endpoint with AWS or any other provider
	- MongoDB::Atlas::PrivateEndpointAWS: the current resource is modified, it is responsible for adding a privateEdnpoint to the Service

As a result, users will now configure a Private Endpoint by defining these resources separately:

Stack:

Atlas Private Endpoint Service
AWS Private Endpoint
Atlas Private Endpoint AWS
