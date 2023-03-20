const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const cloud_backup_snapshot_export_bucket = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-cloud-backup-snapshot-export-bucket',
    module: 'mongodb_cdk_mongodb_cloud_backup_snapshot_export_bucket',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/cloud-backup-snapshot-export-bucket',
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
    'cloud-backup-snapshot-export-bucket'],
  description: 'Retrieves or creates cloud-backup-snapshot-export-buckets in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
  // eslint
});

cloud_backup_snapshot_export_bucket.eslint?.addOverride({
  files: ['src/**'],
  rules: {
    'max-len': [
      'error', {
        code: 180,
      },
    ],
  },
});

cloud_backup_snapshot_export_bucket.synth();
