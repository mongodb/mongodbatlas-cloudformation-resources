const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const private_endpoint_adl = new awscdk.AwsCdkConstructLibrary
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
  name: '@mongodb-cdk/atlas-private-endpoint-adl',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-atlas-private-endpoint-adl',
    module: 'mongodb_cdk_mongodb_atlas_private-endpoint-adl',
  },
  dotnet: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasPrivateEndpointADL',
    packageId: 'MongoDBCdk.MongoDBAtlasPrivateEndpointADL',
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_private-endpoint-adl',
    mavenArtifactId: 'mongodb-atlas-private-endpoint-adl',
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
    'private-endpoint-adl'],
  description: 'Retrieves or creates private-endpoint-adls in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

private_endpoint_adl.synth();
