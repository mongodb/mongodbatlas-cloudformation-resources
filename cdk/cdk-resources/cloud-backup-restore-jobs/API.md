# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCloudBackUpRestoreJobs <a name="CfnCloudBackUpRestoreJobs" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs"></a>

A CloudFormation `MongoDB::Atlas::CloudBackUpRestoreJobs`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer"></a>

```typescript
import { CfnCloudBackUpRestoreJobs } from '@mongodbatlas-awscdk/cloud-backup-restore-jobs'

new CfnCloudBackUpRestoreJobs(scope: Construct, id: string, props: CfnCloudBackUpRestoreJobsProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps">CfnCloudBackUpRestoreJobsProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps">CfnCloudBackUpRestoreJobsProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isConstruct"></a>

```typescript
import { CfnCloudBackUpRestoreJobs } from '@mongodbatlas-awscdk/cloud-backup-restore-jobs'

CfnCloudBackUpRestoreJobs.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isCfnElement"></a>

```typescript
import { CfnCloudBackUpRestoreJobs } from '@mongodbatlas-awscdk/cloud-backup-restore-jobs'

CfnCloudBackUpRestoreJobs.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isCfnResource"></a>

```typescript
import { CfnCloudBackUpRestoreJobs } from '@mongodbatlas-awscdk/cloud-backup-restore-jobs'

CfnCloudBackUpRestoreJobs.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrCreatedAt">attrCreatedAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.CreatedAt`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrDeliveryUrl">attrDeliveryUrl</a></code> | <code>string[]</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.DeliveryUrl`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrExpiresAt">attrExpiresAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.ExpiresAt`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrFinishedAt">attrFinishedAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.FinishedAt`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.Id`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrLinks">attrLinks</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.Links`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrTimestamp">attrTimestamp</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.Timestamp`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps">CfnCloudBackUpRestoreJobsProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrCreatedAt`<sup>Required</sup> <a name="attrCreatedAt" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrCreatedAt"></a>

```typescript
public readonly attrCreatedAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.CreatedAt`.

---

##### `attrDeliveryUrl`<sup>Required</sup> <a name="attrDeliveryUrl" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrDeliveryUrl"></a>

```typescript
public readonly attrDeliveryUrl: string[];
```

- *Type:* string[]

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.DeliveryUrl`.

---

##### `attrExpiresAt`<sup>Required</sup> <a name="attrExpiresAt" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrExpiresAt"></a>

```typescript
public readonly attrExpiresAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.ExpiresAt`.

---

##### `attrFinishedAt`<sup>Required</sup> <a name="attrFinishedAt" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrFinishedAt"></a>

```typescript
public readonly attrFinishedAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.FinishedAt`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.Id`.

---

##### `attrLinks`<sup>Required</sup> <a name="attrLinks" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrLinks"></a>

```typescript
public readonly attrLinks: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.Links`.

---

##### `attrTimestamp`<sup>Required</sup> <a name="attrTimestamp" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.attrTimestamp"></a>

```typescript
public readonly attrTimestamp: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackUpRestoreJobs.Timestamp`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.props"></a>

```typescript
public readonly props: CfnCloudBackUpRestoreJobsProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps">CfnCloudBackUpRestoreJobsProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobs.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/cloud-backup-restore-jobs'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnCloudBackUpRestoreJobsProps <a name="CfnCloudBackUpRestoreJobsProps" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps"></a>

Returns, starts, and cancels Cloud Backup restore jobs.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.Initializer"></a>

```typescript
import { CfnCloudBackUpRestoreJobsProps } from '@mongodbatlas-awscdk/cloud-backup-restore-jobs'

const cfnCloudBackUpRestoreJobsProps: CfnCloudBackUpRestoreJobsProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.projectId">projectId</a></code> | <code>string</code> | The unique identifier of the project for the Atlas cluster. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.cancelled">cancelled</a></code> | <code>boolean</code> | Indicates whether the restore job was canceled. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.clusterName">clusterName</a></code> | <code>string</code> | The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.deliveryType">deliveryType</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType">CfnCloudBackUpRestoreJobsPropsDeliveryType</a></code> | Type of restore job to create.The value can be any one of download,automated or point_in_time. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.expired">expired</a></code> | <code>boolean</code> | Indicates whether the restore job expired. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.instanceName">instanceName</a></code> | <code>string</code> | The instance name of the Serverless cluster whose snapshot you want to restore or you want to retrieve restore jobs. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.opLogInc">opLogInc</a></code> | <code>string</code> | Oplog operation number from which to you want to restore this snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.opLogTs">opLogTs</a></code> | <code>string</code> | Timestamp in the number of seconds that have elapsed since the UNIX epoch from which to you want to restore this snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.pointInTimeUtcSeconds">pointInTimeUtcSeconds</a></code> | <code>number</code> | If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.snapshotId">snapshotId</a></code> | <code>string</code> | Unique identifier of the source snapshot ID of the restore job. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.targetClusterName">targetClusterName</a></code> | <code>string</code> | Name of the target Atlas cluster to which the restore job restores the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.targetProjectId">targetProjectId</a></code> | <code>string</code> | Name of the target Atlas project of the restore job. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

