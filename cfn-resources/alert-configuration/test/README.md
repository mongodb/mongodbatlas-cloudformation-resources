# Alert Configurations 

## Impact 
 - Alert Configuration L1 CDK constructor

## Manual QA

#### Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

### Steps

#### Prerequisites
 - Configure Atlas default profile on your machine as the scripts use Atlas CLI.
 - Ensure aws and cfn CLI are installed and configured correctly for your AWS account.
 - Configure an Atlas profile secret in AWS account with your API keys. Refer README for more information on how to do this.
 - Have Docker running on your machine.

#### Contract tests
   - Run the following command to build the resource and start AWS SAM
```bash
   cd cfn-resources/alert-configuration && make build && sam local start-lambda --skip-pull-image
```

   - Open a new terminal session
   - Generate the cfn test inputs file and required Atlas resources:
```bash
   ./../../cfn-testing-helper.sh <resource-folder>
```
   - Run cfn test
```bash
    cd cfn-resources/alert-configuration 
    cfn test --function-name TestEntrypoint --verbose 
```
#### Testing the resource in a CFN stack
- Publish the resource to the AWS Private registry
```bash
cfn submit --set-default
```
#### Create a stack
- Ensure all steps above are complete
- Getting test parameters: 
  - Option 1: Re-use params from cfn-resources/alert-configuration/inputs/inputs_1_create.template.json generated as part of prerequisites
  - Option 2: Run again  `cfn-testing-helper` to create new parameters.  This will create some required resources for you in your configured Atlas account
- Update template in [examples/alert-configuration](../../../examples/alert-configuration/alert-configuration.json) with your changes (if any) AND required params from last step
- Login to AWS account -> CloudFormation -> click on Create Stack dropdown and select (with new resources)
- Upload the updated template: You can either upload the updated file OR select Create template in Designer
- Add a Stack name on Specify stack details page and provide any required parameters for your resource, if any
- Click on Next and follow UI prompts to configure stack options if needed. Finally click on Submit. Your stack should now be created and resource creation IN_PROGRESS

#### Update the stack
- In AWS CloudFormation, navigate to your stack and click on Update. You can use the current template or edit in designer
- Update any parameters or outputs for your template
- Follow UI prompts and Submit

#### Delete the stack
- In AWS CloudFormation, navigate to your stack and click on Delete

### Success criteria when testing the resource
- Alert Settings for the respective Project in Atlas should be correctly configured:
![image](https://user-images.githubusercontent.com/5663078/226870968-9ef8ae46-b0cf-462b-ac62-7229d2d79ac0.png)
- Stack Outputs should show required data correctly based on the template
- Create: Stack should complete successfully and resource should be created and correctly configured in Atlas account as per template
- Update: Stack should complete successfully and resource should be updated and correctly configured in Atlas account as per template
- Delete: Stack should complete successfully and resource should be deleted from your Atlas account


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)