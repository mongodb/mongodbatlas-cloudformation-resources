const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const third_party_integration = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasThirdPartyIntegration',
    packageId: 'MongoDBCdk.MongoDBAtlasThirdPartyIntegration',
  },
  publishToMaven:{
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_third_party_integration',
    mavenArtifactId: 'mongodb-atlas-third-party-integration',
    mavenGroupId: 'com.github.mongodb.cdk',
  },  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'third-party-integration'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-third-party-integration',
    module: 'mongodb_cdk_mongodb_third_party_integration',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/third-party-integration',
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
    'third-party-integration'],
  description: 'Retrieves or creates third-party-integrations in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

third_party_integration.synth();
