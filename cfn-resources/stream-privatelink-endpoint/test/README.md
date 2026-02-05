# MongoDB::Atlas::StreamPrivatelinkEndpoint

## Prerequisites

### Resources needed to run the manual QA

All resources are created as part of `cfn-test-create-inputs.sh`:

- Atlas Project
- **Note**: Stream Processing Workspace is NOT required for Private Link Endpoints. The endpoint is created at the project level.

### Prerequisites for S3 Vendor:

- ✅ **No S3 bucket required**: The `cfn-test-create-inputs.sh` script uses the standard AWS S3 service endpoint format (`com.amazonaws.<region>.s3`), which is a regional AWS service endpoint that's always available
- AWS account with appropriate permissions for PrivateLink
- **Note**: The S3 service endpoint ID format doesn't require an actual S3 bucket. The endpoint connects to the S3 service itself, not a specific bucket

### Supported Vendors:

**CFN contract tests only support S3 vendor.** Confluent and MSK vendors are not included in automated CFN testing due to complexity and external dependencies.

## Manual QA

Please follow the steps in [TESTING.md](../../../TESTING.md).

### Success criteria when testing the resource

1. The Private Link Endpoint must be created successfully in Atlas:

   - Navigate to **Network Access** → **Stream Processing Private Link Endpoints** in Atlas UI
   - Verify the endpoint appears with vendor type "S3"
   - Verify the State transitions to "AVAILABLE" (may take a few minutes)

2. Verify S3 configuration:

   - Verify `ServiceEndpointId` is in format: `com.amazonaws.<region>.s3`
   - Verify `Region` matches the AWS region where the endpoint is created
   - **Note**: For S3 vendor, `InterfaceEndpointId` and `InterfaceEndpointName` are expected to be null, as S3 uses gateway endpoints rather than interface endpoints

3. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.

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
./test/cfn-test-create-inputs.sh <project_name> > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json
cfn invoke resource DELETE test.request.json
cd -
```

Both CREATE & DELETE tests must pass.
