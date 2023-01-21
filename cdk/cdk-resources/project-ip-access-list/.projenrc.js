const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const project_ip_access_list = new awscdk.AwsCdkConstructLibrary
({
    author: 'MongoDBAtlas',
    authorAddress: 'https://mongodb.com',
    authorName: 'MongoDBAtlas',
    cdkVersion: '2.1.0',
    defaultReleaseBranch: 'INTMDB-548',
    majorVersion: 1,
    releaseToNpm: true,
    releaseTrigger: ReleaseTrigger.manual(),
    docgen: true,
    name: '@mongodbatlas-awscdk/atlas-project-ip-access-list',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
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
        'l1',
        'mongodb',
        'atlas',
        'project-ip-access-list'],
    description: 'Retrieves or creates project-ip-access-lists in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

project_ip_access_list.synth();
