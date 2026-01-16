# Testing
This file contains the steps to follow to test any changes to the CFN resources.


## Manual QA

### Prerequisites
 - Install [AtlasCLI](https://www.mongodb.com/docs/atlas/cli/stable/install-atlas-cli/) and [configure the default profile](https://www.mongodb.com/docs/atlas/cli/stable/connect-atlas-cli/#select-a-connection-method) on your machine as the scripts use Atlas CLI
 - Ensure [AWS CLI](https://aws.amazon.com/cli/), [CFN CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html), [SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html) are installed and configured correctly for your AWS account
 - Configure an Atlas profile secret in AWS account with your API keys. Refer to [README](README.md) for more information on how to do this
 - Have Docker running on your machine
 - Refer to the prerequisites for your resource in the `cfn-resources/[resource-folder]/test`


### Steps

#### Contract tests
   - Run the following command to build the resource and start AWS SAM
```bash
   cd cfn-resources/[resource-folder] && make build && sam local start-lambda --skip-pull-image
```

   - Open a new terminal session
   - Generate the cfn test inputs file and required Atlas resources. Each resource may require specific environment variables to be defined, you will find this under `cfn-resources/[resource-folder]/test/cfn-test-create-inputs.sh`. Generated input files also assume you have a `default` profile defined which will be used.
```bash
   ./../../cfn-testing-helper.sh <resource-folder>
```

   > **Note:** When creating or modifying test input files, non-string values (booleans, integers) must be defined as strings. For example, use `"Enabled": "true"` instead of `"Enabled": true`.

   - Run cfn test
```bash
    cd cfn-resources/[resource-folder]
    cfn test --function-name TestEntrypoint --verbose 
```

#### Testing the resource in a CFN stack

##### Create a stack
- Ensure all steps above are complete
- Publish the resource to the AWS Private registry
  ```bash
  cfn submit --set-default
  ```
- Getting test parameters: 
  - Option 1 [Recommended]: Re-use params from `cfn-resources/[resource-folder]/inputs/inputs_1_create.template.json` generated as part of prerequisites
  - Option 2: Run again  `cfn-testing-helper` to create new parameters.  This will create some required resources for you in your configured Atlas account
- Update template in `examples/[resource-folder]/[resource-name].json` with your changes (if any) AND required params from last step
- Login to AWS account -> CloudFormation -> click on Create Stack dropdown and select (with new resources)
- Upload the updated template: You can either upload the updated file OR select Create template in Designer
- Add a Stack name on Specify stack details page and provide any required parameters for your resource, if any
- Click on Next and follow UI prompts to configure stack options if needed. Finally, click on Submit. Your stack should now be created and resource creation `IN_PROGRESS`

##### Update the stack
- In AWS CloudFormation, navigate to your stack and click on Update. You can use the current template or edit in designer
- Update any parameters or outputs for your template
- Follow UI prompts and Submit

##### Delete the stack
- In AWS CloudFormation, navigate to your stack and click on Delete

### Success criteria when testing the resource
- Refer to the success criteria for your resource first in the `cfn-resources/[resource-folder]/test`.
- Stack Outputs should show required data correctly based on the template
- Create: Stack should complete successfully and resource should be created and correctly configured in Atlas account as per template
- Update: Stack should complete successfully and resource should be updated and correctly configured in Atlas account as per template
- Delete: Stack should complete successfully and resource should be deleted from your Atlas account
