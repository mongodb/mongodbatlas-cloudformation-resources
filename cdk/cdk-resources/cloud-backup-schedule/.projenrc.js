const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const cloud_backup_schedule = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasCloudBackupSchedule',
    packageId: 'MongoDBCdk.MongoDBAtlasCloudBackupSchedule'
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_cloud_backup_schedule',
    mavenArtifactId: 'mongodb-atlas-cloud-backup-schedule',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'cloud-backup-schedule'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-cloud-backup-schedule',
    module: 'mongodb_cdk_mongodb_cloud_backup_schedule'
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/cloud-backup-schedule',
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
    'cloud-backup-schedule'
  ],
  description:
    'Retrieves or creates cloud-backup-schedules in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})

cloud_backup_schedule.synth()
