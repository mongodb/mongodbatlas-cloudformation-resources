# MongoDB::Atlas::AlertConfiguration AlertView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#acknowledgeduntil" title="AcknowledgedUntil">AcknowledgedUntil</a>" : <i>String</i>,
    "<a href="#acknowledgementcomment" title="AcknowledgementComment">AcknowledgementComment</a>" : <i>String</i>,
    "<a href="#acknowledgingusername" title="AcknowledgingUsername">AcknowledgingUsername</a>" : <i>String</i>,
    "<a href="#alertconfigid" title="AlertConfigId">AlertConfigId</a>" : <i>String</i>,
    "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
    "<a href="#created" title="Created">Created</a>" : <i>String</i>,
    "<a href="#currentvalue" title="CurrentValue">CurrentValue</a>" : <i><a href="currentvalue.md">CurrentValue</a></i>,
    "<a href="#eventtypename" title="EventTypeName">EventTypeName</a>" : <i>String</i>,
    "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
    "<a href="#hostnameandport" title="HostnameAndPort">HostnameAndPort</a>" : <i>String</i>,
    "<a href="#id" title="Id">Id</a>" : <i>String</i>,
    "<a href="#lastnotified" title="LastNotified">LastNotified</a>" : <i>String</i>,
    "<a href="#links" title="Links">Links</a>" : <i>[ <a href="link.md">Link</a>, ... ]</i>,
    "<a href="#metricname" title="MetricName">MetricName</a>" : <i>String</i>,
    "<a href="#replicasetname" title="ReplicaSetName">ReplicaSetName</a>" : <i>String</i>,
    "<a href="#resolved" title="Resolved">Resolved</a>" : <i>String</i>,
    "<a href="#status" title="Status">Status</a>" : <i>String</i>,
    "<a href="#typename" title="TypeName">TypeName</a>" : <i>String</i>,
    "<a href="#updated" title="Updated">Updated</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#acknowledgeduntil" title="AcknowledgedUntil">AcknowledgedUntil</a>: <i>String</i>
<a href="#acknowledgementcomment" title="AcknowledgementComment">AcknowledgementComment</a>: <i>String</i>
<a href="#acknowledgingusername" title="AcknowledgingUsername">AcknowledgingUsername</a>: <i>String</i>
<a href="#alertconfigid" title="AlertConfigId">AlertConfigId</a>: <i>String</i>
<a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
<a href="#created" title="Created">Created</a>: <i>String</i>
<a href="#currentvalue" title="CurrentValue">CurrentValue</a>: <i><a href="currentvalue.md">CurrentValue</a></i>
<a href="#eventtypename" title="EventTypeName">EventTypeName</a>: <i>String</i>
<a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
<a href="#hostnameandport" title="HostnameAndPort">HostnameAndPort</a>: <i>String</i>
<a href="#id" title="Id">Id</a>: <i>String</i>
<a href="#lastnotified" title="LastNotified">LastNotified</a>: <i>String</i>
<a href="#links" title="Links">Links</a>: <i>
      - <a href="link.md">Link</a></i>
<a href="#metricname" title="MetricName">MetricName</a>: <i>String</i>
<a href="#replicasetname" title="ReplicaSetName">ReplicaSetName</a>: <i>String</i>
<a href="#resolved" title="Resolved">Resolved</a>: <i>String</i>
<a href="#status" title="Status">Status</a>: <i>String</i>
<a href="#typename" title="TypeName">TypeName</a>: <i>String</i>
<a href="#updated" title="Updated">Updated</a>: <i>String</i>
</pre>

## Properties

#### AcknowledgedUntil

Date and time until which this alert has been acknowledged. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.

- To acknowledge this alert forever, set the parameter value to 100 years in the future.

- To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AcknowledgementComment

Comment that a MongoDB Cloud user submitted when acknowledging the alert.

_Required_: No

_Type_: String

_Maximum_: <code>200</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AcknowledgingUsername

MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AlertConfigId

Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters.

_Required_: No

_Type_: String

_Minimum_: <code>1</code>

_Maximum_: <code>64</code>

_Pattern_: <code>^([a-zA-Z0-9]([a-zA-Z0-9-]){0,21}(?<!-)([\w]{0,42}))$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Created

Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CurrentValue

