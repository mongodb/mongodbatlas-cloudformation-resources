const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const teams = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasTeams',
    packageId: 'MongoDBCdk.MongoDBAtlasTeams'
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_teams',
    mavenArtifactId: 'mongodb-atlas-teams',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'teams'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-teams',
    module: 'mongodb_cdk_mongodb_teams'
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/teams',
  repositoryUrl:
    'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  keywords: [
    'cdk',
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
    'teams'
  ],
  description: 'Retrieves or creates teamss in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})

teams.synth()
