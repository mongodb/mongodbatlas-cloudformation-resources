# How to create a MongoDB::Atlas::StreamProcessor

## Step 1: Activate the stream processor resource in cloudformation

Step a: Create Role using [execution-role.yaml](../../execution-role.yaml) in examples folder.

Step b: Search for Mongodb::Atlas::StreamProcessor resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your StreamProcessor Resource is ready to use.

## Step 2: Choose a template based on your use case

### Example 1: Basic Stream Processor ([stream-processor.json](stream-processor.json))

Creates a stream processor that reads from a source connection and merges data into a cluster connection. This example uses `$merge` to write data to a regular MongoDB collection.

**Use cases:**

- Sample data to cluster (e.g., using `sample_stream_solar`)
- Cluster to cluster data streaming
- Simple data replication

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **WorkspaceName** - Name of your stream instance/workspace
3. **ProcessorName** - Unique name for the stream processor
4. **SourceConnectionName** - Name of the source connection:
   - For sample data: `sample_stream_solar`
   - For cluster source: Your cluster connection name
5. **SinkConnectionName** - Name of the sink cluster connection (must be a cluster connection)
6. **SinkDatabase** - Target database name (optional, default: `test`)
7. **SinkCollection** - Target collection name (optional, default: `output`)
8. **DesiredState** - Desired state of the processor: `CREATED`, `STOPPED`, or `STARTED` (optional, default: `CREATED`)
9. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Pipeline stages:**

- `$source` - Reads from the source connection
- `$merge` - Merges data into the target cluster connection (for regular collections)

### Example 2: Stream Processor with Dead Letter Queue ([stream-processor-dlq.json](stream-processor-dlq.json))

Creates a stream processor with Dead Letter Queue (DLQ) configuration. Failed messages are automatically sent to a DLQ collection for error handling and debugging.

**Additional Parameters (beyond Example 1):**

10. **DlqConnectionName** - Name of the DLQ connection (must be a cluster connection)
11. **DlqDatabase** - DLQ database name (optional, default: `dlq`)
12. **DlqCollection** - DLQ collection name (optional, default: `dlq-messages`)

**Pipeline stages:**

- `$source` - Reads from the source connection
- `$merge` - Merges data into the target cluster connection (for regular collections)
- **Options.Dlq** - Configured to capture failed messages

### Example 3: Kafka to Cluster Stream Processor ([stream-processor-kafka-to-cluster.json](stream-processor-kafka-to-cluster.json))

Creates a stream processor that reads from a Kafka topic and writes to a cluster connection as a time-series collection.

**Use cases:**

- Ingesting data from Kafka into MongoDB Atlas
- Real-time data pipeline from Kafka to MongoDB
- Event streaming from Kafka to time-series collections

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **WorkspaceName** - Name of your stream instance/workspace
3. **ProcessorName** - Unique name for the stream processor
4. **KafkaSourceConnectionName** - Name of the Kafka source connection
5. **KafkaTopic** - Name of the Kafka topic to read from
6. **SinkConnectionName** - Name of the sink cluster connection (must be a cluster connection)
7. **SinkDatabase** - Target database name (optional, default: `kafka`)
8. **SinkCollection** - Target collection name (optional, default: `kafka_messages`)
9. **DesiredState** - Must be `CREATED` or `STOPPED` (cannot be `STARTED` without a working Kafka cluster)
10. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Pipeline stages:**

- `$source` - Reads from Kafka topic (requires `connectionName` and `topic`)
- `$emit` - Writes to cluster connection as time-series collection

**Important Notes:**

- ⚠️ **This processor must be created in `CREATED` state** - it cannot be started (`STARTED`) without a working Kafka cluster that is accessible from MongoDB Atlas Stream Processing infrastructure
- The processor will fail if you attempt to start it without a valid Kafka connection
- To use this processor with a real Kafka cluster, first ensure your Kafka connection is properly configured and accessible, then update the processor's `DesiredState` to `STARTED`

