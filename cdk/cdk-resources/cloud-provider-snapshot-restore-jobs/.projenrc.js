const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const cloud_provider_snapshot_restore_jobs = new awscdk.AwsCdkConstructLibrary
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
  name: '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-atlas-cloud-provider-snapshot-restore-jobs',
    module: 'mongodb_cdk_mongodb_atlas_cloud-provider-snapshot-restore-jobs',
  },
  dotnet: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasCloudProviderSnapshotRestoreJobs',
    packageId: 'MongoDBCdk.MongoDBAtlasCloudProviderSnapshotRestoreJobs',
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_cloud-provider-snapshot-restore-jobs',
    mavenArtifactId: 'mongodb-atlas-cloud-provider-snapshot-restore-jobs',
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
    'cloud-provider-snapshot-restore-jobs'],
  description: 'Retrieves or creates cloud-provider-snapshot-restore-jobss in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

cloud_provider_snapshot_restore_jobs.synth();
