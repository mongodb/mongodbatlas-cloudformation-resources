# MongoDB::Atlas::MaintenanceWindow

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Maintenance window L1 CDK constructor



## Prerequisites 
### Resources needed to manually QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
1. Atlas Project should show configured maintenance window correctly under Project Settings:

![image](https://user-images.githubusercontent.com/122359335/227540482-6f021ea1-7b7e-4fbf-b883-1d9e0e2eea9a.png)

2. General [CFN resource success criteria](../../../TESTING.md.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Maintenance-Windows)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/cluster-maintenance-window/)