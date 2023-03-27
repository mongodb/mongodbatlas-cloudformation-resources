const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const private_endpoint = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'private-endpoint'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-private-endpoint',
    module: 'mongodb_cdk_mongodb_private_endpoint',
  },
  majorVersion: 1,
  releaseToNpm: true,
  npmAccess: 'public',
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/private-endpoint',
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
    'private-endpoint'],
  description: 'Retrieves or creates private-endpoints in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

private_endpoint.synth();
