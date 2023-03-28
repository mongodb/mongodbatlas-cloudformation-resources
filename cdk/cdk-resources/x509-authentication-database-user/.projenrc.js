const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const x509_authentication_database_user = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasX509AuthenticationDatabaseUser',
    packageId: 'MongoDBCdk.MongoDBAtlasX509AuthenticationDatabaseUser'
  },
  publishToMaven: {
    javaPackage:
      'com.github.mongodb.cdk.mongodb_atlas_x509_authentication_database_user',
    mavenArtifactId: 'mongodb-atlas-x509-authentication-database-user',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'x509-authentication-database-user'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-x509-authentication-database-user',
    module: 'mongodb_cdk_mongodb_x509_authentication_database_user'
  },
  majorVersion: 1,
  releaseToNpm: true,
  npmAccess: 'public',
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/x509-authentication-database-user',
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
    'x509-authentication-database-user'
  ],
  description:
    'Retrieves or creates x509-authentication-database-users in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})

x509_authentication_database_user.synth()
