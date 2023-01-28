const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'INTMDB-548',
  name: '@mongodbatlas-awscdk/atlas-basic-privatelink',
  repositoryUrl: "https://github.com/mongodb/mongodbatlas-cloudformation-resources",
  sampleCode: false,
  devDeps: [
    '@mongodbatlas-awscdk/cluster',
    '@mongodbatlas-awscdk/database-user',
    '@mongodbatlas-awscdk/project',
    '@mongodbatlas-awscdk/project-ip-access-list'
  ]
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();