const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const cloud_provider_snapshot_restore_jobs = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/cloud-provider-snapshot-restore-jobs',
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
    'cloud-provider-snapshot-restore-jobs'],
  description: 'Retrieves or creates cloud-provider-snapshot-restore-jobss in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
cloud_provider_snapshot_restore_jobs.eslint?.addOverride({
  files: ['src/**'],
  rules: {
    'max-len': [
      'error', {
        code: 180,
      },
    ],
  },
});

cloud_provider_snapshot_restore_jobs.synth();
