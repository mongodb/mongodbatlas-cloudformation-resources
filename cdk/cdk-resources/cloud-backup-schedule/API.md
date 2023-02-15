# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCloudBackupSchedule <a name="CfnCloudBackupSchedule" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule"></a>

A CloudFormation `MongoDB::Atlas::CloudBackupSchedule`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer"></a>

```typescript
import { CfnCloudBackupSchedule } from '@mongodbatlas-awscdk/cloud-backup-schedule'

new CfnCloudBackupSchedule(scope: Construct, id: string, props: CfnCloudBackupScheduleProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps">CfnCloudBackupScheduleProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps">CfnCloudBackupScheduleProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isConstruct"></a>

```typescript
import { CfnCloudBackupSchedule } from '@mongodbatlas-awscdk/cloud-backup-schedule'

CfnCloudBackupSchedule.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isCfnElement"></a>

```typescript
import { CfnCloudBackupSchedule } from '@mongodbatlas-awscdk/cloud-backup-schedule'

CfnCloudBackupSchedule.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isCfnResource"></a>

```typescript
import { CfnCloudBackupSchedule } from '@mongodbatlas-awscdk/cloud-backup-schedule'

CfnCloudBackupSchedule.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrClusterId">attrClusterId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSchedule.ClusterId`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrClusterName">attrClusterName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSchedule.ClusterName`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrNextSnapshot">attrNextSnapshot</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSchedule.NextSnapshot`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrProjectId">attrProjectId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudBackupSchedule.ProjectId`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps">CfnCloudBackupScheduleProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrClusterId`<sup>Required</sup> <a name="attrClusterId" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrClusterId"></a>

```typescript
public readonly attrClusterId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSchedule.ClusterId`.

---

##### `attrClusterName`<sup>Required</sup> <a name="attrClusterName" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrClusterName"></a>

```typescript
public readonly attrClusterName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSchedule.ClusterName`.

---

##### `attrNextSnapshot`<sup>Required</sup> <a name="attrNextSnapshot" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrNextSnapshot"></a>

```typescript
public readonly attrNextSnapshot: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSchedule.NextSnapshot`.

---

##### `attrProjectId`<sup>Required</sup> <a name="attrProjectId" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.attrProjectId"></a>

```typescript
public readonly attrProjectId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudBackupSchedule.ProjectId`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.props"></a>

```typescript
public readonly props: CfnCloudBackupScheduleProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps">CfnCloudBackupScheduleProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupSchedule.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiAtlasDiskBackupCopySettingView <a name="ApiAtlasDiskBackupCopySettingView" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.Initializer"></a>

```typescript
import { ApiAtlasDiskBackupCopySettingView } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const apiAtlasDiskBackupCopySettingView: ApiAtlasDiskBackupCopySettingView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.cloudProvider">cloudProvider</a></code> | <code>string</code> | A label that identifies the cloud provider that stores the snapshot copy. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.frequencies">frequencies</a></code> | <code>string[]</code> | List that describes which types of snapshots to copy. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.regionName">regionName</a></code> | <code>string</code> | Target region to copy snapshots belonging to replicationSpecId to. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.replicationSpecId">replicationSpecId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.shouldCopyOplogs">shouldCopyOplogs</a></code> | <code>boolean</code> | Flag that indicates whether to copy the oplogs to the target region. |

---

##### `cloudProvider`<sup>Optional</sup> <a name="cloudProvider" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.cloudProvider"></a>

```typescript
public readonly cloudProvider: string;
```

- *Type:* string

A label that identifies the cloud provider that stores the snapshot copy.

---

##### `frequencies`<sup>Optional</sup> <a name="frequencies" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.frequencies"></a>

```typescript
public readonly frequencies: string[];
```

- *Type:* string[]

List that describes which types of snapshots to copy.

---

##### `regionName`<sup>Optional</sup> <a name="regionName" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.regionName"></a>

```typescript
public readonly regionName: string;
```

- *Type:* string

Target region to copy snapshots belonging to replicationSpecId to.

---

##### `replicationSpecId`<sup>Optional</sup> <a name="replicationSpecId" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.replicationSpecId"></a>

```typescript
public readonly replicationSpecId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.

---

##### `shouldCopyOplogs`<sup>Optional</sup> <a name="shouldCopyOplogs" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView.property.shouldCopyOplogs"></a>

```typescript
public readonly shouldCopyOplogs: boolean;
```

- *Type:* boolean

Flag that indicates whether to copy the oplogs to the target region.

---

### ApiDeleteCopiedBackupsView <a name="ApiDeleteCopiedBackupsView" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.Initializer"></a>

