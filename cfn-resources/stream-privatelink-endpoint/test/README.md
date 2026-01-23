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

### Prerequisites for Confluent Cloud Vendor:

- ✅ **Pre-existing Confluent Cloud resources required**: Similar to Terraform, the test script expects pre-existing Confluent Cloud network and private link access resources.
- Confluent Cloud account with a network configured for PrivateLink
- **⚠️ CRITICAL: VPC Endpoint Service Sharing**: The Confluent Cloud VPC endpoint service must be **shared and accepted** in the AWS account where CloudFormation tests run. The VPC endpoint service is in Confluent's AWS account and needs to be accessible from your test AWS account.
  - If the VPC endpoint service is not accessible, endpoint creation will fail with: `Failed to find VPC endpoint service with name <service-endpoint-id>`
  - To fix: Accept the VPC endpoint service connection request in your AWS account's VPC Endpoint Service Connections
- Environment variables (optional, defaults provided):
  - `CONFLUENT_CLOUD_REGION` (default: `us-east-1`)
  - `CONFLUENT_CLOUD_DNS_DOMAIN` (default: `dom4gllez7g.us-east-1.aws.confluent.cloud`)
  - `CONFLUENT_CLOUD_SERVICE_ENDPOINT_ID` (default: `com.amazonaws.vpce.us-east-1.vpce-svc-09f77bf9637bb0090`)
  - `CONFLUENT_CLOUD_DNS_SUB_DOMAIN` (default: comma-separated list of zonal subdomains)
- **Note**: The script uses default values from an existing Confluent Cloud network. To use different values, set the environment variables before running the script.

### Note on MSK Vendor:

MSK implementation is complex and time-consuming. It is available as an example in the `examples/` directory but is not included in the automated test input generation. For testing with MSK, please refer to the example templates.

## Manual QA

Please follow the steps in [TESTING.md](../../../TESTING.md).

### Success criteria when testing the resource

1. The Private Link Endpoint must be created successfully in Atlas:

   - Navigate to **Network Access** → **Stream Processing Private Link Endpoints** in Atlas UI
   - Verify the endpoint appears with vendor type "S3" or "CONFLUENT" (depending on your test)
   - Verify the State transitions to "AVAILABLE" (may take a few minutes)

2. Verify S3 configuration (for S3 vendor):

   - Verify `ServiceEndpointId` is in format: `com.amazonaws.<region>.s3`
   - Verify `Region` matches the AWS region where the endpoint is created

3. Verify Confluent Cloud configuration (for CONFLUENT vendor):

   - Verify `DnsDomain` matches your Confluent Cloud network DNS domain
   - Verify `ServiceEndpointId` matches your Confluent Cloud network VPC endpoint service name
   - Verify `Region` matches the AWS region of your Confluent Cloud network
   - Verify `DnsSubDomain` (if provided) matches your Confluent Cloud network zonal subdomains

4. The AWS Interface Endpoint should be created (for MSK and CONFLUENT vendors):

   - Check AWS VPC Console → Endpoints
   - Verify the interface endpoint is created and in "Available" state
   - Note the `InterfaceEndpointId` and `InterfaceEndpointName` returned by the resource
   - **Note**: For S3 vendors, `InterfaceEndpointId` and `InterfaceEndpointName` are expected to be null, as S3 uses gateway endpoints rather than interface endpoints

5. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.

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
