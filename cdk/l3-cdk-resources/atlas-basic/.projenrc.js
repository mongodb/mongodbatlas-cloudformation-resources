const { awscdk, JsonFile } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'main',
  name: '@mongodbatlas-awscdk/atlas-basic',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  deps: [
    '@mongodbatlas-awscdk/cluster',
    '@mongodbatlas-awscdk/project',
    '@mongodbatlas-awscdk/database-user',
    '@mongodbatlas-awscdk/project-ip-access-list',
  ],
  devDeps: ['aws-cdk', 'ts-node'],
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
    'l3',
    'mongodb',
    'atlas',
    'cluster',
    'dbuser',
    'ip-access-list'],
  description: 'Creates a project, cluster, dbuser and ipaccess list in MongoDB Atlas',

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

new JsonFile(project, 'cdk.json', {
  obj: {
    app: 'npx ts-node --prefer-ts-exts src/integ.default.ts',
  },
});

const common_exclude = ['cdk.out', 'cdk.context.json', 'yarn-error.log'];
project.npmignore.exclude(...common_exclude);
project.gitignore.exclude(...common_exclude);


project.synth();