_Required_: No

_Type_: <a href="currentvalue.md">CurrentValue</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EventTypeName

Incident that triggered this alert.

_Required_: No

_Type_: String

_Allowed Values_: <code>AWS_ENCRYPTION_KEY_NEEDS_ROTATION</code> | <code>AZURE_ENCRYPTION_KEY_NEEDS_ROTATION</code> | <code>CLUSTER_MONGOS_IS_MISSING</code> | <code>CPS_RESTORE_FAILED</code> | <code>CPS_RESTORE_SUCCESSFUL</code> | <code>CPS_SNAPSHOT_BEHIND</code> | <code>CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED</code> | <code>CPS_SNAPSHOT_FALLBACK_FAILED</code> | <code>CPS_SNAPSHOT_FALLBACK_SUCCESSFUL</code> | <code>CPS_SNAPSHOT_SUCCESSFUL</code> | <code>CREDIT_CARD_ABOUT_TO_EXPIRE</code> | <code>DAILY_BILL_OVER_THRESHOLD</code> | <code>GCP_ENCRYPTION_KEY_NEEDS_ROTATION</code> | <code>HOST_DOWN</code> | <code>JOINED_GROUP</code> | <code>NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK</code> | <code>NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK</code> | <code>NO_PRIMARY</code> | <code>OUTSIDE_METRIC_THRESHOLD</code> | <code>OUTSIDE_SERVERLESS_METRIC_THRESHOLD</code> | <code>PENDING_INVOICE_OVER_THRESHOLD</code> | <code>PRIMARY_ELECTED</code> | <code>REMOVED_FROM_GROUP</code> | <code>REPLICATION_OPLOG_WINDOW_RUNNING_OUT</code> | <code>TOO_MANY_ELECTIONS</code> | <code>USERS_WITHOUT_MULTIFACTOR_AUTH</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique 24-hexadecimal digit string that identifies the project that owns this alert.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### HostnameAndPort

Hostname and port of the host to which this alert applies. The resource returns this parameter for alerts of events impacting hosts or replica sets.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Id

Unique 24-hexadecimal digit string that identifies this alert.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### LastNotified

Date and time that any notifications were last sent for this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Links

_Required_: No

_Type_: List of <a href="link.md">Link</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MetricName

Human-readable label that identifies the metric against which MongoDB Cloud checks the alert.

_Required_: No

_Type_: String

