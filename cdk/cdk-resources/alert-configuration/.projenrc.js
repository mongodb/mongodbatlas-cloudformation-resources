const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const alert_configuration = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasAlertConfiguration',
    packageId: 'MongoDBCdk.MongoDBAtlasAlertConfiguration',
  },
  publishToMaven:{
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_alert_configuration',
    mavenArtifactId: 'mongodb-atlas-alert-configuration',
    mavenGroupId: 'com.github.mongodb.cdk',
  },  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'alert-configuration'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-alert-configuration',
    module: 'mongodb_cdk_mongodb_alert_configuration',
  },
  majorVersion: 1,
  releaseToNpm: true,
  npmAccess: 'public',
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/alert-configuration',
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
    'alert-configuration'],
  description: 'Retrieves or creates alert-configurations in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

alert_configuration.synth();
