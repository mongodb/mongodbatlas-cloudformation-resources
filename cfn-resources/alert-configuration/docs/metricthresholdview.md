# MongoDB::Atlas::AlertConfiguration MetricThresholdView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#metricname" title="MetricName">MetricName</a>" : <i>String</i>,
    "<a href="#mode" title="Mode">Mode</a>" : <i>String</i>,
    "<a href="#operator" title="Operator">Operator</a>" : <i>String</i>,
    "<a href="#threshold" title="Threshold">Threshold</a>" : <i>Double</i>,
    "<a href="#units" title="Units">Units</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#metricname" title="MetricName">MetricName</a>: <i>String</i>
<a href="#mode" title="Mode">Mode</a>: <i>String</i>
<a href="#operator" title="Operator">Operator</a>: <i>String</i>
<a href="#threshold" title="Threshold">Threshold</a>: <i>Double</i>
<a href="#units" title="Units">Units</a>: <i>String</i>
</pre>

## Properties

#### MetricName

Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.

_Required_: No

_Type_: String

_Allowed Values_: <code>ASSERT_MSG</code> | <code>ASSERT_REGULAR</code> | <code>ASSERT_USER</code> | <code>ASSERT_WARNING</code> | <code>AVG_COMMAND_EXECUTION_TIME</code> | <code>AVG_READ_EXECUTION_TIME</code> | <code>AVG_WRITE_EXECUTION_TIME</code> | <code>BACKGROUND_FLUSH_AVG</code> | <code>CACHE_BYTES_READ_INTO</code> | <code>CACHE_BYTES_WRITTEN_FROM</code> | <code>CACHE_USAGE_DIRTY</code> | <code>CACHE_USAGE_USED</code> | <code>COMPUTED_MEMORY</code> | <code>CONNECTIONS</code> | <code>CONNECTIONS_MAX</code> | <code>CONNECTIONS_PERCENT</code> | <code>CURSORS_TOTAL_CLIENT_CURSORS_SIZE</code> | <code>CURSORS_TOTAL_OPEN</code> | <code>CURSORS_TOTAL_TIMED_OUT</code> | <code>DB_DATA_SIZE_TOTAL</code> | <code>DB_INDEX_SIZE_TOTAL</code> | <code>DB_STORAGE_TOTAL</code> | <code>DISK_PARTITION_SPACE_USED_DATA</code> | <code>DISK_PARTITION_SPACE_USED_INDEX</code> | <code>DISK_PARTITION_SPACE_USED_JOURNAL</code> | <code>DISK_PARTITION_UTILIZATION_DATA</code> | <code>DISK_PARTITION_UTILIZATION_INDEX</code> | <code>DISK_PARTITION_UTILIZATION_JOURNAL</code> | <code>DOCUMENT_DELETED</code> | <code>DOCUMENT_INSERTED</code> | <code>DOCUMENT_RETURNED</code> | <code>DOCUMENT_UPDATED</code> | <code>EXTRA_INFO_PAGE_FAULTS</code> | <code>FTS_DISK_UTILIZATION</code> | <code>FTS_MEMORY_MAPPED</code> | <code>FTS_MEMORY_RESIDENT</code> | <code>FTS_MEMORY_VIRTUAL</code> | <code>FTS_PROCESS_CPU_KERNEL</code> | <code>FTS_PROCESS_CPU_USER</code> | <code>GLOBAL_ACCESSES_NOT_IN_MEMORY</code> | <code>GLOBAL_LOCK_CURRENT_QUEUE_READERS</code> | <code>GLOBAL_LOCK_CURRENT_QUEUE_TOTAL</code> | <code>GLOBAL_LOCK_CURRENT_QUEUE_WRITERS</code> | <code>GLOBAL_LOCK_PERCENTAGE</code> | <code>GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN</code> | <code>INDEX_COUNTERS_BTREE_ACCESSES</code> | <code>INDEX_COUNTERS_BTREE_HITS</code> | <code>INDEX_COUNTERS_BTREE_MISS_RATIO</code> | <code>INDEX_COUNTERS_BTREE_MISSES</code> | <code>JOURNALING_COMMITS_IN_WRITE_LOCK</code> | <code>JOURNALING_MB</code> | <code>JOURNALING_WRITE_DATA_FILES_MB</code> | <code>LOGICAL_SIZE</code> | <code>MEMORY_MAPPED</code> | <code>MEMORY_RESIDENT</code> | <code>MEMORY_VIRTUAL</code> | <code>MUNIN_CPU_IOWAIT</code> | <code>MUNIN_CPU_IRQ</code> | <code>MUNIN_CPU_NICE</code> | <code>MUNIN_CPU_SOFTIRQ</code> | <code>MUNIN_CPU_STEAL</code> | <code>MUNIN_CPU_SYSTEM</code> | <code>MUNIN_CPU_USER</code> | <code>NETWORK_BYTES_IN</code> | <code>NETWORK_BYTES_OUT</code> | <code>NETWORK_NUM_REQUESTS</code> | <code>NORMALIZED_FTS_PROCESS_CPU_KERNEL</code> | <code>NORMALIZED_FTS_PROCESS_CPU_USER</code> | <code>NORMALIZED_SYSTEM_CPU_STEAL</code> | <code>NORMALIZED_SYSTEM_CPU_USER</code> | <code>OPCOUNTER_CMD</code> | <code>OPCOUNTER_DELETE</code> | <code>OPCOUNTER_GETMORE</code> | <code>OPCOUNTER_INSERT</code> | <code>OPCOUNTER_QUERY</code> | <code>OPCOUNTER_REPL_CMD</code> | <code>OPCOUNTER_REPL_DELETE</code> | <code>OPCOUNTER_REPL_INSERT</code> | <code>OPCOUNTER_REPL_UPDATE</code> | <code>OPCOUNTER_UPDATE</code> | <code>OPERATIONS_SCAN_AND_ORDER</code> | <code>OPLOG_MASTER_LAG_TIME_DIFF</code> | <code>OPLOG_MASTER_TIME</code> | <code>OPLOG_MASTER_TIME_ESTIMATED_TTL</code> | <code>OPLOG_RATE_GB_PER_HOUR</code> | <code>OPLOG_SLAVE_LAG_MASTER_TIME</code> | <code>QUERY_EXECUTOR_SCANNED</code> | <code>QUERY_EXECUTOR_SCANNED_OBJECTS</code> | <code>QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED</code> | <code>QUERY_TARGETING_SCANNED_PER_RETURNED</code> | <code>RESTARTS_IN_LAST_HOUR</code> | <code>SERVERLESS_CONNECTIONS</code> | <code>SERVERLESS_CONNECTIONS_PERCENT</code> | <code>SERVERLESS_DATA_SIZE_TOTAL</code> | <code>SERVERLESS_NETWORK_BYTES_IN</code> | <code>SERVERLESS_NETWORK_BYTES_OUT</code> | <code>SERVERLESS_NETWORK_NUM_REQUESTS</code> | <code>SERVERLESS_OPCOUNTER_CMD</code> | <code>SERVERLESS_OPCOUNTER_DELETE</code> | <code>SERVERLESS_OPCOUNTER_GETMORE</code> | <code>SERVERLESS_OPCOUNTER_INSERT</code> | <code>SERVERLESS_OPCOUNTER_QUERY</code> | <code>SERVERLESS_OPCOUNTER_UPDATE</code> | <code>SERVERLESS_TOTAL_READ_UNITS</code> | <code>SERVERLESS_TOTAL_WRITE_UNITS</code> | <code>SWAP_USAGE_FREE</code> | <code>SWAP_USAGE_USED</code> | <code>SYSTEM_MEMORY_AVAILABLE</code> | <code>SYSTEM_MEMORY_FREE</code> | <code>SYSTEM_MEMORY_USED</code> | <code>SYSTEM_NETWORK_IN</code> | <code>SYSTEM_NETWORK_OUT</code> | <code>TICKETS_AVAILABLE_READS</code> | <code>TICKETS_AVAILABLE_WRITES</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Mode

MongoDB Cloud computes the current metric value as an average.

_Required_: No

_Type_: String

_Allowed Values_: <code>AVERAGE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Operator

Comparison operator to apply when checking the current metric value.

_Required_: No

_Type_: String

_Allowed Values_: <code>GREATER_THAN</code> | <code>LESS_THAN</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Threshold

Value of metric that, when exceeded, triggers an alert.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Units

Element used to express the quantity. This can be an element of time, storage capacity, and the like.

_Required_: No

_Type_: String

_Allowed Values_: <code>BITS</code> | <code>BYTES</code> | <code>DAYS</code> | <code>GIGABITS</code> | <code>GIGABYTES</code> | <code>HOURS</code> | <code>KILOBITS</code> | <code>KILOBYTES</code> | <code>MEGABITS</code> | <code>MEGABYTES</code> | <code>MILLISECONDS</code> | <code>MINUTES</code> | <code>PETABYTES</code> | <code>RAW</code> | <code>SECONDS</code> | <code>TERABYTES</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

