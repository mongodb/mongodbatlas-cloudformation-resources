# MongoDB::Atlas::StreamConnection Examples

This directory contains example CloudFormation templates for creating Stream Connections in MongoDB Atlas.

## Prerequisites

1. **Atlas Project**: You need an existing Atlas project. Get your Project ID from the Atlas UI or using:
   ```bash
   atlas projects list
   ```

2. **Stream Workspace**: You need an existing Stream Workspace (formerly Stream Instance). Create one using:
   ```bash
   atlas streams instances create <workspace-name> --projectId <PROJECT_ID> --region VIRGINIA_USA --provider AWS
   ```

3. **Atlas Cluster** (for Cluster type connections): You need an existing Atlas cluster. Create one using:
   ```bash
   atlas clusters create <cluster-name> --projectId <PROJECT_ID> --provider AWS --region US_EAST_1 --members 3 --tier M10
   ```

4. **AWS Credentials**: Ensure your AWS credentials are configured with permissions to:
   - Create/update/delete CloudFormation stacks
   - Access AWS Secrets Manager (for storing Atlas API keys)

5. **Atlas API Keys**: Store your Atlas API keys in AWS Secrets Manager:
   ```bash
   aws secretsmanager create-secret \
     --name cfn/atlas/profile/default \
     --secret-string '{"PublicKey":"YOUR_PUBLIC_KEY","PrivateKey":"YOUR_PRIVATE_KEY"}'
   ```

## Example Templates

### 1. Cluster Stream Connection (`cluster-stream-connection.json`)

Creates a connection of type `Cluster` that connects a Stream Workspace to an Atlas cluster.

**Parameters:**
- `ProjectId`: Your Atlas project ID (24-hexadecimal characters)
- `WorkspaceName`: Name of the existing Stream Workspace
- `ConnectionName`: Name for the stream connection
- `ClusterName`: Name of the existing Atlas cluster
- `DbRole`: Database role name (e.g., "atlasAdmin", "readWriteAnyDatabase")
- `DbRoleType`: Type of role - "BUILT_IN" or "CUSTOM"
- `Profile`: AWS Secrets Manager profile name (default: "default")

**Deploy:**
```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-connection/cluster-stream-connection.json \
  --stack-name stream-connection-cluster \
  --parameter-overrides \
    ProjectId=YOUR_PROJECT_ID \
    WorkspaceName=YOUR_WORKSPACE_NAME \
    ConnectionName=my-cluster-connection \
    ClusterName=YOUR_CLUSTER_NAME \
    DbRole=atlasAdmin \
    DbRoleType=BUILT_IN \
  --capabilities CAPABILITY_IAM \
  --region us-east-1
```

**Verify with Atlas CLI:**
```bash
# List all stream connections for the workspace
atlas streams connections list <WORKSPACE_NAME> --projectId <PROJECT_ID>

# Get specific connection details
atlas streams connections get <WORKSPACE_NAME> <CONNECTION_NAME> --projectId <PROJECT_ID>
```

**Expected Output:**
- Connection should appear in the list with Type: "Cluster"
- Connection should show ClusterName matching your cluster
- DbRoleToExecute should match the provided role

### 2. Kafka Stream Connection (`kafka-stream-connection.json`)

Creates a connection of type `Kafka` that connects a Stream Workspace to a Kafka cluster.

**Parameters:**
- `ProjectId`: Your Atlas project ID
- `WorkspaceName`: Name of the existing Stream Workspace
- `ConnectionName`: Name for the stream connection
- `BootstrapServers`: Comma-separated list of Kafka broker addresses (e.g., "localhost:9092,localhost:9093")
- `AuthMechanism`: Authentication mechanism - "PLAIN", "SCRAM-256", or "SCRAM-512"
- `AuthUsername`: Kafka username
- `AuthPassword`: Kafka password (will be hidden in console)
- `SecurityProtocol`: "PLAINTEXT" or "SSL"
- `BrokerPublicCertificate`: X.509 certificate for SSL connections (required if SecurityProtocol is SSL)
- `Profile`: AWS Secrets Manager profile name (default: "default")

**Deploy:**
```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-connection/kafka-stream-connection.json \
  --stack-name stream-connection-kafka \
  --parameter-overrides \
    ProjectId=YOUR_PROJECT_ID \
    WorkspaceName=YOUR_WORKSPACE_NAME \
    ConnectionName=my-kafka-connection \
    BootstrapServers=localhost:9092,localhost:9093 \
    AuthMechanism=PLAIN \
    AuthUsername=kafka-user \
    AuthPassword=kafka-password \
    SecurityProtocol=PLAINTEXT \
  --capabilities CAPABILITY_IAM \
  --region us-east-1
```

**Verify with Atlas CLI:**
```bash
# List all stream connections
atlas streams connections list <WORKSPACE_NAME> --projectId <PROJECT_ID>

# Get Kafka connection details
atlas streams connections get <WORKSPACE_NAME> <CONNECTION_NAME> --projectId <PROJECT_ID>
```

**Expected Output:**
- Connection should appear with Type: "Kafka"
- BootstrapServers should match your Kafka cluster
- Authentication mechanism should match the provided value

### 3. Sample Stream Connection (`sample-stream-connection.json`)

Creates a connection of type `Sample` that uses a sample dataset (e.g., `sample_stream_solar`).

