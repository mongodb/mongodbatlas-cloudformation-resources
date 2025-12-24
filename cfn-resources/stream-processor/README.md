# MongoDB::Atlas::StreamProcessor

## Description

Resource for creating and managing [Stream Processors for an Atlas Stream Instance](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createstreamprocessor).

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md). Also refer [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.

## Cloudformation Examples

See the example [CFN Templates](/examples/atlas-streams/stream-processor/) for example resources:
- [Basic Stream Processor](/examples/atlas-streams/stream-processor/stream-processor.json)
- [Stream Processor with DLQ](/examples/atlas-streams/stream-processor/stream-processor-with-dlq.json)

## Prerequisites

Before creating a stream processor, you must have:
- An existing Atlas Project
- An existing Stream Instance/Workspace (created via `MongoDB::Atlas::StreamInstance` resource)
- At least one Stream Connection configured (created via `MongoDB::Atlas::StreamConnection` resource)
  - A source connection (e.g., sample data source, cluster connection, or Kafka connection)
  - A sink connection (must be a cluster connection for merge operations)

## Deployment

### Deploy Basic Stream Processor

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

### Deploy Stream Processor with DLQ

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

## Verification

After deployment, verify the stream processor was created successfully using both Atlas CLI and Atlas UI.

### Atlas CLI Verification

```bash
# List all stream processors for a workspace
atlas streams processors list <WORKSPACE_NAME> --projectId <PROJECT_ID>

# Describe a specific stream processor
atlas streams processors describe <PROCESSOR_NAME> \
  --instance <WORKSPACE_NAME> \
  --projectId <PROJECT_ID>
```

### Expected CLI Output

The `atlas streams processors describe` command should return:
- `id`: Unique identifier of the processor (matches the `Id` attribute in CloudFormation)
- `name`: Processor name (matches `ProcessorName` parameter)
- `state`: Current state (CREATED, STARTED, STOPPED, or FAILED)
- `pipeline`: Array of pipeline stages matching your Pipeline configuration
- `options`: DLQ configuration if provided (should match your Options.Dlq settings)
- `stats`: Processing statistics (available when processor is STARTED)

### Verify Pipeline Configuration

The pipeline should match your CloudFormation template:
- Source connection name should match `SourceConnectionName` parameter
- Merge target connection should match `SinkConnectionName` parameter
- Database and collection should match `SinkDatabase` and `SinkCollection` parameters

### Verify DLQ Configuration (if applicable)

For processors with DLQ:
- `options.dlq.connectionName` should match `DlqConnectionName` parameter
- `options.dlq.db` should match `DlqDatabase` parameter
- `options.dlq.coll` should match `DlqCollection` parameter

### Atlas UI Verification

1. Navigate to your Atlas project in the [Atlas UI](https://cloud.mongodb.com)
2. Go to **Stream Processing** section
3. Select your stream workspace/instance
4. Verify the processor appears in the **Processors** tab with:
   - **Name**: Matches the `ProcessorName` from your CloudFormation template
   - **State**: Matches the `State` parameter (CREATED, STARTED, or STOPPED)
   - **Pipeline**: Click on the processor to view pipeline stages and verify:
     - Source connection matches your `SourceConnectionName` parameter
     - Merge target connection matches your `SinkConnectionName` parameter
     - Target database and collection match your `SinkDatabase` and `SinkCollection` parameters
5. For processors with DLQ:
   - Verify DLQ configuration is displayed in the processor details
   - Check that DLQ connection, database, and collection match your parameters
6. If processor is in STARTED state:
   - Verify processing statistics are available
   - Check that messages are being processed (stats show input/output message counts)

## Notes

- **AWS Only**: This CloudFormation resource is designed for AWS deployments. The provider is effectively AWS.
- **WorkspaceName vs InstanceName**: Use `WorkspaceName` (preferred). `InstanceName` is supported for backward compatibility but is deprecated.
- **State Management**: When creating a processor, specify `State: STARTED` to automatically start processing, or `State: CREATED` to create it in a stopped state.
- **Long-Running Operations**: Creating and starting stream processors can take several minutes. The resource uses callback-based state management to handle these operations asynchronously.
- **Timeout Configuration**: Use `Timeouts.Create` to configure how long to wait for processor creation/startup (default: 20 minutes).
