# MongoDB::Atlas::StreamWorkspace Examples

This directory contains example CloudFormation templates for creating Stream Workspaces in MongoDB Atlas.

## Prerequisites

1. **Atlas Project**: You need an existing Atlas project. Get your Project ID from the Atlas UI or using:
   ```bash
   atlas projects list
   ```

2. **AWS Credentials**: Ensure your AWS credentials are configured with permissions to:
   - Create/update/delete CloudFormation stacks
   - Access AWS Secrets Manager (for storing Atlas API keys)

3. **Atlas API Keys**: Store your Atlas API keys in AWS Secrets Manager:
   ```bash
   aws secretsmanager create-secret \
     --name cfn/atlas/profile/default \
     --secret-string '{"PublicKey":"YOUR_PUBLIC_KEY","PrivateKey":"YOUR_PRIVATE_KEY","BaseURL":"https://cloud.mongodb.com"}' \
     --region eu-west-1
   ```

4. **Resource Type Registered**: Ensure the `MongoDB::Atlas::StreamWorkspace` resource type is registered in your AWS CloudFormation Private Registry:
   ```bash
   aws cloudformation describe-type \
     --type RESOURCE \
     --type-name MongoDB::Atlas::StreamWorkspace \
     --region eu-west-1
   ```

## Example Templates

### 1. Stream Workspace (`stream-workspace.json`)

Creates a Stream Workspace with configurable tier and data processing region.

**Parameters:**
- `ProjectId`: Your Atlas project ID (24-hexadecimal characters, required)
- `WorkspaceName`: Name for the Stream Workspace (optional, will auto-generate if empty)
- `CloudProvider`: Cloud provider for data processing region (default: "AWS", AWS only for CloudFormation)
- `Region`: Region for data processing (default: "VIRGINIA_USA")
- `Tier`: Stream Workspace Tier - "SP30", "SP50", or "SP100" (default: "SP30")
- `Profile`: AWS Secrets Manager profile name (default: "default")

**Deploy:**
```bash
# Setup credentials first
source ./CONVERSION_PROMPTS/setup-credentials.sh CONVERSION_PROMPTS/credPersonalCfnDev.properties

# Deploy the stack
aws cloudformation create-stack \
  --stack-name stream-workspace-example-$(date +%s) \
  --template-body file://examples/atlas-streams/stream-workspace/stream-workspace.json \
  --parameters \
    ParameterKey=ProjectId,ParameterValue=YOUR_PROJECT_ID \
    ParameterKey=WorkspaceName,ParameterValue=my-stream-workspace \
    ParameterKey=CloudProvider,ParameterValue=AWS \
    ParameterKey=Region,ParameterValue=VIRGINIA_USA \
    ParameterKey=Tier,ParameterValue=SP30 \
    ParameterKey=Profile,ParameterValue=default \
  --capabilities CAPABILITY_IAM \
  --region eu-west-1
```

**Monitor Stack Creation:**
```bash
# Check stack status
aws cloudformation describe-stacks \
  --stack-name <stack-name> \
  --region eu-west-1 \
  --query 'Stacks[0].StackStatus' \
  --output text

# Check resource status
aws cloudformation describe-stack-resources \
  --stack-name <stack-name> \
  --region eu-west-1

# Check CloudWatch logs for handler execution
aws logs describe-log-groups \
  --log-group-name-prefix "mongodb-atlas-streamworkspace" \
  --region eu-west-1
```

**Expected Stack Creation Time:**
- Typically 5-10 seconds for stream workspace creation
- Stack status should transition: `CREATE_IN_PROGRESS` → `CREATE_COMPLETE`

**Verify with Atlas CLI:**
```bash
# List all stream workspaces
atlas streams instances list --projectId <PROJECT_ID>

# Get specific workspace details
atlas streams instances describe <WORKSPACE_NAME> --projectId <PROJECT_ID>
```

**Expected Output:**
- Workspace should appear in the list with the specified name
- Workspace should show:
  - `name`: Matches the WorkspaceName parameter
  - `dataProcessRegion.cloudProvider`: "AWS"
  - `dataProcessRegion.region`: Matches the Region parameter (e.g., "VIRGINIA_USA")
  - `streamConfig.tier`: Matches the Tier parameter (e.g., "SP30")
  - `hostnames`: Array of hostnames for connecting to the workspace

**Cross-Reference with CloudFormation:**
```bash
# Get physical resource ID from stack
aws cloudformation describe-stack-resources \
  --stack-name <stack-name> \
  --region eu-west-1 \
  --query 'StackResources[?LogicalResourceId==`StreamWorkspace`].PhysicalResourceId' \
  --output text

# Get stack outputs
aws cloudformation describe-stacks \
  --stack-name <stack-name> \
  --region eu-west-1 \
  --query 'Stacks[0].Outputs' \
  --output json
```

**Stack Outputs:**
- `StreamWorkspaceId`: The unique identifier for the Stream Workspace
- `StreamWorkspaceName`: The name of the Stream Workspace
- `StreamWorkspaceHostnames`: Comma-separated list of hostnames assigned to the stream workspace

**Cleanup:**
```bash
# Delete the stack (will also delete the stream workspace)
aws cloudformation delete-stack \
  --stack-name <stack-name> \
  --region eu-west-1

# Wait for deletion to complete
aws cloudformation wait stack-delete-complete \
  --stack-name <stack-name> \
  --region eu-west-1

# Verify workspace is deleted
atlas streams instances list --projectId <PROJECT_ID>
```

## Notes

- **AWS Only**: This CloudFormation resource is designed for AWS deployments. The CloudProvider parameter is constrained to "AWS" only.
- **Create-Only Properties**: `WorkspaceName`, `ProjectId`, and `Profile` are create-only properties. To change these, you must delete and recreate the stack.
- **Updateable Properties**: `StreamConfig.Tier` and `StreamConfig.MaxTierSize` can be updated after creation.
- **Read-Only Properties**: `Id`, `Hostnames`, and `Connections` are read-only and returned by CloudFormation but cannot be set.

## Troubleshooting

**Stack Creation Fails:**
- Verify Atlas API keys are correctly stored in AWS Secrets Manager
- Check CloudWatch logs for handler execution errors
- Ensure the resource type is registered in your private registry
- Verify your IP address is on the Atlas IP Access List

**Workspace Not Found in Atlas:**
- Wait a few seconds after stack creation completes
- Verify the Project ID is correct
- Check Atlas UI for the workspace

**Handler Execution Errors:**
- Review CloudWatch logs: `aws logs tail /aws/lambda/mongodb-atlas-streamworkspace-role-stack-* --follow --region eu-west-1`
- Verify execution role has `secretsmanager:GetSecretValue` permission
- Check Atlas API key permissions in Atlas UI
