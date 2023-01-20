const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const federated-settings-identity-provider = new awscdk.AwsCdkConstructLibrary
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
    name: 'mongodb-atlas-federated-settings-identity-provider',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-federated-settings-identity-provider',
        module: 'mongodb_cdk_mongodb_atlas_federated-settings-identity-provider',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasFederatedSettingsIdentityProvider',
        packageId: 'MongoDBCdk.MongoDBAtlasFederatedSettingsIdentityProvider',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_federated-settings-identity-provider',
        mavenArtifactId: 'mongodb-atlas-federated-settings-identity-provider',
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
        'federated-settings-identity-provider'],
    description: 'Retrieves or creates federated-settings-identity-providers in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

federated-settings-identity-provider.synth();
