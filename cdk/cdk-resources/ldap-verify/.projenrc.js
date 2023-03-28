const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const ldap_verify = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasLdapVerify',
    packageId: 'MongoDBCdk.MongoDBAtlasLdapVerify'
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_ldap_verify',
    mavenArtifactId: 'mongodb-atlas-ldap-verify',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'ldap-verify'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-ldap-verify',
    module: 'mongodb_cdk_mongodb_ldap_verify'
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/ldap-verify',
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
    'ldap-verify'
  ],
  description:
    'Retrieves or creates ldap-verifys in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})

ldap_verify.synth()
