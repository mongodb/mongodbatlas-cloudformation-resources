import * as cdk from 'aws-cdk-lib';
import { AtlasBasic } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-basic-default', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const orgId = stack.node.tryGetContext('MONGODB_ATLAS_ORG_ID') || process.env.MONGODB_ATLAS_ORG_ID;

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

new AtlasBasic(stack, 'atlas-basic', {
  clusterProps: {
    replicationSpecs: replicationSpecs,
  },
  projectProps: {
    orgId: orgId,
  },
});
