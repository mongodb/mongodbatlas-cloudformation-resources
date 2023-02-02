const { awscdk } = require("projen");
const project = new awscdk.AwsCdkConstructLibrary({
  author: "MongoDB",
  authorAddress: "a@mongodb.com",
  authorName: "MongoDB",
  cdkVersion: "2.1.0",
  defaultReleaseBranch: "main",
  docgen: true,
  name: "client-l2",
  repositoryUrl: "https://github.com/mongodb/mongodbatlas-cloudformation-resources",
  sampleCode: false
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();