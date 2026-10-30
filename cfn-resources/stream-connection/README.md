# MongoDB::Atlas::StreamConnection

Resource for creating and managing [Connections for an Atlas Stream Instance](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createstreamconnection).

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md). Also refer [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.

## CloudFormation Examples

Example templates are available in the [examples directory](/examples/atlas-streams/stream-connection/):

- **Cluster Connection**: [cluster-stream-connection.json](/examples/atlas-streams/stream-connection/cluster-stream-connection.json) - Connects a Stream Workspace to an Atlas cluster
- **Kafka Connection**: [kafka-stream-connection.json](/examples/atlas-streams/stream-connection/kafka-stream-connection.json) - Connects a Stream Workspace to a Kafka cluster
- **Sample Connection**: [sample-stream-connection.json](/examples/atlas-streams/stream-connection/sample-stream-connection.json) - Uses a sample dataset

For detailed deployment and verification instructions, see the [examples README](/examples/atlas-streams/stream-connection/README.md).

## Deployment

### Prerequisites
1. An existing Atlas project
2. An existing Stream Workspace (create using `atlas streams instances create`)
3. For Cluster connections: An existing Atlas cluster
4. AWS credentials configured with CloudFormation permissions
5. Atlas API keys stored in AWS Secrets Manager

### Deploy Example

```bash
aws cloudformation deploy \
  --template-file examples/atlas-streams/stream-connection/cluster-stream-connection.json \
  --stack-name my-stream-connection \
  --parameter-overrides \
    ProjectId=YOUR_PROJECT_ID \
    WorkspaceName=YOUR_WORKSPACE_NAME \
    ConnectionName=my-connection \
    ClusterName=YOUR_CLUSTER_NAME \
    DbRole=atlasAdmin \
    DbRoleType=BUILT_IN \
  --capabilities CAPABILITY_IAM \
  --region us-east-1
```

## Verification

After deployment, verify the connection using Atlas CLI:

```bash
# List all connections for the workspace
atlas streams connections list <WORKSPACE_NAME> --projectId <PROJECT_ID>

# Get specific connection details
atlas streams connections get <WORKSPACE_NAME> <CONNECTION_NAME> --projectId <PROJECT_ID>
```

The connection should appear in the list with the correct type and configuration matching your CloudFormation template parameters.
