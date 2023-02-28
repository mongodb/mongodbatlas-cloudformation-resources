// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { App, Stack } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { AtlasBasic } from '../src';


const RESOURCE_NAME_PROJECT = 'MongoDB::Atlas::Project';
const RESOURCE_NAME_CLUSTER = 'MongoDB::Atlas::Cluster';
const RESOURCE_NAME_DB_USER = 'MongoDB::Atlas::DatabaseUser';
const PROJECT_ID= 'testProjectId';
const ORG_ID= 'testProjectId';
const PROJECT_NAME= 'test';
const INSTANCE_SIZE = 'M30';
const REGION = 'US_EAST_1';
const DATABASE_NAME = 'test';
const DATABASE_USER_NAME = 'atlas-user';
const ADMIN_DB= 'admin';
const ROLE_NAME = 'atlasAdmin';
const PWD = 'test';

test('AtlasBasis construct should contain default properties', () => {
  const mockApp = new App();
  const stack = new Stack(mockApp);

  new AtlasBasic(stack, 'testing-stack', {
    apiKeys: {},
    clusterProps: {
      replicationSpecs: [
        {
          numShards: 3,
          advancedRegionConfigs: [{
            regionName: REGION,
            electableSpecs: {
              instanceSize: INSTANCE_SIZE,
              nodeCount: 3,
            },
          }],
        },
      ],
      name: PROJECT_NAME,
    },
    projectProps: {
      orgId: ORG_ID,
      name: PROJECT_NAME,
    },
    dbUserProps: {
      projectId: PROJECT_ID,
      databaseName: DATABASE_NAME,
      password: PWD,
    },
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties(RESOURCE_NAME_PROJECT, {
    OrgId: ORG_ID,
    Name: PROJECT_NAME,
  });

  template.hasResourceProperties(RESOURCE_NAME_CLUSTER, {
    ClusterType: 'REPLICASET',
    Name: PROJECT_NAME,
    ReplicationSpecs: [{
      NumShards: 3,
      AdvancedRegionConfigs: [{
        RegionName: REGION,
        ElectableSpecs: {
          InstanceSize: INSTANCE_SIZE,
          NodeCount: 3,
        },
      }],
    }],
  });

  template.hasResourceProperties(RESOURCE_NAME_DB_USER, {
    ProjectId: PROJECT_ID,
    DatabaseName: DATABASE_NAME,
    Password: PWD,
    Username: DATABASE_USER_NAME,
    Roles: [
      {
        DatabaseName: ADMIN_DB,
        RoleName: ROLE_NAME,
      },
    ],
  });
});