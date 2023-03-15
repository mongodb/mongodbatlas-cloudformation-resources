# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCloudBackupSnapshot <a name="CfnCloudBackupSnapshot" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot"></a>

A CloudFormation `MongoDB::Atlas::CloudBackupSnapshot`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer"></a>

```typescript
import { CfnCloudBackupSnapshot } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

new CfnCloudBackupSnapshot(scope: Construct, id: string, props: CfnCloudBackupSnapshotProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps">CfnCloudBackupSnapshotProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps">CfnCloudBackupSnapshotProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addMetadata"></a>

```typescript
public addMetadata(key: string, value: any): void
```

Add a value to the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addOverride"></a>

```typescript
public addOverride(path: string, value: any): void
```

Adds an override to the synthesized CloudFormation resource.

To add a
property override, either use `addPropertyOverride` or prefix `path` with
"Properties." (i.e. `Properties.TopicName`).

If the override is nested, separate each nested level using a dot (.) in the path parameter.
If there is an array as part of the nesting, specify the index in the path.

To include a literal `.` in the property name, prefix with a `\`. In most
programming languages you will need to write this as `"\\."` because the
`\` itself will need to be escaped.

For example,
```typescript
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
```
would add the overrides
```json
"Properties": {
   "GlobalSecondaryIndexes": [
     {
       "Projection": {
         "NonKeyAttributes": [ "myattribute" ]
         ...
       }
       ...
     },
     {
       "ProjectionType": "INCLUDE"
       ...
     },
   ]
   ...
}
```

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy?: RemovalPolicy, options?: RemovalPolicyOptions): void
```

Sets the deletion policy of the resource based on the removal policy specified.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.getMetadata"></a>

```typescript
public getMetadata(key: string): any
```

Retrieve a value value from the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isConstruct"></a>

```typescript
import { CfnCloudBackupSnapshot } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

CfnCloudBackupSnapshot.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isCfnElement"></a>

```typescript
import { CfnCloudBackupSnapshot } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

CfnCloudBackupSnapshot.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isCfnResource"></a>

```typescript
import { CfnCloudBackupSnapshot } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

CfnCloudBackupSnapshot.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrCloudProvider">attrCloudProvider</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.CloudProvider`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrCreatedAt">attrCreatedAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.CreatedAt`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrExpiresAt">attrExpiresAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.ExpiresAt`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Id`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMasterKeyUUID">attrMasterKeyUUID</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.MasterKeyUUID`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMongodVersion">attrMongodVersion</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.MongodVersion`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrReplicaSetName">attrReplicaSetName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.ReplicaSetName`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotId">attrSnapshotId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotId`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotIds">attrSnapshotIds</a></code> | <code>string[]</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotIds`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrStatus">attrStatus</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Status`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrStorageSizeBytes">attrStorageSizeBytes</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.StorageSizeBytes`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrType">attrType</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Type`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps">CfnCloudBackupSnapshotProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrCloudProvider`<sup>Required</sup> <a name="attrCloudProvider" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrCloudProvider"></a>

```typescript
public readonly attrCloudProvider: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.CloudProvider`.

---

##### `attrCreatedAt`<sup>Required</sup> <a name="attrCreatedAt" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrCreatedAt"></a>

```typescript
public readonly attrCreatedAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.CreatedAt`.

---

##### `attrExpiresAt`<sup>Required</sup> <a name="attrExpiresAt" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrExpiresAt"></a>

