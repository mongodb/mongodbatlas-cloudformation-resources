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

const { awscdk, JsonFile } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
    author: 'MongoDBAtlas',
    authorAddress: 'https://mongodb.com',
    authorName: 'MongoDBAtlas',
    cdkVersion: '2.1.0',
    defaultReleaseBranch: 'main',
    name: '@mongodbatlas-awscdk/atlas-basic-private-endpoint',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    stability: 'experimental',
    peerDeps: [
        '@mongodbatlas-awscdk/atlas-basic',
        '@mongodbatlas-awscdk/private-endpoint',
        '@mongodbatlas-awscdk/cluster@1.1.4',
        '@mongodbatlas-awscdk/project',
        '@mongodbatlas-awscdk/database-user',
        '@mongodbatlas-awscdk/project-ip-access-list',
    ],
    devDeps: [
        '@mongodbatlas-awscdk/atlas-basic',
        '@mongodbatlas-awscdk/private-endpoint'],

    sampleCode: false,
    keywords: ['cdk',
        'awscdk',
        'aws-cdk',
        'cloudformation',
        'cfn',
        'extensions',
        'constructs',
        'cfn-resources',
        'cloudformation-registry',
        'l3',
        'mongodb',
        'atlas',
        'cluster',
        'dbuser',
        'ip-access-list'],
    description: 'Creates a project, cluster, dbuser, private endpoint and adds an IP entry to the access list in MongoDB Atlas',
});

new JsonFile(project, 'cdk.json', {
    obj: {
        app: 'npx ts-node --prefer-ts-exts src/integ.default.ts',
    },
});

const common_exclude = ['cdk.out', 'cdk.context.json', 'yarn-error.log'];
project.npmignore.exclude(...common_exclude);
project.gitignore.exclude(...common_exclude);


project.synth();
