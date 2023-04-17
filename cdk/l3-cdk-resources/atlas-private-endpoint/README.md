# @mongodbatlas-awscdk/atlas-basic-private-endpoint
## This package is now deprecated and will no longer receive new features. Use [awscdk-resource-mongodbatlas](https://constructs.dev/packages/awscdk-resources-mongodbatlas) instead.
The official [MongoDB Atlas](https://www.mongodb.com/) AWS CDK resource for Node.js.

> This construct uses MongoDB [Level 1 construct](https://constructs.dev/search?q=&offset=0&tags=mongodb-published) and data structures for the [AWS CloudFormation Registry] Level 3 type.

[L1 construct]: https://docs.aws.amazon.com/cdk/latest/guide/constructs.html
[AWS CloudFormation Registry]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html

## Description

Creates a MongoDB Atlas Project, Cluster, DBuser, Private Endpoint and adds an IP entry to the IP Access List.


## MongoDB Atlas API Docs

For more information for each of the specific APIs used in this Level 3 AWS CDK resource please refer to: [Official Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec)

## Usage

In order to use this library, you will need to activate this AWS CloudFormation Registry type in your account. You can do this via the AWS Management Console or using the [AWS CLI](https://aws.amazon.com/cli/) using the following command:

```sh
aws cloudformation activate-type \
  --type-name MongoDB::Atlas::Cluster \
  --publisher-id bb989456c78c398a858fef18f2ca1bfc1fbba082 \
  --type RESOURCE \
  --execution-role-arn ROLE-ARN
  
aws cloudformation activate-type \
  --type-name MongoDB::Atlas::Project \
  --publisher-id bb989456c78c398a858fef18f2ca1bfc1fbba082 \
  --type RESOURCE \
  --execution-role-arn ROLE-ARN

aws cloudformation activate-type \
  --type-name MongoDB::Atlas::DatabaseUser \
  --publisher-id bb989456c78c398a858fef18f2ca1bfc1fbba082 \
  --type RESOURCE \
  --execution-role-arn ROLE-ARN  

aws cloudformation activate-type \
  --type-name MongoDB::Atlas::ProjectIpAccessList \
  --publisher-id bb989456c78c398a858fef18f2ca1bfc1fbba082 \
  --type RESOURCE \
  --execution-role-arn ROLE-ARN  
  
aws cloudformation activate-type \
  --type-name MongoDB::Atlas::PrivateEndpoint \
  --publisher-id bb989456c78c398a858fef18f2ca1bfc1fbba082 \
  --type RESOURCE \
  --execution-role-arn ROLE-ARN
```


### Minimal configuration to use this construct

```typescript
import * as cdk from 'aws-cdk-lib';
import { AtlasBasicProps } from "@mongodbatlas-awscdk/atlas-basic";
import {AtlasBasicPrivateEndpoint, PrivateEndpointProps} from "./index";

const app = new cdk.App();

const stack = new cdk.Stack(app, 'atlas-basic-default', {
    env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const orgId = stack.node.tryGetContext('MONGODB_ATLAS_ORG_ID') || process.env.MONGODB_ATLAS_ORG_ID;
const vpcId = stack.node.tryGetContext('AWS_VPC_ID') || process.env.AWS_VPC_ID;
const subnetId = stack.node.tryGetContext('AWS_SUBNET_ID') || process.env.AWS_SUBNET_ID;
const awsRegion = stack.node.tryGetContext('AWS_REGION') || process.env.AWS_REGION;

const replicationSpecs = [
    {
        numShards: 1,
        advancedRegionConfigs: [
            {
                analyticsSpecs: {
                    ebsVolumeType: 'STANDARD',
                    instanceSize: 'M10',
                    nodeCount: 1,
                },
                electableSpecs: {
                    ebsVolumeType: 'STANDARD',
                    instanceSize: 'M10',
                    nodeCount: 3,
                },
                priority: 7,
                regionName: 'US_EAST_1',
            },
        ],
    },
];

const atlasBasicProps : AtlasBasicProps = {
    clusterProps: {
        replicationSpecs: replicationSpecs,
    },
    projectProps: {
        orgId: orgId,
    },
    ipAccessListProps : {
        accessList: [
            {
                ipAddress: '10.10.0.0/24',
                comment: 'Open Subnets',
            },
        ],
    }
}

const privateEndpointProps : PrivateEndpointProps = {
    privateEndpoints: [{
        vpcId: vpcId,
        subnetIds: [subnetId]
    }],
}

const props   = {
    apiKeys: apiKeys,
    atlasBasicProps: atlasBasicProps,
    privateEndpointProps: privateEndpointProps,
    region: awsRegion
}

new AtlasBasicPrivateEndpoint(stack,'private-endpoint', props)
```
The library also defines some default values for individual L1s.

```typescript

const projectDefaults = {
        projectName: 'atlas-project-{random_num}',
    };

const dbDefaults = {
    dbName: 'admin',
    username: 'atlas-user',
    password: 'atlas-pwd',
    roles: [{
        roleName: 'atlasAdmin',
        databaseName: 'admin',
    }],
}
const clusterDefaults = {
    clusterName: 'atlas-cluster-{random_num}',
    clusterType: 'REPLICASET',
}
```

Default Region is set to us-east-1 region (AWS US East N. Virginia)

You can find more information about activating this type in the [AWS CloudFormation documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html).

## Feedback

* Issues related to this generated library should be [reported here](https://github.com/cdklabs/cdk-cloudformation/issues/new?title=Issue+with+%40cdk-cloudformation%2Fmongodb-atlas-cluster+v1.0.0).
* Issues related to this library should be reported to the [publisher](https://github.com/mongodb/mongodbatlas-cloudformation-resources/issues).
* Feature requests should be [reported here](https://feedback.mongodb.com/forums/924145-atlas?category_id=392596)

[cdklabs/cdk-cloudformation]: https://github.com/cdklabs/cdk-cloudformation

## License

Distributed under the Apache-2.0 License.
