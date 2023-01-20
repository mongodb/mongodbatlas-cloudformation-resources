const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const serverless-instance = new awscdk.AwsCdkConstructLibrary
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
    name: 'mongodb-atlas-serverless-instance',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-serverless-instance',
        module: 'mongodb_cdk_mongodb_atlas_serverless-instance',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasServerlessInstance',
        packageId: 'MongoDBCdk.MongoDBAtlasServerlessInstance',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_serverless-instance',
        mavenArtifactId: 'mongodb-atlas-serverless-instance',
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
        'serverless-instance'],
    description: 'Retrieves or creates serverless-instances in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

serverless-instance.synth();
