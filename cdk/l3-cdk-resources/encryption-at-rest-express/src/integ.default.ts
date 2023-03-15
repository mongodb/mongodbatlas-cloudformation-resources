import * as cdk from 'aws-cdk-lib';
import { AtlasEncryptionAtRestExpress } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-encryption-at-rest-express', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const PROJECT_ID = stack.node.tryGetContext('MONGODB_ATLAS_PROJECT_ID') || process.env.MONGODB_ATLAS_PROJECT_ID;
const ROLE_ID = stack.node.tryGetContext('MONGODB_ATLAS_ROLE_ID') || process.env.MONGODB_ATLAS_ROLE_ID;
const MASTER_KEY_ID = stack.node.tryGetContext('MONGODB_ATLAS_MASTER_KEY_ID') || process.env.MONGODB_ATLAS_MASTER_KEY_ID;

export const integrationDefault = new AtlasEncryptionAtRestExpress(stack, 'atlas-encryption-at-rest-express', {
  projectId: PROJECT_ID,

  encryptionAtRest: {
    roleId: ROLE_ID,
    customerMasterKeyId: MASTER_KEY_ID,
  },
});