```typescript
public readonly attrExpiresAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.ExpiresAt`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Id`.

---

##### `attrMasterKeyUUID`<sup>Required</sup> <a name="attrMasterKeyUUID" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMasterKeyUUID"></a>

```typescript
public readonly attrMasterKeyUUID: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.MasterKeyUUID`.

---

##### `attrMongodVersion`<sup>Required</sup> <a name="attrMongodVersion" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMongodVersion"></a>

```typescript
public readonly attrMongodVersion: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.MongodVersion`.

---

##### `attrReplicaSetName`<sup>Required</sup> <a name="attrReplicaSetName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrReplicaSetName"></a>

```typescript
public readonly attrReplicaSetName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.ReplicaSetName`.

---

##### `attrSnapshotId`<sup>Required</sup> <a name="attrSnapshotId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotId"></a>

```typescript
public readonly attrSnapshotId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotId`.

---

##### `attrSnapshotIds`<sup>Required</sup> <a name="attrSnapshotIds" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotIds"></a>

```typescript
public readonly attrSnapshotIds: string[];
```

- *Type:* string[]

Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotIds`.

---

##### `attrStatus`<sup>Required</sup> <a name="attrStatus" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrStatus"></a>

```typescript
public readonly attrStatus: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Status`.

---

##### `attrStorageSizeBytes`<sup>Required</sup> <a name="attrStorageSizeBytes" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrStorageSizeBytes"></a>

```typescript
public readonly attrStorageSizeBytes: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.StorageSizeBytes`.

---

##### `attrType`<sup>Required</sup> <a name="attrType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrType"></a>

```typescript
public readonly attrType: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Type`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.props"></a>

```typescript
public readonly props: CfnCloudBackupSnapshotProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps">CfnCloudBackupSnapshotProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiAtlasDiskBackupShardedClusterSnapshotMemberView <a name="ApiAtlasDiskBackupShardedClusterSnapshotMemberView" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.Initializer"></a>

```typescript
import { ApiAtlasDiskBackupShardedClusterSnapshotMemberView } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

const apiAtlasDiskBackupShardedClusterSnapshotMemberView: ApiAtlasDiskBackupShardedClusterSnapshotMemberView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.property.cloudProvider">cloudProvider</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider">ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider</a></code> | Human-readable label that identifies the cloud provider that stores this snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.property.id">id</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.property.replicaSetName">replicaSetName</a></code> | <code>string</code> | Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot. |

---

##### `cloudProvider`<sup>Optional</sup> <a name="cloudProvider" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.property.cloudProvider"></a>

```typescript
public readonly cloudProvider: ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider">ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider</a>

Human-readable label that identifies the cloud provider that stores this snapshot.

The resource returns this parameter when `"type": "replicaSet".`

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the snapshot.

---

##### `replicaSetName`<sup>Optional</sup> <a name="replicaSetName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView.property.replicaSetName"></a>

```typescript
public readonly replicaSetName: string;
```

- *Type:* string

Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot.

---

### ApiAtlasDiskBackupShardedClusterSnapshotView <a name="ApiAtlasDiskBackupShardedClusterSnapshotView" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.Initializer"></a>

```typescript
import { ApiAtlasDiskBackupShardedClusterSnapshotView } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

const apiAtlasDiskBackupShardedClusterSnapshotView: ApiAtlasDiskBackupShardedClusterSnapshotView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.createdAt">createdAt</a></code> | <code>string</code> | Date and time when MongoDB Cloud took the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.description">description</a></code> | <code>string</code> | Human-readable phrase or sentence that explains the purpose of the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.expiresAt">expiresAt</a></code> | <code>string</code> | Date and time when MongoDB Cloud deletes the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.frequencyType">frequencyType</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType">ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType</a></code> | Human-readable label that identifies how often this snapshot triggers. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.id">id</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.links">links</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.Link">Link</a>[]</code> | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.masterKeyUuid">masterKeyUuid</a></code> | <code>string</code> | Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.members">members</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>[]</code> | List that includes the snapshots and the cloud provider that stores the snapshots. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.mongodVersion">mongodVersion</a></code> | <code>string</code> | Version of the MongoDB host that this snapshot backs up. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.policyItems">policyItems</a></code> | <code>string[]</code> | List that contains unique identifiers for the policy items. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.snapshotIds">snapshotIds</a></code> | <code>string[]</code> | List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.snapshotType">snapshotType</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType">ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType</a></code> | Human-readable label that identifies when this snapshot triggers. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.status">status</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus">ApiAtlasDiskBackupShardedClusterSnapshotViewStatus</a></code> | Human-readable label that indicates the stage of the backup process for this snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.storageSizeBytes">storageSizeBytes</a></code> | <code>string</code> | Number of bytes taken to store the backup snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.type">type</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType">ApiAtlasDiskBackupShardedClusterSnapshotViewType</a></code> | Human-readable label that categorizes the cluster as a replica set or sharded cluster. |

