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
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrFrequencyType">attrFrequencyType</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.FrequencyType`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Id`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrLinks">attrLinks</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Links`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMasterKeyUUID">attrMasterKeyUUID</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.MasterKeyUUID`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMembers">attrMembers</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Members`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMongodVersion">attrMongodVersion</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.MongodVersion`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrPolicyItems">attrPolicyItems</a></code> | <code>string[]</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.PolicyItems`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrReplicaSetName">attrReplicaSetName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.ReplicaSetName`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrResults">attrResults</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Results`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotId">attrSnapshotId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotId`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotIds">attrSnapshotIds</a></code> | <code>string[]</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotIds`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotType">attrSnapshotType</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotType`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrStatus">attrStatus</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.Status`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrStorageSizeBytes">attrStorageSizeBytes</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.StorageSizeBytes`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrTotalCount">attrTotalCount</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::CloudBackupSnapshot.TotalCount`. |
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

##### `attrFrequencyType`<sup>Required</sup> <a name="attrFrequencyType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrFrequencyType"></a>

```typescript
public readonly attrFrequencyType: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.FrequencyType`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Id`.

---

##### `attrLinks`<sup>Required</sup> <a name="attrLinks" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrLinks"></a>

```typescript
public readonly attrLinks: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Links`.

---

##### `attrMasterKeyUUID`<sup>Required</sup> <a name="attrMasterKeyUUID" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMasterKeyUUID"></a>

```typescript
public readonly attrMasterKeyUUID: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.MasterKeyUUID`.

---

##### `attrMembers`<sup>Required</sup> <a name="attrMembers" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMembers"></a>

```typescript
public readonly attrMembers: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Members`.

---

##### `attrMongodVersion`<sup>Required</sup> <a name="attrMongodVersion" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrMongodVersion"></a>

```typescript
public readonly attrMongodVersion: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.MongodVersion`.

---

##### `attrPolicyItems`<sup>Required</sup> <a name="attrPolicyItems" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrPolicyItems"></a>

```typescript
public readonly attrPolicyItems: string[];
```

- *Type:* string[]

Attribute `MongoDB::Atlas::CloudBackupSnapshot.PolicyItems`.

---

##### `attrReplicaSetName`<sup>Required</sup> <a name="attrReplicaSetName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrReplicaSetName"></a>

```typescript
public readonly attrReplicaSetName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.ReplicaSetName`.

---

##### `attrResults`<sup>Required</sup> <a name="attrResults" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrResults"></a>

```typescript
public readonly attrResults: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::CloudBackupSnapshot.Results`.

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

##### `attrSnapshotType`<sup>Required</sup> <a name="attrSnapshotType" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrSnapshotType"></a>

```typescript
public readonly attrSnapshotType: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSnapshot.SnapshotType`.

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

##### `attrTotalCount`<sup>Required</sup> <a name="attrTotalCount" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshot.property.attrTotalCount"></a>

```typescript
public readonly attrTotalCount: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::CloudBackupSnapshot.TotalCount`.

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

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/cloud-backup-snapshot'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

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
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.clusterName">clusterName</a></code> | <code>string</code> | Human-readable label that identifies the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.description">description</a></code> | <code>string</code> | Human-readable phrase or sentence that explains the purpose of the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.groupId">groupId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.includeCount">includeCount</a></code> | <code>boolean</code> | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.instanceName">instanceName</a></code> | <code>string</code> | Human-readable label that identifies the serverless instance. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.itemsPerPage">itemsPerPage</a></code> | <code>number</code> | Number of items that the response returns per page. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.pageNum">pageNum</a></code> | <code>number</code> | Number of the page that displays the current set of the total objects that the response returns. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.retentionInDays">retentionInDays</a></code> | <code>number</code> | Number of days that MongoDB Cloud should retain the on-demand snapshot. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-snapshot.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `clusterName`<sup>Optional</sup> <a name="clusterName" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

Human-readable label that identifies the cluster.

---

##### `description`<sup>Optional</sup> <a name="description" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

Human-readable phrase or sentence that explains the purpose of the snapshot.

The resource returns this parameter when `"status": "onDemand"`.

---

##### `groupId`<sup>Optional</sup> <a name="groupId" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

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

##### `pageNum`<sup>Optional</sup> <a name="pageNum" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.pageNum"></a>

```typescript
public readonly pageNum: number;
```

- *Type:* number

Number of the page that displays the current set of the total objects that the response returns.

---

##### `retentionInDays`<sup>Optional</sup> <a name="retentionInDays" id="@mongodbatlas-awscdk/cloud-backup-snapshot.CfnCloudBackupSnapshotProps.property.retentionInDays"></a>

```typescript
public readonly retentionInDays: number;
```

- *Type:* number

Number of days that MongoDB Cloud should retain the on-demand snapshot.

Must be at least **1**

---



