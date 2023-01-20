const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const project-ip-access-list = new awscdk.AwsCdkConstructLibrary
({
    author: 'MongoDB',
    authorAddress: 'https://mongodb.com',
    authorName: 'MongoDB',
    cdkVersion: '2.1.0',
    defaultReleaseBranch: 'main',
    majorVersion: 1,
    releaseToNpm: true,
    releaseTrigger: ReleaseTrigger.manual(),
    docgen: true,
    name: 'mongodb-atlas-project-ip-access-list',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-project-ip-access-list',
        module: 'mongodb_cdk_mongodb_atlas_project-ip-access-list',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasProjectIpAccessList',
        packageId: 'MongoDBCdk.MongoDBAtlasProjectIpAccessList',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_project-ip-access-list',
        mavenArtifactId: 'mongodb-atlas-project-ip-access-list',
        mavenGroupId: 'com.github.mongodb.cdk',
    },
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

project-ip-access-list.synth();
