# MongoDB::Atlas::StreamPrivatelinkEndpoint

## Impact

This resource is used by:
- CloudFormation templates for Stream Processing infrastructure
- CDK constructors (if applicable)
- Infrastructure as Code workflows for Atlas Stream Processing

## Prerequisites

### Resources needed to run the manual QA

All resources are created as part of `cfn-testing-helper.sh`:

- Atlas Project
- **Note**: Stream Processing Workspace is NOT required for Private Link Endpoints. The endpoint is created at the project level.

### Prerequisites for S3 Vendor:

- ✅ **Automatically created**: The `cfn-test-create-inputs.sh` script automatically creates an S3 bucket for testing
- AWS account with appropriate permissions for PrivateLink and S3 bucket creation
- The bucket is created with versioning and encryption enabled
- **Note**: Remember to delete the test bucket after testing: `aws s3 rb s3://<bucket-name> --force`

### Note on Other Vendors (MSK, CONFLUENT):

MSK and Confluent implementations are complex and time-consuming. They are available as examples in the `examples/` directory but are not included in the automated test input generation. For testing with MSK or Confluent, please refer to the example templates.

## Manual QA

Please follow the steps in [TESTING.md](../../../TESTING.md).

### Prerequisite steps:

**Recommended approach (using Makefile):**

```bash
# Create test resources (automatically generates project name and creates S3 bucket)
make create-test-resources

# Optional: Set a specific region (defaults to eu-west-1)
export AWS_REGION=us-east-1
make create-test-resources

# Note: The ServiceEndpointId format is: com.amazonaws.<region>.s3
# Example for eu-west-1: com.amazonaws.eu-west-1.s3
# The script automatically generates this based on the region.
```

**Alternative approach (direct script call):**

```bash
# If you need to specify a project name manually:
./test/cfn-test-create-inputs.sh <project_name>

# Optional: Set a specific region (defaults to eu-west-1)
export AWS_REGION=us-east-1
./test/cfn-test-create-inputs.sh <project_name>
```

### Steps to test:

1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources.
2. Generate test inputs:

   ```bash
   # Recommended: Use Makefile target
   make create-test-resources

   # Or call script directly with project name:
   # ./test/cfn-test-create-inputs.sh <project_name>
   ```

3. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.
4. Once the template with required parameters is used to create and delete a stack successfully, validate that success criteria is met.
5. Clean up test resources:

   ```bash
   # Recommended: Use Makefile target
   make delete-test-resources

   # Or call script directly:
   # ./test/cfn-test-delete-inputs.sh
   ```

### Success criteria when testing the resource

1. The Private Link Endpoint must be created successfully in Atlas:

   - Navigate to **Network Access** → **Stream Processing Private Link Endpoints** in Atlas UI
   - Verify the endpoint appears with vendor type "S3"
   - Verify the State transitions to "AVAILABLE" (may take a few minutes)

2. Verify S3 configuration:

   - Verify `ServiceEndpointId` is in format: `com.amazonaws.<region>.s3`
   - Verify `Region` matches the S3 bucket region

3. The AWS Interface Endpoint should be created (for MSK and CONFLUENT vendors):

   - Check AWS VPC Console → Endpoints
   - Verify the interface endpoint is created and in "Available" state
   - Note the `InterfaceEndpointId` and `InterfaceEndpointName` returned by the resource
   - **Note**: For S3 vendors, `InterfaceEndpointId` and `InterfaceEndpointName` are expected to be null, as S3 uses gateway endpoints rather than interface endpoints

4. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links

- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createprivatelinkconnection)
- [Stream Processing Private Link Documentation](https://www.mongodb.com/docs/atlas/atlas-stream-processing/manage-connection-registry/)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```

then in another shell:

```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/stream-privatelink-endpoint

# Generate test inputs (using Makefile or direct script call)
make create-test-resources
# Or: ./test/cfn-test-create-inputs.sh <project_name> > test.request.json

# Use the generated input file
cat ./inputs/inputs_1_create.json > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json
cfn invoke resource DELETE test.request.json

# Clean up
make delete-test-resources
cd -
```

Both CREATE & DELETE tests must pass.

## Notes

- **Long-running operations**: Private Link Endpoint creation can take several minutes. The resource will wait for the state to transition to "AVAILABLE".
- **AWS-only**: This resource currently supports AWS only. Azure and GCP support may be added in the future.
- **S3-only testing**: The test input generation script currently supports S3 only. For MSK and Confluent examples, see the `examples/` directory.
- **Create-only properties**: Most properties (ProjectId, ProviderName, Vendor, Region, ServiceEndpointId) are create-only and cannot be updated. To change these, you must delete and recreate the resource.
- **State monitoring**: The resource automatically monitors the endpoint state and waits for it to become "AVAILABLE" before completing creation.
