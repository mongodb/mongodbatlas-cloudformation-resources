# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCluster <a name="CfnCluster" id="@mongodbatlas-awscdk/cluster.CfnCluster"></a>

A CloudFormation `MongoDB::Atlas::Cluster`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/cluster.CfnCluster.Initializer"></a>

```typescript
import { CfnCluster } from '@mongodbatlas-awscdk/cluster'

new CfnCluster(scope: Construct, id: string, props: CfnClusterProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps">CfnClusterProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/cluster.CfnCluster.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/cluster.CfnCluster.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cluster.CfnCluster.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps">CfnClusterProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/cluster.CfnCluster.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/cluster.CfnCluster.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/cluster.CfnCluster.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/cluster.CfnCluster.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cluster.CfnCluster.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/cluster.CfnCluster.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/cluster.CfnCluster.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/cluster.CfnCluster.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cluster.CfnCluster.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cluster.CfnCluster.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/cluster.CfnCluster.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/cluster.CfnCluster.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cluster.CfnCluster.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/cluster.CfnCluster.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/cluster.CfnCluster.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/cluster.CfnCluster.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/cluster.CfnCluster.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/cluster.CfnCluster.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/cluster.CfnCluster.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/cluster.CfnCluster.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/cluster.CfnCluster.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/cluster.CfnCluster.isConstruct"></a>

```typescript
import { CfnCluster } from '@mongodbatlas-awscdk/cluster'

CfnCluster.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cluster.CfnCluster.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/cluster.CfnCluster.isCfnElement"></a>

```typescript
import { CfnCluster } from '@mongodbatlas-awscdk/cluster'

CfnCluster.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/cluster.CfnCluster.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/cluster.CfnCluster.isCfnResource"></a>

```typescript
import { CfnCluster } from '@mongodbatlas-awscdk/cluster'

CfnCluster.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/cluster.CfnCluster.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.attrCreatedDate">attrCreatedDate</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.CreatedDate`. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.Id`. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.attrMongoDBVersion">attrMongoDBVersion</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.MongoDBVersion`. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.attrStateName">attrStateName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.StateName`. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps">CfnClusterProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrCreatedDate`<sup>Required</sup> <a name="attrCreatedDate" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.attrCreatedDate"></a>

```typescript
public readonly attrCreatedDate: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.CreatedDate`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.Id`.

---

##### `attrMongoDBVersion`<sup>Required</sup> <a name="attrMongoDBVersion" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.attrMongoDBVersion"></a>

```typescript
public readonly attrMongoDBVersion: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.MongoDBVersion`.

---

##### `attrStateName`<sup>Required</sup> <a name="attrStateName" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.attrStateName"></a>

```typescript
public readonly attrStateName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.StateName`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.props"></a>

```typescript
public readonly props: CfnClusterProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps">CfnClusterProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnCluster.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/cluster.CfnCluster.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### AdvancedAutoScaling <a name="AdvancedAutoScaling" id="@mongodbatlas-awscdk/cluster.AdvancedAutoScaling"></a>

AWS Automatic Cluster Scaling.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.AdvancedAutoScaling.Initializer"></a>

```typescript
import { AdvancedAutoScaling } from '@mongodbatlas-awscdk/cluster'

const advancedAutoScaling: AdvancedAutoScaling = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedAutoScaling.property.compute">compute</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.Compute">Compute</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedAutoScaling.property.diskGb">diskGb</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.DiskGb">DiskGb</a></code> | *No description.* |

---

##### `compute`<sup>Optional</sup> <a name="compute" id="@mongodbatlas-awscdk/cluster.AdvancedAutoScaling.property.compute"></a>

```typescript
public readonly compute: Compute;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.Compute">Compute</a>

---

##### `diskGb`<sup>Optional</sup> <a name="diskGb" id="@mongodbatlas-awscdk/cluster.AdvancedAutoScaling.property.diskGb"></a>

```typescript
public readonly diskGb: DiskGb;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.DiskGb">DiskGb</a>

---

### AdvancedRegionConfig <a name="AdvancedRegionConfig" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig"></a>

Hardware specifications for nodes set for a given region.

Each regionConfigs object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each regionConfigs object must have either an analyticsSpecs object, electableSpecs object, or readOnlySpecs object. Tenant clusters only require electableSpecs. Dedicated clusters can specify any of these specifications, but must have at least one electableSpecs object within a replicationSpec. Every hardware specification must use the same instanceSize.

Example:

If you set "replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize" : "M30", set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30"if you have electable nodes and"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize" : "M30" if you have read-only nodes.",

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.Initializer"></a>

```typescript
import { AdvancedRegionConfig } from '@mongodbatlas-awscdk/cluster'

