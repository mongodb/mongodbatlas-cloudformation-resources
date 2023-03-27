const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const ldap_configuration = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasLdapConfiguration',
    packageId: 'MongoDBCdk.MongoDBAtlasLdapConfiguration',
  },
  publishToMaven:{
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_ldap_configuration',
    mavenArtifactId: 'mongodb-atlas-ldap-configuration',
    mavenGroupId: 'com.github.mongodb.cdk',
  },  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'ldap-configuration'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-ldap-configuration',
    module: 'mongodb_cdk_mongodb_ldap_configuration',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/ldap-configuration',
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
    'ldap-configuration'],
  description: 'Retrieves or creates ldap-configurations in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

ldap_configuration.synth();
