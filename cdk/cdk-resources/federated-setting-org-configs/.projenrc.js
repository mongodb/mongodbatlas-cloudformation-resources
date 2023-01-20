const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const federated-setting-org-configs = new awscdk.AwsCdkConstructLibrary
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
    name: 'mongodb-atlas-federated-setting-org-configs',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-federated-setting-org-configs',
        module: 'mongodb_cdk_mongodb_atlas_federated-setting-org-configs',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasFederatedSettingOrgConfigs',
        packageId: 'MongoDBCdk.MongoDBAtlasFederatedSettingOrgConfigs',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_federated-setting-org-configs',
        mavenArtifactId: 'mongodb-atlas-federated-setting-org-configs',
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
        'federated-setting-org-configs'],
    description: 'Retrieves or creates federated-setting-org-configss in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

federated-setting-org-configs.synth();
