# How to create a MongoDB::Atlas::PrivateEndpointAWS

## Step 1: Activate the PrivateEndpointAWS resource in CloudFormation
   Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

   Step b: Search for MongoDB::Atlas::PrivateEndpointAWS resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Your PrivateEndpointAWS Resource is ready to use.

## Step 2: Create template using [private-endpoint-aws.json](private-endpoint-aws.json)
   Note: Make sure you are providing appropriate values for:
   1. MongoDBAtlasProjectId
   2. AtlasPrivateEndpointServiceId (get from: `atlas privateEndpoints aws list --projectId <PROJECT_ID>`)
   3. AWSVPCEndpointId (format: vpce-xxxxxxxxx)
   4. Profile (optional)
   5. EnforceConnectionSuccess (optional)