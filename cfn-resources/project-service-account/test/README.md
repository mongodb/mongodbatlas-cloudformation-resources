# MongoDB::Atlas::ProjectServiceAccount

## Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Currently, no CDK constructors or other components depend on this resource.

## Prerequisites 
### Resources needed to run the manual QA
Resources are created automatically as part of the test scripts:

- MongoDB Atlas Organization (from your Atlas CLI configuration)
- MongoDB Atlas Project (created automatically by `cfn-test-create.sh`)

**Note**: The test scripts will automatically create and manage the project for testing.

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).

### Success criteria when testing the resource
1. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.
2. **Create Operation**:
   - Project service account is created successfully in Atlas
   - Secret is returned in the response (writeOnly property)
   - ClientId is returned (read-only property)
   - Service account is assigned to the specified project
   - All required fields are populated correctly
3. **Read Operation**:
   - Project service account details are retrieved correctly
   - Secret is NOT returned (masked - writeOnly property)
   - MaskedSecretValue is shown instead
   - All read-only fields are populated
4. **Update Operation**:
   - Name, Description, and Roles can be updated
   - Secret is NOT returned in update response (masked)
   - SecretExpiresAfterHours CANNOT be updated (createOnly property)
   - Changes are reflected in Atlas UI
5. **Delete Operation**:
   - Project service account is deleted successfully
   - Service account is removed from the project
6. **List Operation**:
   - All service accounts for the project are listed
   - Secrets are masked in list response
   - Primary identifier fields are set correctly

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Service-Accounts/operation/createGroupServiceAccount)
- [Project Service Accounts Guide](https://www.mongodb.com/docs/atlas/api/service-accounts/)

## Running requests locally

To locally invoke requests, the AWS `sam local` and `cfn invoke` tools can be used:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project-service-account
./test/cfn-test-create-inputs.sh cfn-test-project-name > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke --function-name TestEntrypoint resource CREATE test.request.json 
cfn invoke --function-name TestEntrypoint resource DELETE test.request.json 
cd -
```

