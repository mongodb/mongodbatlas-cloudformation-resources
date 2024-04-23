# MongoDB::Atlas::Cluster

## Description

Resource for managing [Clusters](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters).

**WARNING:** `labels` attribute  is deprecated and will be removed in the future, use the `tags` attribute instead. See https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/cluster/docs/tag.md 

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/cluster/cluster.json) for example resource.
