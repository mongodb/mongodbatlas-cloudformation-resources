# Alert Configurations

## Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Alert Configuration L1 CDK constructor


## Prerequisites
### Resources needed to manually QA:
- Atlas Project

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Alert Settings for the respective Project in Atlas should be correctly configured:
  ![image](https://user-images.githubusercontent.com/5663078/226870968-9ef8ae46-b0cf-462b-ac62-7229d2d79ac0.png)
2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)