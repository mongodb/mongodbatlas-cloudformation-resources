const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Mohit Talniya',
  authorAddress: 'mohittalniya@gmail.com',
  authorName: 'Mohit Talniya',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'main',
  docgen: true,
  name: 'test-atlas-client',
  repositoryUrl: 'https://github.com/SuperMohit',
  sampleCode: false,
  deps: [
    '@mongodbatlas-awscdk/cluster',
    '@mongodbatlas-awscdk/project',
    '@mongodbatlas-awscdk/database-user',
    '@mongodbatlas-awscdk/project-ip-access-list',
  ],
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();