```typescript
import { ApiDeleteCopiedBackupsView } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const apiDeleteCopiedBackupsView: ApiDeleteCopiedBackupsView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.property.cloudProvider">cloudProvider</a></code> | <code>string</code> | A label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.property.regionName">regionName</a></code> | <code>string</code> | Target region for the deleted copy setting whose backup copies you want to delete. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.property.replicationSpecId">replicationSpecId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. |

---

##### `cloudProvider`<sup>Optional</sup> <a name="cloudProvider" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.property.cloudProvider"></a>

```typescript
public readonly cloudProvider: string;
```

- *Type:* string

A label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete.

---

##### `regionName`<sup>Optional</sup> <a name="regionName" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.property.regionName"></a>

```typescript
public readonly regionName: string;
```

- *Type:* string

Target region for the deleted copy setting whose backup copies you want to delete.

---

##### `replicationSpecId`<sup>Optional</sup> <a name="replicationSpecId" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView.property.replicationSpecId"></a>

```typescript
public readonly replicationSpecId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.

---

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### ApiPolicyItemView <a name="ApiPolicyItemView" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.Initializer"></a>

```typescript
import { ApiPolicyItemView } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const apiPolicyItemView: ApiPolicyItemView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.frequencyInterval">frequencyInterval</a></code> | <code>number</code> | Desired frequency of the new backup policy item specified by frequencyType. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.frequencyType">frequencyType</a></code> | <code>string</code> | Frequency associated with the backup policy item. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.id">id</a></code> | <code>string</code> | Unique identifier of the backup policy item. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.retentionUnit">retentionUnit</a></code> | <code>string</code> | Metric of duration of the backup policy item: days, weeks, or months. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.retentionValue">retentionValue</a></code> | <code>number</code> | Duration for which the backup is kept. |

---

##### `frequencyInterval`<sup>Optional</sup> <a name="frequencyInterval" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.frequencyInterval"></a>

```typescript
public readonly frequencyInterval: number;
```

- *Type:* number

Desired frequency of the new backup policy item specified by frequencyType.

---

##### `frequencyType`<sup>Optional</sup> <a name="frequencyType" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.frequencyType"></a>

```typescript
public readonly frequencyType: string;
```

- *Type:* string

Frequency associated with the backup policy item.

One of the following values: hourly, daily, weekly or monthly.

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique identifier of the backup policy item.

---

##### `retentionUnit`<sup>Optional</sup> <a name="retentionUnit" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.retentionUnit"></a>

```typescript
public readonly retentionUnit: string;
```

- *Type:* string

Metric of duration of the backup policy item: days, weeks, or months.

---

##### `retentionValue`<sup>Optional</sup> <a name="retentionValue" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView.property.retentionValue"></a>

```typescript
public readonly retentionValue: number;
```

- *Type:* number

Duration for which the backup is kept.

Associated with retentionUnit.

---

### ApiPolicyView <a name="ApiPolicyView" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView.Initializer"></a>

```typescript
import { ApiPolicyView } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const apiPolicyView: ApiPolicyView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView.property.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView.property.policyItems">policyItems</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView">ApiPolicyItemView</a>[]</code> | *No description.* |

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

---

##### `policyItems`<sup>Optional</sup> <a name="policyItems" id="@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView.property.policyItems"></a>

```typescript
public readonly policyItems: ApiPolicyItemView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyItemView">ApiPolicyItemView</a>[]

---

### CfnCloudBackupScheduleProps <a name="CfnCloudBackupScheduleProps" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps"></a>

An example resource schema demonstrating some basic constructs and validation rules.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.Initializer"></a>

```typescript
import { CfnCloudBackupScheduleProps } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const cfnCloudBackupScheduleProps: CfnCloudBackupScheduleProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.autoExportEnabled">autoExportEnabled</a></code> | <code>boolean</code> | Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.copySettings">copySettings</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView">ApiAtlasDiskBackupCopySettingView</a>[]</code> | List that contains a document for each copy setting item in the desired backup policy. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.deleteCopiedBackups">deleteCopiedBackups</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView">ApiDeleteCopiedBackupsView</a>[]</code> | List that contains a document for each deleted copy setting whose backup copies you want to delete. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.export">export</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Export">Export</a></code> | Policy for automatically exporting cloud backup snapshots. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.id">id</a></code> | <code>string</code> | Unique identifier of the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.links">links</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Link">Link</a>[]</code> | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.policies">policies</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView">ApiPolicyView</a>[]</code> | Rules set for this backup schedule. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.referenceHourOfDay">referenceHourOfDay</a></code> | <code>number</code> | UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.referenceMinuteOfHour">referenceMinuteOfHour</a></code> | <code>number</code> | UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.restoreWindowDays">restoreWindowDays</a></code> | <code>number</code> | Number of days back in time you can restore to with Continuous Cloud Backup accuracy. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.updateSnapshots">updateSnapshots</a></code> | <code>boolean</code> | Flag indicating if updates to retention in the backup policy were applied to snapshots that Atlas took earlier. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.useOrgAndGroupNamesInExportPrefix">useOrgAndGroupNamesInExportPrefix</a></code> | <code>boolean</code> | Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `autoExportEnabled`<sup>Optional</sup> <a name="autoExportEnabled" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.autoExportEnabled"></a>

```typescript
public readonly autoExportEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled.

---

##### `copySettings`<sup>Optional</sup> <a name="copySettings" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.copySettings"></a>

```typescript
public readonly copySettings: ApiAtlasDiskBackupCopySettingView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiAtlasDiskBackupCopySettingView">ApiAtlasDiskBackupCopySettingView</a>[]

