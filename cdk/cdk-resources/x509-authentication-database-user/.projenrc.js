const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const x509-authentication-database-user = new awscdk.AwsCdkConstructLibrary
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
    name: 'mongodb-atlas-x509-authentication-database-user',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-x509-authentication-database-user',
        module: 'mongodb_cdk_mongodb_atlas_x509-authentication-database-user',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasX509AuthenticationDatabaseUser',
        packageId: 'MongoDBCdk.MongoDBAtlasX509AuthenticationDatabaseUser',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_x509-authentication-database-user',
        mavenArtifactId: 'mongodb-atlas-x509-authentication-database-user',
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
        'x509-authentication-database-user'],
    description: 'Retrieves or creates x509-authentication-database-users in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

x509-authentication-database-user.synth();
