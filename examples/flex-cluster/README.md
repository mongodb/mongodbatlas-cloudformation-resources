# How to create a MongoDB::Atlas::FlexCluster 

## Step 1: Activate the cluster resource in cloudformation
   Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

   Step b: Search for Mongodb::Atlas::FlexCluster resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Your FlexCluster Resource is ready to use.

## Step 2: Create template using [flex-cluster.json](flex-cluster.json)
    Note: Make sure you are providing appropriate values for: 
    1. ProjectId
    2. ClusterName
    3. BackingProviderName: AWS, GCP or AZURE
    4. RegionName
    5. Profile (optional)
    6. TerminationProtectionEnabled (optional)