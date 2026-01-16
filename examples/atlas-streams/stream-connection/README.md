# How to create a MongoDB::Atlas::StreamConnection 

## Step 1: Activate the stream connection resource in cloudformation
   Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

   Step b: Search for Mongodb::Atlas::StreamConnection resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Your StreamConnection Resource is ready to use.

## Step 2: Create template using example JSON files
    Examples for each connection type:
    
    **Cluster type** - Connect to an Atlas cluster:
    - [cluster-stream-connection.json](cluster-stream-connection.json)
    
    **Kafka type** - Connect to a Kafka cluster:
    - [kafka-stream-connection.json](kafka-stream-connection.json)
    - [kafka-oauth-stream-connection.json](kafka-oauth-stream-connection.json)
    
    **Sample type** - Use sample datasets:
    - [sample-stream-connection.json](sample-stream-connection.json)
    
    **AWSLambda type** - Connect to AWS Lambda:
    - [aws-lambda-stream-connection.json](aws-lambda-stream-connection.json)
    
    **Https type** - Connect via HTTPS:
    - [https-stream-connection.json](https-stream-connection.json)

    Note: Make sure you are providing appropriate values for: 
    1. ProjectId
    2. WorkspaceName (or InstanceName - deprecated)
    3. ConnectionName
    4. Type: Cluster, Kafka, Sample, AWSLambda, or Https
    5. Profile (optional)
    6. Type-specific fields (ClusterName for Cluster type, BootstrapServers for Kafka type, etc.)
