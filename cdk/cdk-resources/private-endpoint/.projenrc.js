const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const private_endpoint = new awscdk.AwsCdkConstructLibrary
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
  name: '@mongodb-cdk/atlas-private-endpoint',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-atlas-private-endpoint',
    module: 'mongodb_cdk_mongodb_atlas_private-endpoint',
  },
  dotnet: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasPrivateEndpoint',
    packageId: 'MongoDBCdk.MongoDBAtlasPrivateEndpoint',
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_private-endpoint',
    mavenArtifactId: 'mongodb-atlas-private-endpoint',
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
    'private-endpoint'],
  description: 'Retrieves or creates private-endpoints in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

private_endpoint.synth();
