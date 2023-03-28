const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const cloud_backup_restore_jobs = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasCloudBackupRestoreJobs',
    packageId: 'MongoDBCdk.MongoDBAtlasCloudBackupRestoreJobs'
  },
  publishToMaven: {
    javaPackage:
      'com.github.mongodb.cdk.mongodb_atlas_cloud_backup_restore_jobs',
    mavenArtifactId: 'mongodb-atlas-cloud-backup-restore-jobs',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'cloud-backup-restore-jobs'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-cloud-backup-restore-jobs',
    module: 'mongodb_cdk_mongodb_cloud-backup_restore_jobs'
  },
  majorVersion: 1,
  releaseToNpm: true,
  npmAccess: 'public',
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/cloud-backup-restore-jobs',
  repositoryUrl:
    'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  keywords: [
    'cdk',
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
    'cloud-backup-restore-jobs'
  ],
  description:
    'Retrieves or creates cloud-backup-restore-jobss in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})

cloud_backup_restore_jobs.synth()