_Allowed Values_: <code>ASSERT_MSG</code> | <code>ASSERT_REGULAR</code> | <code>ASSERT_USER</code> | <code>ASSERT_WARNING</code> | <code>AVG_COMMAND_EXECUTION_TIME</code> | <code>AVG_READ_EXECUTION_TIME</code> | <code>AVG_WRITE_EXECUTION_TIME</code> | <code>CACHE_BYTES_READ_INTO</code> | <code>CACHE_BYTES_WRITTEN_FROM</code> | <code>CACHE_DIRTY_BYTES</code> | <code>CACHE_USED_BYTES</code> | <code>COMPUTED_MEMORY</code> | <code>CONNECTIONS</code> | <code>CONNECTIONS_PERCENT</code> | <code>CURSORS_TOTAL_OPEN</code> | <code>CURSORS_TOTAL_TIMED_OUT</code> | <code>DB_DATA_SIZE_TOTAL</code> | <code>DB_INDEX_SIZE_TOTAL</code> | <code>DB_STORAGE_TOTAL</code> | <code>DISK_PARTITION_SPACE_USED_DATA</code> | <code>DISK_PARTITION_SPACE_USED_INDEX</code> | <code>DISK_PARTITION_SPACE_USED_JOURNAL</code> | <code>DISK_PARTITION_UTILIZATION_DATA</code> | <code>DISK_PARTITION_UTILIZATION_INDEX</code> | <code>DISK_PARTITION_UTILIZATION_JOURNAL</code> | <code>DOCUMENT_DELETED</code> | <code>DOCUMENT_INSERTED</code> | <code>DOCUMENT_RETURNED</code> | <code>DOCUMENT_UPDATED</code> | <code>EXTRA_INFO_PAGE_FAULTS</code> | <code>FTS_MEMORY_RESIDENT</code> | <code>FTS_MEMORY_SHARED</code> | <code>FTS_MEMORY_VIRTUAL</code> | <code>FTS_PROCESS_CPU_KERNEL</code> | <code>FTS_PROCESS_CPU_USER</code> | <code>FTS_PROCESS_DISK</code> | <code>GLOBAL_LOCK_CURRENT_QUEUE_READERS</code> | <code>GLOBAL_LOCK_CURRENT_QUEUE_TOTAL</code> | <code>GLOBAL_LOCK_CURRENT_QUEUE_WRITERS</code> | <code>LOGICAL_SIZE</code> | <code>MEMORY_RESIDENT</code> | <code>MEMORY_VIRTUAL</code> | <code>NETWORK_BYTES_IN</code> | <code>NETWORK_BYTES_OUT</code> | <code>NETWORK_NUM_REQUESTS</code> | <code>NORMALIZED_FTS_PROCESS_CPU_KERNEL</code> | <code>NORMALIZED_FTS_PROCESS_CPU_USER</code> | <code>NORMALIZED_SYSTEM_CPU_STEAL</code> | <code>NORMALIZED_SYSTEM_CPU_USER</code> | <code>OPCOUNTER_CMD</code> | <code>OPCOUNTER_DELETE</code> | <code>OPCOUNTER_GETMORE</code> | <code>OPCOUNTER_INSERT</code> | <code>OPCOUNTER_QUERY</code> | <code>OPCOUNTER_REPL_CMD</code> | <code>OPCOUNTER_REPL_DELETE</code> | <code>OPCOUNTER_REPL_INSERT</code> | <code>OPCOUNTER_REPL_UPDATE</code> | <code>OPCOUNTER_UPDATE</code> | <code>OPERATIONS_SCAN_AND_ORDER</code> | <code>OPLOG_MASTER_LAG_TIME_DIFF</code> | <code>OPLOG_MASTER_TIME</code> | <code>OPLOG_RATE_GB_PER_HOUR</code> | <code>OPLOG_SLAVE_LAG_MASTER_TIME</code> | <code>QUERY_EXECUTOR_SCANNED</code> | <code>QUERY_EXECUTOR_SCANNED_OBJECTS</code> | <code>QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED</code> | <code>QUERY_TARGETING_SCANNED_PER_RETURNED</code> | <code>RESTARTS_IN_LAST_HOUR</code> | <code>SERVERLESS_CONNECTIONS</code> | <code>SERVERLESS_CONNECTIONS_PERCENT</code> | <code>SERVERLESS_DATA_SIZE_TOTAL</code> | <code>SERVERLESS_NETWORK_BYTES_IN</code> | <code>SERVERLESS_NETWORK_BYTES_OUT</code> | <code>SERVERLESS_NETWORK_NUM_REQUESTS</code> | <code>SERVERLESS_OPCOUNTER_CMD</code> | <code>SERVERLESS_OPCOUNTER_DELETE</code> | <code>SERVERLESS_OPCOUNTER_GETMORE</code> | <code>SERVERLESS_OPCOUNTER_INSERT</code> | <code>SERVERLESS_OPCOUNTER_QUERY</code> | <code>SERVERLESS_OPCOUNTER_UPDATE</code> | <code>SERVERLESS_TOTAL_READ_UNITS</code> | <code>SERVERLESS_TOTAL_WRITE_UNITS</code> | <code>TICKETS_AVAILABLE_READS</code> | <code>TICKETS_AVAILABLE_WRITES</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicaSetName

Name of the replica set to which this alert applies. The response returns this parameter for alerts of events impacting backups, hosts, or replica sets.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Resolved

Date and time that this alert changed to '"status" : "CLOSED"'. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter once '"status" : "CLOSED"'.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

State of this alert at the time you requested its details.

_Required_: No

_Type_: String

_Allowed Values_: <code>CANCELLED</code> | <code>CLOSED</code> | <code>OPEN</code> | <code>TRACKING</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TypeName

Category in which MongoDB Cloud classifies this alert.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Updated

Date and time when someone last updated this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

_Required_: No

_Type_: String

_Pattern_: <code>^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?(?:Z)$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

