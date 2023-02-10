# @mongodbatlas-awscdk/atlas-encryption-at-rest

# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### AtlasEncryptionAtRest <a name="AtlasEncryptionAtRest" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest"></a>

A CloudFormation `MongoDB::Atlas::EncryptionAtRest`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer"></a>

```typescript
import { AtlasEncryptionAtRest } from '@mongodbatlas-awscdk/atlas-encryption-at-rest'

new AtlasEncryptionAtRest(scope: Construct, id: string, props: AtlasEncryptionAtRest)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.props">props</a></code> | <code>AtlasEncryptionAtRest</code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest">AtlasEncryptionAtRest

resource properties.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::EncryptionAtRest.Id`. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.props">props</a></code> | <code><a h@mongodbatlas-awscdk/atlas-encryption-at-restt-rest.AtlasEncryptionAtRest">CfnAtlasEncryptionAtRest></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::EncryptionAtRest.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.CfnEncryptionAtRest.property.props"></a>

```typescript
public readonly props: AtlasEncryptionAtRest;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest">AtlasEncryptionAtRest>

Resource props.

---

## Structs <a name="Structs" id="Structs"></a>

### AtlasEncryptionAtRestProps <a name="AtlasEncryptionAtRestProps" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRestProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.ApiKeyDefinition.Initializer"></a>

```typescript
import { AtlasEncryptionAtRestProps } from '@mongodbatlas-awscdk/atlas-encryption-at-rest'

const atlasEncryptionAtRestProps: AtlasEncryptionAtRestProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | Atlas API Public Key |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | Atlas API Private Key |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.customerMasterKeyId">customerMasterKeyId</a></code> | <code>string</code> | The AWS customer master key used to encrypt and decrypt the MongoDB master keys. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.enabled">enabled</a></code> | <code>boolean</code> | Specifies whether Encryption at Rest is enabled for an Atlas project. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.region">region</a></code> | <code>string</code> | The AWS region in which the AWS customer master key exists. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.roleId">roleId</a></code> | <code>string</code> | ID of an AWS IAM role authorized to manage an AWS customer master key. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the Atlas project to which the user belongs. |
---

##### `privateKey`<sup>Required</sup> <a name="privateKey" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Required</sup> <a name="publicKey" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---


##### `customerMasterKeyId`<sup>Required</sup> <a name="customerMasterKeyId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.customerMasterKeyId"></a>

```typescript
public readonly customerMasterKeyId: string;
```

- *Type:* string

The AWS customer master key used to encrypt and decrypt the MongoDB master keys.

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Specifies whether Encryption at Rest is enabled for an Atlas project.

To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

The AWS region in which the AWS customer master key exists.

---

##### `roleId`<sup>Required</sup> <a name="roleId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AwsKmsConfiguration.property.roleId"></a>

```typescript
public readonly roleId: string;
```

- *Type:* string

ID of an AWS IAM role authorized to manage an AWS customer master key.

---

### AtlasEncryptionAtRest


#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.Initializer"></a>

```typescript
import { AtlasEncryptionAtRest } from '@mongodbatlas-awscdk/atlas-encryption-at-rest'

const atlasEncryptionAtRest: AtlasEncryptionAtRest{ ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.awsKms">awsKms</a></code> | <code><a h@mongodbatlas-awscdk/atlas-encryption-at-restt-rest.AwsKmsConfiguration">AwsKmsConfiguration</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the Atlas project to which the user belongs. |
| <code><a href="#@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.apiKeys">apiKeys</a></code> | <code><a h@mongodbatlas-awscdk/atlas-encryption-at-restt-rest.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |



##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-encryption-at-rest.AtlasEncryptionAtRest.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique identifier of the Atlas project to which the user belongs.



