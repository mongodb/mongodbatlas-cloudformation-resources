const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDB',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDB',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'main',
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: 'mongodb-atlas-project',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-atlas-project',
    module: 'mongodb_cdk_mongodb_atlas_project',
  },
  dotnet: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasProject',
    packageId: 'MongoDBCdk.MongoDBAtlasProject',
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_project',
    mavenArtifactId: 'mongodb-atlas-project',
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
    'project'],
  description: 'Retrieves or creates projects in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

project.synth();