const advancedRegionConfig: AdvancedRegionConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.analyticsAutoScaling">analyticsAutoScaling</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedAutoScaling">AdvancedAutoScaling</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.analyticsSpecs">analyticsSpecs</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.Specs">Specs</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.autoScaling">autoScaling</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedAutoScaling">AdvancedAutoScaling</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.electableSpecs">electableSpecs</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.Specs">Specs</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.priority">priority</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.readOnlySpecs">readOnlySpecs</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.Specs">Specs</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.regionName">regionName</a></code> | <code>string</code> | *No description.* |

---

##### `analyticsAutoScaling`<sup>Optional</sup> <a name="analyticsAutoScaling" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.analyticsAutoScaling"></a>

```typescript
public readonly analyticsAutoScaling: AdvancedAutoScaling;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.AdvancedAutoScaling">AdvancedAutoScaling</a>

---

##### `analyticsSpecs`<sup>Optional</sup> <a name="analyticsSpecs" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.analyticsSpecs"></a>

```typescript
public readonly analyticsSpecs: Specs;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.Specs">Specs</a>

---

##### `autoScaling`<sup>Optional</sup> <a name="autoScaling" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.autoScaling"></a>

```typescript
public readonly autoScaling: AdvancedAutoScaling;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.AdvancedAutoScaling">AdvancedAutoScaling</a>

---

##### `electableSpecs`<sup>Optional</sup> <a name="electableSpecs" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.electableSpecs"></a>

```typescript
public readonly electableSpecs: Specs;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.Specs">Specs</a>

---

##### `priority`<sup>Optional</sup> <a name="priority" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.priority"></a>

```typescript
public readonly priority: number;
```

- *Type:* number

---

##### `readOnlySpecs`<sup>Optional</sup> <a name="readOnlySpecs" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.readOnlySpecs"></a>

```typescript
public readonly readOnlySpecs: Specs;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.Specs">Specs</a>

---

##### `regionName`<sup>Optional</sup> <a name="regionName" id="@mongodbatlas-awscdk/cluster.AdvancedRegionConfig.property.regionName"></a>

```typescript
public readonly regionName: string;
```

- *Type:* string

---

### AdvancedReplicationSpec <a name="AdvancedReplicationSpec" id="@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec"></a>

List of settings that configure your cluster regions.

For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.Initializer"></a>

```typescript
import { AdvancedReplicationSpec } from '@mongodbatlas-awscdk/cluster'

const advancedReplicationSpec: AdvancedReplicationSpec = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.advancedRegionConfigs">advancedRegionConfigs</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig">AdvancedRegionConfig</a>[]</code> | Hardware specifications for nodes set for a given region. |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.id">id</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.numShards">numShards</a></code> | <code>number</code> | Positive integer that specifies the number of shards to deploy in each specified zone. |
| <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.zoneName">zoneName</a></code> | <code>string</code> | Human-readable label that identifies the zone in a Global Cluster. |

---

##### `advancedRegionConfigs`<sup>Optional</sup> <a name="advancedRegionConfigs" id="@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.advancedRegionConfigs"></a>

```typescript
public readonly advancedRegionConfigs: AdvancedRegionConfig[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.AdvancedRegionConfig">AdvancedRegionConfig</a>[]

Hardware specifications for nodes set for a given region.

Each regionConfigs object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each regionConfigs object must have either an analyticsSpecs object, electableSpecs object, or readOnlySpecs object. Tenant clusters only require electableSpecs. Dedicated clusters can specify any of these specifications, but must have at least one electableSpecs object within a replicationSpec. Every hardware specification must use the same instanceSize.

Example:

If you set "replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize" : "M30", set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30"if you have electable nodes and"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize" : "M30" if you have read-only nodes.",

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster.

If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Multi-Cloud Cluster, you may specify this parameter. The request deletes any existing zones in the Multi-Cloud Cluster that you exclude from the request.

---

##### `numShards`<sup>Optional</sup> <a name="numShards" id="@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.numShards"></a>

```typescript
public readonly numShards: number;
```

- *Type:* number

Positive integer that specifies the number of shards to deploy in each specified zone.

If you set this value to 1 and "clusterType" : "SHARDED", MongoDB Cloud deploys a single-shard sharded cluster. Don't create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don't provide the same benefits as multi-shard configurations.

---

##### `zoneName`<sup>Optional</sup> <a name="zoneName" id="@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec.property.zoneName"></a>

```typescript
public readonly zoneName: string;
```

- *Type:* string

Human-readable label that identifies the zone in a Global Cluster.

Provide this value only if "clusterType" : "GEOSHARDED".

---

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/cluster.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/cluster'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/cluster.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/cluster.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnClusterProps <a name="CfnClusterProps" id="@mongodbatlas-awscdk/cluster.CfnClusterProps"></a>

The cluster resource provides access to your cluster configurations.

The resource lets you create, edit and delete clusters. The resource requires your Project ID.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.Initializer"></a>

```typescript
import { CfnClusterProps } from '@mongodbatlas-awscdk/cluster'

