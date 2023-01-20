const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const third-party-integration = new awscdk.AwsCdkConstructLibrary
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
    name: 'mongodb-atlas-third-party-integration',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-third-party-integration',
        module: 'mongodb_cdk_mongodb_atlas_third-party-integration',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasThirdPartyIntegration',
        packageId: 'MongoDBCdk.MongoDBAtlasThirdPartyIntegration',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_third-party-integration',
        mavenArtifactId: 'mongodb-atlas-third-party-integration',
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
        'third-party-integration'],
    description: 'Retrieves or creates third-party-integrations in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

third-party-integration.synth();