List that contains a document for each copy setting item in the desired backup policy.

---

##### `deleteCopiedBackups`<sup>Optional</sup> <a name="deleteCopiedBackups" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.deleteCopiedBackups"></a>

```typescript
public readonly deleteCopiedBackups: ApiDeleteCopiedBackupsView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiDeleteCopiedBackupsView">ApiDeleteCopiedBackupsView</a>[]

List that contains a document for each deleted copy setting whose backup copies you want to delete.

---

##### `export`<sup>Optional</sup> <a name="export" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.export"></a>

```typescript
public readonly export: Export;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Export">Export</a>

Policy for automatically exporting cloud backup snapshots.

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique identifier of the snapshot.

---

##### `links`<sup>Optional</sup> <a name="links" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.links"></a>

```typescript
public readonly links: Link[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Link">Link</a>[]

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.

RFC 5988 outlines these relationships.

---

##### `policies`<sup>Optional</sup> <a name="policies" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.policies"></a>

```typescript
public readonly policies: ApiPolicyView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiPolicyView">ApiPolicyView</a>[]

Rules set for this backup schedule.

---

##### `referenceHourOfDay`<sup>Optional</sup> <a name="referenceHourOfDay" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.referenceHourOfDay"></a>

```typescript
public readonly referenceHourOfDay: number;
```

- *Type:* number

UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.

---

##### `referenceMinuteOfHour`<sup>Optional</sup> <a name="referenceMinuteOfHour" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.referenceMinuteOfHour"></a>

```typescript
public readonly referenceMinuteOfHour: number;
```

- *Type:* number

UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.

---

##### `restoreWindowDays`<sup>Optional</sup> <a name="restoreWindowDays" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.restoreWindowDays"></a>

```typescript
public readonly restoreWindowDays: number;
```

- *Type:* number

Number of days back in time you can restore to with Continuous Cloud Backup accuracy.

Must be a positive, non-zero integer.

---

##### `updateSnapshots`<sup>Optional</sup> <a name="updateSnapshots" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.updateSnapshots"></a>

```typescript
public readonly updateSnapshots: boolean;
```

- *Type:* boolean

Flag indicating if updates to retention in the backup policy were applied to snapshots that Atlas took earlier.

---

##### `useOrgAndGroupNamesInExportPrefix`<sup>Optional</sup> <a name="useOrgAndGroupNamesInExportPrefix" id="@mongodbatlas-awscdk/cloud-backup-schedule.CfnCloudBackupScheduleProps.property.useOrgAndGroupNamesInExportPrefix"></a>

```typescript
public readonly useOrgAndGroupNamesInExportPrefix: boolean;
```

- *Type:* boolean

Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots.

---

### Export <a name="Export" id="@mongodbatlas-awscdk/cloud-backup-schedule.Export"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.Export.Initializer"></a>

```typescript
import { Export } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const export: Export = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Export.property.exportBucketId">exportBucketId</a></code> | <code>string</code> | Unique identifier of the AWS bucket to export the cloud backup snapshot to. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Export.property.frequencyType">frequencyType</a></code> | <code>string</code> | Frequency associated with the export policy. |

---

##### `exportBucketId`<sup>Optional</sup> <a name="exportBucketId" id="@mongodbatlas-awscdk/cloud-backup-schedule.Export.property.exportBucketId"></a>

```typescript
public readonly exportBucketId: string;
```

- *Type:* string

Unique identifier of the AWS bucket to export the cloud backup snapshot to.

---

##### `frequencyType`<sup>Optional</sup> <a name="frequencyType" id="@mongodbatlas-awscdk/cloud-backup-schedule.Export.property.frequencyType"></a>

```typescript
public readonly frequencyType: string;
```

- *Type:* string

Frequency associated with the export policy.

Value can be daily, weekly, or monthly.

---

### Link <a name="Link" id="@mongodbatlas-awscdk/cloud-backup-schedule.Link"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-backup-schedule.Link.Initializer"></a>

```typescript
import { Link } from '@mongodbatlas-awscdk/cloud-backup-schedule'

const link: Link = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Link.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Link.property.href">href</a></code> | <code>string</code> | Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. |
| <code><a href="#@mongodbatlas-awscdk/cloud-backup-schedule.Link.property.rel">rel</a></code> | <code>string</code> | Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/cloud-backup-schedule.Link.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-backup-schedule.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `href`<sup>Optional</sup> <a name="href" id="@mongodbatlas-awscdk/cloud-backup-schedule.Link.property.href"></a>

```typescript
public readonly href: string;
```

- *Type:* string

Uniform Resource Locator (URL) that points another API resource to which this response has some relationship.

This URL often begins with `https://mms.mongodb.com`.

---

##### `rel`<sup>Optional</sup> <a name="rel" id="@mongodbatlas-awscdk/cloud-backup-schedule.Link.property.rel"></a>

```typescript
public readonly rel: string;
```

- *Type:* string

Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource.

This URL often begins with `https://mms.mongodb.com`.

---



