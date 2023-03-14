# @mongodbatlas-awscdk/atlas-basic

The official [MongoDB Atlas](https://www.mongodb.com/) AWS CDK resource for Node.js.

> This construct uses MongoDB [L1 construct](https://constructs.dev/search?q=&offset=0&tags=mongodb-published) and data structures for the [AWS CloudFormation Registry] L3 type.

[L1 construct]: https://docs.aws.amazon.com/cdk/latest/guide/constructs.html
[AWS CloudFormation Registry]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html

## Description

The cluster basic resource provides access to your cluster configurations. 
The resource lets you create Project, Cluster, DB User and configures IpAccesslist. 



## MongoDB Atlas API Docs

For more information about the API refer to: [API Endpoints](https://www.mongodb.com/docs/atlas/reference/api-resources-spec)

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
    
```

Alternatively:

```sh
aws cloudformation activate-type \
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-Cluster \
  --execution-role-arn ROLE-ARN

aws cloudformation activate-type \
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-Project \
  --execution-role-arn ROLE-ARN  

aws cloudformation activate-type \
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-DatabaseUser \
  --execution-role-arn ROLE-ARN

aws cloudformation activate-type \
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-ProjectIpAccessList \
  --execution-role-arn ROLE-ARN
    
```

### Minimal configuration to use this app

```typescript
import * as cdk from 'aws-cdk-lib';
import { AtlasBasic, ApiKeyDefinition } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-basic-default', {
    env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});

const orgId = stack.node.tryGetContext('MONGODB_ATLAS_ORG_ID') || process.env.MONGODB_ATLAS_ORG_ID;

const replicationSpecs = [
    {
        "numShards": 1,
        "advancedRegionConfigs": [
            {
                "analyticsSpecs": {
                    "ebsVolumeType": "STANDARD",
                    "instanceSize": "M10",
                    "nodeCount": 1
                },
                "electableSpecs": {
                    "ebsVolumeType": "STANDARD",
                    "instanceSize": "M10",
                    "nodeCount": 3
                },
                "priority":  7,
                "regionName": "US_EAST_1"
            }
        ]
    }
]

new AtlasBasic(stack, 'atlas-basic', {
    clusterProps: {
        replicationSpecs : replicationSpecs
    },
    projectProps: {
        orgId: orgId,
    }
});
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

You can find more information about activating this type in the [AWS CloudFormation documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html).

## Feedback

* Issues related to this generated library should be [reported here](https://github.com/cdklabs/cdk-cloudformation/issues/new?title=Issue+with+%40cdk-cloudformation%2Fmongodb-atlas-cluster+v1.0.0).
* Issues related to this library should be reported to the [publisher](https://github.com/mongodb/mongodbatlas-cloudformation-resources/issues).
* Feature requests should be [reported here](https://feedback.mongodb.com/forums/924145-atlas?category_id=392596)

[cdklabs/cdk-cloudformation]: https://github.com/cdklabs/cdk-cloudformation

## License

Distributed under the Apache-2.0 License.# replace this
