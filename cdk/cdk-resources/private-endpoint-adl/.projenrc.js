const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const private_endpoint_adl = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasPrivateEndpointAdl',
    packageId: 'MongoDBCdk.MongoDBAtlasPrivateEndpointAdl',
  },
  publishToMaven:{
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_private_endpoint_adl',
    mavenArtifactId: 'mongodb-atlas-private-endpoint-adl',
    mavenGroupId: 'com.github.mongodb.cdk',
  },  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'private-endpoint-adl'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-private-endpoint-adl',
    module: 'mongodb_cdk_mongodb_private_endpoint_adl',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/private-endpoint-adl',
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
    'private-endpoint-adl'],
  description: 'Retrieves or creates private-endpoint-adls in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

private_endpoint_adl.synth();
