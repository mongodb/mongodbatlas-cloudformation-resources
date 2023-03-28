# Maintenance Window

## CFN resource type used
- MongoDB::Atlas::MaintenanceWindow

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Maintenance Window CFN resource](../../../../cfn-resources/maintenance-window/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Atlas Project should show configured maintenance window correctly under Project Settings:

![image](https://user-images.githubusercontent.com/122359335/227540482-6f021ea1-7b7e-4fbf-b883-1d9e0e2eea9a.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Maintenance-Windows)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/cluster-maintenance-window/)
