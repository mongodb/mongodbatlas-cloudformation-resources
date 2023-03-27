const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const project_invitation = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasProjectInvitation',
    packageId: 'MongoDBCdk.MongoDBAtlasProjectInvitation',
  },
  publishToMaven:{
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_project_invitation',
    mavenArtifactId: 'mongodb-atlas-project-invitation',
    mavenGroupId: 'com.github.mongodb.cdk',
  },  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'project-invitation'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-project-invitation',
    module: 'mongodb_cdk_mongodb_project-invitation',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/project-invitation',
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
    'project-invitation'],
  description: 'Retrieves or creates project-invitations in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

project_invitation.synth();