**Parameters:**
- `ProjectId`: Your Atlas project ID
- `WorkspaceName`: Name of the existing Stream Workspace
- `ConnectionName`: Name of the sample dataset (default: "sample_stream_solar")
- `Profile`: AWS Secrets Manager profile name (default: "default")

**Deploy:**
```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-connection/sample-stream-connection.json \
  --stack-name stream-connection-sample \
  --parameter-overrides \
    ProjectId=YOUR_PROJECT_ID \
    WorkspaceName=YOUR_WORKSPACE_NAME \
    ConnectionName=sample_stream_solar \
  --capabilities CAPABILITY_IAM \
  --region us-east-1
```

**Verify with Atlas CLI:**
```bash
# List all stream connections
atlas streams connections list <WORKSPACE_NAME> --projectId <PROJECT_ID>

# Get sample connection details
atlas streams connections get <WORKSPACE_NAME> sample_stream_solar --projectId <PROJECT_ID>
```

**Expected Output:**
- Connection should appear with Type: "Sample"
- ConnectionName should be "sample_stream_solar"

## Field Mapping: CFN Properties → Atlas API

| CFN Property | Atlas API Field | Notes |
|-------------|----------------|-------|
| `ProjectId` | `groupId` | 24-hexadecimal character project ID |
| `WorkspaceName` | `tenantName` | Stream workspace name (preferred over InstanceName) |
| `InstanceName` | `tenantName` | Deprecated - use WorkspaceName instead |
| `ConnectionName` | `name` | Unique connection name within workspace |
| `Type` | `type` | Connection type: Cluster, Kafka, Sample, AWSLambda, Https |
| `ClusterName` | `clusterName` | Required for Cluster type |
| `DbRoleToExecute.Role` | `dbRoleToExecute.role` | Database role name |
| `DbRoleToExecute.Type` | `dbRoleToExecute.type` | BUILT_IN or CUSTOM |
| `BootstrapServers` | `bootstrapServers` | Required for Kafka type |
| `Authentication.Mechanism` | `authentication.mechanism` | PLAIN, SCRAM-256, SCRAM-512, OAUTHBEARER |
| `Security.Protocol` | `security.protocol` | PLAINTEXT or SSL |

## Template Validation

Before deploying, validate the template syntax:

```bash
# Validate cluster connection template
aws cloudformation validate-template \
  --template-body file://examples/atlas-streams/stream-connection/cluster-stream-connection.json

# Validate Kafka connection template
aws cloudformation validate-template \
  --template-body file://examples/atlas-streams/stream-connection/kafka-stream-connection.json

# Validate sample connection template
aws cloudformation validate-template \
  --template-body file://examples/atlas-streams/stream-connection/sample-stream-connection.json
```

If validation succeeds, the command will return JSON with template parameters and description. Any syntax errors will be reported.

## Atlas CLI Validation

After deploying a stack, verify the connection was created correctly:

### 1. List All Connections
```bash
atlas streams connections list <WORKSPACE_NAME> --projectId <PROJECT_ID> --output json
```

**Expected**: Your connection should appear in the list with:
- `name`: Matches your ConnectionName parameter
- `type`: Matches your connection type (Cluster, Kafka, or Sample)

### 2. Get Connection Details
```bash
atlas streams connections get <WORKSPACE_NAME> <CONNECTION_NAME> --projectId <PROJECT_ID> --output json
```

**For Cluster connections, verify:**
- `type` = "Cluster"
- `clusterName` = Your cluster name
- `dbRoleToExecute.role` = Your DbRole parameter
- `dbRoleToExecute.type` = Your DbRoleType parameter

**For Kafka connections, verify:**
- `type` = "Kafka"
- `bootstrapServers` = Your BootstrapServers parameter
- `authentication.mechanism` = Your AuthMechanism parameter
- `security.protocol` = Your SecurityProtocol parameter

**For Sample connections, verify:**
- `type` = "Sample"
- `name` = Your ConnectionName parameter (typically "sample_stream_solar")

### 3. Verify in Atlas UI
1. Navigate to your Atlas project
2. Go to Stream Processing section
3. Select your Stream Workspace
4. View Connections tab
5. Verify the connection appears with correct configuration

## Cleanup

To delete a stream connection created via CloudFormation:

```bash
# Delete the CloudFormation stack (recommended)
aws cloudformation delete-stack --stack-name <STACK_NAME>
aws cloudformation wait stack-delete-complete --stack-name <STACK_NAME>

# Or delete directly using Atlas CLI
atlas streams connections delete <WORKSPACE_NAME> <CONNECTION_NAME> --projectId <PROJECT_ID> --force
```

## Notes

- **AWS-Only**: These templates are designed for AWS CloudFormation. Provider is implicitly AWS.
- **Backward Compatibility**: The resource supports both `WorkspaceName` (preferred) and `InstanceName` (deprecated). If both are provided, `WorkspaceName` takes precedence. CFN does not enforce mutual exclusivity.
- **Primary Identifier**: The resource is uniquely identified by: `ProjectId`, `ConnectionName`, `WorkspaceName`, and `Profile`.
- **Sensitive Fields**: Passwords and secrets should be managed through AWS Secrets Manager or CloudFormation parameters with `NoEcho: true`.
- **Required Resources**: Ensure Stream Workspace and (for Cluster connections) Atlas Cluster exist before deploying connection templates.
