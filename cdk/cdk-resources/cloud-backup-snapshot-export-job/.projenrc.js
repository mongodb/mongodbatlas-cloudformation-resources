const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const cloud-backup-snapshot-export-job = new awscdk.AwsCdkConstructLibrary
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
    name: 'mongodb-atlas-cloud-backup-snapshot-export-job',
    repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
    sampleCode: false,
    publishToPypi: {
        distName: 'mongodb-cdk-mongodb-atlas-cloud-backup-snapshot-export-job',
        module: 'mongodb_cdk_mongodb_atlas_cloud-backup-snapshot-export-job',
    },
    dotnet: {
        dotNetNamespace: 'MongoDBCdk.MongoDBAtlasCloudBackupSnapshotExportJob',
        packageId: 'MongoDBCdk.MongoDBAtlasCloudBackupSnapshotExportJob',
    },
    publishToMaven: {
        javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_cloud-backup-snapshot-export-job',
        mavenArtifactId: 'mongodb-atlas-cloud-backup-snapshot-export-job',
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
        'cloud-backup-snapshot-export-job'],
    description: 'Retrieves or creates cloud-backup-snapshot-export-jobs in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

cloud-backup-snapshot-export-job.synth();
