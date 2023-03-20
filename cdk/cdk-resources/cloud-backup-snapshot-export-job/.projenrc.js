const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const cloud_backup_snapshot_export_job = new awscdk.AwsCdkConstructLibrary
({
    author: 'MongoDBAtlas',
    authorAddress: 'https://mongodb.com',
    authorName: 'MongoDBAtlas',
    cdkVersion: '2.1.0',
    defaultReleaseBranch: 'master',
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-cloud-backup-snapshot-export-job',
    module: 'mongodb_cdk_mongodb_cloud_backup_snapshot_export_job',
  },
    majorVersion: 1,
    releaseToNpm: true,
    releaseTrigger: ReleaseTrigger.manual(),
    docgen: true,
    name: '@mongodbatlas-awscdk/atlas-cloud-backup-snapshot-export-job',
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
        'cloud-backup-snapshot-export-job'],
    description: 'Retrieves or creates cloud-backup-snapshot-export-jobs in any given Atlas organization',
    // deps: [],                /* Runtime dependencies of this module. */
    // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
    // devDeps: [],             /* Build dependencies for this module. */
    // packageName: undefined,  /* The "name" in package.json. */
});

cloud_backup_snapshot_export_job.synth();
