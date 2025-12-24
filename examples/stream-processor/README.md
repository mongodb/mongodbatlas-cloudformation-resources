# How to create a MongoDB::Atlas::StreamProcessor

## Step 1: Activate the stream processor resource in cloudformation

Step a: Create Role using [execution-role.yaml](../../execution-role.yaml) in examples folder.

Step b: Search for Mongodb::Atlas::StreamProcessor resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your StreamProcessor Resource is ready to use.

## Step 2: Create template using [stream-processor.json](stream-processor.json)

    Note: Make sure you are providing appropriate values for:
    1. ProjectId
    2. WorkspaceName (name of your stream instance/workspace)
    3. ProcessorName
    4. SourceConnectionName (e.g., sample_stream_solar for sample data, or your cluster/kafka connection)
    5. SinkConnectionName (must be a cluster connection name - use the connection name, not the cluster name)
    6. SinkDatabase (optional, default: test)
    7. SinkCollection (optional, default: output)
    8. Profile (optional, default: default)

    **Note**: The State property is read-only and cannot be set during creation. The stream processor will default to CREATED state.