The unique identifier of the project for the Atlas cluster.

---

##### `cancelled`<sup>Optional</sup> <a name="cancelled" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.cancelled"></a>

```typescript
public readonly cancelled: boolean;
```

- *Type:* boolean

Indicates whether the restore job was canceled.

---

##### `clusterName`<sup>Optional</sup> <a name="clusterName" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs.

---

##### `deliveryType`<sup>Optional</sup> <a name="deliveryType" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.deliveryType"></a>

```typescript
public readonly deliveryType: CfnCloudBackUpRestoreJobsPropsDeliveryType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType">CfnCloudBackUpRestoreJobsPropsDeliveryType</a>

Type of restore job to create.The value can be any one of download,automated or point_in_time.

---

##### `expired`<sup>Optional</sup> <a name="expired" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.expired"></a>

```typescript
public readonly expired: boolean;
```

- *Type:* boolean

Indicates whether the restore job expired.

---

##### `instanceName`<sup>Optional</sup> <a name="instanceName" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.instanceName"></a>

```typescript
public readonly instanceName: string;
```

- *Type:* string

The instance name of the Serverless cluster whose snapshot you want to restore or you want to retrieve restore jobs.

---

##### `opLogInc`<sup>Optional</sup> <a name="opLogInc" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.opLogInc"></a>

```typescript
public readonly opLogInc: string;
```

- *Type:* string

Oplog operation number from which to you want to restore this snapshot.

This is the second part of an Oplog timestamp.

---

##### `opLogTs`<sup>Optional</sup> <a name="opLogTs" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.opLogTs"></a>

```typescript
public readonly opLogTs: string;
```

- *Type:* string

Timestamp in the number of seconds that have elapsed since the UNIX epoch from which to you want to restore this snapshot.

This is the first part of an Oplog timestamp.

---

##### `pointInTimeUtcSeconds`<sup>Optional</sup> <a name="pointInTimeUtcSeconds" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.pointInTimeUtcSeconds"></a>

```typescript
public readonly pointInTimeUtcSeconds: number;
```

- *Type:* number

If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used.

---

##### `snapshotId`<sup>Optional</sup> <a name="snapshotId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.snapshotId"></a>

```typescript
public readonly snapshotId: string;
```

- *Type:* string

Unique identifier of the source snapshot ID of the restore job.

---

##### `targetClusterName`<sup>Optional</sup> <a name="targetClusterName" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.targetClusterName"></a>

```typescript
public readonly targetClusterName: string;
```

- *Type:* string

Name of the target Atlas cluster to which the restore job restores the snapshot.

Only visible if deliveryType is automated.

---

##### `targetProjectId`<sup>Optional</sup> <a name="targetProjectId" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsProps.property.targetProjectId"></a>

```typescript
public readonly targetProjectId: string;
```

- *Type:* string

Name of the target Atlas project of the restore job.

Only visible if deliveryType is automated.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnCloudBackUpRestoreJobsPropsDeliveryType <a name="CfnCloudBackUpRestoreJobsPropsDeliveryType" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType"></a>

Type of restore job to create.The value can be any one of download,automated or point_in_time.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType.DOWNLOAD">DOWNLOAD</a></code> | download. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType.AUTOMATED">AUTOMATED</a></code> | automated. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType.POINT_IN_TIME">POINT_IN_TIME</a></code> | pointInTime. |

---

##### `DOWNLOAD` <a name="DOWNLOAD" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType.DOWNLOAD"></a>

download.

---


##### `AUTOMATED` <a name="AUTOMATED" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType.AUTOMATED"></a>

automated.

---


##### `POINT_IN_TIME` <a name="POINT_IN_TIME" id="@mongodbatlas-awscdk/cloud-backup-restore-jobs.CfnCloudBackUpRestoreJobsPropsDeliveryType.POINT_IN_TIME"></a>

pointInTime.

---

