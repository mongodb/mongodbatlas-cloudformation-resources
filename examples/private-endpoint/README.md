# How to create a MongoDB::Atlas::PrivateEndpoint 

##Step 1: Activate the private-endpoint resource in cloudformation
   Step a: Create Role using [execution-role.json](execution-role.json) in this folder

   Step b: Search for Mongodb::Atlas::PrivateEndpoint resource.

         (Cloudformation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Private Endpoint Resource is ready to use.

## Step 2: Create template using [privateEndpoint.json](privateEndpoint.json)
    Note: Make sure you are providing appropriate vales for 
    1. OrgId
    2. PublicKey
    3. PrivateKey
    4. Region
    5. VpcId 
    6. SubnetId
