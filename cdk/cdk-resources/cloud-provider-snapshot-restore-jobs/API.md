# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCloudProviderSnapshotRestoreJobs <a name="CfnCloudProviderSnapshotRestoreJobs" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs"></a>

A CloudFormation `MongoDB::Atlas::CloudProviderSnapshotRestoreJobs`.

#### Initializers <a name="Initializers" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer"></a>

```typescript
import { CfnCloudProviderSnapshotRestoreJobs } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

new CfnCloudProviderSnapshotRestoreJobs(scope: Construct, id: string, props: CfnCloudProviderSnapshotRestoreJobsProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps">CfnCloudProviderSnapshotRestoreJobsProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps">CfnCloudProviderSnapshotRestoreJobsProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isConstruct"></a>

```typescript
import { CfnCloudProviderSnapshotRestoreJobs } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

CfnCloudProviderSnapshotRestoreJobs.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isCfnElement"></a>

```typescript
import { CfnCloudProviderSnapshotRestoreJobs } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

CfnCloudProviderSnapshotRestoreJobs.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isCfnResource"></a>

```typescript
import { CfnCloudProviderSnapshotRestoreJobs } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

CfnCloudProviderSnapshotRestoreJobs.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudProviderSnapshotRestoreJobs.Id`. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps">CfnCloudProviderSnapshotRestoreJobsProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudProviderSnapshotRestoreJobs.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.props"></a>

```typescript
public readonly props: CfnCloudProviderSnapshotRestoreJobsProps;
```

- *Type:* <a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps">CfnCloudProviderSnapshotRestoreJobsProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobs.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnCloudProviderSnapshotRestoreJobsProps <a name="CfnCloudProviderSnapshotRestoreJobsProps" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps"></a>

This resource allows you to create, cancel, get one or list all cloud provider snapshot restore jobs.

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.Initializer"></a>

```typescript
import { CfnCloudProviderSnapshotRestoreJobsProps } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

const cfnCloudProviderSnapshotRestoreJobsProps: CfnCloudProviderSnapshotRestoreJobsProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.clusterName">clusterName</a></code> | <code>string</code> | The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.projectId">projectId</a></code> | <code>string</code> | The unique identifier of the project for the Atlas cluster. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.cancelled">cancelled</a></code> | <code>boolean</code> | Indicates whether the restore job was canceled. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.createdAt">createdAt</a></code> | <code>string</code> | UTC ISO 8601 formatted point in time when Atlas created the restore job. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.deliveryType">deliveryType</a></code> | <code>string</code> | Type of restore job to create. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.deliveryUrl">deliveryUrl</a></code> | <code>string[]</code> | One or more URLs for the compressed snapshot files for manual download. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.expired">expired</a></code> | <code>boolean</code> | Indicates whether the restore job expired. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.expiresAt">expiresAt</a></code> | <code>string</code> | UTC ISO 8601 formatted point in time when the restore job expires. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.finishedAt">finishedAt</a></code> | <code>string</code> | UTC ISO 8601 formatted point in time when the restore job completed. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.links">links</a></code> | <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks">CfnCloudProviderSnapshotRestoreJobsPropsLinks</a>[]</code> | One or more links to sub-resources and/or related resources. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.opLogTs">opLogTs</a></code> | <code>string</code> | If you performed a Point-in-Time restores at a time specified by a timestamp from the oplog, oplogTs indicates the timestamp used. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.pointInTimeUtcSeconds">pointInTimeUtcSeconds</a></code> | <code>number</code> | If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.snapshotId">snapshotId</a></code> | <code>string</code> | Unique identifier of the source snapshot ID of the restore job. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.targetClusterName">targetClusterName</a></code> | <code>string</code> | Name of the target Atlas cluster to which the restore job restores the snapshot. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.targetProjectId">targetProjectId</a></code> | <code>string</code> | Name of the target Atlas project of the restore job. |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.timestamp">timestamp</a></code> | <code>string</code> | Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `clusterName`<sup>Required</sup> <a name="clusterName" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs.

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

The unique identifier of the project for the Atlas cluster.

---

##### `cancelled`<sup>Optional</sup> <a name="cancelled" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.cancelled"></a>

```typescript
public readonly cancelled: boolean;
```

- *Type:* boolean

Indicates whether the restore job was canceled.

---

