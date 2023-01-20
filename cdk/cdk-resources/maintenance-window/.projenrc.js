const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const maintenance_window = new awscdk.AwsCdkConstructLibrary
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
  name: '@mongodb-cdk/atlas-maintenance-window',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-atlas-maintenance-window',
    module: 'mongodb_cdk_mongodb_atlas_maintenance-window',
  },
  dotnet: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasMaintenanceWindow',
    packageId: 'MongoDBCdk.MongoDBAtlasMaintenanceWindow',
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_maintenance-window',
    mavenArtifactId: 'mongodb-atlas-maintenance-window',
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
    'maintenance-window'],
  description: 'Retrieves or creates maintenance-windows in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

maintenance_window.synth();