const cfnClusterProps: CfnClusterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the advanced cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the project the cluster belongs to. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.advancedSettings">advancedSettings</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs">ProcessArgs</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.backupEnabled">backupEnabled</a></code> | <code>boolean</code> | Flag that indicates whether the cluster can perform backups. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.biConnector">biConnector</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector">CfnClusterPropsBiConnector</a></code> | Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.clusterType">clusterType</a></code> | <code>string</code> | Configuration of nodes that comprise the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.connectionStrings">connectionStrings</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings">ConnectionStrings</a></code> | Set of connection strings that your applications use to connect to this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.diskSizeGb">diskSizeGb</a></code> | <code>number</code> | Storage capacity that the host's root volume possesses expressed in gigabytes. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.encryptionAtRestProvider">encryptionAtRestProvider</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider">CfnClusterPropsEncryptionAtRestProvider</a></code> | Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.labels">labels</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels">CfnClusterPropsLabels</a>[]</code> | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.mongoDbMajorVersion">mongoDbMajorVersion</a></code> | <code>string</code> | Major MongoDB version of the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.paused">paused</a></code> | <code>boolean</code> | Flag that indicates whether the cluster is paused or not. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.pitEnabled">pitEnabled</a></code> | <code>boolean</code> | Flag that indicates whether the cluster uses continuous cloud backups. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.replicationSpecs">replicationSpecs</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec">AdvancedReplicationSpec</a>[]</code> | List of settings that configure your cluster regions. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.rootCertType">rootCertType</a></code> | <code>string</code> | Root Certificate Authority that MongoDB Cloud cluster uses. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.terminationProtectionEnabled">terminationProtectionEnabled</a></code> | <code>boolean</code> | Flag that indicates whether termination protection is enabled on the cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterProps.property.versionReleaseSystem">versionReleaseSystem</a></code> | <code>string</code> | Method by which the cluster maintains the MongoDB versions. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `name`<sup>Required</sup> <a name="name" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the advanced cluster.

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique identifier of the project the cluster belongs to.

---

##### `advancedSettings`<sup>Optional</sup> <a name="advancedSettings" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.advancedSettings"></a>

```typescript
public readonly advancedSettings: ProcessArgs;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.ProcessArgs">ProcessArgs</a>

---

##### `backupEnabled`<sup>Optional</sup> <a name="backupEnabled" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.backupEnabled"></a>

```typescript
public readonly backupEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether the cluster can perform backups.

If set to true, the cluster can perform backups. You must set this value to true for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to false, the cluster doesn't use backups.

---

##### `biConnector`<sup>Optional</sup> <a name="biConnector" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.biConnector"></a>

```typescript
public readonly biConnector: CfnClusterPropsBiConnector;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector">CfnClusterPropsBiConnector</a>

Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.

---

##### `clusterType`<sup>Optional</sup> <a name="clusterType" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.clusterType"></a>

```typescript
public readonly clusterType: string;
```

- *Type:* string

Configuration of nodes that comprise the cluster.

---

##### `connectionStrings`<sup>Optional</sup> <a name="connectionStrings" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.connectionStrings"></a>

```typescript
public readonly connectionStrings: ConnectionStrings;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings">ConnectionStrings</a>

Set of connection strings that your applications use to connect to this cluster.

