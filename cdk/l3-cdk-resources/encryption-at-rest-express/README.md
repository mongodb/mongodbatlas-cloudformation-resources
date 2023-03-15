# @mongodbatlas-awscdk/encryption-at-rest-express

The official [MongoDB Atlas](https://www.mongodb.com/) AWS CDK resource for Node.js.

> This construct uses MongoDB [L1 construct](https://constructs.dev/search?q=&offset=0&tags=mongodb-published) and data structures for the [AWS CloudFormation Registry] L3 type.

[L1 construct]: https://docs.aws.amazon.com/cdk/latest/guide/constructs.html
[AWS CloudFormation Registry]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html

## Description

The resource lets you create a cluster, dbuser, add an entry to the atlas access list and set encryption-at-rest in your MongoDB Atlas Project.



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
  --type-name MongoDB::Atlas::EncryptionAtRest \
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
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-EncryptionAtRest \
  --execution-role-arn ROLE-ARN  

aws cloudformation activate-type \
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-DatabaseUser \
  --execution-role-arn ROLE-ARN

aws cloudformation activate-type \
  --public-type-arn arn:aws:cloudformation:us-east-1::type/resource/bb989456c78c398a858fef18f2ca1bfc1fbba082/MongoDB-Atlas-ProjectIpAccessList \
  --execution-role-arn ROLE-ARN
    
```

### Impotant Notes

1. Atlas encrypts all cluster storage and snapshot volumes, securing all cluster data on disk: a concept known as encryption at rest, by default.

2. Atlas limits this feature to dedicated cluster tiers of M10 and greater. For more information see: https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Encryption-at-Rest-using-Customer-Key-Management

3. This construnctor supports only [AWS Key Management Service](https://www.mongodb.com/docs/atlas/security-aws-kms/#security-aws-kms) 

4. MongoDB does not support creating and configuring a IAM role via CFN and CDK. **Please follow this guide ([Enable Role-Based Access to Your Encryption Key for a Project](https://www.mongodb.com/docs/atlas/security-aws-kms/#enable-customer-managed-keys-with-aws-kms)) before trying to use this constructor as you must provide** `roleId` **and the** `customerMasterKeyId`.

### Minimal configuration to use this app
In this example, @mongodbatlas-awscdk/encryption-at-rest-express is used to enable encryption at rest for your MongoDB Atlas Project.

```typescript
import * as cdk from 'aws-cdk-lib';
import { AtlasEncryptionAtRestExpress } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-encryption-at-rest-express', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const PROJECT_ID = stack.node.tryGetContext('MONGODB_ATLAS_PROJECT_ID') || process.env.MONGODB_ATLAS_PROJECT_ID;
const ROLE_ID = stack.node.tryGetContext('MONGODB_ATLAS_ROLE_ID') || process.env.MONGODB_ATLAS_ROLE_ID;
const MASTER_KEY_ID = stack.node.tryGetContext('MONGODB_ATLAS_MASTER_KEY_ID') || process.env.MONGODB_ATLAS_MASTER_KEY_ID;

new AtlasEncryptionAtRestExpress(stack, 'atlas-encryption-at-rest-express', {
  projectId: PROJECT_ID,

  encryptionAtRest: {
    roleId: ROLE_ID,
    customerMasterKeyId: MASTER_KEY_ID,
  },
});
```

### Minimal configuration to create a cluster, database User, configure atlas access list and enable encryption at rest

In this example we are providing the minimal configuration to the constructor to define an Atlas Cluster, DB user, add an entry to the Atlas Access list, and enable encryption at rest for the cluster and project.

```typescript
import * as cdk from 'aws-cdk-lib';
import { AtlasEncryptionAtRestExpress } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-encryption-at-rest-express', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const PROJECT_ID = stack.node.tryGetContext('MONGODB_ATLAS_PROJECT_ID') || process.env.MONGODB_ATLAS_PROJECT_ID;
const ROLE_ID = stack.node.tryGetContext('MONGODB_ATLAS_ROLE_ID') || process.env.MONGODB_ATLAS_ROLE_ID;
const MASTER_KEY_ID = stack.node.tryGetContext('MONGODB_ATLAS_MASTER_KEY_ID') || process.env.MONGODB_ATLAS_MASTER_KEY_ID;
const DB_USER_PASSWORD = stack.node.tryGetContext('MONGODB_ATLAS_DB_USER_PASSWORD') || process.env.MONGODB_ATLAS_DB_USER_PASSWORD;

new AtlasEncryptionAtRestExpress(this, 'EncryptionAtRestExpress', {
      cluster: {
        name: "ClusterName",
      },

      accessList: {
        accessList: [{
          ipAddress: "192.0.0.1"
        }]
      },

      encryptionAtRest: {
        customerMasterKeyId: MONGODB_ATLAS_MASTER_KEY_ID,
        roleId: ROLE_ID,
      },

      databaseUser: {
        password: DB_USER_PASSWORD
      },

      projectId: PROJECT_ID,
    })
  }
```
The constructor will deploy the following resources using default values:

- Cluster
    - name: [provided name]
    - provider: AWS
    - region: us-east-1
    - cluster type: replica-set 
    - elegible nodes: 3
    - instance size: M30
    - analytics node: 1
    - analytics instance size: M30
    - backup: enabled
    - encryption at rest provider: AWS

- DataBase User
    - username: cdkUser
    - password: [provided password]

- Atlas Access list
    - ip address: [provided ip address]

### Configuration to create a cluster, database User, configure atlas access list and enable encryption at rest
In this example we override the default values for the atlas cluster and creates a replica-set with 6 nodes, instance size M10 in the AWS region US_WEST_2. 
```typescript
import * as cdk from 'aws-cdk-lib';
import { AtlasEncryptionAtRestExpress } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-encryption-at-rest-express', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const PROJECT_ID = stack.node.tryGetContext('MONGODB_ATLAS_PROJECT_ID') || process.env.MONGODB_ATLAS_PROJECT_ID;
const ROLE_ID = stack.node.tryGetContext('MONGODB_ATLAS_ROLE_ID') || process.env.MONGODB_ATLAS_ROLE_ID;
const MASTER_KEY_ID = stack.node.tryGetContext('MONGODB_ATLAS_MASTER_KEY_ID') || process.env.MONGODB_ATLAS_MASTER_KEY_ID;
const DB_USER_PASSWORD = stack.node.tryGetContext('MONGODB_ATLAS_DB_USER_PASSWORD') || process.env.MONGODB_ATLAS_DB_USER_PASSWORD;

new AtlasEncryptionAtRestExpress(this, 'EncryptionAtRestExpress', {
      cluster: {
        name: "Cluster0",
        backupEnabled: false,
        mongoDbMajorVersion: "6.0",
        replicationSpecs: [
          {
            numShards: 1,
            advancedRegionConfigs: [
              {
                regionName: "US_WEAST_2",
                electableSpecs: {
                  instanceSize: "M10",
                  nodeCount: 6,
                  ebsVolumeType: "STANDARD"
                },
                priority: 7
            }
          ]}
        ]
      },

      accessList: {
        accessList: [{
          ipAddress: "192.0.0.1"
        }]
      },

      encryptionAtRest: {
        customerMasterKeyId: MONGODB_ATLAS_MASTER_KEY_ID,
        roleId: ROLE_ID,
      },

      databaseUser: {
        password: DB_USER_PASSWORD
      },

      projectId: PROJECT_ID,
    })
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
