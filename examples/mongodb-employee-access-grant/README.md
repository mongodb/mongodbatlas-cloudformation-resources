# How to create a MongoDB::Atlas::MongoDbEmployeeAccessGrant

## Step 1: Activate the mongodb employee access grant resource in cloudformation

Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

Step b: Search for Mongodb::Atlas::MongoDbEmployeeAccessGrant resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your MongoDbEmployeeAccessGrant Resource is ready to use.

## Step 2: Create template using [mongodb-employee-access-grant.json](mongodb-employee-access-grant.json)

    Note: Make sure you are providing appropriate values for:
    1. ProjectId
    2. ClusterName
    3. GrantType: CLUSTER_DATABASE_LOGS, CLUSTER_INFRASTRUCTURE, or CLUSTER_INFRASTRUCTURE_AND_APP_SERVICES_SYNC_DATA
    4. ExpirationTime: RFC3339 format (e.g., 2025-08-01T12:00:00Z)
    5. Profile (optional)
