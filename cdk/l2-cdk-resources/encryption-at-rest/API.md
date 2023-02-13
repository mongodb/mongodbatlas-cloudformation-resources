# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### AtlasEncryptionAtRest <a name="AtlasEncryptionAtRest" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer"></a>

```typescript
import { AtlasEncryptionAtRest } from '@mongodbatlas-awscdk/atlas-encryption-at-rest'

new AtlasEncryptionAtRest(scope: Construct, id: string, props: AtlasEncryptionAtRestProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps">AtlasEncryptionAtRestProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps">AtlasEncryptionAtRestProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.isConstruct"></a>

```typescript
import { AtlasEncryptionAtRest } from '@mongodbatlas-awscdk/atlas-encryption-at-rest'

AtlasEncryptionAtRest.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.cfnEncryptionAtRest">cfnEncryptionAtRest</a></code> | <code>@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `cfnEncryptionAtRest`<sup>Required</sup> <a name="cfnEncryptionAtRest" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.cfnEncryptionAtRest"></a>

```typescript
public readonly cfnEncryptionAtRest: CfnEncryptionAtRest;
```

- *Type:* @mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest

---


## Structs <a name="Structs" id="Structs"></a>

### AtlasEncryptionAtRestProps <a name="AtlasEncryptionAtRestProps" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.Initializer"></a>

```typescript
import { AtlasEncryptionAtRestProps } from '@mongodbatlas-awscdk/atlas-encryption-at-rest'

const atlasEncryptionAtRestProps: AtlasEncryptionAtRestProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.customerMasterKeyId">customerMasterKeyId</a></code> | <code>string</code> | The AWS customer master key used to encrypt and decrypt the MongoDB master keys. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the Atlas project to which the user belongs. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.roleId">roleId</a></code> | <code>string</code> | ID of an AWS IAM role authorized to manage an AWS customer master key. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.enabled">enabled</a></code> | <code>boolean</code> | Specifies whether Encryption at Rest is enabled for an Atlas project. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.region">region</a></code> | <code>string</code> | The AWS region in which the AWS customer master key exists. |

---

##### `customerMasterKeyId`<sup>Required</sup> <a name="customerMasterKeyId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.customerMasterKeyId"></a>

```typescript
public readonly customerMasterKeyId: string;
```

- *Type:* string

The AWS customer master key used to encrypt and decrypt the MongoDB master keys.

---

##### `privateKey`<sup>Required</sup> <a name="privateKey" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique identifier of the Atlas project to which the user belongs.

---

##### `publicKey`<sup>Required</sup> <a name="publicKey" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

##### `roleId`<sup>Required</sup> <a name="roleId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.roleId"></a>

```typescript
public readonly roleId: string;
```

- *Type:* string

ID of an AWS IAM role authorized to manage an AWS customer master key.

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Specifies whether Encryption at Rest is enabled for an Atlas project.

To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
Default Value: true

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

The AWS region in which the AWS customer master key exists.

Default Value: US_EAST_1

---



