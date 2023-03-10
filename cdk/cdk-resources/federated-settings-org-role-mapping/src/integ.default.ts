import * as cdk from 'aws-cdk-lib';
import { CfnFederatedSettingsOrgRoleMapping } from './index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'atlas-fedOrgRoleMapping-default', {
  env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const orgId = stack.node.tryGetContext('MONGODB_ATLAS_ORG_ID') || process.env.MONGODB_ATLAS_ORG_ID;
const fedSettingsId = stack.node.tryGetContext('ATLAS_FEDERATED_SETTINGS_ID') || process.env.ATLAS_FEDERATED_SETTINGS_ID;

new CfnFederatedSettingsOrgRoleMapping(stack, 'federatedSettingsOrgRoleMapping', {
  profile: 'federation',
  externalGroupName: 'RG-01',
  orgId: orgId,
  federationSettingsId: fedSettingsId,
  roleAssignments: [
    {
      role: 'GROUP_DATA_ACCESS_READ_WRITE',
      projectId: '<project Id>',
    },
  ],

});
