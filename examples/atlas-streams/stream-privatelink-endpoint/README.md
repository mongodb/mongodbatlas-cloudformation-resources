# Stream Processing Private Link Endpoint Examples

This directory contains CloudFormation example templates for creating MongoDB Atlas Stream Processing Private Link Endpoints.

## Overview

Stream Processing Private Link Endpoints enable secure, private connectivity between Atlas Stream Processing and streaming services (AWS MSK, Confluent Cloud, or AWS S3) over AWS PrivateLink. This resource supports AWS only.

## Prerequisites

1. **Atlas Project**: You must have an existing MongoDB Atlas project.
2. **AWS Account**: An AWS account with appropriate permissions for PrivateLink and VPC endpoint creation.
3. **Atlas Credentials**: Set up an AWS profile to securely give CloudFormation access to your Atlas credentials. For instructions, see [MongoDB Atlas API Keys Credential Management](/README.md#mongodb-atlas-api-keys-credential-management).
4. **AWS Credentials**: Configure AWS credentials with permissions to create VPC endpoints and manage PrivateLink connections.

## Example Templates

### S3 Private Link Endpoint

**File**: `stream-privatelink-endpoint-s3.json`

This template creates a Private Link Endpoint for AWS S3, enabling secure connectivity between Atlas Stream Processing and S3 buckets.

#### Parameters

- `Profile` (optional, default: "default"): Secret Manager Profile that contains the Atlas Programmatic keys
- `ProjectId` (required): Unique 24-hexadecimal digit string that identifies your Atlas project
- `Region` (optional, default: "eu-west-1"): AWS region where the S3 bucket is located (e.g., us-east-1, eu-west-1)
- `ServiceEndpointId` (optional, default: "com.amazonaws.eu-west-1.s3"): S3 service endpoint ID in the format `com.amazonaws.<region>.s3`

#### Deployment

```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-privatelink-endpoint/stream-privatelink-endpoint-s3.json \
  --stack-name atlas-stream-privatelink-endpoint-s3 \
  --parameter-overrides \
    ProjectId=<YOUR_PROJECT_ID> \
    Region=us-east-1 \
    ServiceEndpointId=com.amazonaws.us-east-1.s3 \
    Profile=default \
  --capabilities CAPABILITY_IAM \
  --region eu-west-1
```

**Important Notes**:

- The `--region` flag (eu-west-1) must match the region where the `MongoDB::Atlas::StreamPrivatelinkEndpoint` resource type is registered in your CloudFormation Private Registry.
- The `Region` parameter (us-east-1) should match where your S3 bucket is located.
- The `ServiceEndpointId` must match the `Region` parameter. For example:
  - `us-east-1` → `com.amazonaws.us-east-1.s3`
  - `eu-west-1` → `com.amazonaws.eu-west-1.s3`
  - `ap-southeast-1` → `com.amazonaws.ap-southeast-1.s3`

#### Template Validation

Before deploying, validate the template syntax:

```bash
aws cloudformation validate-template \
  --template-body file://examples/atlas-streams/stream-privatelink-endpoint/stream-privatelink-endpoint-s3.json
```

## Verification

### Using Atlas CLI

After deployment, verify the Private Link Endpoint was created:

```bash
# List all Private Link Endpoints for the project
atlas streams privatelink list --projectId <PROJECT_ID>

# Describe a specific Private Link Endpoint
atlas streams privatelink describe <ENDPOINT_ID> --projectId <PROJECT_ID>
```

**Expected Output Fields**:

- `id`: The Private Link connection ID
- `providerName`: Should be "AWS"
- `vendor`: Should be "S3"
- `region`: The AWS region
- `serviceEndpointId`: The S3 service endpoint ID
- `state`: Should be "AVAILABLE" when ready

### Using Atlas UI

1. Navigate to your Atlas project
2. Go to **Network Access** → **Stream Processing Private Link Endpoints**
3. Verify the endpoint appears with:
   - Vendor type: "S3"
   - State: "AVAILABLE" (may take a few minutes to transition)
   - Correct region and service endpoint ID

### Using AWS Console

1. Navigate to AWS VPC Console → **Endpoints**
2. Find the interface endpoint created by the resource
3. Verify it's in "Available" state
4. Note the endpoint ID and name match the values returned by the CloudFormation stack

## Important Notes

- **AWS-only**: This resource currently supports AWS only. Azure and GCP support may be added in the future.
- **Long-running operations**: Private Link Endpoint creation can take several minutes. The resource automatically waits for the state to transition to "AVAILABLE" before completing.
- **Create-only properties**: Most properties (ProjectId, ProviderName, Vendor, Region, ServiceEndpointId) are create-only and cannot be updated. To change these, you must delete and recreate the resource.
- **Primary identifier**: The resource ID is the combination of `ProjectId` and the connection ID returned by the Atlas API, formatted as `ProjectId|ConnectionId`.
- **Sensitive fields**: The `Profile` property references AWS Secrets Manager, which securely stores Atlas API credentials. Never hardcode credentials in templates.
- **No Stream Workspace required**: Private Link Endpoints are created at the project level and do not require a Stream Processing Workspace to be created first.
- **S3-specific behavior**: For S3 vendors, `InterfaceEndpointId` and `InterfaceEndpointName` are expected to be null, as S3 uses gateway endpoints rather than interface endpoints.

## Field Mapping

The following table maps CloudFormation properties to Atlas API fields:

| CloudFormation Property | Atlas API Field     | Required | Notes                                  |
| ----------------------- | ------------------- | -------- | -------------------------------------- |
| `ProjectId`             | `projectId`         | Yes      | Unique 24-hexadecimal digit string     |
| `ProviderName`          | `providerName`      | Yes      | Always "AWS" for CloudFormation        |
| `Vendor`                | `vendor`            | Yes      | Valid values: "MSK", "CONFLUENT", "S3" |
| `Region`                | `region`            | No       | AWS region or Confluent domain name    |
| `ServiceEndpointId`     | `serviceEndpointId` | No       | VPC endpoint service name or S3 format |
| `Arn`                   | `arn`               | No       | Required for MSK vendor                |
| `DnsDomain`             | `dnsDomain`         | No       | Required for CONFLUENT vendor          |
| `DnsSubDomain`          | `dnsSubDomain`      | No       | Required for CONFLUENT vendor (array)  |
| `Profile`               | N/A                 | No       | AWS Secrets Manager profile name       |

## Cleanup

To delete the stack and all resources:

```bash
aws cloudformation delete-stack \
  --stack-name atlas-stream-privatelink-endpoint-s3 \
  --region eu-west-1
```

**Note**: The Private Link Endpoint will be automatically deleted when the stack is deleted. However, any AWS VPC endpoints created by the resource may need to be manually cleaned up if they are not automatically removed.

## Related Resources

- [MongoDB::Atlas::StreamInstance](../stream-instance/stream-instance.json): Create a Stream Processing instance
- [MongoDB::Atlas::StreamConnection](../stream-connection/): Create connections to data sources and sinks
- [Stream Processing Documentation](https://www.mongodb.com/docs/atlas/atlas-stream-processing/manage-connection-registry/)
