# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCloudProviderSnapshots <a name="CfnCloudProviderSnapshots" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots"></a>

A CloudFormation `MongoDB::Atlas::CloudProviderSnapshots`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer"></a>

```typescript
import { CfnCloudProviderSnapshots } from '@mongodbatlas-awscdk/cloud-provider-snapshots'

new CfnCloudProviderSnapshots(scope: Construct, id: string, props: CfnCloudProviderSnapshotsProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps">CfnCloudProviderSnapshotsProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps">CfnCloudProviderSnapshotsProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isConstruct"></a>

```typescript
import { CfnCloudProviderSnapshots } from '@mongodbatlas-awscdk/cloud-provider-snapshots'

CfnCloudProviderSnapshots.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isCfnElement"></a>

```typescript
import { CfnCloudProviderSnapshots } from '@mongodbatlas-awscdk/cloud-provider-snapshots'

CfnCloudProviderSnapshots.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isCfnResource"></a>

```typescript
import { CfnCloudProviderSnapshots } from '@mongodbatlas-awscdk/cloud-provider-snapshots'

CfnCloudProviderSnapshots.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::CloudProviderSnapshots.Id`. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps">CfnCloudProviderSnapshotsProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::CloudProviderSnapshots.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.props"></a>

```typescript
public readonly props: CfnCloudProviderSnapshotsProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps">CfnCloudProviderSnapshotsProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshots.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/cloud-provider-snapshots'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnCloudProviderSnapshotsProps <a name="CfnCloudProviderSnapshotsProps" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps"></a>

This resource allows to take one on-demand snapshot, get one or all cloud provider snapshot and delete one cloud provider snapshot.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.Initializer"></a>

```typescript
import { CfnCloudProviderSnapshotsProps } from '@mongodbatlas-awscdk/cloud-provider-snapshots'

const cfnCloudProviderSnapshotsProps: CfnCloudProviderSnapshotsProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.clusterName">clusterName</a></code> | <code>string</code> | The name of the Atlas cluster that contains the snapshots you want to retrieve. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.projectId">projectId</a></code> | <code>string</code> | The unique identifier of the project for the Atlas cluster. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.createdAt">createdAt</a></code> | <code>Date</code> | UTC ISO 8601, formatted point in time when Atlas took the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.description">description</a></code> | <code>string</code> | Description of the on-demand snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.masterKeyUuid">masterKeyUuid</a></code> | <code>string</code> | Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.mongoVersion">mongoVersion</a></code> | <code>string</code> | Version of the MongoDB server. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.retentionInDays">retentionInDays</a></code> | <code>number</code> | The number of days that Atlas should retain the on-demand snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.snapshotType">snapshotType</a></code> | <code>string</code> | Specified the type of snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.status">status</a></code> | <code>string</code> | Current status of the snapshot. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.storageSizeBytes">storageSizeBytes</a></code> | <code>number</code> | Specifies the size of the snapshot in bytes. |
| <code><a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.type">type</a></code> | <code>string</code> | Specifies the type of cluster. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cloud-provider-snapshots.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `clusterName`<sup>Required</sup> <a name="clusterName" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

The name of the Atlas cluster that contains the snapshots you want to retrieve.

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

The unique identifier of the project for the Atlas cluster.

---

##### `createdAt`<sup>Optional</sup> <a name="createdAt" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.createdAt"></a>

```typescript
public readonly createdAt: Date;
```

- *Type:* Date

UTC ISO 8601, formatted point in time when Atlas took the snapshot.

---

##### `description`<sup>Optional</sup> <a name="description" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

Description of the on-demand snapshot.

---

##### `masterKeyUuid`<sup>Optional</sup> <a name="masterKeyUuid" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.masterKeyUuid"></a>

```typescript
public readonly masterKeyUuid: string;
```

- *Type:* string

Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot.

---

##### `mongoVersion`<sup>Optional</sup> <a name="mongoVersion" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.mongoVersion"></a>

```typescript
public readonly mongoVersion: string;
```

- *Type:* string

Version of the MongoDB server.

---

##### `retentionInDays`<sup>Optional</sup> <a name="retentionInDays" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.retentionInDays"></a>

```typescript
public readonly retentionInDays: number;
```

- *Type:* number

The number of days that Atlas should retain the on-demand snapshot.

---

##### `snapshotType`<sup>Optional</sup> <a name="snapshotType" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.snapshotType"></a>

```typescript
public readonly snapshotType: string;
```

- *Type:* string

Specified the type of snapshot.

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.status"></a>

```typescript
public readonly status: string;
```

- *Type:* string

Current status of the snapshot.

---

##### `storageSizeBytes`<sup>Optional</sup> <a name="storageSizeBytes" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.storageSizeBytes"></a>

```typescript
public readonly storageSizeBytes: number;
```

- *Type:* number

Specifies the size of the snapshot in bytes.

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/cloud-provider-snapshots.CfnCloudProviderSnapshotsProps.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string

Specifies the type of cluster.

---



