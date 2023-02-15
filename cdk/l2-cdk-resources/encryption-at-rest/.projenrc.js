const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const atlasEncryptionAtRest = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.64.0',
  defaultReleaseBranch: 'master',
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/atlas-encryption-at-rest',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  peerDeps: ['@mongodbatlas-awscdk/encryption-at-rest@1.0.1'],
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
    'encryption-at-rest'],
  description: 'Returns and edits the Encryption at Rest using Customer Key Management configuration.',
  stability: 'experimental',
  deps: [
    '@mongodbatlas-awscdk/encryption-at-rest@1.0.1',
  ],
  devDeps: [
    '@mongodbatlas-awscdk/encryption-at-rest@1.0.1',
  ],
});
atlasEncryptionAtRest.synth();
