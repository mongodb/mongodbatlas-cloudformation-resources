# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnGlobalClusterConfig <a name="CfnGlobalClusterConfig" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig"></a>

A CloudFormation `MongoDB::Atlas::GlobalClusterConfig`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer"></a>

```typescript
import { CfnGlobalClusterConfig } from '@mongodbatlas-awscdk/global-cluster-config'

new CfnGlobalClusterConfig(scope: Construct, id: string, props: CfnGlobalClusterConfigProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps">CfnGlobalClusterConfigProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps">CfnGlobalClusterConfigProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isConstruct"></a>

```typescript
import { CfnGlobalClusterConfig } from '@mongodbatlas-awscdk/global-cluster-config'

CfnGlobalClusterConfig.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isCfnElement"></a>

```typescript
import { CfnGlobalClusterConfig } from '@mongodbatlas-awscdk/global-cluster-config'

CfnGlobalClusterConfig.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isCfnResource"></a>

```typescript
import { CfnGlobalClusterConfig } from '@mongodbatlas-awscdk/global-cluster-config'

CfnGlobalClusterConfig.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.attrRemoveAllZoneMapping">attrRemoveAllZoneMapping</a></code> | <code>aws-cdk-lib.IResolvable</code> | Attribute `MongoDB::Atlas::GlobalClusterConfig.RemoveAllZoneMapping`. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps">CfnGlobalClusterConfigProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrRemoveAllZoneMapping`<sup>Required</sup> <a name="attrRemoveAllZoneMapping" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.attrRemoveAllZoneMapping"></a>

```typescript
public readonly attrRemoveAllZoneMapping: IResolvable;
```

- *Type:* aws-cdk-lib.IResolvable

Attribute `MongoDB::Atlas::GlobalClusterConfig.RemoveAllZoneMapping`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.props"></a>

```typescript
public readonly props: CfnGlobalClusterConfigProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps">CfnGlobalClusterConfigProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfig.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnGlobalClusterConfigProps <a name="CfnGlobalClusterConfigProps" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps"></a>

Returns, adds, and removes Global Cluster managed namespaces and custom zone mappings.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.Initializer"></a>

