# MongoDB::Atlas::OrgServiceAccount

## Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Currently, no CDK constructors or other components depend on this resource.

## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- MongoDB Atlas Organization (OrgId from environment variable `MONGODB_ATLAS_ORG_ID`)

**Note**: Service Account is an organization-level resource. No project or cluster resources are required.

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).

### Success criteria when testing the resource
1. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.
2. **Create Operation**:
   - Service account is created successfully in Atlas
   - Secret is returned in the response (writeOnly property)
   - ClientId is returned (read-only property)
   - All required fields are populated correctly
3. **Read Operation**:
   - Service account details are retrieved correctly
   - Secret is NOT returned (masked - writeOnly property)
   - MaskedSecretValue is shown instead
   - All read-only fields are populated
4. **Update Operation**:
   - Name, Description, and Roles can be updated
   - Secret is NOT returned in update response (masked)
   - Changes are reflected in Atlas UI
5. **Delete Operation**:
   - Service account is deleted successfully
   - Resource is removed from Atlas
6. **List Operation**:
   - All service accounts for the organization are listed
   - Secrets are masked in list response
   - Primary identifier fields are set correctly

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Service-Accounts)

## Running requests locally

To locally invoke requests, the AWS `sam local` and `cfn invoke` tools can be used:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/org-service-account
export MONGODB_ATLAS_ORG_ID="your-org-id"
./test/cfn-test-create-inputs.sh cfn-test-service-account > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke --function-name TestEntrypoint resource CREATE test.request.json 
cfn invoke --function-name TestEntrypoint resource DELETE test.request.json 
cd -
```