---

##### `createdAt`<sup>Optional</sup> <a name="createdAt" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.createdAt"></a>

```typescript
public readonly createdAt: string;
```

- *Type:* string

Date and time when MongoDB Cloud took the snapshot.

This parameter expresses its value in the ISO 8601 timestamp format in UTC.

---

##### `description`<sup>Optional</sup> <a name="description" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

Human-readable phrase or sentence that explains the purpose of the snapshot.

The resource returns this parameter when `"status": "onDemand"`.

---

##### `expiresAt`<sup>Optional</sup> <a name="expiresAt" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.expiresAt"></a>

```typescript
public readonly expiresAt: string;
```

- *Type:* string

Date and time when MongoDB Cloud deletes the snapshot.

This parameter expresses its value in the ISO 8601 timestamp format in UTC.

---

##### `frequencyType`<sup>Optional</sup> <a name="frequencyType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.frequencyType"></a>

```typescript
public readonly frequencyType: ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType">ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType</a>

Human-readable label that identifies how often this snapshot triggers.

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the snapshot.

---

##### `links`<sup>Optional</sup> <a name="links" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.links"></a>

```typescript
public readonly links: Link[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.Link">Link</a>[]

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.

RFC 5988 outlines these relationships.

---

##### `masterKeyUuid`<sup>Optional</sup> <a name="masterKeyUuid" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.masterKeyUuid"></a>

```typescript
public readonly masterKeyUuid: string;
```

- *Type:* string

Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot.

The resource returns this value when `"encryptionEnabled" : true`.

---

##### `members`<sup>Optional</sup> <a name="members" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.members"></a>

```typescript
public readonly members: ApiAtlasDiskBackupShardedClusterSnapshotMemberView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>[]

List that includes the snapshots and the cloud provider that stores the snapshots.

The resource returns this parameter when `"type" : "SHARDED_CLUSTER"`.

---

##### `mongodVersion`<sup>Optional</sup> <a name="mongodVersion" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.mongodVersion"></a>

```typescript
public readonly mongodVersion: string;
```

- *Type:* string

Version of the MongoDB host that this snapshot backs up.

---

##### `policyItems`<sup>Optional</sup> <a name="policyItems" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.policyItems"></a>

```typescript
public readonly policyItems: string[];
```

- *Type:* string[]

List that contains unique identifiers for the policy items.

---

##### `snapshotIds`<sup>Optional</sup> <a name="snapshotIds" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.snapshotIds"></a>

```typescript
public readonly snapshotIds: string[];
```

- *Type:* string[]

List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster.

The resource returns this parameter when `"type": "SHARDED_CLUSTER"`. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name.

---

##### `snapshotType`<sup>Optional</sup> <a name="snapshotType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.snapshotType"></a>

```typescript
public readonly snapshotType: ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType">ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType</a>

Human-readable label that identifies when this snapshot triggers.

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.status"></a>

```typescript
public readonly status: ApiAtlasDiskBackupShardedClusterSnapshotViewStatus;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus">ApiAtlasDiskBackupShardedClusterSnapshotViewStatus</a>

Human-readable label that indicates the stage of the backup process for this snapshot.

---

##### `storageSizeBytes`<sup>Optional</sup> <a name="storageSizeBytes" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.storageSizeBytes"></a>

```typescript
public readonly storageSizeBytes: string;
```

- *Type:* string

Number of bytes taken to store the backup snapshot.

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView.property.type"></a>

```typescript
public readonly type: ApiAtlasDiskBackupShardedClusterSnapshotViewType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType">ApiAtlasDiskBackupShardedClusterSnapshotViewType</a>

Human-readable label that categorizes the cluster as a replica set or sharded cluster.

---

### CfnCloudBackupSnapshotProps <a name="CfnCloudBackupSnapshotProps" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps"></a>

Returns, takes, and removes Cloud Backup snapshots.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.Initializer"></a>

```typescript
import { CfnCloudBackupSnapshotProps } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

const cfnCloudBackupSnapshotProps: CfnCloudBackupSnapshotProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.clusterName">clusterName</a></code> | <code>string</code> | Human-readable label that identifies the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.description">description</a></code> | <code>string</code> | Human-readable phrase or sentence that explains the purpose of the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.frequencyType">frequencyType</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType">CfnCloudBackupSnapshotPropsFrequencyType</a></code> | Human-readable label that identifies how often this snapshot triggers. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.includeCount">includeCount</a></code> | <code>boolean</code> | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.instanceName">instanceName</a></code> | <code>string</code> | Human-readable label that identifies the serverless instance. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.itemsPerPage">itemsPerPage</a></code> | <code>number</code> | Number of items that the response returns per page. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.links">links</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.Link">Link</a>[]</code> | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.members">members</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>[]</code> | List that includes the snapshots and the cloud provider that stores the snapshots. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.pageNum">pageNum</a></code> | <code>number</code> | Number of the page that displays the current set of the total objects that the response returns. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.policyItems">policyItems</a></code> | <code>string[]</code> | List that contains unique identifiers for the policy items. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.results">results</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView">ApiAtlasDiskBackupShardedClusterSnapshotView</a>[]</code> | List of returned documents that MongoDB Cloud provides when completing this request. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.retentionInDays">retentionInDays</a></code> | <code>number</code> | Number of days that MongoDB Cloud should retain the on-demand snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.snapshotType">snapshotType</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType">CfnCloudBackupSnapshotPropsSnapshotType</a></code> | Human-readable label that identifies when this snapshot triggers. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.totalCount">totalCount</a></code> | <code>number</code> | Number of documents returned in this response. |

---

##### `clusterName`<sup>Required</sup> <a name="clusterName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

Human-readable label that identifies the cluster.

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `description`<sup>Optional</sup> <a name="description" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

Human-readable phrase or sentence that explains the purpose of the snapshot.

The resource returns this parameter when `"status": "onDemand"`.

---

##### `frequencyType`<sup>Optional</sup> <a name="frequencyType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.frequencyType"></a>

```typescript
public readonly frequencyType: CfnCloudBackupSnapshotPropsFrequencyType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType">CfnCloudBackupSnapshotPropsFrequencyType</a>

Human-readable label that identifies how often this snapshot triggers.

---

##### `includeCount`<sup>Optional</sup> <a name="includeCount" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.includeCount"></a>

```typescript
public readonly includeCount: boolean;
```

- *Type:* boolean

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

---

##### `instanceName`<sup>Optional</sup> <a name="instanceName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.instanceName"></a>

```typescript
public readonly instanceName: string;
```

- *Type:* string

Human-readable label that identifies the serverless instance.

---

##### `itemsPerPage`<sup>Optional</sup> <a name="itemsPerPage" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.itemsPerPage"></a>

```typescript
public readonly itemsPerPage: number;
```

- *Type:* number

Number of items that the response returns per page.

---

##### `links`<sup>Optional</sup> <a name="links" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.links"></a>

```typescript
public readonly links: Link[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.Link">Link</a>[]

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.

RFC 5988 outlines these relationships.

---

##### `members`<sup>Optional</sup> <a name="members" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.members"></a>

```typescript
public readonly members: ApiAtlasDiskBackupShardedClusterSnapshotMemberView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberView">ApiAtlasDiskBackupShardedClusterSnapshotMemberView</a>[]

List that includes the snapshots and the cloud provider that stores the snapshots.

The resource returns this parameter when `"type" : "SHARDED_CLUSTER"`.

---

##### `pageNum`<sup>Optional</sup> <a name="pageNum" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.pageNum"></a>

```typescript
public readonly pageNum: number;
```

- *Type:* number

Number of the page that displays the current set of the total objects that the response returns.

---

##### `policyItems`<sup>Optional</sup> <a name="policyItems" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.policyItems"></a>

```typescript
public readonly policyItems: string[];
```

- *Type:* string[]

List that contains unique identifiers for the policy items.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---

##### `results`<sup>Optional</sup> <a name="results" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.results"></a>

```typescript
public readonly results: ApiAtlasDiskBackupShardedClusterSnapshotView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotView">ApiAtlasDiskBackupShardedClusterSnapshotView</a>[]

List of returned documents that MongoDB Cloud provides when completing this request.

---

##### `retentionInDays`<sup>Optional</sup> <a name="retentionInDays" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.retentionInDays"></a>

```typescript
public readonly retentionInDays: number;
```

- *Type:* number

Number of days that MongoDB Cloud should retain the on-demand snapshot.

Must be at least **1**

---

##### `snapshotType`<sup>Optional</sup> <a name="snapshotType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.snapshotType"></a>

```typescript
public readonly snapshotType: CfnCloudBackupSnapshotPropsSnapshotType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType">CfnCloudBackupSnapshotPropsSnapshotType</a>

Human-readable label that identifies when this snapshot triggers.

---

##### `totalCount`<sup>Optional</sup> <a name="totalCount" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.totalCount"></a>

```typescript
public readonly totalCount: number;
```

- *Type:* number

Number of documents returned in this response.

---

### Link <a name="Link" id="@mongodbatlas-awscdk/cloud-backup-snapshot.Link"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-snapshot.Link.Initializer"></a>

```typescript
import { Link } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

const link: Link = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.Link.property.href">href</a></code> | <code>string</code> | Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.Link.property.rel">rel</a></code> | <code>string</code> | Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. |

---

##### `href`<sup>Optional</sup> <a name="href" id="@mongodbatlas-awscdk/cloud-backup-snapshot.Link.property.href"></a>

```typescript
public readonly href: string;
```

- *Type:* string

Uniform Resource Locator (URL) that points another API resource to which this response has some relationship.

This URL often begins with `https://mms.mongodb.com`.

---

##### `rel`<sup>Optional</sup> <a name="rel" id="@mongodbatlas-awscdk/cloud-backup-snapshot.Link.property.rel"></a>

```typescript
public readonly rel: string;
```

- *Type:* string

Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource.

This URL often begins with `https://mms.mongodb.com`.

---



## Enums <a name="Enums" id="Enums"></a>

### ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider <a name="ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider"></a>

Human-readable label that identifies the cloud provider that stores this snapshot.

The resource returns this parameter when `"type": "replicaSet".`

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider.AWS">AWS</a></code> | AWS. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider.AZURE">AZURE</a></code> | AZURE. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider.GCP">GCP</a></code> | GCP. |

---

##### `AWS` <a name="AWS" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider.AWS"></a>

AWS.

---


##### `AZURE` <a name="AZURE" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider.AZURE"></a>

AZURE.

---


##### `GCP` <a name="GCP" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider.GCP"></a>

GCP.

---


### ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType <a name="ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType"></a>

Human-readable label that identifies how often this snapshot triggers.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.HOURLY">HOURLY</a></code> | hourly. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.DAILY">DAILY</a></code> | daily. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.WEEKLY">WEEKLY</a></code> | weekly. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.MONTHLY">MONTHLY</a></code> | monthly. |

---

##### `HOURLY` <a name="HOURLY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.HOURLY"></a>

hourly.

---


##### `DAILY` <a name="DAILY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.DAILY"></a>

daily.

---


##### `WEEKLY` <a name="WEEKLY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.WEEKLY"></a>

weekly.

---


##### `MONTHLY` <a name="MONTHLY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType.MONTHLY"></a>

monthly.

---


### ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType <a name="ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType"></a>

Human-readable label that identifies when this snapshot triggers.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType.ON_DEMAND">ON_DEMAND</a></code> | onDemand. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType.SCHEDULED">SCHEDULED</a></code> | scheduled. |

---

##### `ON_DEMAND` <a name="ON_DEMAND" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType.ON_DEMAND"></a>

onDemand.

---


##### `SCHEDULED` <a name="SCHEDULED" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType.SCHEDULED"></a>

scheduled.

---


### ApiAtlasDiskBackupShardedClusterSnapshotViewStatus <a name="ApiAtlasDiskBackupShardedClusterSnapshotViewStatus" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus"></a>

Human-readable label that indicates the stage of the backup process for this snapshot.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.QUEUED">QUEUED</a></code> | queued. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.IN_PROGRESS">IN_PROGRESS</a></code> | inProgress. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.COMPLETED">COMPLETED</a></code> | completed. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.FAILED">FAILED</a></code> | failed. |

---

##### `QUEUED` <a name="QUEUED" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.QUEUED"></a>

queued.

---


##### `IN_PROGRESS` <a name="IN_PROGRESS" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.IN_PROGRESS"></a>

inProgress.

---


##### `COMPLETED` <a name="COMPLETED" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.COMPLETED"></a>

completed.

---


##### `FAILED` <a name="FAILED" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus.FAILED"></a>

failed.

---


### ApiAtlasDiskBackupShardedClusterSnapshotViewType <a name="ApiAtlasDiskBackupShardedClusterSnapshotViewType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType"></a>

Human-readable label that categorizes the cluster as a replica set or sharded cluster.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType.REPLICA_SET">REPLICA_SET</a></code> | REPLICA_SET. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType.SHARDED_CLUSTER">SHARDED_CLUSTER</a></code> | SHARDED_CLUSTER. |

---

##### `REPLICA_SET` <a name="REPLICA_SET" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType.REPLICA_SET"></a>

REPLICA_SET.

---


##### `SHARDED_CLUSTER` <a name="SHARDED_CLUSTER" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiAtlasDiskBackupShardedClusterSnapshotViewType.SHARDED_CLUSTER"></a>

SHARDED_CLUSTER.

---


### CfnCloudBackupSnapshotPropsFrequencyType <a name="CfnCloudBackupSnapshotPropsFrequencyType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType"></a>

Human-readable label that identifies how often this snapshot triggers.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.HOURLY">HOURLY</a></code> | hourly. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.DAILY">DAILY</a></code> | daily. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.WEEKLY">WEEKLY</a></code> | weekly. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.MONTHLY">MONTHLY</a></code> | monthly. |

---

##### `HOURLY` <a name="HOURLY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.HOURLY"></a>

hourly.

---


##### `DAILY` <a name="DAILY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.DAILY"></a>

daily.

---


##### `WEEKLY` <a name="WEEKLY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.WEEKLY"></a>

weekly.

---


##### `MONTHLY` <a name="MONTHLY" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsFrequencyType.MONTHLY"></a>

monthly.

---


### CfnCloudBackupSnapshotPropsSnapshotType <a name="CfnCloudBackupSnapshotPropsSnapshotType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType"></a>

Human-readable label that identifies when this snapshot triggers.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType.ON_DEMAND">ON_DEMAND</a></code> | onDemand. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType.SCHEDULED">SCHEDULED</a></code> | scheduled. |

---

##### `ON_DEMAND` <a name="ON_DEMAND" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType.ON_DEMAND"></a>

onDemand.

---


##### `SCHEDULED` <a name="SCHEDULED" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotPropsSnapshotType.SCHEDULED"></a>

scheduled.

---