### Example 4: Cluster to Kafka Stream Processor ([stream-processor-cluster-to-kafka.json](stream-processor-cluster-to-kafka.json))

Creates a stream processor that reads from a cluster connection and writes to a Kafka topic.

**Use cases:**

- Streaming MongoDB data to Kafka
- Real-time data export from Atlas to Kafka
- Event streaming from MongoDB to Kafka topics

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **WorkspaceName** - Name of your stream instance/workspace
3. **ProcessorName** - Unique name for the stream processor
4. **SourceConnectionName** - Name of the source cluster connection
5. **KafkaSinkConnectionName** - Name of the Kafka sink connection
6. **KafkaTopic** - Name of the Kafka topic to write to
7. **DesiredState** - Must be `CREATED` or `STOPPED` (cannot be `STARTED` without a working Kafka cluster)
8. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Pipeline stages:**

- `$source` - Reads from cluster connection
- `$emit` - Writes to Kafka topic (requires `connectionName` and `topic`)

**Important Notes:**

- ⚠️ **This processor must be created in `CREATED` state** - it cannot be started (`STARTED`) without a working Kafka cluster that is accessible from MongoDB Atlas Stream Processing infrastructure
- The processor will fail if you attempt to start it without a valid Kafka connection
- To use this processor with a real Kafka cluster, first ensure your Kafka connection is properly configured and accessible, then update the processor's `DesiredState` to `STARTED`

## Pipeline Stage Options

### $source

Reads data from a source connection. Supported sources:

- **Sample connections**: `sample_stream_solar` (for testing)
- **Cluster connections**: Read from MongoDB collections
- **Kafka connections**: Read from Kafka topics (requires `topic` parameter)

### $emit

Writes data to a target connection. Options:

- **Cluster**: Write to MongoDB collections
  - `connectionName` - Target cluster connection name
  - `db` - Target database
  - `coll` - Target collection
  - `timeseries` (optional) - For time-series collections
    - `timeField` - Field name containing timestamp
- **Kafka**: Write to Kafka topics
  - `connectionName` - Target Kafka connection name
  - `topic` - Target Kafka topic name

### $merge

Merges data into regular MongoDB collections. Use `$merge` for standard collections (non-timeseries).

- **Cluster**: Merge into MongoDB collections
  - `connectionName` - Target cluster connection name
  - `db` - Target database
  - `coll` - Target collection
  - `into` - Object containing connection, database, and collection details

**Note:** Use `$merge` for regular collections. Use `$emit` only for time-series collections (requires `timeseries` option).

## State Management

The `DesiredState` parameter controls the desired processor lifecycle:

- **CREATED** - Processor is created but not running (default)
- **STARTED** - Processor is actively processing data
- **STOPPED** - Processor is stopped (can be restarted)

The `State` output (read-only) reflects the actual current state of the processor as returned by the Atlas API. Common states include `CREATED`, `STARTED`, `STOPPED`, and `FAILED`.

**Note:** When updating a processor, if the current state is `STARTED`, the processor will be stopped, updated, and then restarted if the `DesiredState` is `STARTED`.

## Kafka Integration Notes

When working with Kafka-based stream processors:

1. **Connection Validation**: The stream processor validates that the connection name exists in the workspace, but does not validate Kafka connectivity at creation time.

2. **State Management**: Kafka processors should be created in `CREATED` state. They can only be started (`STARTED`) when:
   - A valid Kafka connection exists
   - The Kafka cluster is accessible from MongoDB Atlas Stream Processing infrastructure
   - Authentication credentials are correct
   - Network connectivity is established (public access or VPC peering)

3. **Failure Handling**: If you attempt to start a Kafka processor without a working Kafka cluster, the processor will enter `FAILED` state. You can check the processor state via the `ProcessorState` output.

4. **Testing**: The examples provided (Examples 3 and 4) are designed to be created successfully even without a working Kafka cluster, allowing you to validate the CloudFormation template structure. To actually process data, you'll need a properly configured Kafka cluster.
