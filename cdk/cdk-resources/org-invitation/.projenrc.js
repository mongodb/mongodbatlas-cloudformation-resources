const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const org_invitation = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'org-invitation'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-org-invitation',
    module: 'mongodb_cdk_mongodb_org_invitation',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/org-invitation',
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
    'org-invitation'],
  description: 'Retrieves or creates org-invitations in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

org_invitation.synth();
