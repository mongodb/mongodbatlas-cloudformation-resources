# MongoDB::Atlas::StreamPrivatelinkEndpoint

## Description

Resource for creating and managing [Stream Processing Private Link Endpoints](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createprivatelinkconnection). This resource enables secure, private connectivity between Atlas Stream Processing and streaming services (AWS MSK, Confluent Cloud, or AWS S3) over AWS PrivateLink. This resource supports AWS only.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md). Also refer [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.

## CloudFormation Examples

See the example [CFN Template](/examples/atlas-streams/stream-privatelink-endpoint/stream-privatelink-endpoint-s3.json) for example resource.

## Deployment

### Prerequisites

- **Atlas Project**: You must have an existing MongoDB Atlas project.
- **AWS Account**: An AWS account with appropriate permissions for PrivateLink and VPC endpoint creation.
- **Atlas Credentials**: Set up an AWS profile to securely give CloudFormation access to your Atlas credentials. For instructions, see [MongoDB Atlas API Keys Credential Management](/README.md#mongodb-atlas-api-keys-credential-management).
- **AWS Credentials**: Configure AWS credentials with permissions to create VPC endpoints and manage PrivateLink connections.

### Basic Deployment

```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-privatelink-endpoint/stream-privatelink-endpoint-s3.json \
  --stack-name atlas-stream-privatelink-endpoint-s3 \
  --parameter-overrides \
    ProjectId=<YOUR_PROJECT_ID> \
    Region=eu-west-1 \
    ServiceEndpointId=com.amazonaws.eu-west-1.s3 \
    Profile=default \
  --capabilities CAPABILITY_IAM \
  --region eu-west-1
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

### Using Atlas UI

1. Navigate to your Atlas project
2. Go to **Network Access** → **Stream Processing Private Link Endpoints**
3. Verify the endpoint appears with:
   - Vendor type: "S3", "MSK", or "CONFLUENT"
   - State: "AVAILABLE" (may take a few minutes to transition)
   - Correct region and service endpoint ID
