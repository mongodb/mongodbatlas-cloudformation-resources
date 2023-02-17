# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### AtlasBasic <a name="AtlasBasic" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer"></a>

```typescript
import { AtlasBasic } from '@mongodbatlas-awscdk/atlas-basic'

new AtlasBasic(scope: Construct, id: string, props: AtlasBasicProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps">AtlasBasicProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps">AtlasBasicProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isConstruct"></a>

```typescript
import { AtlasBasic } from '@mongodbatlas-awscdk/atlas-basic'

AtlasBasic.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.ipAccessList">ipAccessList</a></code> | <code>@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.mCluster">mCluster</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnCluster</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.mDBUser">mDBUser</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUser</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.mProject">mProject</a></code> | <code>@mongodbatlas-awscdk/project.CfnProject</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `ipAccessList`<sup>Required</sup> <a name="ipAccessList" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.ipAccessList"></a>

```typescript
public readonly ipAccessList: CfnProjectIpAccessList;
```

- *Type:* @mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList

---

##### `mCluster`<sup>Required</sup> <a name="mCluster" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.mCluster"></a>

```typescript
public readonly mCluster: CfnCluster;
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnCluster

---

##### `mDBUser`<sup>Required</sup> <a name="mDBUser" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.mDBUser"></a>

```typescript
public readonly mDBUser: CfnDatabaseUser;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUser

---

##### `mProject`<sup>Required</sup> <a name="mProject" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.mProject"></a>

```typescript
public readonly mProject: CfnProject;
```

- *Type:* @mongodbatlas-awscdk/project.CfnProject

---


## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/atlas-basic'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### AtlasBasicProps <a name="AtlasBasicProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.Initializer"></a>

```typescript
import { AtlasBasicProps } from '@mongodbatlas-awscdk/atlas-basic'

const atlasBasicProps: AtlasBasicProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.clusterProps">clusterProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps">ClusterProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.projectProps">projectProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps">ProjectProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.dbUserProps">dbUserProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps">DatabaseUserProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.ipAccessListProps">ipAccessListProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps">IpAccessListProps</a></code> | *No description.* |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `clusterProps`<sup>Required</sup> <a name="clusterProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.clusterProps"></a>

```typescript
public readonly clusterProps: ClusterProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps">ClusterProps</a>

---

##### `projectProps`<sup>Required</sup> <a name="projectProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.projectProps"></a>

```typescript
public readonly projectProps: ProjectProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps">ProjectProps</a>

---

##### `dbUserProps`<sup>Optional</sup> <a name="dbUserProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.dbUserProps"></a>

```typescript
public readonly dbUserProps: DatabaseUserProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps">DatabaseUserProps</a>

---

##### `ipAccessListProps`<sup>Optional</sup> <a name="ipAccessListProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.ipAccessListProps"></a>

```typescript
public readonly ipAccessListProps: IpAccessListProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps">IpAccessListProps</a>

---

### ClusterProps <a name="ClusterProps" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.Initializer"></a>

```typescript
import { ClusterProps } from '@mongodbatlas-awscdk/atlas-basic'

const clusterProps: ClusterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.advancedSettings">advancedSettings</a></code> | <code>@mongodbatlas-awscdk/cluster.ProcessArgs</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.apiKeys">apiKeys</a></code> | <code>@mongodbatlas-awscdk/cluster.ApiKeyDefinition</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.backupEnabled">backupEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.biConnector">biConnector</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.clusterType">clusterType</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.connectionStrings">connectionStrings</a></code> | <code>@mongodbatlas-awscdk/cluster.ConnectionStrings</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.diskSizeGb">diskSizeGb</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.encryptionAtRestProvider">encryptionAtRestProvider</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.labels">labels</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.mongoDbMajorVersion">mongoDbMajorVersion</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.paused">paused</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.pitEnabled">pitEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.replicationSpecs">replicationSpecs</a></code> | <code>@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.rootCertType">rootCertType</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.terminationProtectionEnabled">terminationProtectionEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.versionReleaseSystem">versionReleaseSystem</a></code> | <code>string</code> | *No description.* |

---

##### `advancedSettings`<sup>Optional</sup> <a name="advancedSettings" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.advancedSettings"></a>

```typescript
public readonly advancedSettings: ProcessArgs;
```

- *Type:* @mongodbatlas-awscdk/cluster.ProcessArgs

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* @mongodbatlas-awscdk/cluster.ApiKeyDefinition

---

##### `backupEnabled`<sup>Optional</sup> <a name="backupEnabled" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.backupEnabled"></a>

