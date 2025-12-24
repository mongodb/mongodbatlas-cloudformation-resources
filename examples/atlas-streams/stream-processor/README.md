# Stream Processor CloudFormation Examples

This directory contains example CloudFormation templates for creating MongoDB Atlas Stream Processors.

## Prerequisites

Before deploying these templates, ensure you have:

1. **AWS Account**: An active AWS account with appropriate permissions
2. **AWS Credentials**: Configured AWS credentials with permissions to:
   - Deploy CloudFormation stacks
   - Access AWS Secrets Manager (for Atlas API keys)
   - Create IAM roles (if using `CAPABILITY_IAM`)
3. **Atlas API Keys**: Stored in AWS Secrets Manager under the path `cfn/atlas/profile/{Profile}` (default profile is `default`)
   - For instructions on setting up credentials, see the [main README](/README.md#mongodb-atlas-api-keys-credential-management)
4. **Atlas Project**: An existing MongoDB Atlas project
5. **Stream Instance/Workspace**: An existing stream instance (created via `MongoDB::Atlas::StreamInstance` resource)
6. **Stream Connections**:
   - A source connection (e.g., `sample_stream_solar` for sample data, or a cluster/Kafka connection)
   - A sink connection (must be a cluster connection for merge operations)
   - For DLQ examples: A separate cluster connection for the dead letter queue

## Example Templates

### 1. Basic Stream Processor (`stream-processor.json`)

Creates a basic stream processor that:
- Reads from a source connection
- Merges data into a sink connection (cluster)
- Supports configurable state (CREATED, STARTED, or STOPPED)

**Parameters:**
- `ProjectId`: Your Atlas project ID (24 hexadecimal characters)
- `WorkspaceName`: Name of your stream workspace/instance
- `ProcessorName`: Name for the stream processor
- `SourceConnectionName`: Name of the source connection (e.g., `sample_stream_solar`)
- `SinkConnectionName`: Name of the sink cluster connection
- `SinkDatabase`: Target database name (default: `test`)
- `SinkCollection`: Target collection name (default: `output`)
- `State`: Initial processor state (default: `CREATED`)

**Deploy:**
```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-processor/stream-processor.json \
  --stack-name stream-processor-stack \
  --parameter-overrides \
    ProjectId=<YOUR_PROJECT_ID> \
    WorkspaceName=<YOUR_WORKSPACE_NAME> \
    ProcessorName=my-processor \
    SourceConnectionName=sample_stream_solar \
    SinkConnectionName=<YOUR_CLUSTER_CONNECTION_NAME> \
    SinkDatabase=test \
    SinkCollection=output \
    State=CREATED \
  --capabilities CAPABILITY_IAM \
  --region us-east-1
```

### 2. Stream Processor with DLQ (`stream-processor-with-dlq.json`)

Creates a stream processor with Dead Letter Queue (DLQ) configuration:
- All features of the basic processor
- Failed messages are sent to a DLQ collection
- Useful for error handling and message recovery

**Additional Parameters:**
- `DlqConnectionName`: Name of the DLQ cluster connection
- `DlqDatabase`: DLQ database name (default: `dlq`)
- `DlqCollection`: DLQ collection name (default: `dlq-messages`)

**Deploy:**
```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-processor/stream-processor-with-dlq.json \
  --stack-name stream-processor-dlq-stack \
  --parameter-overrides \
    ProjectId=<YOUR_PROJECT_ID> \
    WorkspaceName=<YOUR_WORKSPACE_NAME> \
    ProcessorName=my-processor-dlq \
    SourceConnectionName=sample_stream_solar \
    SinkConnectionName=<YOUR_CLUSTER_CONNECTION_NAME> \
    SinkDatabase=test \
    SinkCollection=output \
    DlqConnectionName=<YOUR_DLQ_CLUSTER_CONNECTION_NAME> \
    DlqDatabase=dlq \
    DlqCollection=dlq-messages \
    State=CREATED \
  --capabilities CAPABILITY_IAM \
  --region us-east-1
```

## Template Validation

Validate template syntax before deployment:

```bash
# Validate basic template
aws cloudformation validate-template \
  --template-body file://examples/atlas-streams/stream-processor/stream-processor.json

# Validate DLQ template
aws cloudformation validate-template \
  --template-body file://examples/atlas-streams/stream-processor/stream-processor-with-dlq.json
```

## Verification

After deployment, verify the stream processor:

```bash
# List all processors for a workspace
atlas streams processors list <WORKSPACE_NAME> --projectId <PROJECT_ID>

# Describe specific processor
atlas streams processors describe <PROCESSOR_NAME> \
  --instance <WORKSPACE_NAME> \
  --projectId <PROJECT_ID>
```

### Expected Output Fields

- `id`: Unique processor identifier (matches CloudFormation `Id` attribute)
- `name`: Processor name (matches `ProcessorName` parameter)
- `state`: Current state (CREATED, STARTED, STOPPED, or FAILED)
- `pipeline`: Array of pipeline stages
- `options.dlq`: DLQ configuration (if provided)
- `stats`: Processing statistics (when processor is STARTED)

## Field Mapping

The following table maps CloudFormation properties to Atlas API fields:

| CloudFormation Property | Atlas API Field | Notes |
|------------------------|-----------------|-------|
| `Profile` | N/A | Used for AWS Secrets Manager credential lookup |
| `ProjectId` | `groupId` | 24-character hexadecimal project identifier |
| `WorkspaceName` | `tenantName` | Preferred field name (replaces deprecated `InstanceName`) |
| `InstanceName` | `tenantName` | Deprecated, use `WorkspaceName` instead |
| `ProcessorName` | `name` | Unique processor name within the workspace |
| `Pipeline` | `pipeline` | JSON-encoded array of aggregation pipeline stages |
| `State` | `state` | Valid values: CREATED, STARTED, STOPPED |
| `Options.Dlq.ConnectionName` | `options.dlq.connectionName` | DLQ connection name (must be cluster connection) |
| `Options.Dlq.Db` | `options.dlq.db` | DLQ database name |
| `Options.Dlq.Coll` | `options.dlq.coll` | DLQ collection name |
| `Id` | `id` | Read-only: Unique processor identifier |
| `Stats` | `stats` | Read-only: Processing statistics (JSON string) |
| `Timeouts.Create` | N/A | CloudFormation-specific timeout configuration |
| `DeleteOnCreateTimeout` | N/A | Write-only: Controls cleanup behavior on timeout |

### Primary Identifier

The primary identifier for this resource consists of:
- `ProjectId`
- `WorkspaceName` (or `InstanceName` for backward compatibility)
- `ProcessorName`
- `Profile`

All of these fields must be present in every returned model to ensure CloudFormation can properly track the resource.

## Important Notes

- **AWS Only**: These templates are designed for AWS CloudFormation deployments. The provider is effectively AWS.
- **WorkspaceName vs InstanceName**:
  - Use `WorkspaceName` parameter (preferred)
  - `InstanceName` is supported for backward compatibility but is deprecated
  - CloudFormation does not enforce mutual exclusivity - both can be present, with `WorkspaceName` taking precedence
- **State Management**:
  - `CREATED`: Processor is created but not started
  - `STARTED`: Processor is created and immediately started (long-running operation)
  - `STOPPED`: Processor is created in stopped state
- **Long-Running Operations**: Processor creation/startup can take several minutes. The resource uses callback-based state management to handle these operations asynchronously.
- **Pipeline Format**: The pipeline is a JSON-encoded array of aggregation stages. See [MongoDB Stream Processing Documentation](https://www.mongodb.com/docs/atlas/atlas-sp/overview/) for pipeline syntax
- **Sensitive Fields**: No sensitive fields are exposed in this resource. All credentials are managed via AWS Secrets Manager using the `Profile` parameter.

## Cleanup

To delete the stack and remove the stream processor:

```bash
aws cloudformation delete-stack --stack-name stream-processor-stack --region us-east-1
```

Note: The stream processor will be deleted from Atlas when the CloudFormation stack is deleted.
