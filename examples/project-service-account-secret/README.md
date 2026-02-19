# How to create a MongoDB::Atlas::ProjectServiceAccountSecret

## Step 1: Activate the resource in CloudFormation

Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

Step b: Search for MongoDB::Atlas::ProjectServiceAccountSecret resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your ProjectServiceAccountSecret Resource is ready to use.

## Step 2: Create template using [project-service-account-secret.json](project-service-account-secret.json)

    Note: Make sure you are providing appropriate values for:
    1. ProjectId
    2. ClientId (of an existing Project Service Account)
    3. SecretExpiresAfterHours (optional)
    4. Profile (optional)

## Important Notes

- **Existing Service Account Required**: You must have an existing Project Service Account. Create one using the MongoDB Atlas UI or API.
- **Secret Value**: The actual secret value is only returned during creation and is not available in subsequent reads.
- **No Updates**: This resource does not support updates. Any property change will trigger a replacement (delete + create).