```typescript
public readonly backupEnabled: boolean;
```

- *Type:* boolean

---

##### `biConnector`<sup>Optional</sup> <a name="biConnector" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.biConnector"></a>

```typescript
public readonly biConnector: CfnClusterPropsBiConnector;
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector

---

##### `clusterType`<sup>Optional</sup> <a name="clusterType" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.clusterType"></a>

```typescript
public readonly clusterType: string;
```

- *Type:* string

---

##### `connectionStrings`<sup>Optional</sup> <a name="connectionStrings" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.connectionStrings"></a>

```typescript
public readonly connectionStrings: ConnectionStrings;
```

- *Type:* @mongodbatlas-awscdk/cluster.ConnectionStrings
- *Default:* REPLICASET

---

##### `diskSizeGb`<sup>Optional</sup> <a name="diskSizeGb" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.diskSizeGb"></a>

```typescript
public readonly diskSizeGb: number;
```

- *Type:* number

---

##### `encryptionAtRestProvider`<sup>Optional</sup> <a name="encryptionAtRestProvider" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.encryptionAtRestProvider"></a>

```typescript
public readonly encryptionAtRestProvider: CfnClusterPropsEncryptionAtRestProvider;
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.labels"></a>

```typescript
public readonly labels: CfnClusterPropsLabels[];
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnClusterPropsLabels[]

---

##### `mongoDbMajorVersion`<sup>Optional</sup> <a name="mongoDbMajorVersion" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.mongoDbMajorVersion"></a>

```typescript
public readonly mongoDbMajorVersion: string;
```

- *Type:* string

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `paused`<sup>Optional</sup> <a name="paused" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.paused"></a>

```typescript
public readonly paused: boolean;
```

- *Type:* boolean
- *Default:* auto-generated

---

##### `pitEnabled`<sup>Optional</sup> <a name="pitEnabled" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.pitEnabled"></a>

```typescript
public readonly pitEnabled: boolean;
```

- *Type:* boolean

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

---

##### `replicationSpecs`<sup>Optional</sup> <a name="replicationSpecs" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.replicationSpecs"></a>

```typescript
public readonly replicationSpecs: AdvancedReplicationSpec[];
```

- *Type:* @mongodbatlas-awscdk/cluster.AdvancedReplicationSpec[]

---

##### `rootCertType`<sup>Optional</sup> <a name="rootCertType" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.rootCertType"></a>

```typescript
public readonly rootCertType: string;
```

- *Type:* string

---

##### `terminationProtectionEnabled`<sup>Optional</sup> <a name="terminationProtectionEnabled" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.terminationProtectionEnabled"></a>

```typescript
public readonly terminationProtectionEnabled: boolean;
```

- *Type:* boolean

---

##### `versionReleaseSystem`<sup>Optional</sup> <a name="versionReleaseSystem" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.versionReleaseSystem"></a>

```typescript
public readonly versionReleaseSystem: string;
```

- *Type:* string

---

### DatabaseUserProps <a name="DatabaseUserProps" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.Initializer"></a>

```typescript
import { DatabaseUserProps } from '@mongodbatlas-awscdk/atlas-basic'

const databaseUserProps: DatabaseUserProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.awsiamType">awsiamType</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsAwsiamType</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.databaseName">databaseName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.deleteAfterDate">deleteAfterDate</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.labels">labels</a></code> | <code>@mongodbatlas-awscdk/database-user.LabelDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.ldapAuthType">ldapAuthType</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsLdapAuthType</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.password">password</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.roles">roles</a></code> | <code>@mongodbatlas-awscdk/database-user.RoleDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.scopes">scopes</a></code> | <code>@mongodbatlas-awscdk/database-user.ScopeDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.username">username</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.x509Type">x509Type</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsX509Type</code> | *No description.* |

---

##### `awsiamType`<sup>Optional</sup> <a name="awsiamType" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.awsiamType"></a>

```typescript
public readonly awsiamType: CfnDatabaseUserPropsAwsiamType;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsAwsiamType

---

##### `databaseName`<sup>Optional</sup> <a name="databaseName" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.databaseName"></a>

```typescript
public readonly databaseName: string;
```

- *Type:* string

---

##### `deleteAfterDate`<sup>Optional</sup> <a name="deleteAfterDate" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.deleteAfterDate"></a>

```typescript
public readonly deleteAfterDate: string;
```

- *Type:* string

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.labels"></a>

```typescript
public readonly labels: LabelDefinition[];
```

- *Type:* @mongodbatlas-awscdk/database-user.LabelDefinition[]
- *Default:* admin

---

##### `ldapAuthType`<sup>Optional</sup> <a name="ldapAuthType" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.ldapAuthType"></a>

```typescript
public readonly ldapAuthType: CfnDatabaseUserPropsLdapAuthType;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsLdapAuthType

