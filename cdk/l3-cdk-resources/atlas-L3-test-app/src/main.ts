
import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as atlas_basic from 'test-atlas-client';
import * as defaults from './defaults.json';
import {App} from "aws-cdk-lib";
import * as project from '@mongodbatlas-awscdk/project';


export class CdkMigrateFromCfnStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // if (env.ORG_ID == undefined) {
    //   throw 'ORG_ID is missing. Please, set the ORG_ID as an environment variable';
    // }
    //
    // if (env.PUBLIC_KEY == undefined) {
    //   throw 'PUBLIC_KEY is missing. Please, set the PUBLIC_KEY as an environment variable';
    // }
    //
    // if (env.PRIVATE_KEY == undefined) {
    //   throw 'PRIVATE_KEY is missing. Please, set the PRIVATE_KEY as an environment variable';
    // }

    const apiKey = { privateKey: 'e6c4bac8-8312-4add-bfca-ee750d4798e4', publicKey: 'hynkfzcw' };
    // Create a new MongoDB Atlas Project
      const mProject = new project.CfnProject(this, 'project-'.concat(id), {
          orgId: '5fe4ea50d1a2b617175ee3d4',
          name: defaults.projectName,
          apiKeys: apiKey
      });

    new atlas_basic.AtlasBasic(this,
      'test-app',
      {
        apiKeys: apiKey,
        clusterProps: {
          apiKeys: apiKey,
          projectId: mProject.ref,
          name: defaults.name,
        },
        dbUserProps: {
          databaseName: defaults.databaseName,
          projectId: mProject.ref,
          roles: defaults.roles,
          username: defaults.username,
            apiKeys: apiKey,
        },
        ipAccessListProps: {
          apiKeys: apiKey,
          projectId: mProject.ref,
          accessList: defaults.accessList,
        },
        projectId: mProject.ref,
      });
  }

}

const app = new App();

new CdkMigrateFromCfnStack(app, "test-basic-app", {  });

app.synth();