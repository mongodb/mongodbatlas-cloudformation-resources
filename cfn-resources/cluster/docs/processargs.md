# MongoDB::Atlas::Cluster processArgs

Advanced configuration details to add for one cluster in the specified project.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#defaultreadconcern" title="DefaultReadConcern">DefaultReadConcern</a>" : <i>String</i>,
    "<a href="#defaultwriteconcern" title="DefaultWriteConcern">DefaultWriteConcern</a>" : <i>String</i>,
    "<a href="#failindexkeytoolong" title="FailIndexKeyTooLong">FailIndexKeyTooLong</a>" : <i>Boolean</i>,
    "<a href="#javascriptenabled" title="JavascriptEnabled">JavascriptEnabled</a>" : <i>Boolean</i>,
    "<a href="#minimumenabledtlsprotocol" title="MinimumEnabledTLSProtocol">MinimumEnabledTLSProtocol</a>" : <i>String</i>,
    "<a href="#tlscipherconfigmode" title="TlsCipherConfigMode">TlsCipherConfigMode</a>" : <i>String</i>,
    "<a href="#customopensslcipherconfigtls12" title="CustomOpensslCipherConfigTls12">CustomOpensslCipherConfigTls12</a>" : <i>[ String, ... ]</i>,
    "<a href="#notablescan" title="NoTableScan">NoTableScan</a>" : <i>Boolean</i>,
    "<a href="#oplogsizemb" title="OplogSizeMB">OplogSizeMB</a>" : <i>Integer</i>,
    "<a href="#samplesizebiconnector" title="SampleSizeBIConnector">SampleSizeBIConnector</a>" : <i>Integer</i>,
    "<a href="#samplerefreshintervalbiconnector" title="SampleRefreshIntervalBIConnector">SampleRefreshIntervalBIConnector</a>" : <i>Integer</i>,
    "<a href="#oplogminretentionhours" title="OplogMinRetentionHours">OplogMinRetentionHours</a>" : <i>Double</i>,
    "<a href="#transactionlifetimelimitseconds" title="TransactionLifetimeLimitSeconds">TransactionLifetimeLimitSeconds</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#defaultreadconcern" title="DefaultReadConcern">DefaultReadConcern</a>: <i>String</i>
<a href="#defaultwriteconcern" title="DefaultWriteConcern">DefaultWriteConcern</a>: <i>String</i>
<a href="#failindexkeytoolong" title="FailIndexKeyTooLong">FailIndexKeyTooLong</a>: <i>Boolean</i>
<a href="#javascriptenabled" title="JavascriptEnabled">JavascriptEnabled</a>: <i>Boolean</i>
<a href="#minimumenabledtlsprotocol" title="MinimumEnabledTLSProtocol">MinimumEnabledTLSProtocol</a>: <i>String</i>
<a href="#tlscipherconfigmode" title="TlsCipherConfigMode">TlsCipherConfigMode</a>: <i>String</i>
<a href="#customopensslcipherconfigtls12" title="CustomOpensslCipherConfigTls12">CustomOpensslCipherConfigTls12</a>: <i>
      - String</i>
<a href="#notablescan" title="NoTableScan">NoTableScan</a>: <i>Boolean</i>
<a href="#oplogsizemb" title="OplogSizeMB">OplogSizeMB</a>: <i>Integer</i>
<a href="#samplesizebiconnector" title="SampleSizeBIConnector">SampleSizeBIConnector</a>: <i>Integer</i>
<a href="#samplerefreshintervalbiconnector" title="SampleRefreshIntervalBIConnector">SampleRefreshIntervalBIConnector</a>: <i>Integer</i>
<a href="#oplogminretentionhours" title="OplogMinRetentionHours">OplogMinRetentionHours</a>: <i>Double</i>
<a href="#transactionlifetimelimitseconds" title="TransactionLifetimeLimitSeconds">TransactionLifetimeLimitSeconds</a>: <i>Integer</i>
</pre>

## Properties

#### DefaultReadConcern

Default level of acknowledgment requested from MongoDB for read operations set for this cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DefaultWriteConcern

Default level of acknowledgment requested from MongoDB for write operations set for this cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FailIndexKeyTooLong

Flag that indicates whether you can insert or update documents where all indexed entries don't exceed 1024 bytes. If you set this to false, mongod writes documents that exceed this limit but doesn't index them.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### JavascriptEnabled

Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MinimumEnabledTLSProtocol

Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TlsCipherConfigMode

The TLS cipher suite configuration mode. Valid values include `CUSTOM` or `DEFAULT`. The `DEFAULT` mode uses the default cipher suites. The `CUSTOM` mode allows you to specify custom cipher suites for both TLS 1.2 and TLS 1.3. To unset, this should be set back to `DEFAULT`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomOpensslCipherConfigTls12

The custom OpenSSL cipher suite list for TLS 1.2. This field is only valid when `tls_cipher_config_mode` is set to `CUSTOM`.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NoTableScan

Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OplogSizeMB

Storage limit of cluster's oplog expressed in megabytes. A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SampleSizeBIConnector

Interval in seconds at which the mongosqld process re-samples data to create its relational schema.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SampleRefreshIntervalBIConnector

Number of documents per database to sample when gathering schema information.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OplogMinRetentionHours

Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TransactionLifetimeLimitSeconds

Lifetime, in seconds, of multi-document transactions. Atlas considers the transactions that exceed this limit as expired and so aborts them through a periodic cleanup process.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

