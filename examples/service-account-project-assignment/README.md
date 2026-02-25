# How to create a MongoDB::Atlas::ServiceAccountProjectAssignment

## Step 1: Activate the service account project assignment resource in cloudformation

Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

Step b: Search for Mongodb::Atlas::ServiceAccountProjectAssignment resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your ServiceAccountProjectAssignment Resource is ready to use.

## Step 2: Create template using [service-account-project-assignment.json](service-account-project-assignment.json)

    Note: Make sure you are providing appropriate values for:
    1. OrgId
    2. ClientId
    3. ProjectId
    4. Roles
    5. Profile (optional)
