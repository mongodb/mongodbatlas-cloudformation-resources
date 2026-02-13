# How to create a MongoDB::Atlas::ServiceAccountSecret 

## Step 1: Activate the resource in CloudFormation
   Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

   Step b: Search for MongoDB::Atlas::ServiceAccountSecret resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Your ServiceAccountSecret Resource is ready to use.

## Step 2: Create template using [service-account-secret.json](service-account-secret.json)
    Note: Make sure you are providing appropriate values for: 
    1. OrgId
    2. ClientId (of an existing Service Account)
    3. SecretExpiresAfterHours (optional)
    4. Profile (optional)

## Important Notes

- **Existing Service Account Required**: You must have an existing Service Account. Create one using the [service-account example](../service-account/) or via the MongoDB Atlas UI.
- **Secret Value**: The actual secret value is only returned during creation and is not available in subsequent reads.
- **No Updates**: This resource does not support updates. Any property change will trigger a replacement (delete + create).
