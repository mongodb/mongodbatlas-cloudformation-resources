import * as atlas from '@mongodbatlas-awscdk/cluster';
import * as user from '@mongodbatlas-awscdk/database-user';
import * as project from '@mongodbatlas-awscdk/project';
import * as ipAccessList from '@mongodbatlas-awscdk/project-ip-access-list';
import { App, Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
// import * as privateEndpoint from '@mongodbatlas-awscdk/private-endpoint';
import * as request from './request.json';

export class ClusterL3 extends Stack {
  constructor(scope: Construct, id: string, props: StackProps = {}) {
    super(scope, id, props);

    // Create a new MongoDB Atlas Project
    const myProject = new project.CfnProject(this, 'MyProject', {
      name: request.projectName,
      orgId: request.orgId,
      apiKeys: {
        publicKey: request.publicKey,
        privateKey: request.privateKey,
      },
    });
    // Create a new MongoDB Atlas Cluster and pass project ID
    new atlas.CfnCluster(this, 'MyCluster', {
      apiKeys: {
        publicKey: request.publicKey,
        privateKey: request.privateKey,
      },
      name: request.clusterName,
      projectId: myProject.ref,
      clusterType: request.clusterType,
      backupEnabled: request.backupEnabled,
      replicationSpecs: request.replicationSpecs,
    });

    // Create a new MongoDB Atlas Database User
    new user.CfnDatabaseUser(this, 'MyUser', {
      databaseName: request.databaseName,
      projectId: myProject.ref,
      roles: request.roles,
      username: request.username,
      password: request.password,
      apiKeys: {
        publicKey: request.publicKey,
        privateKey: request.privateKey,
      },
    });

    // Create a new MongoDB Atlas Project IP Access List
    new ipAccessList.CfnProjectIpAccessList(this, 'MyIpAccessList', {
      accessList: request.accessList,
      apiKeys: {
        publicKey: request.publicKey,
        privateKey: request.privateKey,
      },
      projectId: myProject.ref,
    });

    // new privateEndpoint.CfnPrivateEndpoint(this, 'MyPrivateEndpoint',{
    //     groupId: myProject.ref,
    //     region: "us-east-1",
    //     endpointServiceName:"",
    //     errorMessage: "L3 CDK : Private EndPoint creation failed.",
    //     privateEndpoints: [],
    //     status: "",
    //     apiKeys: {
    //       publicKey:"",
    //       privateKey:""
    //     },
    // });
  }
}

// for development, use account/region from cdk cli
const devEnv = {
  account: '', //process.env.CDK_DEFAULT_ACCOUNT,
  region: '', //process.env.CDK_DEFAULT_REGION,
};

const app = new App();

new ClusterL3(app, request.name, { env: devEnv });

app.synth();