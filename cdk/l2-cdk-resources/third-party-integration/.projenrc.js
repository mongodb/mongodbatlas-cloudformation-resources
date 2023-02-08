const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const project = new awscdk.AwsCdkConstructLibrary({
  // author: 'Aastha Mahendru',
  // authorAddress: 'aastha.mahendru@mongodb.com',
  // cdkVersion: '2.1.0',
  // defaultReleaseBranch: 'main',
  // name: 'third-party-integration',
  // repositoryUrl: 'git@github.com:mongodb/mongodbatlas-cloudformation-resources.git',
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'INTMDB-548',
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/atlas-integrations',
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
    'l2',
    'mongodb',
    'atlas',
    'third-party-integration'],
  description: 'Retrieves or creates third-party-integrations in any given Atlas organization',
  stability: 'experimental',
  deps: [
    '@mongodbatlas-awscdk/third-party-integration',
  ], /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();