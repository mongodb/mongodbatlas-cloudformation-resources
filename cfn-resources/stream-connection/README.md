# MongoDB::Atlas::StreamConnection

Resource for creating and managing [Connections for an Atlas Stream Instance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Streams/operation/createStreamConnection).


> **NOTE**: 
> - **Atlas Streams functionality is currently in [Public Preview](https://www.mongodb.com/blog/post/atlas-stream-processing-now-in-public-preview).** 
> - Please review [Limitations](https://www.mongodb.com/docs/atlas/atlas-sp/limitations/#std-label-atlas-sp-limitations) of Atlas Streams Processing during this preview period.
> - Please refer [our example section](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/atlas-streams/README.md) for further details.


## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md). Also refer [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.

## Cloudformation Examples

See the example [CFN Template](/examples/stream-connection/kafka-stream-connection.json) for example resource.
