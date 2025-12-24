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

### Example 3: Sample to Cluster with Time-Series Collection ([stream-processor-sample-emit.json](stream-processor-sample-emit.json))

Creates a stream processor that reads from a sample connection and emits data to a time-series collection in a cluster connection using `$emit`.

**Use cases:**

- Sample data to time-series collection
- Real-time time-series data ingestion
- IoT sensor data streaming

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **WorkspaceName** - Name of your stream instance/workspace
3. **ProcessorName** - Unique name for the stream processor
4. **SourceConnectionName** - Name of the source connection (default: `sample_stream_solar`)
5. **SinkConnectionName** - Name of the sink cluster connection (must be a cluster connection)
6. **SinkDatabase** - Target database name (optional, default: `sample`)
7. **SinkCollection** - Target time-series collection name (optional, default: `solar`)
8. **TimeField** - Field name containing timestamp (optional, default: `_ts`)
9. **DesiredState** - Desired state of the processor: `CREATED`, `STOPPED`, or `STARTED` (optional, default: `STARTED`)
10. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Pipeline stages:**

- `$source` - Reads from the sample connection
- `$emit` - Emits data to a time-series collection with timeseries configuration

**Note:** The target collection must be a time-series collection. Ensure the collection exists with the correct time-series configuration before starting the processor.

## Pipeline Stage Options

### $source

Reads data from a source connection. Supported sources:

- **Sample connections**: `sample_stream_solar` (for testing)
- **Cluster connections**: Read from MongoDB collections

### $emit

Writes data to a target connection. Options:

- **Cluster**: Write to MongoDB collections
  - `connectionName` - Target cluster connection name
  - `db` - Target database
  - `coll` - Target collection
  - `timeseries` (optional) - For time-series collections
    - `timeField` - Field name containing timestamp

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