Use the parameters in this object to connect your applications to this cluster. See the MongoDB [Connection String URI Format](https://docs.mongodb.com/manual/reference/connection-string/) reference for further details.

---

##### `diskSizeGb`<sup>Optional</sup> <a name="diskSizeGb" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.diskSizeGb"></a>

```typescript
public readonly diskSizeGb: number;
```

- *Type:* number

Storage capacity that the host's root volume possesses expressed in gigabytes.

Increase this number to add capacity. MongoDB Cloud requires this parameter if you set replicationSpecs. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier.

---

##### `encryptionAtRestProvider`<sup>Optional</sup> <a name="encryptionAtRestProvider" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.encryptionAtRestProvider"></a>

```typescript
public readonly encryptionAtRestProvider: CfnClusterPropsEncryptionAtRestProvider;
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider">CfnClusterPropsEncryptionAtRestProvider</a>

Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster.

To enable customer key management for encryption at rest, the cluster replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize setting must be M10 or higher and "backupEnabled" : false or omitted entirely.

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.labels"></a>

```typescript
public readonly labels: CfnClusterPropsLabels[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels">CfnClusterPropsLabels</a>[]

Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster.

The MongoDB Cloud console doesn't display your labels.

---

##### `mongoDbMajorVersion`<sup>Optional</sup> <a name="mongoDbMajorVersion" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.mongoDbMajorVersion"></a>

```typescript
public readonly mongoDbMajorVersion: string;
```

- *Type:* string

Major MongoDB version of the cluster.

MongoDB Cloud deploys the cluster with the latest stable release of the specified version.

---

##### `paused`<sup>Optional</sup> <a name="paused" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.paused"></a>

```typescript
public readonly paused: boolean;
```

- *Type:* boolean

Flag that indicates whether the cluster is paused or not.

---

##### `pitEnabled`<sup>Optional</sup> <a name="pitEnabled" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.pitEnabled"></a>

```typescript
public readonly pitEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether the cluster uses continuous cloud backups.

---

##### `replicationSpecs`<sup>Optional</sup> <a name="replicationSpecs" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.replicationSpecs"></a>

```typescript
public readonly replicationSpecs: AdvancedReplicationSpec[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec">AdvancedReplicationSpec</a>[]

List of settings that configure your cluster regions.

For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.

---

##### `rootCertType`<sup>Optional</sup> <a name="rootCertType" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.rootCertType"></a>

```typescript
public readonly rootCertType: string;
```

- *Type:* string

Root Certificate Authority that MongoDB Cloud cluster uses.

MongoDB Cloud supports Internet Security Research Group.

---

##### `terminationProtectionEnabled`<sup>Optional</sup> <a name="terminationProtectionEnabled" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.terminationProtectionEnabled"></a>

```typescript
public readonly terminationProtectionEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether termination protection is enabled on the cluster.

If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.

---

##### `versionReleaseSystem`<sup>Optional</sup> <a name="versionReleaseSystem" id="@mongodbatlas-awscdk/cluster.CfnClusterProps.property.versionReleaseSystem"></a>

```typescript
public readonly versionReleaseSystem: string;
```

- *Type:* string

Method by which the cluster maintains the MongoDB versions.

If value is CONTINUOUS, you must not specify mongoDBMajorVersion

---

### CfnClusterPropsBiConnector <a name="CfnClusterPropsBiConnector" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector"></a>

Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector.Initializer"></a>

```typescript
import { CfnClusterPropsBiConnector } from '@mongodbatlas-awscdk/cluster'

const cfnClusterPropsBiConnector: CfnClusterPropsBiConnector = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector.property.enabled">enabled</a></code> | <code>boolean</code> | Flag that indicates whether MongoDB Connector for Business Intelligence is enabled on the specified cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector.property.readPreference">readPreference</a></code> | <code>string</code> | Data source node designated for the MongoDB Connector for Business Intelligence on MongoDB Cloud. |

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Flag that indicates whether MongoDB Connector for Business Intelligence is enabled on the specified cluster.

---

##### `readPreference`<sup>Optional</sup> <a name="readPreference" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector.property.readPreference"></a>

```typescript
public readonly readPreference: string;
```

- *Type:* string
- *Default:* ANALYTICS node, or SECONDARY if there are no ANALYTICS nodes.

Data source node designated for the MongoDB Connector for Business Intelligence on MongoDB Cloud.

The MongoDB Connector for Business Intelligence on MongoDB Cloud reads data from the primary, secondary, or analytics node based on your read preferences. Defaults to ANALYTICS node, or SECONDARY if there are no ANALYTICS nodes.

---

### CfnClusterPropsLabels <a name="CfnClusterPropsLabels" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels.Initializer"></a>

```typescript
import { CfnClusterPropsLabels } from '@mongodbatlas-awscdk/cluster'

const cfnClusterPropsLabels: CfnClusterPropsLabels = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels.property.key">key</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels.property.value">value</a></code> | <code>string</code> | *No description.* |

---

##### `key`<sup>Optional</sup> <a name="key" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels.property.key"></a>

```typescript
public readonly key: string;
```

- *Type:* string

---

##### `value`<sup>Optional</sup> <a name="value" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels.property.value"></a>

```typescript
public readonly value: string;
```

- *Type:* string

---

### Compute <a name="Compute" id="@mongodbatlas-awscdk/cluster.Compute"></a>

Automatic Compute Scaling.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.Compute.Initializer"></a>

```typescript
import { Compute } from '@mongodbatlas-awscdk/cluster'

const compute: Compute = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.Compute.property.enabled">enabled</a></code> | <code>boolean</code> | Flag that indicates whether someone enabled instance size auto-scaling. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Compute.property.maxInstanceSize">maxInstanceSize</a></code> | <code>string</code> | Maximum instance size to which your cluster can automatically scale. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Compute.property.minInstanceSize">minInstanceSize</a></code> | <code>string</code> | Minimum instance size to which your cluster can automatically scale. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Compute.property.scaleDownEnabled">scaleDownEnabled</a></code> | <code>boolean</code> | Flag that indicates whether the instance size may scale down. |

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/cluster.Compute.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Flag that indicates whether someone enabled instance size auto-scaling.

Set to true to enable instance size auto-scaling. If enabled, you must specify a value for replicationSpecs[n].regionConfigs[m].autoScaling.compute.maxInstanceSize.
Set to false to disable instance size automatic scaling.

---

##### `maxInstanceSize`<sup>Optional</sup> <a name="maxInstanceSize" id="@mongodbatlas-awscdk/cluster.Compute.property.maxInstanceSize"></a>

```typescript
public readonly maxInstanceSize: string;
```

- *Type:* string

Maximum instance size to which your cluster can automatically scale.

MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true.

---

##### `minInstanceSize`<sup>Optional</sup> <a name="minInstanceSize" id="@mongodbatlas-awscdk/cluster.Compute.property.minInstanceSize"></a>

```typescript
public readonly minInstanceSize: string;
```

- *Type:* string

Minimum instance size to which your cluster can automatically scale.

MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true.

---

##### `scaleDownEnabled`<sup>Optional</sup> <a name="scaleDownEnabled" id="@mongodbatlas-awscdk/cluster.Compute.property.scaleDownEnabled"></a>

```typescript
public readonly scaleDownEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether the instance size may scale down.

MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true. If you enable this option, specify a value for replicationSpecs[n].regionConfigs[m].autoScaling.compute.minInstanceSize.

---

### ConnectionStrings <a name="ConnectionStrings" id="@mongodbatlas-awscdk/cluster.ConnectionStrings"></a>

Collection of Uniform Resource Locators that point to the MongoDB database.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.Initializer"></a>

```typescript
import { ConnectionStrings } from '@mongodbatlas-awscdk/cluster'

const connectionStrings: ConnectionStrings = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.awsPrivateLink">awsPrivateLink</a></code> | <code>string</code> | Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.awsPrivateLinkSrv">awsPrivateLinkSrv</a></code> | <code>string</code> | Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.private">private</a></code> | <code>string</code> | Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.privateEndpoint">privateEndpoint</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.PrivateEndpoint">PrivateEndpoint</a>[]</code> | List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.privateSrv">privateSrv</a></code> | <code>string</code> | Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.standard">standard</a></code> | <code>string</code> | Public connection string that you can use to connect to this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ConnectionStrings.property.standardSrv">standardSrv</a></code> | <code>string</code> | Public connection string that you can use to connect to this cluster. |

---

##### `awsPrivateLink`<sup>Optional</sup> <a name="awsPrivateLink" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.awsPrivateLink"></a>

```typescript
public readonly awsPrivateLink: string;
```

- *Type:* string

Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink.

Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related mongodb:// connection string that you use to connect to MongoDB Cloud through the interface endpoint that the key names.

---

##### `awsPrivateLinkSrv`<sup>Optional</sup> <a name="awsPrivateLinkSrv" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.awsPrivateLinkSrv"></a>

```typescript
public readonly awsPrivateLinkSrv: string;
```

- *Type:* string

Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink.

Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related mongodb:// connection string that you use to connect to Atlas through the interface endpoint that the key names.

---

##### `private`<sup>Optional</sup> <a name="private" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.private"></a>

```typescript
public readonly private: string;
```

- *Type:* string

Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster.

This connection string uses the mongodb+srv:// protocol. The resource returns this parameter once someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this resource returns this parameter only if you enable custom DNS.

---

##### `privateEndpoint`<sup>Optional</sup> <a name="privateEndpoint" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.privateEndpoint"></a>

```typescript
public readonly privateEndpoint: PrivateEndpoint[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.PrivateEndpoint">PrivateEndpoint</a>[]

List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint.

This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.

---

##### `privateSrv`<sup>Optional</sup> <a name="privateSrv" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.privateSrv"></a>

```typescript
public readonly privateSrv: string;
```

- *Type:* string

Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster.

This connection string uses the mongodb+srv:// protocol. The resource returns this parameter when someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this parameter returns only if you enable custom DNS.

---

##### `standard`<sup>Optional</sup> <a name="standard" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.standard"></a>

```typescript
public readonly standard: string;
```

- *Type:* string

Public connection string that you can use to connect to this cluster.

This connection string uses the mongodb:// protocol.

---

##### `standardSrv`<sup>Optional</sup> <a name="standardSrv" id="@mongodbatlas-awscdk/cluster.ConnectionStrings.property.standardSrv"></a>

```typescript
public readonly standardSrv: string;
```

- *Type:* string

Public connection string that you can use to connect to this cluster.

This connection string uses the mongodb+srv:// protocol.

---

### DiskGb <a name="DiskGb" id="@mongodbatlas-awscdk/cluster.DiskGb"></a>

Automatic cluster storage settings that apply to this cluster.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.DiskGb.Initializer"></a>

```typescript
import { DiskGb } from '@mongodbatlas-awscdk/cluster'

const diskGb: DiskGb = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.DiskGb.property.enabled">enabled</a></code> | <code>boolean</code> | Flag that indicates whether this cluster enables disk auto-scaling. |

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/cluster.DiskGb.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Flag that indicates whether this cluster enables disk auto-scaling.

The maximum memory allowed for the selected cluster tier and the oplog size can limit storage auto-scaling.

---

### Endpoint <a name="Endpoint" id="@mongodbatlas-awscdk/cluster.Endpoint"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.Endpoint.Initializer"></a>

```typescript
import { Endpoint } from '@mongodbatlas-awscdk/cluster'

const endpoint: Endpoint = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.Endpoint.property.endpointId">endpointId</a></code> | <code>string</code> | Unique string that the cloud provider uses to identify the private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Endpoint.property.providerName">providerName</a></code> | <code>string</code> | Cloud provider in which MongoDB Cloud deploys the private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Endpoint.property.region">region</a></code> | <code>string</code> | Region in which MongoDB Cloud deploys the private endpoint. |

---

##### `endpointId`<sup>Optional</sup> <a name="endpointId" id="@mongodbatlas-awscdk/cluster.Endpoint.property.endpointId"></a>

```typescript
public readonly endpointId: string;
```

- *Type:* string

Unique string that the cloud provider uses to identify the private endpoint.

---

##### `providerName`<sup>Optional</sup> <a name="providerName" id="@mongodbatlas-awscdk/cluster.Endpoint.property.providerName"></a>

```typescript
public readonly providerName: string;
```

- *Type:* string

Cloud provider in which MongoDB Cloud deploys the private endpoint.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/cluster.Endpoint.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

Region in which MongoDB Cloud deploys the private endpoint.

---

### PrivateEndpoint <a name="PrivateEndpoint" id="@mongodbatlas-awscdk/cluster.PrivateEndpoint"></a>

List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint.

This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.PrivateEndpoint.Initializer"></a>

```typescript
import { PrivateEndpoint } from '@mongodbatlas-awscdk/cluster'

const privateEndpoint: PrivateEndpoint = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.connectionString">connectionString</a></code> | <code>string</code> | Private endpoint-aware connection string that uses the mongodb:// protocol to connect to MongoDB Cloud through a private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.endpoints">endpoints</a></code> | <code><a href="#@mongodbatlas-awscdk/cluster.Endpoint">Endpoint</a>[]</code> | List that contains the private endpoints through which you connect to MongoDB Cloud when you use connectionStrings.privateEndpoint[n].connectionString or connectionStrings.privateEndpoint[n].srvConnectionString. |
| <code><a href="#@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.srvConnectionString">srvConnectionString</a></code> | <code>string</code> | Private endpoint-aware connection string that uses the mongodb+srv:// protocol to connect to MongoDB Cloud through a private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.type">type</a></code> | <code>string</code> | Enum: "MONGOD" "MONGOS" MongoDB process type to which your application connects. |

---

##### `connectionString`<sup>Optional</sup> <a name="connectionString" id="@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.connectionString"></a>

```typescript
public readonly connectionString: string;
```

- *Type:* string

Private endpoint-aware connection string that uses the mongodb:// protocol to connect to MongoDB Cloud through a private endpoint.

---

##### `endpoints`<sup>Optional</sup> <a name="endpoints" id="@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.endpoints"></a>

```typescript
public readonly endpoints: Endpoint[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/cluster.Endpoint">Endpoint</a>[]

List that contains the private endpoints through which you connect to MongoDB Cloud when you use connectionStrings.privateEndpoint[n].connectionString or connectionStrings.privateEndpoint[n].srvConnectionString.

---

##### `srvConnectionString`<sup>Optional</sup> <a name="srvConnectionString" id="@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.srvConnectionString"></a>

```typescript
public readonly srvConnectionString: string;
```

- *Type:* string

Private endpoint-aware connection string that uses the mongodb+srv:// protocol to connect to MongoDB Cloud through a private endpoint.

The mongodb+srv protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your application supports it. If it doesn't, use connectionStrings.privateEndpoint[n].connectionString.

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/cluster.PrivateEndpoint.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string

Enum: "MONGOD" "MONGOS" MongoDB process type to which your application connects.

Use MONGOD for replica sets and MONGOS for sharded clusters.

---

### ProcessArgs <a name="ProcessArgs" id="@mongodbatlas-awscdk/cluster.ProcessArgs"></a>

Advanced configuration details to add for one cluster in the specified project.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.ProcessArgs.Initializer"></a>

```typescript
import { ProcessArgs } from '@mongodbatlas-awscdk/cluster'

const processArgs: ProcessArgs = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.defaultReadConcern">defaultReadConcern</a></code> | <code>string</code> | Default level of acknowledgment requested from MongoDB for read operations set for this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.defaultWriteConcern">defaultWriteConcern</a></code> | <code>string</code> | Default level of acknowledgment requested from MongoDB for write operations set for this cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.failIndexKeyTooLong">failIndexKeyTooLong</a></code> | <code>boolean</code> | Flag that indicates whether you can insert or update documents where all indexed entries don't exceed 1024 bytes. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.javascriptEnabled">javascriptEnabled</a></code> | <code>boolean</code> | Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.minimumEnabledTlsProtocol">minimumEnabledTlsProtocol</a></code> | <code>string</code> | Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.noTableScan">noTableScan</a></code> | <code>boolean</code> | Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.oplogSizeMb">oplogSizeMb</a></code> | <code>number</code> | Storage limit of cluster's oplog expressed in megabytes. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.sampleRefreshIntervalBiConnector">sampleRefreshIntervalBiConnector</a></code> | <code>number</code> | Number of documents per database to sample when gathering schema information. |
| <code><a href="#@mongodbatlas-awscdk/cluster.ProcessArgs.property.sampleSizeBiConnector">sampleSizeBiConnector</a></code> | <code>number</code> | Interval in seconds at which the mongosqld process re-samples data to create its relational schema. |

---

##### `defaultReadConcern`<sup>Optional</sup> <a name="defaultReadConcern" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.defaultReadConcern"></a>

```typescript
public readonly defaultReadConcern: string;
```

- *Type:* string

Default level of acknowledgment requested from MongoDB for read operations set for this cluster.

---

##### `defaultWriteConcern`<sup>Optional</sup> <a name="defaultWriteConcern" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.defaultWriteConcern"></a>

```typescript
public readonly defaultWriteConcern: string;
```

- *Type:* string

Default level of acknowledgment requested from MongoDB for write operations set for this cluster.

---

##### `failIndexKeyTooLong`<sup>Optional</sup> <a name="failIndexKeyTooLong" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.failIndexKeyTooLong"></a>

```typescript
public readonly failIndexKeyTooLong: boolean;
```

- *Type:* boolean

Flag that indicates whether you can insert or update documents where all indexed entries don't exceed 1024 bytes.

If you set this to false, mongod writes documents that exceed this limit but doesn't index them.

---

##### `javascriptEnabled`<sup>Optional</sup> <a name="javascriptEnabled" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.javascriptEnabled"></a>

```typescript
public readonly javascriptEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript.

---

##### `minimumEnabledTlsProtocol`<sup>Optional</sup> <a name="minimumEnabledTlsProtocol" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.minimumEnabledTlsProtocol"></a>

```typescript
public readonly minimumEnabledTlsProtocol: string;
```

- *Type:* string

Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections.

Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version.

---

##### `noTableScan`<sup>Optional</sup> <a name="noTableScan" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.noTableScan"></a>

```typescript
public readonly noTableScan: boolean;
```

- *Type:* boolean

Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results.

---

##### `oplogSizeMb`<sup>Optional</sup> <a name="oplogSizeMb" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.oplogSizeMb"></a>

```typescript
public readonly oplogSizeMb: number;
```

- *Type:* number

Storage limit of cluster's oplog expressed in megabytes.

A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates.

---

##### `sampleRefreshIntervalBiConnector`<sup>Optional</sup> <a name="sampleRefreshIntervalBiConnector" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.sampleRefreshIntervalBiConnector"></a>

```typescript
public readonly sampleRefreshIntervalBiConnector: number;
```

- *Type:* number

Number of documents per database to sample when gathering schema information.

---

##### `sampleSizeBiConnector`<sup>Optional</sup> <a name="sampleSizeBiConnector" id="@mongodbatlas-awscdk/cluster.ProcessArgs.property.sampleSizeBiConnector"></a>

```typescript
public readonly sampleSizeBiConnector: number;
```

- *Type:* number

Interval in seconds at which the mongosqld process re-samples data to create its relational schema.

---

### Specs <a name="Specs" id="@mongodbatlas-awscdk/cluster.Specs"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/cluster.Specs.Initializer"></a>

```typescript
import { Specs } from '@mongodbatlas-awscdk/cluster'

const specs: Specs = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.Specs.property.diskIops">diskIops</a></code> | <code>string</code> | Target throughput desired for storage attached to your AWS-provisioned cluster. Only change this parameter if you:. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Specs.property.ebsVolumeType">ebsVolumeType</a></code> | <code>string</code> | Type of storage you want to attach to your AWS-provisioned cluster. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Specs.property.instanceSize">instanceSize</a></code> | <code>string</code> | Hardware specification for the instance sizes in this region. |
| <code><a href="#@mongodbatlas-awscdk/cluster.Specs.property.nodeCount">nodeCount</a></code> | <code>number</code> | Number of read-only nodes for MongoDB Cloud deploys to the region. |

---

##### `diskIops`<sup>Optional</sup> <a name="diskIops" id="@mongodbatlas-awscdk/cluster.Specs.property.diskIops"></a>

```typescript
public readonly diskIops: string;
```

- *Type:* string

Target throughput desired for storage attached to your AWS-provisioned cluster. Only change this parameter if you:.

set "replicationSpecs[n].regionConfigs[m].providerName" : "AWS".
set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30" or greater not including Mxx_NVME tiers.
The maximum input/output operations per second (IOPS) depend on the selected .instanceSize and .diskSizeGB. This parameter defaults to the cluster tier's standard IOPS value. Changing this value impacts cluster cost. MongoDB Cloud enforces minimum ratios of storage capacity to system memory for given cluster tiers. This keeps cluster performance consistent with large datasets.

Instance sizes M10 to M40 have a ratio of disk capacity to system memory of 60:1.
Instance sizes greater than M40 have a ratio of 120:1.

---

##### `ebsVolumeType`<sup>Optional</sup> <a name="ebsVolumeType" id="@mongodbatlas-awscdk/cluster.Specs.property.ebsVolumeType"></a>

```typescript
public readonly ebsVolumeType: string;
```

- *Type:* string

Type of storage you want to attach to your AWS-provisioned cluster.

STANDARD volume types can't exceed the default input/output operations per second (IOPS) rate for the selected volume size.

PROVISIONED volume types must fall within the allowable IOPS range for the selected volume size."

---

##### `instanceSize`<sup>Optional</sup> <a name="instanceSize" id="@mongodbatlas-awscdk/cluster.Specs.property.instanceSize"></a>

```typescript
public readonly instanceSize: string;
```

- *Type:* string

Hardware specification for the instance sizes in this region.

Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. If you deploy a Global Cluster, you must choose a instance size of M30 or greater.

---

##### `nodeCount`<sup>Optional</sup> <a name="nodeCount" id="@mongodbatlas-awscdk/cluster.Specs.property.nodeCount"></a>

```typescript
public readonly nodeCount: number;
```

- *Type:* number

Number of read-only nodes for MongoDB Cloud deploys to the region.

Read-only nodes can never become the primary, but can enable local reads.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnClusterPropsEncryptionAtRestProvider <a name="CfnClusterPropsEncryptionAtRestProvider" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider"></a>

Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster.

To enable customer key management for encryption at rest, the cluster replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize setting must be M10 or higher and "backupEnabled" : false or omitted entirely.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.AWS">AWS</a></code> | AWS. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.GCP">GCP</a></code> | GCP. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.AZURE">AZURE</a></code> | AZURE. |
| <code><a href="#@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.NONE">NONE</a></code> | NONE. |

---

##### `AWS` <a name="AWS" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.AWS"></a>

AWS.

---


##### `GCP` <a name="GCP" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.GCP"></a>

GCP.

---


##### `AZURE` <a name="AZURE" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.AZURE"></a>

AZURE.

---


##### `NONE` <a name="NONE" id="@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider.NONE"></a>

NONE.

---

