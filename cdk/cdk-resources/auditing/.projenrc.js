const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Mohit Talniya',
  authorAddress: 'mohittalniya@gmail.com',
  authorName: 'MongoDB',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'main',
  docgen: true,
  name: 'path',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();