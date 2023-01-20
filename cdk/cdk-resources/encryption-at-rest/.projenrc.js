const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const encryption_at_rest = new awscdk.AwsCdkConstructLibrary
({
    author: 'MongoDB',
    authorAddress: 'https://mongodb.com',
    authorName: 'MongoDB',
    cdkVersion: '2.1.0',
    defaultReleaseBranch: 'INTMDB-548',
    majorVersion: 1,
    releaseToNpm: true,
    releaseTrigger: ReleaseTrigger.manual(),
    docgen: true,
    name: '@mongodb-cdk/atlas-encryption-at-rest',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-encryption-at-rest',
        module: 'mongodb_cdk_mongodb_atlas_encryption-at-rest',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasEncryptionAtRest',
        packageId: 'MongoDBCdk.MongoDBAtlasEncryptionAtRest',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_encryption-at-rest',
        mavenArtifactId: 'mongodb-atlas-encryption-at-rest',
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
        'encryption-at-rest'],
    description: 'Retrieves or creates encryption-at-rests in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

encryption_at_rest.synth();
