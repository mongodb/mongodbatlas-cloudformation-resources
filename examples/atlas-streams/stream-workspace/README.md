# How to create a MongoDB::Atlas::StreamWorkspace 

## Step 1: Activate the stream workspace resource in cloudformation
   Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

   Step b: Search for Mongodb::Atlas::StreamWorkspace resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Your StreamWorkspace Resource is ready to use.

## Step 2: Create template using [stream-workspace.json](stream-workspace.json)
    Note: Make sure you are providing appropriate values for: 
    1. ProjectId
    2. WorkspaceName (optional)
    3. CloudProvider: AWS (optional, default: AWS)
    4. Region (optional, default: VIRGINIA_USA)
    5. Tier: SP2, SP5, SP10, SP30, or SP50 (optional, default: SP30)
    6. MaxTierSize (optional, default: SP50)
    7. Profile (optional)