---

##### `password`<sup>Optional</sup> <a name="password" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.password"></a>

```typescript
public readonly password: string;
```

- *Type:* string

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string
- *Default:* cdk-pwd

---

##### `roles`<sup>Optional</sup> <a name="roles" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.roles"></a>

```typescript
public readonly roles: RoleDefinition[];
```

- *Type:* @mongodbatlas-awscdk/database-user.RoleDefinition[]

---

##### `scopes`<sup>Optional</sup> <a name="scopes" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.scopes"></a>

```typescript
public readonly scopes: ScopeDefinition[];
```

- *Type:* @mongodbatlas-awscdk/database-user.ScopeDefinition[]

---

##### `username`<sup>Optional</sup> <a name="username" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.username"></a>

```typescript
public readonly username: string;
```

- *Type:* string
- *Default:* cdk-user

---

##### `x509Type`<sup>Optional</sup> <a name="x509Type" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.x509Type"></a>

```typescript
public readonly x509Type: CfnDatabaseUserPropsX509Type;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsX509Type

---

### IpAccessListProps <a name="IpAccessListProps" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.Initializer"></a>

```typescript
import { IpAccessListProps } from '@mongodbatlas-awscdk/atlas-basic'

const ipAccessListProps: IpAccessListProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.accessList">accessList</a></code> | <code>@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.listOptions">listOptions</a></code> | <code>@mongodbatlas-awscdk/project-ip-access-list.ListOptions</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.totalCount">totalCount</a></code> | <code>number</code> | *No description.* |

---

##### `accessList`<sup>Required</sup> <a name="accessList" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.accessList"></a>

```typescript
public readonly accessList: AccessListDefinition[];
```

- *Type:* @mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition[]

---

##### `listOptions`<sup>Optional</sup> <a name="listOptions" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.listOptions"></a>

```typescript
public readonly listOptions: ListOptions;
```

- *Type:* @mongodbatlas-awscdk/project-ip-access-list.ListOptions

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string
- *Default:* allow-all

---

##### `totalCount`<sup>Optional</sup> <a name="totalCount" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.totalCount"></a>

```typescript
public readonly totalCount: number;
```

- *Type:* number

---

### ProjectProps <a name="ProjectProps" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.Initializer"></a>

```typescript
import { ProjectProps } from '@mongodbatlas-awscdk/atlas-basic'

const projectProps: ProjectProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.orgId">orgId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.clusterCount">clusterCount</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectApiKeys">projectApiKeys</a></code> | <code>@mongodbatlas-awscdk/project.ProjectApiKey[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectOwnerId">projectOwnerId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectSettings">projectSettings</a></code> | <code>@mongodbatlas-awscdk/project.ProjectSettings</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectTeams">projectTeams</a></code> | <code>@mongodbatlas-awscdk/project.ProjectTeam[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.withDefaultAlertsSettings">withDefaultAlertsSettings</a></code> | <code>boolean</code> | *No description.* |

---

##### `orgId`<sup>Required</sup> <a name="orgId" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string
- *Default:* auto-generated

---

##### `clusterCount`<sup>Optional</sup> <a name="clusterCount" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.clusterCount"></a>

```typescript
public readonly clusterCount: number;
```

- *Type:* number

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `projectApiKeys`<sup>Optional</sup> <a name="projectApiKeys" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectApiKeys"></a>

```typescript
public readonly projectApiKeys: ProjectApiKey[];
```

- *Type:* @mongodbatlas-awscdk/project.ProjectApiKey[]

---

##### `projectOwnerId`<sup>Optional</sup> <a name="projectOwnerId" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectOwnerId"></a>

```typescript
public readonly projectOwnerId: string;
```

- *Type:* string

---

##### `projectSettings`<sup>Optional</sup> <a name="projectSettings" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectSettings"></a>

```typescript
public readonly projectSettings: ProjectSettings;
```

- *Type:* @mongodbatlas-awscdk/project.ProjectSettings

---

##### `projectTeams`<sup>Optional</sup> <a name="projectTeams" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectTeams"></a>

```typescript
public readonly projectTeams: ProjectTeam[];
```

- *Type:* @mongodbatlas-awscdk/project.ProjectTeam[]

---

##### `withDefaultAlertsSettings`<sup>Optional</sup> <a name="withDefaultAlertsSettings" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.withDefaultAlertsSettings"></a>

```typescript
public readonly withDefaultAlertsSettings: boolean;
```

- *Type:* boolean

---



