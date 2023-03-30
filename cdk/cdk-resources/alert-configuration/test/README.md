# Alert Configurations

## CFN resource type used
- MongoDB::Atlas::AlertConfiguration

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Alert Configuration CFN resource](../../../../cfn-resources/alert-configuration/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Alert Settings for the respective Project in Atlas should be correctly configured:
  ![image](https://user-images.githubusercontent.com/5663078/226870968-9ef8ae46-b0cf-462b-ac62-7229d2d79ac0.png)
2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)