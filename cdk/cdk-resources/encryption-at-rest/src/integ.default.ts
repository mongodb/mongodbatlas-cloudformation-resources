import * as cdk from 'aws-cdk-lib';
import { CfnEncryptionAtRest } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-EncAtRest-cdk-test', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const projectId = stack.node.tryGetContext('MONGODB_PROJECT_ID') || process.env.MONGODB_ATLAS_ORG_ID;
const customerMasterKeyId = stack.node.tryGetContext('CUSTOMER_MASTER_KEY_ID') || process.env.CUSTOMER_MASTER_KEY_ID;
const roleID = stack.node.tryGetContext('ROLE_ID') || process.env.ROLE_ID;

new CfnEncryptionAtRest(stack, 'encryptionAtRest', {
  projectId: projectId,
  awsKms: {
    roleId: roleID,
    customerMasterKeyId: customerMasterKeyId,
    enabled: true,
    region: 'AP_NORTHEAST_2',
  },
});