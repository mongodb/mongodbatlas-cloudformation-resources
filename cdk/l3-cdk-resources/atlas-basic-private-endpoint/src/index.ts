import * as privateEndpoint from '@mongodbatlas-awscdk/private-endpoint';
import * as cdk from 'aws-cdk-lib';
import { App } from 'aws-cdk-lib';
import * as atlas_basic from 'test-atlas-client';

export interface AtlasBasicPropsPrivateEndpoint extends cdk.StackProps{
  readonly projectId : string;
  readonly vpcId : string;
  readonly subnetId : string;
  readonly region : string;
  readonly apiKeys : privateEndpoint.ApiKey;
}

class AtlasBasePrivateEndpoint extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props: AtlasBasicPropsPrivateEndpoint) {
    super(scope, id, props);

    new atlas_basic.AtlasBasic(this, 'QuickStartProject', this.getAtlasBasicInput(props));

    new privateEndpoint.CfnPrivateEndpoint(this, 'QuickStartProject',
      {
        groupId: props.projectId,
        region: props.region,
        apiKeys: { privateKey: props.apiKeys.privateKey, publicKey: props.apiKeys.publicKey },
        privateEndpoints: [{ vpcId: props.vpcId, subnetIds: [props.subnetId] }],
      });
  }

  getAtlasBasicInput(a: AtlasBasicPropsPrivateEndpoint) : atlas_basic.AtlasBasicProps {
    return {
      apiKeys: a.apiKeys,
      clusterProps: {
        apiKeys: a.apiKeys,
        projectId: a.projectId,
        name: 'Cluster_1',
      },
      dbUserProps: {
        databaseName: 'database_1',
        projectId: a.projectId,
        roles: [{
          roleName: 'atlasAdmin',
          databaseName: 'admin',
        }],
        username: 'testUser',
        apiKeys: a.apiKeys,
      },
      ipAccessListProps: {
        apiKeys: a.apiKeys,
        projectId: a.projectId,
        accessList: [
          {
            ipAddress: '0.0.0.0/1',
            comment: 'Testing open all ips',
          },
        ],
      },
      projectId: a.projectId,
    };
  }
}

const modal: AtlasBasicPropsPrivateEndpoint = {
  projectId: '',
  vpcId: 'vpc-',
  subnetId: 'subnet-',
  region: 'us-east-1',
  apiKeys: { privateKey: '', publicKey: '' },
};

const app = new App();

let greeter = new AtlasBasePrivateEndpoint(app, 'test-something', modal );

app.synth();