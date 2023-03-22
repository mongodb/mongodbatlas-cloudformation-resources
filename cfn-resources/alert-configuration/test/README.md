# Alert Configurations 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Alert Configuration L1 CDK constructor


## Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [MANUALQA.md](../../../MANUALQA.md).


### Success criteria when testing the resource
- Alert Settings for the respective Project in Atlas should be correctly configured:
![image](https://user-images.githubusercontent.com/5663078/226870968-9ef8ae46-b0cf-462b-ac62-7229d2d79ac0.png)

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)