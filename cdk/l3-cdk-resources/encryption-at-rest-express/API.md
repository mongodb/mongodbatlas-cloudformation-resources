# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### AtlasEncryptionAtRestExpress <a name="AtlasEncryptionAtRestExpress" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer"></a>

```typescript
import { AtlasEncryptionAtRestExpress } from '@mongodbatlas-awscdk/encryption-at-rest-express'

new AtlasEncryptionAtRestExpress(scope: Construct, id: string, props: AtlasEncryptionAtRestExpressProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps">AtlasEncryptionAtRestExpressProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps">AtlasEncryptionAtRestExpressProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.isConstruct"></a>

```typescript
import { AtlasEncryptionAtRestExpress } from '@mongodbatlas-awscdk/encryption-at-rest-express'

AtlasEncryptionAtRestExpress.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.encryptionAtRest">encryptionAtRest</a></code> | <code>@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.accessList">accessList</a></code> | <code>@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.cluster">cluster</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnCluster</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.databaseUser">databaseUser</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUser</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `encryptionAtRest`<sup>Required</sup> <a name="encryptionAtRest" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.encryptionAtRest"></a>

```typescript
public readonly encryptionAtRest: CfnEncryptionAtRest;
```

- *Type:* @mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest

---

##### `accessList`<sup>Optional</sup> <a name="accessList" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.accessList"></a>

```typescript
public readonly accessList: CfnProjectIpAccessList;
```

- *Type:* @mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList

---

##### `cluster`<sup>Optional</sup> <a name="cluster" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.cluster"></a>

```typescript
public readonly cluster: CfnCluster;
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnCluster

---

##### `databaseUser`<sup>Optional</sup> <a name="databaseUser" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpress.property.databaseUser"></a>

```typescript
public readonly databaseUser: CfnDatabaseUser;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUser

---


## Structs <a name="Structs" id="Structs"></a>

### AtlasEncryptionAtRestExpressProps <a name="AtlasEncryptionAtRestExpressProps" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.Initializer"></a>

```typescript
import { AtlasEncryptionAtRestExpressProps } from '@mongodbatlas-awscdk/encryption-at-rest-express'

const atlasEncryptionAtRestExpressProps: AtlasEncryptionAtRestExpressProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.encryptionAtRest">encryptionAtRest</a></code> | <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps">AtlasEncryptionAtRestProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.accessList">accessList</a></code> | <code>@mongodbatlas-awscdk/atlas-basic.IpAccessListProps</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.cluster">cluster</a></code> | <code>@mongodbatlas-awscdk/atlas-basic.ClusterProps</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.databaseUser">databaseUser</a></code> | <code>@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.profile">profile</a></code> | <code>string</code> | *No description.* |

---

##### `encryptionAtRest`<sup>Required</sup> <a name="encryptionAtRest" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.encryptionAtRest"></a>

```typescript
public readonly encryptionAtRest: AtlasEncryptionAtRestProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps">AtlasEncryptionAtRestProps</a>

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

---

##### `accessList`<sup>Optional</sup> <a name="accessList" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.accessList"></a>

```typescript
public readonly accessList: IpAccessListProps;
```

- *Type:* @mongodbatlas-awscdk/atlas-basic.IpAccessListProps

---

##### `cluster`<sup>Optional</sup> <a name="cluster" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.cluster"></a>

```typescript
public readonly cluster: ClusterProps;
```

- *Type:* @mongodbatlas-awscdk/atlas-basic.ClusterProps

---

##### `databaseUser`<sup>Optional</sup> <a name="databaseUser" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.databaseUser"></a>

```typescript
public readonly databaseUser: DatabaseUserProps;
```

- *Type:* @mongodbatlas-awscdk/atlas-basic.DatabaseUserProps

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestExpressProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

---

### AtlasEncryptionAtRestProps <a name="AtlasEncryptionAtRestProps" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.Initializer"></a>

```typescript
import { AtlasEncryptionAtRestProps } from '@mongodbatlas-awscdk/encryption-at-rest-express'

const atlasEncryptionAtRestProps: AtlasEncryptionAtRestProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.customerMasterKeyId">customerMasterKeyId</a></code> | <code>string</code> | The AWS customer master key used to encrypt and decrypt the MongoDB master keys. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.roleId">roleId</a></code> | <code>string</code> | ID of an AWS IAM role authorized to manage an AWS customer master key. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.enabledEncryptionAtRest">enabledEncryptionAtRest</a></code> | <code>boolean</code> | Specifies whether Encryption at Rest is enabled for an Atlas project. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.region">region</a></code> | <code>string</code> | The AWS region in which the AWS customer master key exists. |

---

##### `customerMasterKeyId`<sup>Required</sup> <a name="customerMasterKeyId" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.customerMasterKeyId"></a>

```typescript
public readonly customerMasterKeyId: string;
```

- *Type:* string

The AWS customer master key used to encrypt and decrypt the MongoDB master keys.

---

##### `roleId`<sup>Required</sup> <a name="roleId" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.roleId"></a>

```typescript
public readonly roleId: string;
```

- *Type:* string

ID of an AWS IAM role authorized to manage an AWS customer master key.

---

##### `enabledEncryptionAtRest`<sup>Optional</sup> <a name="enabledEncryptionAtRest" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.enabledEncryptionAtRest"></a>

```typescript
public readonly enabledEncryptionAtRest: boolean;
```

- *Type:* boolean

Specifies whether Encryption at Rest is enabled for an Atlas project.

To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
Default Value: true

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/encryption-at-rest-express.AtlasEncryptionAtRestProps.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

The AWS region in which the AWS customer master key exists.

---