##### `createdAt`<sup>Optional</sup> <a name="createdAt" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.createdAt"></a>

```typescript
public readonly createdAt: string;
```

- *Type:* string

UTC ISO 8601 formatted point in time when Atlas created the restore job.

---

##### `deliveryType`<sup>Optional</sup> <a name="deliveryType" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.deliveryType"></a>

```typescript
public readonly deliveryType: string;
```

- *Type:* string

Type of restore job to create.

---

##### `deliveryUrl`<sup>Optional</sup> <a name="deliveryUrl" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.deliveryUrl"></a>

```typescript
public readonly deliveryUrl: string[];
```

- *Type:* string[]

One or more URLs for the compressed snapshot files for manual download.

Only visible if deliveryType is download.

---

##### `expired`<sup>Optional</sup> <a name="expired" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.expired"></a>

```typescript
public readonly expired: boolean;
```

- *Type:* boolean

Indicates whether the restore job expired.

---

##### `expiresAt`<sup>Optional</sup> <a name="expiresAt" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.expiresAt"></a>

```typescript
public readonly expiresAt: string;
```

- *Type:* string

UTC ISO 8601 formatted point in time when the restore job expires.

---

##### `finishedAt`<sup>Optional</sup> <a name="finishedAt" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.finishedAt"></a>

```typescript
public readonly finishedAt: string;
```

- *Type:* string

UTC ISO 8601 formatted point in time when the restore job completed.

---

##### `links`<sup>Optional</sup> <a name="links" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.links"></a>

```typescript
public readonly links: CfnCloudProviderSnapshotRestoreJobsPropsLinks[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks">CfnCloudProviderSnapshotRestoreJobsPropsLinks</a>[]

One or more links to sub-resources and/or related resources.

---

##### `opLogTs`<sup>Optional</sup> <a name="opLogTs" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.opLogTs"></a>

```typescript
public readonly opLogTs: string;
```

- *Type:* string

If you performed a Point-in-Time restores at a time specified by a timestamp from the oplog, oplogTs indicates the timestamp used.

---

##### `pointInTimeUtcSeconds`<sup>Optional</sup> <a name="pointInTimeUtcSeconds" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.pointInTimeUtcSeconds"></a>

```typescript
public readonly pointInTimeUtcSeconds: number;
```

- *Type:* number

If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used.

---

##### `snapshotId`<sup>Optional</sup> <a name="snapshotId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.snapshotId"></a>

```typescript
public readonly snapshotId: string;
```

- *Type:* string

Unique identifier of the source snapshot ID of the restore job.

---

##### `targetClusterName`<sup>Optional</sup> <a name="targetClusterName" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.targetClusterName"></a>

```typescript
public readonly targetClusterName: string;
```

- *Type:* string

Name of the target Atlas cluster to which the restore job restores the snapshot.

Only visible if deliveryType is automated.

---

##### `targetProjectId`<sup>Optional</sup> <a name="targetProjectId" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.targetProjectId"></a>

```typescript
public readonly targetProjectId: string;
```

- *Type:* string

Name of the target Atlas project of the restore job.

Only visible if deliveryType is automated.

---

##### `timestamp`<sup>Optional</sup> <a name="timestamp" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsProps.property.timestamp"></a>

```typescript
public readonly timestamp: string;
```

- *Type:* string

Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.

---

### CfnCloudProviderSnapshotRestoreJobsPropsLinks <a name="CfnCloudProviderSnapshotRestoreJobsPropsLinks" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks.Initializer"></a>

```typescript
import { CfnCloudProviderSnapshotRestoreJobsPropsLinks } from '@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs'

const cfnCloudProviderSnapshotRestoreJobsPropsLinks: CfnCloudProviderSnapshotRestoreJobsPropsLinks = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks.property.href">href</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks.property.rel">rel</a></code> | <code>string</code> | *No description.* |

---

##### `href`<sup>Optional</sup> <a name="href" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks.property.href"></a>

```typescript
public readonly href: string;
```

- *Type:* string

---

##### `rel`<sup>Optional</sup> <a name="rel" id="@mongodb-cdk/atlas-cloud-provider-snapshot-restore-jobs.CfnCloudProviderSnapshotRestoreJobsPropsLinks.property.rel"></a>

```typescript
public readonly rel: string;
```

- *Type:* string

---



