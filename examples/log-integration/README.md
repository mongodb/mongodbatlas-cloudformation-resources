# How to create a MongoDB::Atlas::LogIntegration

## Step 1: Activate the LogIntegration resource in CloudFormation

### Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder

### Step b: Search for MongoDB::Atlas::LogIntegration resource

```
CloudFormation > Public extensions > choose 'Third party' > Search with "Execution name prefix = MongoDB"
```

### Step c: Select and activate

Enter the RoleArn that is created in step 1.

Your LogIntegration Resource is ready to use.

## Step 2: Create template using [log-integration.json](log-integration.json)

Note: Make sure you are providing appropriate values for:

1. **ProjectId** (required) - Your MongoDB Atlas project ID
2. **BucketName** (required) - S3 bucket name for log export
3. **IamRoleId** (required) - 24-character hex string from Atlas Cloud Provider Access
4. **PrefixPath** (optional) - S3 prefix path (default: "mongodb/logs")
5. **LogType1** (required) - First log type to export (MONGOD, MONGOS, MONGOD_AUDIT, MONGOS_AUDIT)
6. **LogType2** (optional) - Second log type to export
7. **KmsKey** (optional) - AWS KMS key ARN for encryption
8. **Profile** (optional) - Secret Manager profile name (default: "default")