```typescript
import { CfnGlobalClusterConfigProps } from '@mongodbatlas-awscdk/global-cluster-config'

const cfnGlobalClusterConfigProps: CfnGlobalClusterConfigProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.clusterName">clusterName</a></code> | <code>string</code> | The name of the Atlas cluster that contains the snapshots you want to retrieve. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.customZoneMappings">customZoneMappings</a></code> | <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ZoneMapping">ZoneMapping</a>[]</code> | List that contains comma-separated key value pairs to map zones to geographic regions. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.managedNamespaces">managedNamespaces</a></code> | <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace">ManagedNamespace</a>[]</code> | List that contains comma-separated key value pairs to map zones to geographic regions. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.projectId">projectId</a></code> | <code>string</code> | The unique identifier of the project for the Atlas cluster. |

---

##### `clusterName`<sup>Optional</sup> <a name="clusterName" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

The name of the Atlas cluster that contains the snapshots you want to retrieve.

---

##### `customZoneMappings`<sup>Optional</sup> <a name="customZoneMappings" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.customZoneMappings"></a>

```typescript
public readonly customZoneMappings: ZoneMapping[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/global-cluster-config.ZoneMapping">ZoneMapping</a>[]

List that contains comma-separated key value pairs to map zones to geographic regions.

These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.

This parameter returns an empty object if no custom zones exist.

---

##### `managedNamespaces`<sup>Optional</sup> <a name="managedNamespaces" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.managedNamespaces"></a>

```typescript
public readonly managedNamespaces: ManagedNamespace[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace">ManagedNamespace</a>[]

List that contains comma-separated key value pairs to map zones to geographic regions.

These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.

This parameter returns an empty object if no custom zones exist.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/global-cluster-config.CfnGlobalClusterConfigProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

The unique identifier of the project for the Atlas cluster.

---

### ManagedNamespace <a name="ManagedNamespace" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.Initializer"></a>

```typescript
import { ManagedNamespace } from '@mongodbatlas-awscdk/global-cluster-config'

const managedNamespace: ManagedNamespace = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.collection">collection</a></code> | <code>string</code> | Human-readable label of the collection to manage for this Global Cluster. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.customShardKey">customShardKey</a></code> | <code>string</code> | Database parameter used to divide the *collection* into shards. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.db">db</a></code> | <code>string</code> | Human-readable label of the database to manage for this Global Cluster. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.isCustomShardKeyHashed">isCustomShardKeyHashed</a></code> | <code>boolean</code> | Flag that indicates whether someone hashed the custom shard key for the specified collection. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.isShardKeyUnique">isShardKeyUnique</a></code> | <code>boolean</code> | Flag that indicates whether someone [hashed](https://www.mongodb.com/docs/manual/reference/method/sh.shardCollection/#hashed-shard-keys) the custom shard key. If this parameter returns `false`, this cluster uses [ranged sharding](https://www.mongodb.com/docs/manual/core/ranged-sharding/). |

---

##### `collection`<sup>Optional</sup> <a name="collection" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.collection"></a>

```typescript
public readonly collection: string;
```

- *Type:* string

Human-readable label of the collection to manage for this Global Cluster.

---

##### `customShardKey`<sup>Optional</sup> <a name="customShardKey" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.customShardKey"></a>

```typescript
public readonly customShardKey: string;
```

- *Type:* string

Database parameter used to divide the *collection* into shards.

Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key.

---

##### `db`<sup>Optional</sup> <a name="db" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.db"></a>

```typescript
public readonly db: string;
```

- *Type:* string

Human-readable label of the database to manage for this Global Cluster.

---

##### `isCustomShardKeyHashed`<sup>Optional</sup> <a name="isCustomShardKeyHashed" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.isCustomShardKeyHashed"></a>

```typescript
public readonly isCustomShardKeyHashed: boolean;
```

- *Type:* boolean

Flag that indicates whether someone hashed the custom shard key for the specified collection.

If you set this value to `false`, MongoDB Cloud uses ranged sharding.

---

##### `isShardKeyUnique`<sup>Optional</sup> <a name="isShardKeyUnique" id="@mongodbatlas-awscdk/global-cluster-config.ManagedNamespace.property.isShardKeyUnique"></a>

```typescript
public readonly isShardKeyUnique: boolean;
```

- *Type:* boolean

Flag that indicates whether someone [hashed](https://www.mongodb.com/docs/manual/reference/method/sh.shardCollection/#hashed-shard-keys) the custom shard key. If this parameter returns `false`, this cluster uses [ranged sharding](https://www.mongodb.com/docs/manual/core/ranged-sharding/).

---

### ZoneMapping <a name="ZoneMapping" id="@mongodbatlas-awscdk/global-cluster-config.ZoneMapping"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/global-cluster-config.ZoneMapping.Initializer"></a>

```typescript
import { ZoneMapping } from '@mongodbatlas-awscdk/global-cluster-config'

const zoneMapping: ZoneMapping = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ZoneMapping.property.location">location</a></code> | <code>string</code> | Code that represents a location that maps to a zone in your global cluster. |
| <code><a href="#@mongodbatlas-awscdk/global-cluster-config.ZoneMapping.property.zone">zone</a></code> | <code>string</code> | Human-readable label that identifies the zone in your global cluster. |

---

##### `location`<sup>Optional</sup> <a name="location" id="@mongodbatlas-awscdk/global-cluster-config.ZoneMapping.property.location"></a>

```typescript
public readonly location: string;
```

- *Type:* string

Code that represents a location that maps to a zone in your global cluster.

MongoDB Cloud represents this location with a ISO 3166-2 location and subdivision codes when possible.

---

##### `zone`<sup>Optional</sup> <a name="zone" id="@mongodbatlas-awscdk/global-cluster-config.ZoneMapping.property.zone"></a>

```typescript
public readonly zone: string;
```

- *Type:* string

Human-readable label that identifies the zone in your global cluster.

This zone maps to a location code.

---



