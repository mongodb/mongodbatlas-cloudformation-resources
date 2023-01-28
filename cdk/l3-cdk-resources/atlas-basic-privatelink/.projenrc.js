const { awscdk } = require("projen");
const project = new awscdk.AwsCdkConstructLibrary({
  author: "sangramPeerI",
  authorAddress: "sangram.konde@peerislands.io",
  cdkVersion: "2.1.0",
  defaultReleaseBranch: "main",
  name: "atlas-basic-privatelink",
  repositoryUrl: "https://github.com/sangramPeerI/mongodbatlas-cloudformation-resources.git",

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();